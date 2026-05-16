package scheme

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// Base16 represents a base16 color scheme
type Base16 struct {
	System  string `yaml:"system"`
	Name    string `yaml:"name"`
	Author  string `yaml:"author"`
	Variant string `yaml:"variant"` // "light" or "dark"
	Accent  string `yaml:"accent"`  // optional override hex (without #), falls back to base0D
	Palette Colors `yaml:"palette"`
}

// Colors holds the 16 base colors
type Colors struct {
	Base00 string `yaml:"base00"` // Default Background
	Base01 string `yaml:"base01"` // Lighter Background (status bars)
	Base02 string `yaml:"base02"` // Selection Background
	Base03 string `yaml:"base03"` // Comments, Invisibles
	Base04 string `yaml:"base04"` // Dark Foreground (status bars)
	Base05 string `yaml:"base05"` // Default Foreground
	Base06 string `yaml:"base06"` // Light Foreground
	Base07 string `yaml:"base07"` // Lightest Foreground
	Base08 string `yaml:"base08"` // Red
	Base09 string `yaml:"base09"` // Orange
	Base0A string `yaml:"base0A"` // Yellow
	Base0B string `yaml:"base0B"` // Green
	Base0C string `yaml:"base0C"` // Cyan
	Base0D string `yaml:"base0D"` // Blue
	Base0E string `yaml:"base0E"` // Purple
	Base0F string `yaml:"base0F"` // Brown
}

// Parse reads a base16 or Gogh YAML scheme file (auto-detects format)
func Parse(path string) (*Base16, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read scheme: %w", err)
	}

	var scheme Base16
	if err := yaml.Unmarshal(data, &scheme); err != nil {
		return nil, fmt.Errorf("parse scheme: %w", err)
	}

	// Normalize colors (remove # prefix if present, lowercase)
	scheme.Palette.normalize()
	scheme.Accent = normalizeColor(scheme.Accent)

	// Auto-detect Gogh format: missing base09 (orange) or base0F (brown).
	// Backup heuristic: Gogh uses .yml, Base16 uses .yaml (not enforced here).
	if scheme.Palette.Base09 == "" || scheme.Palette.Base0F == "" {
		gogh, err := parseGogh(path)
		if err == nil && gogh.Color01 != "" {
			b16 := gogh.ToBase16()
			b16.Accent = scheme.Accent
			b16.autoPromoteAccent()
			return b16, nil
		}
	}

	scheme.autoPromoteAccent()
	return &scheme, nil
}

// autoPromoteAccent runs after Parse() and replaces the scheme's accent if it
// fails WCAG AA against base00 or base01 — see internal/scheme/contrast.go.
// If accent is empty, default to base0D first so the AA check has something
// to evaluate (this matches the existing ToMap() fallback). Stderr-warns on
// any promotion so the user sees it during a render.
func (s *Base16) autoPromoteAccent() {
	if s == nil {
		return
	}
	original := s.Accent
	if original == "" {
		original = s.Palette.Base0D
	}
	promoted, _, replaced := promoteAccent(original, s.Palette)
	if replaced {
		warnAccentPromoted(s.Name, original, promoted)
		s.Accent = promoted
	}
}

func (c *Colors) normalize() {
	c.Base00 = normalizeColor(c.Base00)
	c.Base01 = normalizeColor(c.Base01)
	c.Base02 = normalizeColor(c.Base02)
	c.Base03 = normalizeColor(c.Base03)
	c.Base04 = normalizeColor(c.Base04)
	c.Base05 = normalizeColor(c.Base05)
	c.Base06 = normalizeColor(c.Base06)
	c.Base07 = normalizeColor(c.Base07)
	c.Base08 = normalizeColor(c.Base08)
	c.Base09 = normalizeColor(c.Base09)
	c.Base0A = normalizeColor(c.Base0A)
	c.Base0B = normalizeColor(c.Base0B)
	c.Base0C = normalizeColor(c.Base0C)
	c.Base0D = normalizeColor(c.Base0D)
	c.Base0E = normalizeColor(c.Base0E)
	c.Base0F = normalizeColor(c.Base0F)
}

