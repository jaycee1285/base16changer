package targets

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/jaycee1285/base16changer/internal/orchis"
	"github.com/jaycee1285/base16changer/internal/scheme"
	"github.com/jaycee1285/base16changer/internal/template"
)

// Config holds paths and settings for theme application
type Config struct {
	// Scheme directories
	SchemesDir string // Path to base16 schemes (YAML files)
	SelectedScheme string

	// Target config paths
	KittyThemeConf string // ~/.config/kitty/current-theme.conf
	FuzzelIni      string // ~/.config/fuzzel/fuzzel.ini
	SyntectThemeDir string // ~/.local/share/themes/tmThemes
	SyntectCurrent  string // ~/.config/syntect/current.tmTheme
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

	// Mako notification config
	MakoConfig string

	// LibreWolf/Firefox colors.css
	LibrewolfCSS string

	// Firefox theme payload for Firefox Color/theme API consumers
	FirefoxThemeJSON string

	// Packaged Firefox/LibreWolf theme extension
	FirefoxExtensionDir string
	FirefoxExtensionXPI string
	FirefoxExtensionID  string

	// Sidebery sidebar tab extension CSS
	SideberyCSS string

	// Ferritebar config to touch after apply
	FerritebarConfig string

	// Orchis theme output directory (e.g. ~/.local/share/themes)
	OrchisDestDir string

	// DarkMode selects Orchis-Dark-Compact vs Orchis-Light-Compact
	DarkMode bool
}

