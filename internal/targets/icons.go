package targets

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jaycee1285/base16changer/internal/scheme"
)

var (
	colorSchemeTextRE      = regexp.MustCompile(`(?i)(\.ColorScheme-Text\s*\{\s*color:\s*)#?[0-9a-f]{6}`)
	colorSchemeHighlightRE = regexp.MustCompile(`(?i)(\.ColorScheme-Highlight\s*\{\s*color:\s*)#?[0-9a-f]{6}`)
)

func recolorIconTheme(cfg *Config, s *scheme.Base16) error {
	themeDir, err := FindIconThemeDir(cfg.IconTheme)
	if err != nil {
		return err
	}

	if cfg.DryRun {
		logf(cfg, "  Would recolor SVG palette in: %s\n", themeDir)
		return nil
	}

	foreground := "#" + s.Palette.Base05
	background := "#" + s.Palette.Base02
	updated := 0

	err = filepath.WalkDir(themeDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() || strings.ToLower(filepath.Ext(path)) != ".svg" {
			return nil
		}

		changed, err := recolorIconSVG(path, foreground, background)
		if err != nil {
			return err
		}
		if changed {
			updated++
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("recolor icon theme: %w", err)
	}

	logf(cfg, "  Recolored %d SVGs in %s\n", updated, themeDir)
	return nil
}

func recolorIconSVG(path, foreground, background string) (bool, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("read %s: %w", path, err)
	}

	original := string(content)
	updated := colorSchemeTextRE.ReplaceAllString(original, "${1}"+foreground)
	updated = colorSchemeHighlightRE.ReplaceAllString(updated, "${1}"+background)

	if updated == original {
		return false, nil
	}

	if err := os.WriteFile(path, []byte(updated), 0644); err != nil {
		return false, fmt.Errorf("write %s: %w", path, err)
	}
	return true, nil
}