func normalizeColor(c string) string {
	c = strings.TrimPrefix(c, "#")
	c = strings.ToLower(c)
	return c
}

// Hex returns the color with # prefix
func (c *Colors) Hex(name string) string {
	switch name {
	case "base00":
		return "#" + c.Base00
	case "base01":
		return "#" + c.Base01
	case "base02":
		return "#" + c.Base02
	case "base03":
		return "#" + c.Base03
	case "base04":
		return "#" + c.Base04
	case "base05":
		return "#" + c.Base05
	case "base06":
		return "#" + c.Base06
	case "base07":
		return "#" + c.Base07
	case "base08":
		return "#" + c.Base08
	case "base09":
		return "#" + c.Base09
	case "base0A", "base0a":
		return "#" + c.Base0A
	case "base0B", "base0b":
		return "#" + c.Base0B
	case "base0C", "base0c":
		return "#" + c.Base0C
	case "base0D", "base0d":
		return "#" + c.Base0D
	case "base0E", "base0e":
		return "#" + c.Base0E
	case "base0F", "base0f":
		return "#" + c.Base0F
	default:
		return ""
	}
}

// hexToDec converts a 6-char hex color string to decimal R, G, B strings
func hexToDec(hex string) (string, string, string) {
	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)
	return strconv.FormatUint(r, 10), strconv.FormatUint(g, 10), strconv.FormatUint(b, 10)
}

// ToMap returns colors as a map for template rendering
func (s *Base16) ToMap() map[string]string {
	m := map[string]string{
		"scheme-name":   s.Name,
		"scheme-author": s.Author,
		"scheme-slug":   slugify(s.Name),
		"base00-hex":    s.Palette.Base00,
		"base01-hex":    s.Palette.Base01,
		"base02-hex":    s.Palette.Base02,
		"base03-hex":    s.Palette.Base03,
		"base04-hex":    s.Palette.Base04,
		"base05-hex":    s.Palette.Base05,
		"base06-hex":    s.Palette.Base06,
		"base07-hex":    s.Palette.Base07,
		"base08-hex":    s.Palette.Base08,
		"base09-hex":    s.Palette.Base09,
		"base0A-hex":    s.Palette.Base0A,
		"base0B-hex":    s.Palette.Base0B,
		"base0C-hex":    s.Palette.Base0C,
		"base0D-hex":    s.Palette.Base0D,
		"base0E-hex":    s.Palette.Base0E,
		"base0F-hex":    s.Palette.Base0F,
	}

	// Add decimal R, G, B values for each base color
	bases := []struct {
		name string
		hex  string
	}{
		{"base00", s.Palette.Base00}, {"base01", s.Palette.Base01},
		{"base02", s.Palette.Base02}, {"base03", s.Palette.Base03},
		{"base04", s.Palette.Base04}, {"base05", s.Palette.Base05},
		{"base06", s.Palette.Base06}, {"base07", s.Palette.Base07},
		{"base08", s.Palette.Base08}, {"base09", s.Palette.Base09},
		{"base0A", s.Palette.Base0A}, {"base0B", s.Palette.Base0B},
		{"base0C", s.Palette.Base0C}, {"base0D", s.Palette.Base0D},
		{"base0E", s.Palette.Base0E}, {"base0F", s.Palette.Base0F},
	}
	for _, b := range bases {
		r, g, bl := hexToDec(b.hex)
		m[b.name+"-dec-r"] = r
		m[b.name+"-dec-g"] = g
		m[b.name+"-dec-b"] = bl
	}

	// Accent color: use explicit accent if set, otherwise fall back to base0D
	accent := s.Palette.Base0D
	if s.Accent != "" {
		accent = s.Accent
	}
	m["accent-hex"] = accent
	ar, ag, ab := hexToDec(accent)
	m["accent-dec-r"] = ar
	m["accent-dec-g"] = ag
	m["accent-dec-b"] = ab

	neutralRamp := s.neutralRamp()
	for key, value := range neutralRamp {
		m[key] = value
	}

	// Fuzzel selection background: for a light scheme, base02 (darker
	// neutral) keeps dark text readable. For a dark scheme, base02 is too
	// close to the background, so step out to base04 — unless base04 vs
	// base05 (text) fails AA, in which case fall back to base03.
	switch {
	case s.isLightVariant():
		m["fuzzel-selection-hex"] = s.Palette.Base02
	case ContrastRatio(s.Palette.Base05, s.Palette.Base04) < AAThreshold:
		m["fuzzel-selection-hex"] = s.Palette.Base03
	default:
		m["fuzzel-selection-hex"] = s.Palette.Base04
	}

	// AA-normalized variants of chromatic slots, for use as syntax-color
	// foregrounds in gtk-sourceview5. If a slot fails WCAG AA against either
	// the view background (base00) or the current-line background (base01),
	// substitute the default foreground (base05). Mirrors the contrast
	// normalization already performed by the tmTheme target.
	chromatic := []struct{ name, hex string }{
		{"base09", s.Palette.Base09},
		{"base0A", s.Palette.Base0A},
		{"base0B", s.Palette.Base0B},
		{"base0C", s.Palette.Base0C},
		{"base0D", s.Palette.Base0D},
		{"base0E", s.Palette.Base0E},
	}
	for _, c := range chromatic {
		aa := c.hex
		if MinRatio(c.hex, s.Palette.Base00, s.Palette.Base01) < AAThreshold {
			aa = s.Palette.Base05
		}
		m[c.name+"-aa-hex"] = aa
	}

	return m
}