// DefaultConfig returns config with standard paths
func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	return &Config{
		// SchemesDir is used for CLI --schemes-dir override only
		// ScanSchemesDirs() returns the actual search paths
		KittyThemeConf:   filepath.Join(home, ".config/kitty/current-theme.conf"),
		FuzzelIni:        filepath.Join(home, ".config/fuzzel/fuzzel.ini"),
		SyntectThemeDir:  filepath.Join(home, ".local/share/themes/tmThemes"),
		SyntectCurrent:   filepath.Join(home, ".config/syntect/current.tmTheme"),
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
		MakoConfig:       filepath.Join(home, ".config/mako/config"),
		LibrewolfCSS:     filepath.Join(home, ".config/base16changer/librewolf/colors.css"),
		FirefoxThemeJSON: filepath.Join(home, ".config/base16changer/firefox/theme.json"),
		FirefoxExtensionDir: filepath.Join(home, ".config/base16changer/firefox/extension"),
		FirefoxExtensionXPI: filepath.Join(home, ".config/base16changer/firefox/base16changer-theme.xpi"),
		FirefoxExtensionID:  "base16changer-theme@john",
		SideberyCSS:      filepath.Join(home, ".config/base16changer/sidebery/styles.css"),
		FerritebarConfig: filepath.Join(home, ".config/ferritebar/config.toml"),
		OrchisDestDir:    filepath.Join(home, ".local/share/themes"),
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

	// 2b. Syntect current tmTheme
	if err := applySyntectCurrentTheme(cfg, s); err != nil {
		logf(cfg, "  [WARN] syntect: %v\n", err)
	} else {
		logln(cfg, "  [OK] syntect")
	}

	// 3. Orchis theme (GTK 3/4 compiled from SCSS)
	var orchisVariant orchis.Variant
	if cfg.DarkMode {
		orchisVariant = orchis.Dark
	} else {
		orchisVariant = orchis.Light
	}
	cfg.GtkThemeName = orchis.ThemeName(orchisVariant)

	if err := orchis.Build(s, orchisVariant, cfg.OrchisDestDir); err != nil {
		logf(cfg, "  [WARN] orchis: %v\n", err)
	} else {
		logln(cfg, "  [OK] orchis ("+string(orchisVariant)+")")
	}

	// 4. GTK-2
	if err := applyGtk2(cfg, s); err != nil {
		logf(cfg, "  [WARN] gtk-2: %v\n", err)
	} else {
		logln(cfg, "  [OK] gtk-2")
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

	// 7. Mako
	if err := applyMako(cfg, s); err != nil {
		logf(cfg, "  [WARN] mako: %v\n", err)
	} else {
		logln(cfg, "  [OK] mako")
	}

	// 7b. LibreWolf/Firefox
	if err := applyLibrewolf(cfg, s); err != nil {
		logf(cfg, "  [WARN] librewolf: %v\n", err)
	} else {
		logln(cfg, "  [OK] librewolf")
	}

	// 7c. Firefox theme payload (preferred integration point for Sidebery)
	if err := applyFirefoxTheme(cfg, s); err != nil {
		logf(cfg, "  [WARN] firefox theme: %v\n", err)
	} else {
		logln(cfg, "  [OK] firefox theme")
	}

	// 7d. Packaged Firefox theme extension (.xpi)
	if err := applyFirefoxThemeExtension(cfg, s); err != nil {
		logf(cfg, "  [WARN] firefox theme xpi: %v\n", err)
	} else {
		logln(cfg, "  [OK] firefox theme xpi")
	}

	// 7e. Sidebery (sidebar tab extension CSS fallback)
	if err := applySidebery(cfg, s); err != nil {
		logf(cfg, "  [WARN] sidebery: %v\n", err)
	} else {
		logln(cfg, "  [OK] sidebery")
	}

	// 8. Icon theme (if specified)
	if cfg.IconTheme != "" {
		if err := applyIconTheme(cfg); err != nil {
			logf(cfg, "  [WARN] icon theme: %v\n", err)
		} else {
			logln(cfg, "  [OK] icon theme")
		}
		if err := recolorIconTheme(cfg, s); err != nil {
			logf(cfg, "  [WARN] icon recolor: %v\n", err)
		} else {
			logln(cfg, "  [OK] icon recolor")
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

func applySyntectCurrentTheme(cfg *Config, s *scheme.Base16) error {
	src, err := findSyntectThemeSource(cfg, s)
	if err != nil {
		return err
	}
	if cfg.DryRun {
		logf(cfg, "  Would copy %s -> %s\n", src, cfg.SyntectCurrent)
		return nil
	}

	content, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("read tmTheme %s: %w", src, err)
	}

	dir := filepath.Dir(cfg.SyntectCurrent)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}

	if err := os.WriteFile(cfg.SyntectCurrent, content, 0644); err != nil {
		return fmt.Errorf("write current tmTheme %s: %w", cfg.SyntectCurrent, err)
	}

	return nil
}

func findSyntectThemeSource(cfg *Config, s *scheme.Base16) (string, error) {
	themeDirs := syntectThemeDirs(cfg)
	candidates := syntectThemeCandidates(cfg, s)

	for _, dir := range themeDirs {
		for _, candidate := range candidates {
			path := filepath.Join(dir, candidate+".tmTheme")
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
		}
	}

	return "", fmt.Errorf(
		"tmTheme not found for scheme %q (tried names: %s in dirs: %s)",
		cfg.SelectedScheme,
		strings.Join(candidates, ", "),
		strings.Join(themeDirs, ", "),
	)
}

func syntectThemeDirs(cfg *Config) []string {
	dirs := []string{cfg.SyntectThemeDir}
	if strings.Contains(cfg.SyntectThemeDir, "tmThemes") {
		dirs = append(dirs, strings.Replace(cfg.SyntectThemeDir, "tmThemes", "tmthemes", 1))
	}
	if strings.Contains(cfg.SyntectThemeDir, "tmthemes") {
		dirs = append(dirs, strings.Replace(cfg.SyntectThemeDir, "tmthemes", "tmThemes", 1))
	}
	return dedupeStrings(dirs)
}

func syntectThemeCandidates(cfg *Config, s *scheme.Base16) []string {
	var raw []string

	add := func(name string) {
		name = strings.TrimSpace(name)
		if name == "" {
			return
		}
		raw = append(raw, name)
		raw = append(raw, strings.ReplaceAll(name, " ", "-"))
		raw = append(raw, strings.ReplaceAll(name, "-", " "))
		raw = append(raw, strings.ReplaceAll(name, " ", "_"))
		raw = append(raw, strings.ReplaceAll(name, "_", " "))
		raw = append(raw, strings.ToLower(name))
	}

	add(cfg.SelectedScheme)
	if s != nil {
		add(s.Name)
	}

	return dedupeStrings(raw)
}

func dedupeStrings(items []string) []string {
	set := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	for _, item := range items {
		if _, exists := set[item]; exists {
			continue
		}
		set[item] = struct{}{}
		out = append(out, item)
	}
	sort.Strings(out)
	return out
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

func applyMako(cfg *Config, s *scheme.Base16) error {
	m := s.ToMap()
	colors := map[string]string{
		"background-color": "#" + m["accent-hex"] + "FF", // placeholder, overwritten below
		"text-color":       "#" + m["accent-hex"] + "FF",
		"border-color":     "#" + m["accent-hex"] + "FF",
	}
	// Render the template to get the actual values
	rendered, err := template.RenderString(makoTemplate, m)
	if err != nil {
		return err
	}
	// Parse rendered template for the color values
	for _, line := range strings.Split(rendered, "\n") {
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		if idx := strings.Index(line, "="); idx > 0 {
			key := strings.TrimSpace(line[:idx])
			val := strings.TrimSpace(line[idx+1:])
			if _, ok := colors[key]; ok {
				colors[key] = val
			}
		}
	}

	if cfg.DryRun {
		logf(cfg, "  Would update colors in: %s\n", cfg.MakoConfig)
		return nil
	}

	existing, err := os.ReadFile(cfg.MakoConfig)
	if err != nil {
		// File doesn't exist, write the full rendered template
		return writeFile(cfg, cfg.MakoConfig, rendered)
	}

	// Update existing keys or append missing ones
	lines := strings.Split(string(existing), "\n")
	found := make(map[string]bool)
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if idx := strings.Index(trimmed, "="); idx > 0 {
			key := strings.TrimSpace(trimmed[:idx])
			if val, ok := colors[key]; ok {
				lines[i] = key + "=" + val
				found[key] = true
			}
		}
	}
	// Append any keys not already in the file
	for key, val := range colors {
		if !found[key] {
			lines = append(lines, key+"="+val)
		}
	}

	return os.WriteFile(cfg.MakoConfig, []byte(strings.Join(lines, "\n")), 0644)
}

func applyLibrewolf(cfg *Config, s *scheme.Base16) error {
	content, err := template.RenderString(librewolfTemplate, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.LibrewolfCSS, content)
}

func applyFirefoxTheme(cfg *Config, s *scheme.Base16) error {
	content, err := template.RenderString(firefoxThemeTemplate, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.FirefoxThemeJSON, content)
}

func applyFirefoxThemeExtension(cfg *Config, s *scheme.Base16) error {
	data := s.ToMap()
	data["extension-id"] = cfg.FirefoxExtensionID
	data["extension-version"] = firefoxExtensionVersion()

	manifest, err := template.RenderString(firefoxManifestTemplate, data)
	if err != nil {
		return err
	}
	background, err := template.RenderString(firefoxBackgroundTemplate, data)
	if err != nil {
		return err
	}
	themeJSON, err := template.RenderString(firefoxThemeTemplate, data)
	if err != nil {
		return err
	}

	if cfg.DryRun {
		logf(cfg, "  Would write Firefox extension files to: %s\n", cfg.FirefoxExtensionDir)
		logf(cfg, "  Would package XPI at: %s\n", cfg.FirefoxExtensionXPI)
		return nil
	}

	if err := os.RemoveAll(cfg.FirefoxExtensionDir); err != nil {
		return fmt.Errorf("clean firefox extension dir: %w", err)
	}
	if err := os.MkdirAll(cfg.FirefoxExtensionDir, 0755); err != nil {
		return fmt.Errorf("mkdir firefox extension dir: %w", err)
	}

	files := map[string]string{
		"manifest.json": manifest,
		"background.js": background,
		"theme.json":    themeJSON,
	}
	for name, content := range files {
		path := filepath.Join(cfg.FirefoxExtensionDir, name)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("write firefox extension file %s: %w", path, err)
		}
	}

	if err := os.Remove(cfg.FirefoxExtensionXPI); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove old firefox xpi: %w", err)
	}

	cmd := exec.Command(
		"7z", "a", "-tzip", cfg.FirefoxExtensionXPI,
		"manifest.json", "background.js", "theme.json",
	)
	cmd.Dir = cfg.FirefoxExtensionDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("package firefox xpi: %w", err)
	}

	return nil
}

