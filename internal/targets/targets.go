package targets

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

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
	Gtk3CSS        string // ~/.config/gtk-3.0/gtk.css
	Gtk4CSS        string // ~/.config/gtk-4.0/gtk.css
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
}

// DefaultConfig returns config with standard paths
func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	return &Config{
		// SchemesDir is used for CLI --schemes-dir override only
		// ScanSchemesDirs() returns the actual search paths
		KittyThemeConf: filepath.Join(home, ".config/kitty/current-theme.conf"),
		FuzzelIni:      filepath.Join(home, ".config/fuzzel/fuzzel.ini"),
		Gtk3CSS:        filepath.Join(home, ".config/gtk-3.0/gtk.css"),
		Gtk4CSS:          filepath.Join(home, ".config/gtk-4.0/gtk.css"),
		OpenboxThemerc:   filepath.Join(home, ".themes/Base16/openbox-3/themerc"),
		LabwcRcXml:       filepath.Join(home, ".config/labwc/rc.xml"),
		OpenboxThemeName: "Base16",
		GtkThemeName:     "Base16",
		WallpaperDir:     filepath.Join(home, "Pictures/walls"),
		DryRun:           false,
	}
}

// Apply applies a base16 scheme to all targets
func Apply(cfg *Config, s *scheme.Base16) error {
	fmt.Printf("Applying scheme: %s\n", s.Name)

	// 1. Kitty
	if err := applyKitty(cfg, s); err != nil {
		fmt.Printf("  [WARN] kitty: %v\n", err)
	} else {
		fmt.Println("  [OK] kitty")
	}

	// 2. Fuzzel
	if err := applyFuzzel(cfg, s); err != nil {
		fmt.Printf("  [WARN] fuzzel: %v\n", err)
	} else {
		fmt.Println("  [OK] fuzzel")
	}

	// 3. GTK-4
	if err := applyGtk4(cfg, s); err != nil {
		fmt.Printf("  [WARN] gtk-4: %v\n", err)
	} else {
		fmt.Println("  [OK] gtk-4")
	}

	// 4. GTK-3
	if err := applyGtk3(cfg, s); err != nil {
		fmt.Printf("  [WARN] gtk-3: %v\n", err)
	} else {
		fmt.Println("  [OK] gtk-3")
	}

	// 5. LabWC/Openbox themerc
	if err := applyOpenbox(cfg, s); err != nil {
		fmt.Printf("  [WARN] openbox: %v\n", err)
	} else {
		fmt.Println("  [OK] openbox")
	}

	// 6. LabWC rc.xml (set theme name and icon theme)
	if err := updateLabwcRcXml(cfg); err != nil {
		fmt.Printf("  [WARN] labwc rc.xml: %v\n", err)
	} else {
		fmt.Println("  [OK] labwc rc.xml")
	}

	// 7. Icon theme (if specified)
	if cfg.IconTheme != "" {
		if err := applyIconTheme(cfg); err != nil {
			fmt.Printf("  [WARN] icon theme: %v\n", err)
		} else {
			fmt.Println("  [OK] icon theme")
		}
	}

	// 8. Wallpaper (if specified)
	if cfg.Wallpaper != "" {
		if err := applyWallpaper(cfg); err != nil {
			fmt.Printf("  [WARN] wallpaper: %v\n", err)
		} else {
			fmt.Println("  [OK] wallpaper")
		}
	}

	// 9. Trigger reloads
	fmt.Println("\nTriggering reloads...")
	triggerReloads(cfg)

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
		fmt.Printf("  Would update [colors] in: %s\n", cfg.FuzzelIni)
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
	return writeFile(cfg, cfg.Gtk4CSS, content)
}

func applyGtk3(cfg *Config, s *scheme.Base16) error {
	content, err := template.RenderString(gtk3Template, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.Gtk3CSS, content)
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
		fmt.Printf("  Would update theme name in: %s\n", cfg.LabwcRcXml)
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
		fmt.Printf("  Would set icon theme: %s\n", cfg.IconTheme)
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
		fmt.Printf("  Would set wallpaper: %s\n", cfg.Wallpaper)
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
		fmt.Printf("  Would write to: %s\n", path)
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
		fmt.Println("  Would run: pkill -SIGUSR1 kitty")
		fmt.Println("  Would run: labwc -r")
		fmt.Println("  Would run: dconf toggle gtk-theme")
		return
	}

	// Kitty - SIGUSR1 tells kitty to reload its config
	if err := run("pkill", "-SIGUSR1", "kitty"); err != nil {
		fmt.Printf("  [WARN] kitty reload: %v\n", err)
	} else {
		fmt.Println("  [OK] kitty reload")
	}

	// LabWC
	if err := run("labwc", "-r"); err != nil {
		fmt.Printf("  [WARN] labwc reconfigure: %v\n", err)
	} else {
		fmt.Println("  [OK] labwc reconfigure")
	}

	// GTK reload via dconf toggle
	_ = run("dconf", "write", "/org/gnome/desktop/interface/gtk-theme", "'dummy'")
	if err := run("dconf", "write", "/org/gnome/desktop/interface/gtk-theme", fmt.Sprintf("'%s'", cfg.GtkThemeName)); err != nil {
		fmt.Printf("  [WARN] gtk reload: %v\n", err)
	} else {
		fmt.Println("  [OK] gtk reload")
	}
}

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
