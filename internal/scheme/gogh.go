package scheme

import (
	"os"

	"github.com/lucasb-eyer/go-colorful"
	"gopkg.in/yaml.v3"
)

// Gogh represents a Gogh terminal color scheme
type Gogh struct {
	Name    string `yaml:"name"`
	Author  string `yaml:"author"`
	Variant string `yaml:"variant"`

	// ANSI colors
	Color01 string `yaml:"color_01"` // Black
	Color02 string `yaml:"color_02"` // Red
	Color03 string `yaml:"color_03"` // Green
	Color04 string `yaml:"color_04"` // Yellow
	Color05 string `yaml:"color_05"` // Blue
	Color06 string `yaml:"color_06"` // Magenta
	Color07 string `yaml:"color_07"` // Cyan
	Color08 string `yaml:"color_08"` // White
	Color09 string `yaml:"color_09"` // Bright Black
	Color10 string `yaml:"color_10"` // Bright Red
	Color11 string `yaml:"color_11"` // Bright Green
	Color12 string `yaml:"color_12"` // Bright Yellow
	Color13 string `yaml:"color_13"` // Bright Blue
	Color14 string `yaml:"color_14"` // Bright Magenta
	Color15 string `yaml:"color_15"` // Bright Cyan
	Color16 string `yaml:"color_16"` // Bright White

	Background string `yaml:"background"`
	Foreground string `yaml:"foreground"`
	Cursor     string `yaml:"cursor"`
}

// parseGogh reads a Gogh YAML scheme file
func parseGogh(path string) (*Gogh, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var g Gogh
	if err := yaml.Unmarshal(data, &g); err != nil {
		return nil, err
	}

	return &g, nil
}

// ToBase16 converts a Gogh scheme to Base16 format
func (g *Gogh) ToBase16() *Base16 {
	bg := normalizeColor(g.Background)
	fg := normalizeColor(g.Foreground)

	// Direct mappings
	red := normalizeColor(g.Color02)
	yellow := normalizeColor(g.Color04)

	// Derive orange: blend red and yellow
	orange := blendColors(red, yellow, 0.5)

	// Derive brown: darken orange toward background
	brown := blendColors(orange, bg, 0.4)

	return &Base16{
		System:  "base16",
		Name:    g.Name,
		Author:  g.Author,
		Variant: g.Variant,
		Palette: Colors{
			Base00: bg,                                    // Background
			Base01: interpolate(bg, fg, 0.1),              // Lighter bg
			Base02: interpolate(bg, fg, 0.2),              // Selection bg
			Base03: normalizeColor(g.Color09),             // Bright black (comments)
			Base04: interpolate(bg, fg, 0.4),              // Dark fg
			Base05: fg,                                    // Foreground
			Base06: interpolate(bg, fg, 0.8),              // Light fg
			Base07: normalizeColor(g.Color16),             // Bright white
			Base08: red,                                   // Red
			Base09: orange,                                // Orange (derived)
			Base0A: yellow,                                // Yellow
			Base0B: normalizeColor(g.Color03),             // Green
			Base0C: normalizeColor(g.Color07),             // Cyan
			Base0D: normalizeColor(g.Color05),             // Blue
			Base0E: normalizeColor(g.Color06),             // Magenta
			Base0F: brown,                                 // Brown (derived)
		},
	}
}

// interpolate blends two hex colors by factor t (0.0 = c1, 1.0 = c2)
func interpolate(c1, c2 string, t float64) string {
	col1, err1 := colorful.Hex("#" + c1)
	col2, err2 := colorful.Hex("#" + c2)
	if err1 != nil || err2 != nil {
		return c1 // fallback
	}

	blended := col1.BlendLab(col2, t)
	return normalizeColor(blended.Hex())
}

// blendColors blends two hex colors in Lab space
func blendColors(c1, c2 string, t float64) string {
	return interpolate(c1, c2, t)
}
