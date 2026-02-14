package targets

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/jaycee1285/base16changer/internal/scheme"
	"github.com/jaycee1285/base16changer/internal/template"
)

// Config holds paths and settings for theme application
type Config struct {
	// Scheme directories
	SchemesDir string // Path to base16 schemes (YAML files)

	// Target config paths
	KittyThemeConf string // ~/.config/kitty/current-theme.conf
	FuzzelIni      string // ~/.config/fuzzel/fuzzel.ini
	Gtk2RC         string // ~/.themes/Base16/gtk-2.0/gtkrc
	Gtk3CSS        string // ~/.themes/Base16/gtk-3.0/gtk.css
	Gtk4CSS        string // ~/.config/gtk-4.0/gtk.css (libadwaita)
	Gtk4ThemeCSS   string // ~/.themes/Base16/gtk-4.0/gtk.css
	IndexTheme     string // ~/.themes/Base16/index.theme
	OpenboxThemerc string // ~/.themes/Base16/openbox-3/themerc
	LabwcRcXml     string // ~/.config/labwc/rc.xml

	// Openbox theme name (written to rc.xml)
	OpenboxThemeName string

	// GTK base theme name (for dconf toggle trick)
	GtkThemeName string

	// Icon theme (optional, set via flag)
	IconTheme string

	// Wallpaper (optional, set via flag)
	Wallpaper    string
	WallpaperDir string

	// Dry run mode - print what would be done
	DryRun bool

	// Quiet mode - suppress stdout logging (useful for TUI)
	Quiet bool

	// Ferritebar config to touch after apply
	FerritebarConfig string
}

// DefaultConfig returns config with standard paths
func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	return &Config{
		// SchemesDir is used for CLI --schemes-dir override only
		// ScanSchemesDirs() returns the actual search paths
		KittyThemeConf:   filepath.Join(home, ".config/kitty/current-theme.conf"),
		FuzzelIni:        filepath.Join(home, ".config/fuzzel/fuzzel.ini"),
		Gtk2RC:           filepath.Join(home, ".themes/Base16/gtk-2.0/gtkrc"),
		Gtk3CSS:          filepath.Join(home, ".themes/Base16/gtk-3.0/gtk.css"),
		Gtk4CSS:          filepath.Join(home, ".config/gtk-4.0/gtk.css"),
		Gtk4ThemeCSS:     filepath.Join(home, ".themes/Base16/gtk-4.0/gtk.css"),
		IndexTheme:       filepath.Join(home, ".themes/Base16/index.theme"),
		OpenboxThemerc:   filepath.Join(home, ".themes/Base16/openbox-3/themerc"),
		LabwcRcXml:       filepath.Join(home, ".config/labwc/rc.xml"),
		OpenboxThemeName: "Base16",
		GtkThemeName:     "Base16",
		WallpaperDir:     filepath.Join(home, "Pictures/walls"),
		DryRun:           false,
		Quiet:            false,
		FerritebarConfig: filepath.Join(home, ".config/ferritebar/config.toml"),
	}
}

// Apply applies a base16 scheme to all targets
func Apply(cfg *Config, s *scheme.Base16) error {
	logf(cfg, "Applying scheme: %s\n", s.Name)

	// 1. Kitty
	if err := applyKitty(cfg, s); err != nil {
		logf(cfg, "  [WARN] kitty: %v\n", err)
	} else {
		logln(cfg, "  [OK] kitty")
	}

	// 2. Fuzzel
	if err := applyFuzzel(cfg, s); err != nil {
		logf(cfg, "  [WARN] fuzzel: %v\n", err)
	} else {
		logln(cfg, "  [OK] fuzzel")
	}

	// 3. GTK-4
	if err := applyGtk4(cfg, s); err != nil {
		logf(cfg, "  [WARN] gtk-4: %v\n", err)
	} else {
		logln(cfg, "  [OK] gtk-4")
	}

	// 4. GTK-2
	if err := applyGtk2(cfg, s); err != nil {
		logf(cfg, "  [WARN] gtk-2: %v\n", err)
	} else {
		logln(cfg, "  [OK] gtk-2")
	}

	// 5. GTK-3 (theme directory)
	if err := applyGtk3(cfg, s); err != nil {
		logf(cfg, "  [WARN] gtk-3: %v\n", err)
	} else {
		logln(cfg, "  [OK] gtk-3")
	}
	// Clean up old user CSS that would override theme colors
	cleanupOldGtkCSS(cfg)

	// 4b. Theme index.theme
	if err := applyIndexTheme(cfg); err != nil {
		logf(cfg, "  [WARN] index.theme: %v\n", err)
	} else {
		logln(cfg, "  [OK] index.theme")
	}

	// 4c. GTK settings.ini (set theme name)
	home, _ := os.UserHomeDir()
	for _, iniPath := range []string{
		filepath.Join(home, ".config/gtk-3.0/settings.ini"),
		filepath.Join(home, ".config/gtk-4.0/settings.ini"),
	} {
		if err := updateGtkSettingsIni(cfg, iniPath); err != nil {
			logf(cfg, "  [WARN] %s: %v\n", iniPath, err)
		}
	}

	// 5. LabWC/Openbox themerc
	if err := applyOpenbox(cfg, s); err != nil {
		logf(cfg, "  [WARN] openbox: %v\n", err)
	} else {
		logln(cfg, "  [OK] openbox")
	}

	// 6. LabWC rc.xml (set theme name and icon theme)
	if err := updateLabwcRcXml(cfg); err != nil {
		logf(cfg, "  [WARN] labwc rc.xml: %v\n", err)
	} else {
		logln(cfg, "  [OK] labwc rc.xml")
	}

	// 7. Icon theme (if specified)
	if cfg.IconTheme != "" {
		if err := applyIconTheme(cfg); err != nil {
			logf(cfg, "  [WARN] icon theme: %v\n", err)
		} else {
			logln(cfg, "  [OK] icon theme")
		}
	}

	// 8. Wallpaper (if specified)
	if cfg.Wallpaper != "" {
		if err := applyWallpaper(cfg); err != nil {
			logf(cfg, "  [WARN] wallpaper: %v\n", err)
		} else {
			logln(cfg, "  [OK] wallpaper")
		}
	}

	// 9. Trigger reloads
	logln(cfg, "\nTriggering reloads...")
	triggerReloads(cfg)

	// 10. Touch ferritebar config (final step)
	if err := touchFerritebarConfig(cfg); err != nil {
		logf(cfg, "  [WARN] ferritebar config: %v\n", err)
	} else {
		logln(cfg, "  [OK] ferritebar config")
	}

	return nil
}

