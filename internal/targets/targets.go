package targets

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jaycee1285/base16changer/internal/gtksourceview"
	"github.com/jaycee1285/base16changer/internal/orchis"
	"github.com/jaycee1285/base16changer/internal/scheme"
	"github.com/jaycee1285/base16changer/internal/template"
)

// Config holds paths and settings for theme application
type Config struct {
	// Scheme directories
	SchemesDir     string // Path to base16 schemes (YAML files)
	SelectedScheme string

	// Target config paths
	KittyThemeConf        string // ~/.config/kitty/current-theme.conf
	FuzzelIni             string // ~/.config/fuzzel/fuzzel.ini
	SyntectThemeDir       string // ~/.local/share/themes/tmThemes
	SyntectCurrent        string // ~/.config/syntect/current.tmTheme
	GtkSourceview5Current string // ~/.config/gtksourceview-5/current.xml
	Gtk2RC                string // ~/.themes/Base16/gtk-2.0/gtkrc
	Gtk3CSS               string // ~/.themes/Base16/gtk-3.0/gtk.css
	Gtk4CSS               string // ~/.config/gtk-4.0/gtk.css (libadwaita)
	Gtk4ThemeCSS          string // ~/.themes/Base16/gtk-4.0/gtk.css
	IndexTheme            string // ~/.themes/Base16/index.theme
	LabwcRcXml            string // config repo labwc rc.xml
	GowallThemeJSON       string // throwaway gowall theme JSON for current scheme
	WallpaperCurrent      string // current generated wallpaper image

	// GTK base theme name (for dconf toggle trick + openbox path resolution).
	// Set per-apply to Orchis-Light-Compact or Orchis-Dark-Compact based on
	// the scheme's variant; openbox themerc is written to
	// <OrchisDestDir>/<GtkThemeName>/openbox-3/themerc.
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

	// Tiny JSON payload consumed by the `base16-accent` WebExtension's
	// native-messaging host to colour LibreWolf private windows. See
	// tools/firefox-extension/ for the extension + host install steps.
	FirefoxAccentJSON string

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
		KittyThemeConf:        filepath.Join(home, ".config/kitty/current-theme.conf"),
		FuzzelIni:             filepath.Join(home, ".config/fuzzel/fuzzel.ini"),
		SyntectThemeDir:       filepath.Join(home, ".local/share/themes/tmThemes"),
		SyntectCurrent:        filepath.Join(home, ".config/syntect/current.tmTheme"),
		GtkSourceview5Current: filepath.Join(home, ".config/gtksourceview-5/current.xml"),
		Gtk2RC:                filepath.Join(home, ".themes/Base16/gtk-2.0/gtkrc"),
		Gtk3CSS:               filepath.Join(home, ".themes/Base16/gtk-3.0/gtk.css"),
		Gtk4CSS:               filepath.Join(home, ".config/gtk-4.0/gtk.css"),
		Gtk4ThemeCSS:          filepath.Join(home, ".themes/Base16/gtk-4.0/gtk.css"),
		IndexTheme:            filepath.Join(home, ".themes/Base16/index.theme"),
		LabwcRcXml:            filepath.Join(home, "repos/config/home/labwc/rc.xml"),
		GtkThemeName:          "Orchis-Light-Compact", // overwritten per-apply by Apply()
		WallpaperDir:          WallpaperDir(),
		GowallThemeJSON:       filepath.Join(os.TempDir(), "current-gowall.json"),
		WallpaperCurrent:      filepath.Join(home, ".cache/base16changer/wallpaper-current.png"),
		DryRun:                false,
		Quiet:                 false,
		MakoConfig:            filepath.Join(home, ".config/mako/config"),
		LibrewolfCSS:          filepath.Join(home, ".config/base16changer/librewolf/colors.css"),
		FirefoxAccentJSON:     filepath.Join(home, ".config/base16changer/accent.json"),
		SideberyCSS:           filepath.Join(home, ".config/base16changer/sidebery/styles.css"),
		FerritebarConfig:      filepath.Join(home, ".config/ferritebar/config.toml"),
		OrchisDestDir:         filepath.Join(home, ".local/share/themes"),
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

	// 2c. GtkSourceView 5 current style scheme
	if err := applyGtkSourceview5Current(cfg, s); err != nil {
		logf(cfg, "  [WARN] gtksourceview-5: %v\n", err)
	} else {
		logln(cfg, "  [OK] gtksourceview-5")
	}

	// 3. Orchis theme (GTK 3/4 compiled from SCSS)
	var orchisVariant orchis.Variant
	if cfg.DarkMode {
		orchisVariant = orchis.Dark
	} else {
		orchisVariant = orchis.Light
	}
	cfg.GtkThemeName = orchis.ThemeName(orchisVariant)

	if cfg.DryRun {
		logf(cfg, "  Would build orchis (%s) in: %s\n", string(orchisVariant), cfg.OrchisDestDir)
		logln(cfg, "  [OK] orchis ("+string(orchisVariant)+")")
	} else if err := orchis.Build(s, orchisVariant, cfg.OrchisDestDir); err != nil {
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

	// 7c. accent.json — consumed by the base16-accent WebExtension via
	// native messaging to theme LibreWolf private windows. The extension
	// itself is a one-time install (see tools/firefox-extension/).
	if err := applyFirefoxAccent(cfg, s); err != nil {
		logf(cfg, "  [WARN] firefox accent: %v\n", err)
	} else {
		logln(cfg, "  [OK] firefox accent")
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
		if err := applyWallpaper(cfg, s); err != nil {
			logf(cfg, "  [WARN] wallpaper: %v\n", err)
		} else {
			logln(cfg, "  [OK] wallpaper")
		}
	}

	// 9. Trigger reloads
	logln(cfg, "\nTriggering reloads...")
	triggerReloads(cfg)

	// 10. Touch ferritebar config
	if err := touchFerritebarConfig(cfg); err != nil {
		logf(cfg, "  [WARN] ferritebar config: %v\n", err)
	} else {
		logln(cfg, "  [OK] ferritebar config")
	}

	// 11. Restart ferritebar (same effect as Alt-k, then Alt-b in labwc)
	if err := restartFerritebar(cfg); err != nil {
		logf(cfg, "  [WARN] ferritebar restart: %v\n", err)
	} else {
		logln(cfg, "  [OK] ferritebar restart")
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
	content = normalizeTmThemeContrast(content, 4.5)

	dir := filepath.Dir(cfg.SyntectCurrent)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}

	if err := os.WriteFile(cfg.SyntectCurrent, content, 0644); err != nil {
		return fmt.Errorf("write current tmTheme %s: %w", cfg.SyntectCurrent, err)
	}

	return nil
}

func normalizeTmThemeContrast(content []byte, minRatio float64) []byte {
	text := string(content)
	bg := firstPlistColor(text, "background")
	defaultFg := firstPlistColor(text, "foreground")
	if bg == "" || defaultFg == "" {
		return content
	}

	foregroundRe := regexp.MustCompile(`(?s)(<key>foreground</key>\s*<string>)(#[0-9A-Fa-f]{6})(</string>)`)
	normalized := foregroundRe.ReplaceAllStringFunc(text, func(match string) string {
		parts := foregroundRe.FindStringSubmatch(match)
		if len(parts) != 4 {
			return match
		}
		if contrastRatio(parts[2], bg) >= minRatio {
			return match
		}
		return parts[1] + defaultFg + parts[3]
	})

	return []byte(normalized)
}

func firstPlistColor(content, key string) string {
	re := regexp.MustCompile(`(?s)<key>` + regexp.QuoteMeta(key) + `</key>\s*<string>(#[0-9A-Fa-f]{6})</string>`)
	match := re.FindStringSubmatch(content)
	if len(match) != 2 {
		return ""
	}
	return match[1]
}

func contrastRatio(fg, bg string) float64 {
	l1 := relativeLuminance(fg)
	l2 := relativeLuminance(bg)
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	return (l1 + 0.05) / (l2 + 0.05)
}

func relativeLuminance(hex string) float64 {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0
	}
	r := linearRGB(hex[0:2])
	g := linearRGB(hex[2:4])
	b := linearRGB(hex[4:6])
	return 0.2126*r + 0.7152*g + 0.0722*b
}

