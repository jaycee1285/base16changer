# base16changer Taskboard

## Completed

- [x] Core base16 scheme parser (YAML → colors)
- [x] Template engine for mustache-style rendering
- [x] Kitty theme generation + SIGUSR1 reload
- [x] Fuzzel colors (section-only update, preserves [main])
- [x] GTK-3 CSS generation
- [x] GTK-4 CSS generation (libadwaita vars)
- [x] LabWC/Openbox themerc generation
- [x] LabWC rc.xml theme name update
- [x] Icon theme support (--icon flag, dconf)
- [x] Wallpaper support (--wallpaper flag, swww)
- [x] Multi-directory scheme search (~/.local/share/themes + /run/current-system/sw/share/themes)
- [x] CLI mode with flags
- [x] TUI mode (Bubble Tea) with expandable panels
- [x] Hot reload for all targets (kitty, labwc, gtk via dconf toggle)
- [x] Gogh theme auto-conversion
  - Auto-detects Gogh format (missing base09/base0F)
  - Maps ANSI colors to base16 slots
  - Derives orange (base09) and brown (base0F) via Lab-space blending
  - Supports both .yaml and .yml extensions
  - Drop Gogh themes in ~/.local/share/themes/ alongside base16 schemes

## In Progress

None

## Planned

### v2: Waybar Integration

- [ ] Generate waybar colors.css with base16 vars
- [ ] @import from existing style.css
- [ ] SIGUSR2 reload or restart

### v2: Kvantum Support

- [ ] Research Kvantum theme format
- [ ] Generate Kvantum color scheme
- [ ] Document Qt app restart requirement

### Future Ideas

- [ ] Live preview in TUI before applying
- [ ] Theme favorites/history
- [ ] Auto-detect dark/light based on time
- [ ] Nix flake with overlay for easy installation