func applyKitty(cfg *Config, s *scheme.Base16) error {
	content, err := template.RenderString(kittyTemplate, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.KittyThemeConf, content)
}

func applyFuzzel(cfg *Config, s *scheme.Base16) error {
	colorsSection, err := template.RenderString(fuzzelTemplate, s.ToMap())
	if err != nil {
		return err
	}

	if cfg.DryRun {
		logf(cfg, "  Would update [colors] in: %s\n", cfg.FuzzelIni)
		return nil
	}

	// Read existing file
	existing, err := os.ReadFile(cfg.FuzzelIni)
	if err != nil {
		// File doesn't exist, create with just colors
		return writeFileForce(cfg.FuzzelIni, colorsSection)
	}

	// Replace or append [colors] section
	newContent := replaceIniSection(string(existing), "colors", colorsSection)
	return writeFileForce(cfg.FuzzelIni, newContent)
}

// replaceIniSection replaces a [section] in INI content, or appends if not found
func replaceIniSection(content, sectionName, newSection string) string {
	lines := strings.Split(content, "\n")
	var result []string
	inSection := false
	sectionFound := false
	sectionHeader := "[" + sectionName + "]"

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check if we're entering the target section
		if strings.EqualFold(trimmed, sectionHeader) {
			inSection = true
			sectionFound = true
			// Add the new section content (without trailing newline)
			result = append(result, strings.TrimSuffix(newSection, "\n"))
			continue
		}

		// Check if we're entering a different section (leaving target section)
		if inSection && strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			inSection = false
		}

		// Skip lines while in the target section (we already added new content)
		if inSection {
			continue
		}

		result = append(result, line)
	}

	// If section wasn't found, append it
	if !sectionFound {
		result = append(result, "")
		result = append(result, strings.TrimSuffix(newSection, "\n"))
	}

	return strings.Join(result, "\n")
}