func (s *Base16) neutralRamp() map[string]string {
	light := s.isLightVariant()

	var ramp []string
	if light {
		ramp = []string{
			s.Palette.Base00,
			s.Palette.Base01,
			s.Palette.Base01,
			s.Palette.Base02,
			s.Palette.Base02,
			s.Palette.Base03,
			s.Palette.Base03,
			s.Palette.Base04,
			s.Palette.Base04,
			s.Palette.Base05,
			s.Palette.Base05,
			s.Palette.Base06,
			s.Palette.Base06,
			s.Palette.Base07,
			s.Palette.Base07,
			s.Palette.Base07,
			s.Palette.Base07,
			s.Palette.Base07,
		}
	} else {
		ramp = []string{
			s.Palette.Base07,
			s.Palette.Base06,
			s.Palette.Base06,
			s.Palette.Base05,
			s.Palette.Base05,
			s.Palette.Base04,
			s.Palette.Base04,
			s.Palette.Base03,
			s.Palette.Base03,
			s.Palette.Base02,
			s.Palette.Base02,
			s.Palette.Base01,
			s.Palette.Base01,
			s.Palette.Base00,
			s.Palette.Base00,
			s.Palette.Base00,
			s.Palette.Base00,
			s.Palette.Base00,
		}
	}

	keys := []string{
		"grey-050-hex",
		"grey-100-hex",
		"grey-150-hex",
		"grey-200-hex",
		"grey-250-hex",
		"grey-300-hex",
		"grey-350-hex",
		"grey-400-hex",
		"grey-450-hex",
		"grey-500-hex",
		"grey-550-hex",
		"grey-600-hex",
		"grey-650-hex",
		"grey-700-hex",
		"grey-750-hex",
		"grey-800-hex",
		"grey-850-hex",
		"grey-900-hex",
	}

	out := make(map[string]string, len(keys)+1)
	for i, key := range keys {
		out[key] = ramp[i]
	}
	out["grey-950-hex"] = ramp[len(ramp)-1]
	return out
}

func (s *Base16) isLightVariant() bool {
	switch strings.ToLower(strings.TrimSpace(s.Variant)) {
	case "light":
		return true
	case "dark":
		return false
	}

	bgR, bgG, bgB := hexToDec(s.Palette.Base00)
	fgR, fgG, fgB := hexToDec(s.Palette.Base05)

	bg := atoi(bgR) + atoi(bgG) + atoi(bgB)
	fg := atoi(fgR) + atoi(fgG) + atoi(fgB)
	return bg > fg
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, ",", "")
	return s
}