func applySidebery(cfg *Config, s *scheme.Base16) error {
	content, err := template.RenderString(sideberyTemplate, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.SideberyCSS, content)
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
	if err := run("awww", "img", wpPath); err != nil {
		return fmt.Errorf("awww: %w", err)
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
		logln(cfg, "  Would run: makoctl reload")
		logln(cfg, "  Would run: dconf toggle gtk-theme")
		return
	}

	// Kitty - SIGUSR1 tells kitty to reload its config
	if err := run("pkill", "-SIGUSR1", "kitty"); err != nil {
		logf(cfg, "  [WARN] kitty reload: %v\n", err)
	} else {
		logln(cfg, "  [OK] kitty reload")
	}

	// LabWC / SartWC
	for _, compositor := range []string{"labwc", "sartwc"} {
		if err := run("pgrep", compositor); err == nil {
			if err := run(compositor, "-r"); err != nil {
				logf(cfg, "  [WARN] %s reconfigure: %v\n", compositor, err)
			} else {
				logf(cfg, "  [OK] %s reconfigure\n", compositor)
			}
			break
		}
	}

	// Mako
	if err := run("makoctl", "reload"); err != nil {
		logf(cfg, "  [WARN] mako reload: %v\n", err)
	} else {
		logln(cfg, "  [OK] mako reload")
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

func firefoxExtensionVersion() string {
	now := time.Now()
	return fmt.Sprintf("%d.%d.%d.%d", now.Year(), now.Month(), now.Day(), now.Hour()*100+now.Minute())
}
