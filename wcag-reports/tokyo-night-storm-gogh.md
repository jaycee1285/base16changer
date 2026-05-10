# WCAG AA report — Tokyo Night Storm

- **source**: `/home/john/.local/share/themes/Tokyo Night Storm.yml`
- **format**: gogh
- **variant**: dark
- **accent**: #f7768e (explicit)
- **threshold**: AA strict (≥ 4.5:1 for all roles)
- **tmTheme**: `/home/john/.local/share/themes/tmThemes/Tokyo Night Storm.tmTheme` (bg=#24283b, fg=#c0caf5)
- **summary**: 122 pass / 29 fail / 151 total

## Palette

| Slot | Hex |
|------|-----|
| base00 | `#24283b` |
| base01 | `#34384e` |
| base02 | `#434860` |
| base03 | `#414868` |
| base04 | `#626985` |
| base05 | `#c0caf5` |
| base06 | `#a1aad0` |
| base07 | `#c0caf5` |
| base08 | `#f7768e` |
| base09 | `#ec937b` |
| base0A | `#e0af68` |
| base0B | `#9ece6a` |
| base0C | `#7dcfff` |
| base0D | `#7aa2f7` |
| base0E | `#bb9af7` |
| base0F | `#9c6861` |
| accent | `#f7768e` |

## Failures

| Group | Target | Role | fg | bg | Ratio | Suggested swap | Notes |
|-------|--------|------|----|----|-------|----------------|-------|
| intra | kitty | url_color on background | `base04` #626985 | `base00` #24283b | **2.69:1** | `base05` (#c0caf5, 9.02:1) |  |
| intra | kitty | active border on bg | `base03` #414868 | `base00` #24283b | **1.63:1** | `base05` (#c0caf5, 9.02:1) |  |
| intra | kitty | bright black (comments) on bg | `base03` #414868 | `base00` #24283b | **1.63:1** | `base05` (#c0caf5, 9.02:1) |  |
| intra | fuzzel | match on background | `accent` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) | user-flagged: accent often disappears when accent==base0D and palette has unsaturated blue |
| intra | fuzzel | selection-match on selection | `accent` #f7768e | `base02` #434860 | **3.40:1** | `base0C` (#7dcfff, 5.24:1) |  |
| intra | fuzzel | border on bg (visual indicator) | `accent` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | mako | border on bg | `accent` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | openbox | inactive button icon on inactive btn bg | `base03` #414868 | `base00` #24283b | **1.63:1** | `base05` (#c0caf5, 9.02:1) |  |
| intra | openbox | active border (outline) on active title | `accent` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | openbox | close button (red) on active title | `base08` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | gtk3 | border (visual UI) on bg | `base02` #434860 | `base00` #24283b | **1.62:1** | — |  |
| intra | gtk4 | accent_color on headerbar_bg | `accent` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | gtk4 | accent_color on sidebar_bg | `accent` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | gtk4 | accent_color on card_bg | `accent` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | gtk4 | scrollbar outline on view_bg | `base02` #434860 | `base00` #24283b | **1.62:1** | — |  |
| intra | gsv | bracket-mismatch fg on bg | `base08` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | firefox | selected suggestion url on hover bg | `accent` #f7768e | `base02` #434860 | **3.40:1** | `base0C` (#7dcfff, 5.24:1) |  |
| intra | firefox | danger (.close-icon hover) on toolbar | `base08` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| intra | sidebery | inactive tab fg on frame bg | `base04` #626985 | `base00` #24283b | **2.69:1** | `base05` (#c0caf5, 9.02:1) | user-flagged equivalent: muted inactive tab text |
| stack | code-in-gtk4 | syntect HEADING on GTK-4 selection_bg | `tm:heading` #7aa2f7 | `base02` #434860 | **3.57:1** | — |  |
| stack | code-in-gtk4 | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #24283b | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #24283b | `base00` #24283b | **1.00:1** | — | if tmTheme bg differs from GTK-4 view bg, the gutter strip seam fails |
| stack | ff-private | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #f7768e | `FF private toolbar_field #42414d` #42414d | **3.79:1** | `base0C` (#7dcfff, 5.84:1) |  |
| stack | fuzzel-on-gtk4 | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #434860 | `base01` #34384e | **1.28:1** | — | USER-FLAGGED #3b (post-T3): stripe is now base02 vs base01 |
| stack | fuzzel-on-gtk4 | selection_match (accent) on selected_bg (base02) | `accent` #f7768e | `base02` #434860 | **3.40:1** | `base0C` (#7dcfff, 5.24:1) | USER-FLAGGED #3c (post-T3): typed-letter highlight on the new base02 selection |
| stack | fuzzel-on-gtk4 | match (accent) on fuzzel bg (base01) | `accent` #f7768e | `base01` #34384e | **4.36:1** | `base0C` (#7dcfff, 6.72:1) |  |
| stack | fuzzel-on-gtk4 | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #434860 | `base00` #24283b | **1.62:1** | — | GTK-4 analog of fuzzel #3b |
| stack | fuzzel-on-gtk4 | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #434860 | `base01` #34384e | **1.28:1** | — | GTK-3 analog of fuzzel #3b |
| stack | openbox-on-gtk4 | inactive border (base02) on GTK-4 view_bg | `base02` #434860 | `base00` #24283b | **1.62:1** | — |  |

