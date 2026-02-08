package targets

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// SchemesDirs returns directories to scan for base16 scheme YAML files
func SchemesDirs() []string {
	home, _ := os.UserHomeDir()
	return []string{
		filepath.Join(home, ".local/share/themes"), // user custom schemes first
		"/run/current-system/sw/share/themes",      // system-installed base16-schemes
	}
}

// ScanSchemesDir returns available base16 scheme names from a single directory
func ScanSchemesDir(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var schemes []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(name, ".yaml") {
			schemes = append(schemes, strings.TrimSuffix(name, ".yaml"))
		}
	}
	sort.Strings(schemes)
	return schemes, nil
}

// ScanAllSchemes returns schemes from all SchemesDirs, deduped
func ScanAllSchemes() []string {
	set := make(map[string]struct{})
	for _, dir := range SchemesDirs() {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			name := e.Name()
			if strings.HasSuffix(name, ".yaml") {
				set[strings.TrimSuffix(name, ".yaml")] = struct{}{}
			}
		}
	}

	schemes := make([]string, 0, len(set))
	for s := range set {
		schemes = append(schemes, s)
	}
	sort.Strings(schemes)
	return schemes
}

// FindScheme searches SchemesDirs for a scheme and returns its full path
func FindScheme(name string) (string, error) {
	for _, dir := range SchemesDirs() {
		path := filepath.Join(dir, name+".yaml")
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("scheme not found: %s", name)
}


// dirEntryIsDir returns true for real directories AND symlinks that point to directories.
// NixOS commonly exposes themes/icons under /run/current-system/sw as symlink entries.
func dirEntryIsDir(parent string, e os.DirEntry) bool {
	if e.IsDir() {
		return true
	}
	if e.Type()&os.ModeSymlink == 0 {
		return false
	}
	fi, err := os.Stat(filepath.Join(parent, e.Name())) // follows symlink
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func exists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// IconDirs returns directories to scan for icon themes
func IconDirs() []string {
	home, _ := os.UserHomeDir()
	return []string{
		"/usr/share/icons",
		filepath.Join(home, ".local/share/icons"),
		"/run/current-system/sw/share/icons",
		filepath.Join(home, ".nix-profile/share/icons"),
	}
}

// WallpaperDir returns the wallpaper directory
func WallpaperDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Pictures/walls")
}

// ScanIconThemes returns available icon themes
func ScanIconThemes() []string {
	set := map[string]struct{}{}
	for _, dir := range IconDirs() {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !dirEntryIsDir(dir, e) {
				continue
			}
			name := e.Name()
			if strings.HasPrefix(name, ".") {
				continue
			}
			index := filepath.Join(dir, name, "index.theme")
			if exists(index) {
				set[name] = struct{}{}
			}
		}
	}
	out := make([]string, 0, len(set))
	for k := range set {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// ScanWallpapers returns available wallpapers
func ScanWallpapers() []string {
	dir := WallpaperDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}
	out := []string{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		ext := strings.ToLower(filepath.Ext(name))
		switch ext {
		case ".jpg", ".jpeg", ".png", ".webp":
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}
