# Changelog

## 2026-02-08 - Initial Development

### Created base16changer

A hot-reload base16 theme switcher for LabWC/Wayland desktops, inspired by Stylix but with runtime theme switching instead of build-time.

### Features Implemented

**Core**
- Base16 YAML scheme parser with full palette support
- Mustache-style template engine for config generation
- Multi-directory scheme search (~/.local/share/themes + system paths)

**Targets**
- **Kitty**: Writes current-theme.conf, reloads via SIGUSR1
- **Fuzzel**: Updates only [colors] section, preserves other config
- **GTK-3**: Generates gtk.css with theme color vars
- **GTK-4**: Comprehensive libadwaita color definitions (89 vars from Stylix)
- **LabWC/Openbox**: Generates themerc with window decorations, buttons, menus
- **LabWC rc.xml**: Updates theme name reference

**Extras**
- Icon theme switching via dconf
- Wallpaper switching via swww

**Hot Reload**
- Kitty: `pkill -SIGUSR1 kitty`
- LabWC: `labwc -r`
- GTK: dconf theme toggle trick (dummy → actual theme)

**Interface**
- CLI with flags: `base16changer [--icon X] [--wallpaper Y] <scheme>`
- TUI (Bubble Tea): Expandable panels for Schemes, Icons, Wallpapers
- Filtering support in TUI lists

### Architecture

```
cmd/base16changer/main.go    # CLI + TUI entry point
internal/
  scheme/parse.go            # YAML parser
  template/render.go         # Mustache renderer
  targets/
    targets.go               # Apply logic + reload triggers
    templates.go             # Embedded templates (kitty, fuzzel, gtk, openbox)
    scan.go                  # Directory scanning for schemes/icons/wallpapers
  ui/model.go                # Bubble Tea TUI
```

### Dependencies

- gopkg.in/yaml.v3
- github.com/charmbracelet/bubbletea
- github.com/charmbracelet/bubbles
- github.com/charmbracelet/lipgloss

### Usage

```bash
# TUI mode (default)
./base16changer

# CLI mode
./base16changer gruvbox-dark-medium
./base16changer --icon Tela --wallpaper sunset.jpg catppuccin-mocha

# List available
./base16changer --list
./base16changer --list-icons
./base16changer --list-wallpapers

# Testing with local schemes repo
./base16changer --schemes-dir ~/repos/schemes/base16 --dry-run nord
```

### Scheme Locations

1. `~/.local/share/themes/*.yaml` - User curated schemes (checked first)
2. `/run/current-system/sw/share/themes/*.yaml` - NixOS base16-schemes package

### Reference Repos Used

- https://github.com/tinted-theming/schemes - Base16 scheme definitions
- https://github.com/tinted-theming/tinted-terminal - Kitty template reference
- https://github.com/tinted-theming/base16-gtk-flatcolor - GTK approach
- https://github.com/danth/stylix - GTK-4/libadwaita template, architecture inspiration
