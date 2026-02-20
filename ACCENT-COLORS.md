# Accent Color Discovery: Non-Blue Primaries Hidden in Base16/Gogh Themes

Most Base16 themes default to blue (`base0D`) as the primary accent. But many theme authors chose more interesting colors — buried in cursor values, ANSI slots, or non-standard base16 mappings. This document captures what was found and how to use it.

---

## The Good Ones

### Pinks & Magentas

| Theme | Accent | Hex | Where It Hides | base0D (blue) |
|-------|--------|-----|----------------|---------------|
| **Hemisu Dark** | hot pink | `#ff0054` | color_02 (red slot) | `#67bee3` |
| **Hemisu Light** | hot pink | `#ff0054` | cursor | `#538091` |
| **Jup** | magenta | `#dd006f` | color_02 | `#006fdd` |
| **Rosé Pine Moon** | rose pink | `#eb6f92` | color_02 | `#3e8fb0` |
| **Papercolor Dark** | deep magenta | `#af005f` | color_02 | `#5fafd7` |
| **Pencil Dark** | vivid pink | `#c30771` | color_02 | `#008ec4` |
| **Pencil Light** | vivid pink | `#c30771` | color_02 | `#008ec4` |
| **One Light** | purple-pink | `#930092` | color_06 (magenta) | `#315eee` |
| **Catppuccin Latte** | rosewater | `#dc8a78` | base06 (repurposed) | `#1e66f5` |

### Oranges & Ambers

| Theme | Accent | Hex | Where It Hides | base0D (blue) |
|-------|--------|-----|----------------|---------------|
| **Ayu Light** | warm orange | `#ff9940` | cursor | `#399ee6` |
| **Ayu Mirage** | golden orange | `#ffcc66` | cursor | `#73d0ff` |
| **Rosé Pine Dawn** | marigold | `#ea9d34` | color_04 (yellow) | `#286983` |
| **Solarized Light** | amber | `#cb4b16` | active_border_color | `#268bd2` |
| **Solarized Dark HC** | amber | `#a57706` | color_04 | `#2176c7` |
| **Monokai Pro Light** | warm orange | `#cc7a0a` | color_04 | `#e16032` |
| **Monokai Pro Light Sun** | burnt orange | `#b16803` | color_04 | `#d4572b` |
| **Monokai Pro Ristretto** | golden | `#edce73` | color_04 | `#dc9373` |
| **Seoul256** | amber | `#d8af5f` | color_04 | `#85add4` |
| **Tokyo Night Storm** | honey | `#e0af68` | color_04 | `#7aa2f7` |
| **Everforest Light Hard** | deep gold | `#dfa000` | color_04 | `#3a94c5` |
| **Gruvbox Material Light** | ochre | `#b47109` | color_04 | `#45707a` |
| **Gruvbox Material Dark** | warm gold | `#d8a657` | color_04 | `#7daea3` |

### Reds & Corals

| Theme | Accent | Hex | Where It Hides | base0D (blue) |
|-------|--------|-----|----------------|---------------|
| **Kanagawa Dragon** | muted coral | `#c4746e` | color_02 | `#8ba4b0` |
| **Kanagawa Wave** | crimson | `#c34043` | color_02 | `#7e9cd8` |
| **Kanagawa Lotus** | rose red | `#c84053` | color_02 | `#4d699b` |
| **Everforest Dark Hard** | soft red | `#e67e80` | color_02 | `#7fbbb3` |
| **Gotham** | signal red | `#c33027` | color_02 | `#195465` |
| **Modus Operandi** | deep red | `#a60000` | color_02 | `#0031a9` |
| **Oceanic Next** | coral red | `#e44754` | color_02 | `#5486c0` |

### Greens & Teals

| Theme | Accent | Hex | Where It Hides | base0D (blue) |
|-------|--------|-----|----------------|---------------|
| **Arc Light** | neon green | `#00ff00` | cursor | `#2455c3` |
| **Selenized Light** | leaf green | `#489100` | color_03 | `#0072d4` |
| **Tomorrow** | olive green | `#718c00` | color_03 | `#4271ae` |
| **Sea Shells** | deep teal | `#027c9b` | color_03 | `#1e4950` |
| **3024 Day** | emerald | `#01a252` | color_03 | `#01a0e4` |
| **Tomorrow Night Blue** | bright aqua | `#99ffff` | color_07 (cyan) | `#bbdaff` |

### Yellows

| Theme | Accent | Hex | Where It Hides | base0D (blue) |
|-------|--------|-----|----------------|---------------|
| **Flexoki Dark** | mustard | `#ad8301` | color_04 | `#205ea6` |
| **Flexoki Light** | mustard | `#ad8301` | color_04 | `#205ea6` |
| **Modus Vivendi** | bright yellow | `#d0bc00` | color_04 | `#2fafff` |
| **Selenized Dark** | gold | `#dbb32d` | color_04 | `#4695f7` |
| **Paper** | solarized yellow | `#b58900` | color_04 | `#1e6fcc` |