## All pairs

### intra — firefox (15/17 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar text on toolbar bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | field (URL bar) text on field bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | field focus border on field bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | field highlight on field bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | field highlight text on highlight | `base00` #24283b | `accent` #f7768e | 5.51:1 |
| ✅ | popup text on popup bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | popup border on popup bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ❌ | selected suggestion url on hover bg | `accent` #f7768e | `base02` #434860 | 3.40:1 |
| ✅ | sidebar text on sidebar bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | sidebar border on sidebar bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | muted text on toolbar bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | tab selected fg on tab selected bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | tab line (accent underline) on tab bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | tab hover fg on tab hover bg | `base05` #c0caf5 | `base02` #434860 | 5.57:1 |
| ❌ | danger (.close-icon hover) on toolbar | `base08` #f7768e | `base01` #34384e | 4.36:1 |
| ✅ | warning (attention) on toolbar | `base0A` #e0af68 | `base01` #34384e | 5.76:1 |
| ✅ | success (download badge) on toolbar | `base0B` #9ece6a | `base01` #34384e | 6.31:1 |

### intra — fuzzel (3/6 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ❌ | match on background | `accent` #f7768e | `base01` #34384e | 4.36:1 |
| ✅ | selection-text on selection | `base05` #c0caf5 | `base02` #434860 | 5.57:1 |
| ❌ | selection-match on selection | `accent` #f7768e | `base02` #434860 | 3.40:1 |
| ❌ | border on bg (visual indicator) | `accent` #f7768e | `base01` #34384e | 4.36:1 |
| ✅ | prompt/text on bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |

### intra — gsv (23/24 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on view bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | text on current_line bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | comment on view bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | string on view bg | `base0B` #9ece6a | `base00` #24283b | 7.97:1 |
| ✅ | string on current_line | `base0B` #9ece6a | `base01` #34384e | 6.31:1 |
| ✅ | number on view bg | `base09` #ec937b | `base00` #24283b | 6.26:1 |
| ✅ | function on view bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | function on current_line | `base0D` #7aa2f7 | `base01` #34384e | 4.58:1 |
| ✅ | keyword on view bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | keyword on current_line | `base0E` #bb9af7 | `base01` #34384e | 4.98:1 |
| ✅ | type on view bg | `base0B` #9ece6a | `base00` #24283b | 7.97:1 |
| ✅ | builtin on view bg | `base0C` #7dcfff | `base00` #24283b | 8.49:1 |
| ✅ | constant/special on view bg | `base0A` #e0af68 | `base00` #24283b | 7.28:1 |
| ✅ | operator on view bg | `base0B` #9ece6a | `base00` #24283b | 7.97:1 |
| ✅ | preprocessor on view bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | heading on view bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | list-marker on view bg | `base09` #ec937b | `base00` #24283b | 6.26:1 |
| ✅ | link-text on view bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | link-destination on view bg | `base0C` #7dcfff | `base00` #24283b | 8.49:1 |
| ✅ | search match on match bg | `base05` #c0caf5 | `base02` #434860 | 5.57:1 |
| ✅ | bracket match on view bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ❌ | bracket-mismatch fg on bg | `base08` #f7768e | `base01` #34384e | 4.36:1 |
| ✅ | right_margin text on margin bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | line-numbers on gutter | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |

### intra — gtk3 (10/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | fg on bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | text on base (content area) | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | disabled text on bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | selected text on selected bg | `base05` #c0caf5 | `base02` #434860 | 5.57:1 |
| ✅ | tooltip text on tooltip bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | link on bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | link on base (content) | `base0D` #7aa2f7 | `base01` #34384e | 4.58:1 |
| ✅ | success on bg | `base0B` #9ece6a | `base00` #24283b | 7.97:1 |
| ✅ | warning on bg | `base0A` #e0af68 | `base00` #24283b | 7.28:1 |
| ✅ | error on bg | `base08` #f7768e | `base00` #24283b | 5.51:1 |
| ❌ | border (visual UI) on bg | `base02` #434860 | `base00` #24283b | 1.62:1 |

### intra — gtk4 (15/19 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | window_fg on window_bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | view_fg on view_bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | headerbar_fg on headerbar_bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | sidebar_fg on sidebar_bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | card_fg on card_bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | dialog_fg on dialog_bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | popover_fg on popover_bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | accent_color on window_bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | accent_color on view_bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ❌ | accent_color on headerbar_bg | `accent` #f7768e | `base01` #34384e | 4.36:1 |
| ❌ | accent_color on sidebar_bg | `accent` #f7768e | `base01` #34384e | 4.36:1 |
| ❌ | accent_color on card_bg | `accent` #f7768e | `base01` #34384e | 4.36:1 |
| ✅ | accent_fg_color on accent_bg_color | `base00` #24283b | `accent` #f7768e | 5.51:1 |
| ✅ | destructive_fg on destructive_bg | `base00` #24283b | `base08` #f7768e | 5.51:1 |
| ✅ | success_fg on success_bg | `base00` #24283b | `base0B` #9ece6a | 7.97:1 |
| ✅ | warning_fg on warning_bg | `base00` #24283b | `base0A` #e0af68 | 7.28:1 |
| ✅ | error_fg on error_bg | `base00` #24283b | `base08` #f7768e | 5.51:1 |
| ✅ | blue_3 link on view_bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ❌ | scrollbar outline on view_bg | `base02` #434860 | `base00` #24283b | 1.62:1 |

### intra — kitty (12/15 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | default text on background | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | selection text on selection bg | `base02` #434860 | `base05` #c0caf5 | 5.57:1 |
| ✅ | cursor on background | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | cursor_text on cursor | `base00` #24283b | `base05` #c0caf5 | 9.02:1 |
| ❌ | url_color on background | `base04` #626985 | `base00` #24283b | 2.69:1 |
| ❌ | active border on bg | `base03` #414868 | `base00` #24283b | 1.63:1 |
| ✅ | active tab fg on active tab bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | inactive tab fg on inactive tab bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | ANSI red on bg (color_01 logs) | `base08` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | ANSI green on bg | `base0B` #9ece6a | `base00` #24283b | 7.97:1 |
| ✅ | ANSI yellow on bg | `base0A` #e0af68 | `base00` #24283b | 7.28:1 |
| ✅ | ANSI blue on bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | ANSI magenta on bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | ANSI cyan on bg | `base0C` #7dcfff | `base00` #24283b | 8.49:1 |
| ❌ | bright black (comments) on bg | `base03` #414868 | `base00` #24283b | 1.63:1 |

### intra — mako (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ❌ | border on bg | `accent` #f7768e | `base01` #34384e | 4.36:1 |

### intra — openbox (8/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active title text on active title bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | inactive title text on inactive title bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | active button icon on button bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | active button icon hover | `base07` #c0caf5 | `base02` #434860 | 5.57:1 |
| ❌ | inactive button icon on inactive btn bg | `base03` #414868 | `base00` #24283b | 1.63:1 |
| ❌ | active border (outline) on active title | `accent` #f7768e | `base01` #34384e | 4.36:1 |
| ✅ | menu items on menu bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | menu disabled on menu bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | menu active item on highlight | `base05` #c0caf5 | `base02` #434860 | 5.57:1 |
| ✅ | OSD label on osd bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ❌ | close button (red) on active title | `base08` #f7768e | `base01` #34384e | 4.36:1 |

### intra — sidebery (6/7 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active tab fg on active tab bg | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ❌ | inactive tab fg on frame bg | `base04` #626985 | `base00` #24283b | 2.69:1 |
| ✅ | toolbar fg on toolbar bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | tab hover fg on hover bg | `base05` #c0caf5 | `base02` #434860 | 5.57:1 |
| ✅ | popup fg on popup bg | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | scroll progress (accent) on frame bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | active-tab outline (accent) on tab | `accent` #f7768e | `base00` #24283b | 5.51:1 |