func linearRGB(hex string) float64 {
	v, err := strconv.ParseUint(hex, 16, 8)
	if err != nil {
		return 0
	}
	c := float64(v) / 255
	if c <= 0.03928 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}

func applyGtkSourceview5Current(cfg *Config, s *scheme.Base16) error {
	content, err := gtksourceview.RenderV5Current(s)
	if err != nil {
		return err
	}
	return writeFile(cfg, cfg.GtkSourceview5Current, content)
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

// applyFirefoxAccent writes a tiny JSON payload that the base16-accent
// WebExtension reads (via a native-messaging host) when a private window is
// opened. Only the accent and variant are exposed; FF's forced private chrome
// surfaces are dark, so the accent is the only colour we need to inject.
//
// Contains both the regular accent (used in normal-window targets) and a
// "private" variant: if the chosen accent fails AA on FF's forced private
// field bg #42414d, fall back to base07 (light foreground) which clears AA
// against both #42414d and #1c1b22.
func applyFirefoxAccent(cfg *Config, s *scheme.Base16) error {
	m := s.ToMap()
	accent := "#" + m["accent-hex"]
	private := accent
	const ffPrivField = "#42414d"
	const ffPrivChrome = "#1c1b22"
	if scheme.MinRatio(accent, ffPrivField, ffPrivChrome) < scheme.AAThreshold {
		// Pick the slot with the best worst-case ratio against FF's two
		// forced private surfaces. base05/06/07 are nominally "fg" slots
		// but some themes (e.g. Catppuccin) repurpose base07 as a chromatic
		// accent — so we can't blindly assume base07 is light. Scanning
		// gives the right answer for any palette.
		candidates := []string{
			"#" + m["base07-hex"], "#" + m["base06-hex"], "#" + m["base05-hex"],
			"#" + m["base0A-hex"], "#" + m["base0B-hex"], "#" + m["base0C-hex"],
			"#" + m["base0D-hex"], "#" + m["base0E-hex"],
		}
		best := private
		bestMin := scheme.MinRatio(private, ffPrivField, ffPrivChrome)
		for _, c := range candidates {
			r := scheme.MinRatio(c, ffPrivField, ffPrivChrome)
			if r > bestMin {
				bestMin = r
				best = c
			}
		}
		private = best
	}
	variant := s.Variant
	if variant == "" {
		// Older schemes may omit the field; infer from base00 luminance.
		if scheme.ContrastRatio("#000000", "#"+s.Palette.Base00) > scheme.ContrastRatio("#ffffff", "#"+s.Palette.Base00) {
			variant = "light"
		} else {
			variant = "dark"
		}
	}
	// Manual JSON marshal — payload is trivial and we want stable key order.
	payload := fmt.Sprintf(
		"{\n  \"accent\": %q,\n  \"private_accent\": %q,\n  \"variant\": %q,\n  \"scheme\": %q\n}\n",
		accent, private, variant, s.Name,
	)
	return writeFile(cfg, cfg.FirefoxAccentJSON, payload)
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
	// LabWC resolves its Openbox-style decorations from the selected theme
	// name. Keep that authority with Orchis, but generate a Base16-safe
	// openbox-3/themerc under the selected Orchis variant.
	themercPath := filepath.Join(cfg.OrchisDestDir, cfg.GtkThemeName, "openbox-3", "themerc")
	if cfg.DryRun {
		logf(cfg, "  Would write openbox theme: %s\n", themercPath)
		return nil
	}

	content, err := template.RenderString(openboxTemplate, s.ToMap())
	if err != nil {
		return err
	}
	return writeFile(cfg, themercPath, content)
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
	newContent := re.ReplaceAllString(string(content), "${1}"+cfg.GtkThemeName+"${2}")

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

func applyWallpaper(cfg *Config, s *scheme.Base16) error {
	wpPath := resolveWallpaperPath(cfg)
	if cfg.DryRun {
		logf(cfg, "  Would read wallpaper source: %s\n", wpPath)
		logf(cfg, "  Would render gowall theme: %s\n", cfg.GowallThemeJSON)
		logf(cfg, "  Would convert wallpaper to: %s\n", cfg.WallpaperCurrent)
		logf(cfg, "  Would run: awww img %s\n", cfg.WallpaperCurrent)
		return nil
	}

	if _, err := os.Stat(wpPath); err != nil {
		return fmt.Errorf("wallpaper source missing at %s: %w", wpPath, err)
	}

	themeJSON, err := renderGowallTheme(s)
	if err != nil {
		return err
	}
	if err := writeFile(cfg, cfg.GowallThemeJSON, themeJSON); err != nil {
		return fmt.Errorf("gowall theme json: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(cfg.WallpaperCurrent), 0755); err != nil {
		return fmt.Errorf("mkdir wallpaper output: %w", err)
	}
	if err := cleanupGeneratedWallpaper(cfg); err != nil {
		return err
	}
	if err := run("gowall", "convert", wpPath, "--output", cfg.WallpaperCurrent, "--theme", cfg.GowallThemeJSON, "--format", "png", "--yes"); err != nil {
		return fmt.Errorf("gowall convert: %w", err)
	}
	if err := run("awww", "img", cfg.WallpaperCurrent); err != nil {
		return fmt.Errorf("awww: %w", err)
	}

	return nil
}

func resolveWallpaperPath(cfg *Config) string {
	if filepath.IsAbs(cfg.Wallpaper) || strings.ContainsRune(cfg.Wallpaper, os.PathSeparator) {
		return cfg.Wallpaper
	}
	for _, dir := range WallpaperDirs() {
		path := filepath.Join(dir, cfg.Wallpaper)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return filepath.Join(cfg.WallpaperDir, cfg.Wallpaper)
}

func cleanupGeneratedWallpaper(cfg *Config) error {
	if err := os.Remove(cfg.WallpaperCurrent); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove old generated wallpaper: %w", err)
	}
	return nil
}

func renderGowallTheme(s *scheme.Base16) (string, error) {
	return template.RenderString(`{
  "name": "{{scheme-name}}",
  "colors": [
    "#{{base00-hex}}",
    "#{{base01-hex}}",
    "#{{base02-hex}}",
    "#{{base03-hex}}",
    "#{{base04-hex}}",
    "#{{base05-hex}}",
    "#{{base06-hex}}",
    "#{{base07-hex}}",
    "#{{base08-hex}}",
    "#{{base09-hex}}",
    "#{{base0A-hex}}",
    "#{{base0B-hex}}",
    "#{{base0C-hex}}",
    "#{{base0D-hex}}",
    "#{{base0E-hex}}",
    "#{{base0F-hex}}"
  ]
}
`, s.ToMap())
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

func restartFerritebar(cfg *Config) error {
	if cfg.DryRun {
		logln(cfg, "  Would run: pkill ferritebar")
		logln(cfg, "  Would run: ferritebar")
		return nil
	}

	_ = run("pkill", "ferritebar")
	time.Sleep(700 * time.Millisecond)

	cmd := exec.Command("ferritebar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start ferritebar: %w", err)
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