### Already-Interesting base0D (not blue by default)

These themes already have non-blue values in their base0D slot — no changes needed for base16changer:

| Theme | base0D color | Hue |
|-------|-------------|-----|
| **Gruvbox dark** | `#458588` | teal |
| **Gruvbox dark, hard** | `#83a598` | sage teal |
| **Gruvbox dark, pale** | `#83adad` | muted teal |
| **Gruvbox light** | `#458588` | teal |
| **Gruvbox light, hard** | `#076678` | deep teal |
| **Gruvbox Material Dark, Hard** | `#7daea3` | seafoam |
| **Gruvbox Material Light, Hard** | `#45707a` | slate teal |
| **Tokyo Night Dark** | `#2ac3de` | bright cyan |

---

## How the Detection Works

Scan script at `/tmp/scan-accents.ts` (run with `bun run /tmp/scan-accents.ts`).

**For Gogh themes** (`.yml` with `color_01`-`color_16` + `cursor`/`background`/`foreground`):
1. Check `cursor` against distinctiveness filter (distance from bg > 60, distance from fg > 60, saturation > 35%)
2. If cursor fails, check all ANSI color slots, pick most saturated that passes
3. If nothing passes, fall back to color_05 (blue)

**For Base16 themes** (`.yaml` with `palette.base00`-`base0F`):
1. Check base06/base07 (sometimes repurposed — Catppuccin puts rosewater and lavender here)
2. Check base08-base0F for most saturated non-blue
3. Report base0D and best alternative

---

## Implementation: Adding `accent` to base16changer

### 1. Add optional `accent` field to YAML parsing

**File**: `internal/scheme/parse.go`

In the `Scheme` or palette struct, add:

```go
Accent string `yaml:"accent"` // optional override, hex without #
```

In `ParseFile()`, after loading the palette, check for accent:

```go
if scheme.Accent != "" {
    data["accent-hex"] = scheme.Accent
    // also generate accent-rgb-r, accent-rgb-g, accent-rgb-b
} else {
    // fall back to base0D
    data["accent-hex"] = data["base0D-hex"]
}
```

### 2. Update templates to use `accent-hex`

**File**: `internal/targets/templates.go`

Replace `base0D-hex` with `accent-hex` in these specific accent/primary roles only:

**GTK-4** (lines ~80-82):
```css
@define-color accent_color #{{accent-hex}};
@define-color accent_bg_color #{{accent-hex}};
@define-color accent_fg_color #{{base00-hex}};
```

**LabWC/Openbox** (active window border):
```
window.active.border.color: #{{accent-hex}}
```

**Fuzzel** (match highlight and border):
```
match={{accent-hex}}ff
selection-match={{accent-hex}}ff
border={{accent-hex}}ff
```

**Leave these unchanged** (they should stay as literal ANSI blue):
- Kitty `color4` / `color12` → keep `base0D-hex`
- GTK-3/GTK-2 `link_color` → keep `base0D-hex` (or change to accent, your call)
- GTK-4 `blue_1` through `blue_5` → keep `base0D-hex`

### 3. Add `accent` field to the 4-5 theme YAMLs you want to test

Example for Catppuccin Latte (`catppuccin-latte.yaml`):

```yaml
system: "base16"
name: "Catppuccin Latte"
author: "https://github.com/catppuccin/catppuccin"
variant: "light"
accent: "dc8a78"  # rosewater — from base06, cursor in kitty conf
palette:
  base00: "eff1f5"
  # ... rest unchanged
```

The field is optional. Themes without it use base0D as before. Zero breakage.

### 4. For Gogh `.yml` files

Same approach — add `accent:` field. The Gogh parser in `parse.go` already handles extra YAML fields; it just needs to forward the accent value into the template data map the same way.

### 5. For kitty `.conf` files (DayLight/Skeleton pipeline)

No changes needed — the Skeleton generator already has `findAccentColor()` which does this detection. The accent colors in the tables above are what it finds (or would find with the expanded ANSI check).

---

## Starter Set: 5 Themes to Test First

Pick from different hue families to validate the full range:

1. **Catppuccin Latte** — rosewater `#dc8a78` (pink/coral)
2. **Rosé Pine Dawn** — marigold `#ea9d34` (warm orange) — use from Gogh cursor or color_04
3. **Hemisu Light** — hot pink `#ff0054` (magenta)
4. **Gruvbox Material Light, Hard** — already teal `#45707a` in base0D (no change needed, control group)
5. **Monokai Pro Light** — warm orange `#cc7a0a`

---

*Generated 2026-02-19 from scan of 85 themes across `/home/john/repos/Gogh/themes/` and `/home/john/.local/share/themes/`.*