func writeFileForce(path, content string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func applyGtk4(cfg *Config, s *scheme.Base16) error {
	content, err := template.RenderString(gtk4Template, s.ToMap())
	if err != nil {
		return err
	}
	// Write to theme directory for plain GTK-4 apps
	if err := writeFile(cfg, cfg.Gtk4ThemeCSS, content); err != nil {
		return err
	}
	// Write to user CSS for libadwaita apps (they ignore theme directories)
	return writeFile(cfg, cfg.Gtk4CSS, content)
}

func applyGtk3(cfg *Config, s *scheme.Base16) error {
	content, err := template.RenderString(gtk3Template, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.Gtk3CSS, content)
}

func applyGtk2(cfg *Config, s *scheme.Base16) error {
	content, err := template.RenderString(gtk2Template, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.Gtk2RC, content)
}

func cleanupOldGtkCSS(cfg *Config) {
	home, _ := os.UserHomeDir()
	old := filepath.Join(home, ".config/gtk-3.0/gtk.css")
	if _, err := os.Stat(old); err == nil {
		if cfg.DryRun {
			logf(cfg, "  Would remove old user CSS: %s\n", old)
			return
		}
		if err := os.Remove(old); err != nil {
			logf(cfg, "  [WARN] remove old gtk-3 css: %v\n", err)
		}
	}
}

func applyIndexTheme(cfg *Config) error {
	return writeFile(cfg, cfg.IndexTheme, indexThemeTemplate)
}

func updateGtkSettingsIni(cfg *Config, path string) error {
	if cfg.DryRun {
		logf(cfg, "  Would update gtk-theme-name in: %s\n", path)
		return nil
	}

	existing, err := os.ReadFile(path)
	if err != nil {
		// File doesn't exist, skip (managed by NixOS/home-manager)
		return nil
	}

	lines := strings.Split(string(existing), "\n")
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, "gtk-theme-name=") {
			lines[i] = "gtk-theme-name=" + cfg.GtkThemeName
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

func applyOpenbox(cfg *Config, s *scheme.Base16) error {
	// Ensure theme directory exists
	themeDir := filepath.Dir(cfg.OpenboxThemerc)
	if !cfg.DryRun {
		if err := os.MkdirAll(themeDir, 0755); err != nil {
			return err
		}
	}

	content, err := template.RenderString(openboxTemplate, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.OpenboxThemerc, content)
}

func updateLabwcRcXml(cfg *Config) error {
	if cfg.DryRun {
		logf(cfg, "  Would update theme name in: %s\n", cfg.LabwcRcXml)
		return nil
	}

	content, err := os.ReadFile(cfg.LabwcRcXml)
	if err != nil {
		return fmt.Errorf("read rc.xml: %w", err)
	}

	// Replace <theme><name>...</name> with our theme name
	// Match: <theme>...<name>something</name>
	re := regexp.MustCompile(`(<theme>\s*\n\s*<name>)[^<]*(</name>)`)
	newContent := re.ReplaceAllString(string(content), "${1}"+cfg.OpenboxThemeName+"${2}")

	if string(content) == newContent {
		// No change needed or pattern not found
		return nil
	}

	return os.WriteFile(cfg.LabwcRcXml, []byte(newContent), 0644)
}

func applyIconTheme(cfg *Config) error {
	if cfg.DryRun {
		logf(cfg, "  Would set icon theme: %s\n", cfg.IconTheme)
		return nil
	}

	// Update via dconf
	if err := run("dconf", "write", "/org/gnome/desktop/interface/icon-theme", fmt.Sprintf("'%s'", cfg.IconTheme)); err != nil {
		return fmt.Errorf("dconf icon-theme: %w", err)
	}

	return nil
}

func applyWallpaper(cfg *Config) error {
	if cfg.DryRun {
		logf(cfg, "  Would set wallpaper: %s\n", cfg.Wallpaper)
		return nil
	}

	wpPath := filepath.Join(cfg.WallpaperDir, cfg.Wallpaper)
	if err := run("swww", "img", wpPath); err != nil {
		return fmt.Errorf("swww: %w", err)
	}

	return nil
}

func writeFile(cfg *Config, path, content string) error {
	if cfg.DryRun {
		logf(cfg, "  Would write to: %s\n", path)
		return nil
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}

	return os.WriteFile(path, []byte(content), 0644)
}

func triggerReloads(cfg *Config) {
	if cfg.DryRun {
		logln(cfg, "  Would run: pkill -SIGUSR1 kitty")
		logln(cfg, "  Would run: labwc -r")
		logln(cfg, "  Would run: dconf toggle gtk-theme")
		return
	}

	// Kitty - SIGUSR1 tells kitty to reload its config
	if err := run("pkill", "-SIGUSR1", "kitty"); err != nil {
		logf(cfg, "  [WARN] kitty reload: %v\n", err)
	} else {
		logln(cfg, "  [OK] kitty reload")
	}

	// LabWC
	if err := run("labwc", "-r"); err != nil {
		logf(cfg, "  [WARN] labwc reconfigure: %v\n", err)
	} else {
		logln(cfg, "  [OK] labwc reconfigure")
	}

	// GTK reload via dconf toggle
	_ = run("dconf", "write", "/org/gnome/desktop/interface/gtk-theme", "'dummy'")
	if err := run("dconf", "write", "/org/gnome/desktop/interface/gtk-theme", fmt.Sprintf("'%s'", cfg.GtkThemeName)); err != nil {
		logf(cfg, "  [WARN] gtk reload: %v\n", err)
	} else {
		logln(cfg, "  [OK] gtk reload")
	}
}

func touchFerritebarConfig(cfg *Config) error {
	if cfg.DryRun {
		logf(cfg, "  Would touch: %s\n", cfg.FerritebarConfig)
		return nil
	}

	dir := filepath.Dir(cfg.FerritebarConfig)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}

	f, err := os.OpenFile(cfg.FerritebarConfig, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("touch ferritebar config: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close ferritebar config: %w", err)
	}

	now := time.Now()
	if err := os.Chtimes(cfg.FerritebarConfig, now, now); err != nil {
		return fmt.Errorf("chtimes ferritebar config: %w", err)
	}

	return nil
}

func logf(cfg *Config, format string, args ...any) {
	if cfg != nil && cfg.Quiet {
		return
	}
	fmt.Printf(format, args...)
}

func logln(cfg *Config, args ...any) {
	if cfg != nil && cfg.Quiet {
		return
	}
	fmt.Println(args...)
}

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
