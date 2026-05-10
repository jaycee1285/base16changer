# WCAG AA report — Everforest Dark Hard

- **source**: `/home/john/.local/share/themes/everforest-dark-hard.yaml`
- **format**: base16
- **variant**: dark
- **accent**: #7fbbb3 (explicit)
- **threshold**: AA strict (≥ 4.5:1 for all roles)
- **tmTheme**: `/home/john/.local/share/themes/tmThemes/Everforest Dark Hard.tmTheme` (bg=#272e33, fg=#d3c6aa)
- **summary**: 133 pass / 18 fail / 151 total

## Palette

| Slot | Hex |
|------|-----|
| base00 | `#272e33` |
| base01 | `#2e383c` |
| base02 | `#414b50` |
| base03 | `#859289` |
| base04 | `#9da9a0` |
| base05 | `#d3c6aa` |
| base06 | `#edeada` |
| base07 | `#fffbef` |
| base08 | `#e67e80` |
| base09 | `#e69875` |
| base0A | `#dbbc7f` |
| base0B | `#a7c080` |
| base0C | `#83c092` |
| base0D | `#7fbbb3` |
| base0E | `#d699b6` |
| base0F | `#9da9a0` |
| accent | `#7fbbb3` |

## Failures

| Group | Target | Role | fg | bg | Ratio | Suggested swap | Notes |
|-------|--------|------|----|----|-------|----------------|-------|
| intra | kitty | active border on bg | `base03` #859289 | `base00` #272e33 | **4.24:1** | `base06` (#edeada, 11.40:1) |  |
| intra | kitty | bright black (comments) on bg | `base03` #859289 | `base00` #272e33 | **4.24:1** | `base06` (#edeada, 11.40:1) |  |
| intra | fuzzel | selection-match on selection | `accent` #7fbbb3 | `base02` #414b50 | **4.12:1** | `base0A` (#dbbc7f, 4.90:1) |  |
| intra | openbox | inactive button icon on inactive btn bg | `base03` #859289 | `base00` #272e33 | **4.24:1** | `base06` (#edeada, 11.40:1) |  |
| intra | openbox | close button (red) on active title | `base08` #e67e80 | `base01` #2e383c | **4.38:1** | `base0A` (#dbbc7f, 6.59:1) |  |
| intra | gtk3 | border (visual UI) on bg | `base02` #414b50 | `base00` #272e33 | **1.54:1** | — |  |
| intra | gtk4 | scrollbar outline on view_bg | `base02` #414b50 | `base00` #272e33 | **1.54:1** | — |  |
| intra | gsv | bracket-mismatch fg on bg | `base08` #e67e80 | `base01` #2e383c | **4.38:1** | `base0A` (#dbbc7f, 6.59:1) |  |
| intra | firefox | selected suggestion url on hover bg | `accent` #7fbbb3 | `base02` #414b50 | **4.12:1** | `base0A` (#dbbc7f, 4.90:1) |  |
| intra | firefox | danger (.close-icon hover) on toolbar | `base08` #e67e80 | `base01` #2e383c | **4.38:1** | `base0A` (#dbbc7f, 6.59:1) |  |
| stack | code-in-gtk4 | syntect HEADING on GTK-4 selection_bg | `tm:heading` #7fbbb3 | `base02` #414b50 | **4.12:1** | — |  |
| stack | code-in-gtk4 | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #272e33 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #272e33 | `base00` #272e33 | **1.00:1** | — | if tmTheme bg differs from GTK-4 view bg, the gutter strip seam fails |
| stack | fuzzel-on-gtk4 | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #414b50 | `base01` #2e383c | **1.34:1** | — | USER-FLAGGED #3b (post-T3): stripe is now base02 vs base01 |
| stack | fuzzel-on-gtk4 | selection_match (accent) on selected_bg (base02) | `accent` #7fbbb3 | `base02` #414b50 | **4.12:1** | `base0A` (#dbbc7f, 4.90:1) | USER-FLAGGED #3c (post-T3): typed-letter highlight on the new base02 selection |
| stack | fuzzel-on-gtk4 | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #414b50 | `base00` #272e33 | **1.54:1** | — | GTK-4 analog of fuzzel #3b |
| stack | fuzzel-on-gtk4 | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #414b50 | `base01` #2e383c | **1.34:1** | — | GTK-3 analog of fuzzel #3b |
| stack | openbox-on-gtk4 | inactive border (base02) on GTK-4 view_bg | `base02` #414b50 | `base00` #272e33 | **1.54:1** | — |  |

## All pairs

### intra — firefox (15/17 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar text on toolbar bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | field (URL bar) text on field bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | field focus border on field bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | field highlight on field bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | field highlight text on highlight | `base00` #272e33 | `accent` #7fbbb3 | 6.34:1 |
| ✅ | popup text on popup bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | popup border on popup bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ❌ | selected suggestion url on hover bg | `accent` #7fbbb3 | `base02` #414b50 | 4.12:1 |
| ✅ | sidebar text on sidebar bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | sidebar border on sidebar bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | muted text on toolbar bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | tab selected fg on tab selected bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | tab line (accent underline) on tab bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | tab hover fg on tab hover bg | `base05` #d3c6aa | `base02` #414b50 | 5.30:1 |
| ❌ | danger (.close-icon hover) on toolbar | `base08` #e67e80 | `base01` #2e383c | 4.38:1 |
| ✅ | warning (attention) on toolbar | `base0A` #dbbc7f | `base01` #2e383c | 6.59:1 |
| ✅ | success (download badge) on toolbar | `base0B` #a7c080 | `base01` #2e383c | 6.01:1 |

### intra — fuzzel (5/6 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | match on background | `accent` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | selection-text on selection | `base05` #d3c6aa | `base02` #414b50 | 5.30:1 |
| ❌ | selection-match on selection | `accent` #7fbbb3 | `base02` #414b50 | 4.12:1 |
| ✅ | border on bg (visual indicator) | `accent` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | prompt/text on bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |

### intra — gsv (23/24 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on view bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | text on current_line bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | comment on view bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | string on view bg | `base0B` #a7c080 | `base00` #272e33 | 6.88:1 |
| ✅ | string on current_line | `base0B` #a7c080 | `base01` #2e383c | 6.01:1 |
| ✅ | number on view bg | `base09` #e69875 | `base00` #272e33 | 5.98:1 |
| ✅ | function on view bg | `base0D` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | function on current_line | `base0D` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | keyword on view bg | `base0E` #d699b6 | `base00` #272e33 | 5.96:1 |
| ✅ | keyword on current_line | `base0E` #d699b6 | `base01` #2e383c | 5.20:1 |
| ✅ | type on view bg | `base0B` #a7c080 | `base00` #272e33 | 6.88:1 |
| ✅ | builtin on view bg | `base0C` #83c092 | `base00` #272e33 | 6.51:1 |
| ✅ | constant/special on view bg | `base0A` #dbbc7f | `base00` #272e33 | 7.55:1 |
| ✅ | operator on view bg | `base0B` #a7c080 | `base00` #272e33 | 6.88:1 |
| ✅ | preprocessor on view bg | `base0E` #d699b6 | `base00` #272e33 | 5.96:1 |
| ✅ | heading on view bg | `base0D` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | list-marker on view bg | `base09` #e69875 | `base00` #272e33 | 5.98:1 |
| ✅ | link-text on view bg | `base0E` #d699b6 | `base00` #272e33 | 5.96:1 |
| ✅ | link-destination on view bg | `base0C` #83c092 | `base00` #272e33 | 6.51:1 |
| ✅ | search match on match bg | `base05` #d3c6aa | `base02` #414b50 | 5.30:1 |
| ✅ | bracket match on view bg | `base0E` #d699b6 | `base00` #272e33 | 5.96:1 |
| ❌ | bracket-mismatch fg on bg | `base08` #e67e80 | `base01` #2e383c | 4.38:1 |
| ✅ | right_margin text on margin bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | line-numbers on gutter | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |

### intra — gtk3 (10/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | fg on bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | text on base (content area) | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | disabled text on bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | selected text on selected bg | `base05` #d3c6aa | `base02` #414b50 | 5.30:1 |
| ✅ | tooltip text on tooltip bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | link on bg | `base0D` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | link on base (content) | `base0D` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | success on bg | `base0B` #a7c080 | `base00` #272e33 | 6.88:1 |
| ✅ | warning on bg | `base0A` #dbbc7f | `base00` #272e33 | 7.55:1 |
| ✅ | error on bg | `base08` #e67e80 | `base00` #272e33 | 5.02:1 |
| ❌ | border (visual UI) on bg | `base02` #414b50 | `base00` #272e33 | 1.54:1 |

### intra — gtk4 (18/19 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | window_fg on window_bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | view_fg on view_bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | headerbar_fg on headerbar_bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | sidebar_fg on sidebar_bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | card_fg on card_bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | dialog_fg on dialog_bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | popover_fg on popover_bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | accent_color on window_bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | accent_color on view_bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | accent_color on headerbar_bg | `accent` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | accent_color on sidebar_bg | `accent` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | accent_color on card_bg | `accent` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | accent_fg_color on accent_bg_color | `base00` #272e33 | `accent` #7fbbb3 | 6.34:1 |
| ✅ | destructive_fg on destructive_bg | `base00` #272e33 | `base08` #e67e80 | 5.02:1 |
| ✅ | success_fg on success_bg | `base00` #272e33 | `base0B` #a7c080 | 6.88:1 |
| ✅ | warning_fg on warning_bg | `base00` #272e33 | `base0A` #dbbc7f | 7.55:1 |
| ✅ | error_fg on error_bg | `base00` #272e33 | `base08` #e67e80 | 5.02:1 |
| ✅ | blue_3 link on view_bg | `base0D` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ❌ | scrollbar outline on view_bg | `base02` #414b50 | `base00` #272e33 | 1.54:1 |

### intra — kitty (13/15 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | default text on background | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | selection text on selection bg | `base02` #414b50 | `base05` #d3c6aa | 5.30:1 |
| ✅ | cursor on background | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | cursor_text on cursor | `base00` #272e33 | `base05` #d3c6aa | 8.15:1 |
| ✅ | url_color on background | `base04` #9da9a0 | `base00` #272e33 | 5.65:1 |
| ❌ | active border on bg | `base03` #859289 | `base00` #272e33 | 4.24:1 |
| ✅ | active tab fg on active tab bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | inactive tab fg on inactive tab bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | ANSI red on bg (color_01 logs) | `base08` #e67e80 | `base00` #272e33 | 5.02:1 |
| ✅ | ANSI green on bg | `base0B` #a7c080 | `base00` #272e33 | 6.88:1 |
| ✅ | ANSI yellow on bg | `base0A` #dbbc7f | `base00` #272e33 | 7.55:1 |
| ✅ | ANSI blue on bg | `base0D` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | ANSI magenta on bg | `base0E` #d699b6 | `base00` #272e33 | 5.96:1 |
| ✅ | ANSI cyan on bg | `base0C` #83c092 | `base00` #272e33 | 6.51:1 |
| ❌ | bright black (comments) on bg | `base03` #859289 | `base00` #272e33 | 4.24:1 |

### intra — mako (2/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | border on bg | `accent` #7fbbb3 | `base01` #2e383c | 5.53:1 |

### intra — openbox (9/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active title text on active title bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | inactive title text on inactive title bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | active button icon on button bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | active button icon hover | `base07` #fffbef | `base02` #414b50 | 8.65:1 |
| ❌ | inactive button icon on inactive btn bg | `base03` #859289 | `base00` #272e33 | 4.24:1 |
| ✅ | active border (outline) on active title | `accent` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | menu items on menu bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | menu disabled on menu bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | menu active item on highlight | `base05` #d3c6aa | `base02` #414b50 | 5.30:1 |
| ✅ | OSD label on osd bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ❌ | close button (red) on active title | `base08` #e67e80 | `base01` #2e383c | 4.38:1 |

### intra — sidebery (7/7 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active tab fg on active tab bg | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | inactive tab fg on frame bg | `base04` #9da9a0 | `base00` #272e33 | 5.65:1 |
| ✅ | toolbar fg on toolbar bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | tab hover fg on hover bg | `base05` #d3c6aa | `base02` #414b50 | 5.30:1 |
| ✅ | popup fg on popup bg | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | scroll progress (accent) on frame bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | active-tab outline (accent) on tab | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |

### stack — code-in-gtk4 (15/18 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | syntect HEADING on GTK-4 view_bg | `tm:heading` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | syntect HEADING on GTK-4 popover_bg | `tm:heading` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | syntect HEADING on GTK-4 sidebar_bg | `tm:heading` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ❌ | syntect HEADING on GTK-4 selection_bg | `tm:heading` #7fbbb3 | `base02` #414b50 | 4.12:1 |
| ✅ | syntect comment on GTK-4 view_bg | `tm:comment` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | syntect comment on GTK-4 popover_bg | `tm:comment` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | syntect string on GTK-4 view_bg | `tm:string` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | syntect string on GTK-4 popover_bg | `tm:string` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | syntect keyword on GTK-4 view_bg | `tm:keyword` #d699b6 | `base00` #272e33 | 5.96:1 |
| ✅ | syntect keyword on GTK-4 popover_bg | `tm:keyword` #d699b6 | `base01` #2e383c | 5.20:1 |
| ✅ | syntect function on GTK-4 view_bg | `tm:function` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | syntect function on GTK-4 popover_bg | `tm:function` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | syntect type on GTK-4 view_bg | `tm:type` #dbbc7f | `base00` #272e33 | 7.55:1 |
| ✅ | syntect number on GTK-4 view_bg | `tm:number` #e29e80 | `base00` #272e33 | 6.19:1 |
| ❌ | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #272e33 | NaN:1 |
| ✅ | syntect default fg on GTK-4 view_bg | `tm:__fg` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | syntect default fg on GTK-4 popover_bg | `tm:__fg` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ❌ | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #272e33 | `base00` #272e33 | 1.00:1 |

### stack — ff-in-gtk4 (4/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar_text on GTK-4 headerbar | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ✅ | toolbar_field_text on toolbar_field | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ✅ | toolbar_field_border_focus on field bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | tab_line (accent) on GTK-4 view_bg below | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |

### stack — ff-private (4/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #7fbbb3 | `FF private toolbar_field #42414d` #42414d | 4.61:1 |
| ✅ | accent / focus_border (scheme accent) on chrome bg (FF forced) | `accent` #7fbbb3 | `FF private chrome bg #1c1b22` #1c1b22 | 7.86:1 |
| ✅ | toolbar_field_text (dark scheme base05) on toolbar_field (FF forced) — control | `base05` #d3c6aa | `FF private toolbar_field #42414d` #42414d | 5.93:1 |
| ✅ | toolbar_text (dark scheme base05) on chrome bg (FF forced) — control | `base05` #d3c6aa | `FF private chrome bg #1c1b22` #1c1b22 | 10.11:1 |

### stack — fuzzel-on-gtk4 (4/8 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | selected_fg (base05) on selected_bg (base02) | `base05` #d3c6aa | `base02` #414b50 | 5.30:1 |
| ❌ | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #414b50 | `base01` #2e383c | 1.34:1 |
| ❌ | selection_match (accent) on selected_bg (base02) | `accent` #7fbbb3 | `base02` #414b50 | 4.12:1 |
| ✅ | match (accent) on fuzzel bg (base01) | `accent` #7fbbb3 | `base01` #2e383c | 5.53:1 |
| ✅ | border (accent) vs GTK-4 window_bg behind | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | text (base05) on fuzzel bg, GTK-4 window behind | `base05` #d3c6aa | `base01` #2e383c | 7.11:1 |
| ❌ | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #414b50 | `base00` #272e33 | 1.54:1 |
| ❌ | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #414b50 | `base01` #2e383c | 1.34:1 |

### stack — mako-on-gtk4 (2/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | border (accent) vs GTK-4 window_bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | text (base05) vs GTK-4 window_bg seam | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |

### stack — openbox-on-gtk4 (2/3 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active border (accent) on GTK-4 view_bg | `accent` #7fbbb3 | `base00` #272e33 | 6.34:1 |
| ✅ | active title text on GTK-4 view_bg seam | `base05` #d3c6aa | `base00` #272e33 | 8.15:1 |
| ❌ | inactive border (base02) on GTK-4 view_bg | `base02` #414b50 | `base00` #272e33 | 1.54:1 |