### stack — code-in-gtk4 (15/18 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | syntect HEADING on GTK-4 view_bg | `tm:heading` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | syntect HEADING on GTK-4 popover_bg | `tm:heading` #7aa2f7 | `base01` #34384e | 4.58:1 |
| ✅ | syntect HEADING on GTK-4 sidebar_bg | `tm:heading` #7aa2f7 | `base01` #34384e | 4.58:1 |
| ❌ | syntect HEADING on GTK-4 selection_bg | `tm:heading` #7aa2f7 | `base02` #434860 | 3.57:1 |
| ✅ | syntect comment on GTK-4 view_bg | `tm:comment` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | syntect comment on GTK-4 popover_bg | `tm:comment` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | syntect string on GTK-4 view_bg | `tm:string` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | syntect string on GTK-4 popover_bg | `tm:string` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | syntect keyword on GTK-4 view_bg | `tm:keyword` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | syntect keyword on GTK-4 popover_bg | `tm:keyword` #bb9af7 | `base01` #34384e | 4.98:1 |
| ✅ | syntect function on GTK-4 view_bg | `tm:function` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | syntect function on GTK-4 popover_bg | `tm:function` #7aa2f7 | `base01` #34384e | 4.58:1 |
| ✅ | syntect type on GTK-4 view_bg | `tm:type` #e0af68 | `base00` #24283b | 7.28:1 |
| ✅ | syntect number on GTK-4 view_bg | `tm:number` #ed957c | `base00` #24283b | 6.38:1 |
| ❌ | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #24283b | NaN:1 |
| ✅ | syntect default fg on GTK-4 view_bg | `tm:__fg` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | syntect default fg on GTK-4 popover_bg | `tm:__fg` #c0caf5 | `base01` #34384e | 7.14:1 |
| ❌ | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #24283b | `base00` #24283b | 1.00:1 |

### stack — ff-in-gtk4 (4/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar_text on GTK-4 headerbar | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ✅ | toolbar_field_text on toolbar_field | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ✅ | toolbar_field_border_focus on field bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | tab_line (accent) on GTK-4 view_bg below | `accent` #f7768e | `base00` #24283b | 5.51:1 |

### stack — ff-private (3/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #f7768e | `FF private toolbar_field #42414d` #42414d | 3.79:1 |
| ✅ | accent / focus_border (scheme accent) on chrome bg (FF forced) | `accent` #f7768e | `FF private chrome bg #1c1b22` #1c1b22 | 6.46:1 |
| ✅ | toolbar_field_text (dark scheme base05) on toolbar_field (FF forced) — control | `base05` #c0caf5 | `FF private toolbar_field #42414d` #42414d | 6.21:1 |
| ✅ | toolbar_text (dark scheme base05) on chrome bg (FF forced) — control | `base05` #c0caf5 | `FF private chrome bg #1c1b22` #1c1b22 | 10.58:1 |

### stack — fuzzel-on-gtk4 (3/8 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | selected_fg (base05) on selected_bg (base02) | `base05` #c0caf5 | `base02` #434860 | 5.57:1 |
| ❌ | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #434860 | `base01` #34384e | 1.28:1 |
| ❌ | selection_match (accent) on selected_bg (base02) | `accent` #f7768e | `base02` #434860 | 3.40:1 |
| ❌ | match (accent) on fuzzel bg (base01) | `accent` #f7768e | `base01` #34384e | 4.36:1 |
| ✅ | border (accent) vs GTK-4 window_bg behind | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | text (base05) on fuzzel bg, GTK-4 window behind | `base05` #c0caf5 | `base01` #34384e | 7.14:1 |
| ❌ | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #434860 | `base00` #24283b | 1.62:1 |
| ❌ | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #434860 | `base01` #34384e | 1.28:1 |

### stack — mako-on-gtk4 (2/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | border (accent) vs GTK-4 window_bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | text (base05) vs GTK-4 window_bg seam | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |

### stack — openbox-on-gtk4 (2/3 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active border (accent) on GTK-4 view_bg | `accent` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | active title text on GTK-4 view_bg seam | `base05` #c0caf5 | `base00` #24283b | 9.02:1 |
| ❌ | inactive border (base02) on GTK-4 view_bg | `base02` #434860 | `base00` #24283b | 1.62:1 |

