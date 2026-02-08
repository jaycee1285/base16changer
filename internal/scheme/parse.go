package scheme

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Base16 represents a base16 color scheme
type Base16 struct {
	System  string `yaml:"system"`
	Name    string `yaml:"name"`
	Author  string `yaml:"author"`
	Variant string `yaml:"variant"` // "light" or "dark"
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

// Parse reads a base16 YAML scheme file
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

	return &scheme, nil
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

// ToMap returns colors as a map for template rendering
func (s *Base16) ToMap() map[string]string {
	return map[string]string{
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
}

func slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, ",", "")
	return s
}
