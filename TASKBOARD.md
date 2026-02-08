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

## In Progress

None

## Planned

### v1.1: Gogh Theme Converter

- [ ] `gogh2base16` converter tool
  - Parse Gogh YAML format (color_01-16, background, foreground, cursor)
  - Map direct matches:
    - base00 ← background
    - base03 ← color_09 (bright black)
    - base05 ← foreground
    - base07 ← color_16 (bright white)
    - base08 ← color_02 (red)
    - base0A ← color_04 (yellow)
    - base0B ← color_03 (green)
    - base0C ← color_07 (cyan)
    - base0D ← color_05 (blue)
    - base0E ← color_06 (magenta)
  - Derive missing shades:
    - base01 ← interpolate(bg, fg, 0.1)
    - base02 ← interpolate(bg, fg, 0.2)
    - base04 ← interpolate(bg, fg, 0.4)
    - base06 ← interpolate(bg, fg, 0.8)
  - Derive missing accents:
    - base09 (orange) ← blend(red, yellow) or use yellow
    - base0F (brown) ← darken(orange) or use red
  - Output base16 YAML to ~/.local/share/themes/
  - Batch convert all Gogh themes

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
