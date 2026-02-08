package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jaycee1285/base16changer/internal/scheme"
	"github.com/jaycee1285/base16changer/internal/targets"
	"github.com/jaycee1285/base16changer/internal/ui"
)

func main() {
	var (
		schemeName     string
		schemePath     string
		schemesDir     string
		iconTheme      string
		wallpaper      string
		listFlag       bool
		listIcons      bool
		listWallpapers bool
		dryRun         bool
	)

	flag.StringVar(&schemeName, "scheme", "", "Name of the scheme (e.g., gruvbox-dark-medium)")
	flag.StringVar(&schemePath, "path", "", "Direct path to scheme YAML file")
	flag.StringVar(&schemesDir, "schemes-dir", "", "Directory containing scheme YAML files")
	flag.StringVar(&iconTheme, "icon", "", "Icon theme to apply")
	flag.StringVar(&wallpaper, "wallpaper", "", "Wallpaper filename to apply (from ~/Pictures/walls)")
	flag.BoolVar(&listFlag, "list", false, "List available schemes")
	flag.BoolVar(&listIcons, "list-icons", false, "List available icon themes")
	flag.BoolVar(&listWallpapers, "list-wallpapers", false, "List available wallpapers")
	flag.BoolVar(&dryRun, "dry-run", false, "Show what would be done without making changes")
	flag.Parse()

	// Also accept scheme name as positional arg
	if schemeName == "" && len(flag.Args()) > 0 {
		schemeName = flag.Args()[0]
	}

	cfg := targets.DefaultConfig()
	cfg.DryRun = dryRun
	cfg.IconTheme = iconTheme
	cfg.Wallpaper = wallpaper
	if schemesDir != "" {
		cfg.SchemesDir = schemesDir
	}

	// Handle list commands
	if listFlag {
		if schemesDir != "" {
			listSchemesFromDir(schemesDir)
		} else {
			listAllSchemes()
		}
		return
	}
	if listIcons {
		listIconThemes()
		return
	}
	if listWallpapers {
		listWallpaperFiles()
		return
	}

	// If no scheme specified, launch TUI
	if schemeName == "" && schemePath == "" {
		runTUI(cfg)
		return
	}

	// CLI mode
	runCLI(cfg, schemeName, schemePath)
}

func runTUI(cfg *targets.Config) {
	m := ui.New(cfg)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "base16changer error:", err)
		os.Exit(1)
	}
}

func runCLI(cfg *targets.Config, schemeName, schemePath string) {
	// Resolve scheme path
	var schemeFile string
	var err error
	if schemePath != "" {
		schemeFile = schemePath
	} else if cfg.SchemesDir != "" {
		// Try both extensions
		for _, ext := range []string{".yaml", ".yml"} {
			path := filepath.Join(cfg.SchemesDir, schemeName+ext)
			if _, err := os.Stat(path); err == nil {
				schemeFile = path
				break
			}
		}
		if schemeFile == "" {
			schemeFile = filepath.Join(cfg.SchemesDir, schemeName+".yaml") // fallback for error msg
		}
	} else {
		// Search standard locations
		schemeFile, err = targets.FindScheme(schemeName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Fprintf(os.Stderr, "Searched: %v\n", targets.SchemesDirs())
			os.Exit(1)
		}
	}

	// Parse scheme
	s, err := scheme.Parse(schemeFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Apply
	if err := targets.Apply(cfg, s); err != nil {
		fmt.Fprintf(os.Stderr, "Error applying scheme: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nDone!")
}

func listAllSchemes() {
	schemes := targets.ScanAllSchemes()
	fmt.Printf("Available schemes (from %v):\n\n", targets.SchemesDirs())
	printColumns(schemes, 3)
	fmt.Printf("\nTotal: %d schemes\n", len(schemes))
}

func listSchemesFromDir(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading schemes directory: %v\n", err)
		fmt.Fprintf(os.Stderr, "Expected: %s\n", dir)
		os.Exit(1)
	}

	fmt.Printf("Available schemes in %s:\n\n", dir)

	var schemes []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(name, ".yaml") {
			schemes = append(schemes, strings.TrimSuffix(name, ".yaml"))
		} else if strings.HasSuffix(name, ".yml") {
			schemes = append(schemes, strings.TrimSuffix(name, ".yml"))
		}
	}

	printColumns(schemes, 3)
	fmt.Printf("\nTotal: %d schemes\n", len(schemes))
}

func listIconThemes() {
	icons := targets.ScanIconThemes()
	fmt.Println("Available icon themes:\n")
	printColumns(icons, 3)
	fmt.Printf("\nTotal: %d icon themes\n", len(icons))
}

func listWallpaperFiles() {
	walls := targets.ScanWallpapers()
	fmt.Printf("Available wallpapers in %s:\n\n", targets.WallpaperDir())
	printColumns(walls, 2)
	fmt.Printf("\nTotal: %d wallpapers\n", len(walls))
}

func printColumns(items []string, cols int) {
	for i, s := range items {
		fmt.Printf("%-35s", s)
		if (i+1)%cols == 0 {
			fmt.Println()
		}
	}
	if len(items)%cols != 0 {
		fmt.Println()
	}
}
