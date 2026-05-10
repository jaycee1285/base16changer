# WCAG AA report — Kanagawa Wave

- **source**: `/home/john/.local/share/themes/Kanagawa Wave.yml`
- **format**: gogh
- **variant**: dark
- **accent**: #e82424 (explicit)
- **threshold**: AA strict (≥ 4.5:1 for all roles)
- **tmTheme**: `/home/john/.local/share/themes/tmThemes/Kanagawa Wave.tmTheme` (bg=#1f1f28, fg=#dcd7ba)
- **summary**: 99 pass / 52 fail / 151 total

## Palette

| Slot | Hex |
|------|-----|
| base00 | `#1f1f28` |
| base01 | `#323137` |
| base02 | `#454445` |
| base03 | `#727169` |
| base04 | `#6b6962` |
| base05 | `#dcd7ba` |
| base06 | `#b6b29d` |
| base07 | `#dcd7ba` |
| base08 | `#c34043` |
| base09 | `#c27259` |
| base0A | `#c0a36e` |
| base0B | `#76946a` |
| base0C | `#6a9589` |
| base0D | `#7e9cd8` |
| base0E | `#957fb8` |
| base0F | `#815145` |
| accent | `#e82424` |

## Failures

| Group | Target | Role | fg | bg | Ratio | Suggested swap | Notes |
|-------|--------|------|----|----|-------|----------------|-------|
| intra | kitty | url_color on background | `base04` #6b6962 | `base00` #1f1f28 | **2.97:1** | `base05` (#dcd7ba, 11.26:1) |  |
| intra | kitty | active border on bg | `base03` #727169 | `base00` #1f1f28 | **3.33:1** | `base05` (#dcd7ba, 11.26:1) |  |
| intra | kitty | ANSI red on bg (color_01 logs) | `base08` #c34043 | `base00` #1f1f28 | **3.22:1** | `base0A` (#c0a36e, 6.78:1) |  |
| intra | kitty | bright black (comments) on bg | `base03` #727169 | `base00` #1f1f28 | **3.33:1** | `base05` (#dcd7ba, 11.26:1) |  |
| intra | fuzzel | match on background | `accent` #e82424 | `base01` #323137 | **2.89:1** | `base0A` (#c0a36e, 5.34:1) | user-flagged: accent often disappears when accent==base0D and palette has unsaturated blue |
| intra | fuzzel | selection-match on selection | `accent` #e82424 | `base02` #454445 | **2.17:1** | — |  |
| intra | fuzzel | border on bg (visual indicator) | `accent` #e82424 | `base01` #323137 | **2.89:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | mako | border on bg | `accent` #e82424 | `base01` #323137 | **2.89:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | openbox | inactive button icon on inactive btn bg | `base03` #727169 | `base00` #1f1f28 | **3.33:1** | `base05` (#dcd7ba, 11.26:1) |  |
| intra | openbox | active border (outline) on active title | `accent` #e82424 | `base01` #323137 | **2.89:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | openbox | close button (red) on active title | `base08` #c34043 | `base01` #323137 | **2.54:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | gtk3 | error on bg | `base08` #c34043 | `base00` #1f1f28 | **3.22:1** | `base0A` (#c0a36e, 6.78:1) |  |
| intra | gtk3 | border (visual UI) on bg | `base02` #454445 | `base00` #1f1f28 | **1.69:1** | — |  |
| intra | gtk4 | accent_color on window_bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) | this is the pair the accent: field is meant to fix |
| intra | gtk4 | accent_color on view_bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| intra | gtk4 | accent_color on headerbar_bg | `accent` #e82424 | `base01` #323137 | **2.89:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | gtk4 | accent_color on sidebar_bg | `accent` #e82424 | `base01` #323137 | **2.89:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | gtk4 | accent_color on card_bg | `accent` #e82424 | `base01` #323137 | **2.89:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | gtk4 | accent_fg_color on accent_bg_color | `base00` #1f1f28 | `accent` #e82424 | **3.66:1** | — | text rendered on solid accent button |
| intra | gtk4 | destructive_fg on destructive_bg | `base00` #1f1f28 | `base08` #c34043 | **3.22:1** | — |  |
| intra | gtk4 | error_fg on error_bg | `base00` #1f1f28 | `base08` #c34043 | **3.22:1** | — |  |
| intra | gtk4 | scrollbar outline on view_bg | `base02` #454445 | `base00` #1f1f28 | **1.69:1** | — |  |
| intra | gsv | string on current_line | `base0B` #76946a | `base01` #323137 | **3.81:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | gsv | keyword on current_line | `base0E` #957fb8 | `base01` #323137 | **3.68:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | gsv | bracket-mismatch fg on bg | `base08` #c34043 | `base01` #323137 | **2.54:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | firefox | field focus border on field bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| intra | firefox | field highlight on field bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| intra | firefox | field highlight text on highlight | `base00` #1f1f28 | `accent` #e82424 | **3.66:1** | — |  |
| intra | firefox | selected suggestion url on hover bg | `accent` #e82424 | `base02` #454445 | **2.17:1** | — |  |
| intra | firefox | tab line (accent underline) on tab bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| intra | firefox | danger (.close-icon hover) on toolbar | `base08` #c34043 | `base01` #323137 | **2.54:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | firefox | success (download badge) on toolbar | `base0B` #76946a | `base01` #323137 | **3.81:1** | `base0A` (#c0a36e, 5.34:1) |  |
| intra | sidebery | inactive tab fg on frame bg | `base04` #6b6962 | `base00` #1f1f28 | **2.97:1** | `base05` (#dcd7ba, 11.26:1) | user-flagged equivalent: muted inactive tab text |
| intra | sidebery | scroll progress (accent) on frame bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| intra | sidebery | active-tab outline (accent) on tab | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| stack | code-in-gtk4 | syntect HEADING on GTK-4 selection_bg | `tm:heading` #7e9cd8 | `base02` #454445 | **3.52:1** | — |  |
| stack | code-in-gtk4 | syntect keyword on GTK-4 popover_bg | `tm:keyword` #957fb8 | `base01` #323137 | **3.68:1** | — |  |
| stack | code-in-gtk4 | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #1f1f28 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #1f1f28 | `base00` #1f1f28 | **1.00:1** | — | if tmTheme bg differs from GTK-4 view bg, the gutter strip seam fails |
| stack | ff-in-gtk4 | toolbar_field_border_focus on field bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| stack | ff-in-gtk4 | tab_line (accent) on GTK-4 view_bg below | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| stack | ff-private | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #e82424 | `FF private toolbar_field #42414d` #42414d | **2.25:1** | — |  |
| stack | ff-private | accent / focus_border (scheme accent) on chrome bg (FF forced) | `accent` #e82424 | `FF private chrome bg #1c1b22` #1c1b22 | **3.83:1** | `base0A` (#c0a36e, 7.08:1) |  |
| stack | fuzzel-on-gtk4 | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #454445 | `base01` #323137 | **1.33:1** | — | USER-FLAGGED #3b (post-T3): stripe is now base02 vs base01 |
| stack | fuzzel-on-gtk4 | selection_match (accent) on selected_bg (base02) | `accent` #e82424 | `base02` #454445 | **2.17:1** | — | USER-FLAGGED #3c (post-T3): typed-letter highlight on the new base02 selection |
| stack | fuzzel-on-gtk4 | match (accent) on fuzzel bg (base01) | `accent` #e82424 | `base01` #323137 | **2.89:1** | `base0A` (#c0a36e, 5.34:1) |  |
| stack | fuzzel-on-gtk4 | border (accent) vs GTK-4 window_bg behind | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| stack | fuzzel-on-gtk4 | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #454445 | `base00` #1f1f28 | **1.69:1** | — | GTK-4 analog of fuzzel #3b |
| stack | fuzzel-on-gtk4 | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #454445 | `base01` #323137 | **1.33:1** | — | GTK-3 analog of fuzzel #3b |
| stack | openbox-on-gtk4 | active border (accent) on GTK-4 view_bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |
| stack | openbox-on-gtk4 | inactive border (base02) on GTK-4 view_bg | `base02` #454445 | `base00` #1f1f28 | **1.69:1** | — |  |
| stack | mako-on-gtk4 | border (accent) vs GTK-4 window_bg | `accent` #e82424 | `base00` #1f1f28 | **3.66:1** | `base0A` (#c0a36e, 6.78:1) |  |

## All pairs

### intra — firefox (10/17 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar text on toolbar bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | field (URL bar) text on field bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ❌ | field focus border on field bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ❌ | field highlight on field bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ❌ | field highlight text on highlight | `base00` #1f1f28 | `accent` #e82424 | 3.66:1 |
| ✅ | popup text on popup bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | popup border on popup bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ❌ | selected suggestion url on hover bg | `accent` #e82424 | `base02` #454445 | 2.17:1 |
| ✅ | sidebar text on sidebar bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | sidebar border on sidebar bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | muted text on toolbar bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | tab selected fg on tab selected bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ❌ | tab line (accent underline) on tab bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ✅ | tab hover fg on tab hover bg | `base05` #dcd7ba | `base02` #454445 | 6.68:1 |
| ❌ | danger (.close-icon hover) on toolbar | `base08` #c34043 | `base01` #323137 | 2.54:1 |
| ✅ | warning (attention) on toolbar | `base0A` #c0a36e | `base01` #323137 | 5.34:1 |
| ❌ | success (download badge) on toolbar | `base0B` #76946a | `base01` #323137 | 3.81:1 |

### intra — fuzzel (3/6 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ❌ | match on background | `accent` #e82424 | `base01` #323137 | 2.89:1 |
| ✅ | selection-text on selection | `base05` #dcd7ba | `base02` #454445 | 6.68:1 |
| ❌ | selection-match on selection | `accent` #e82424 | `base02` #454445 | 2.17:1 |
| ❌ | border on bg (visual indicator) | `accent` #e82424 | `base01` #323137 | 2.89:1 |
| ✅ | prompt/text on bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |

### intra — gsv (21/24 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on view bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | text on current_line bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | comment on view bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | string on view bg | `base0B` #76946a | `base00` #1f1f28 | 4.84:1 |
| ❌ | string on current_line | `base0B` #76946a | `base01` #323137 | 3.81:1 |
| ✅ | number on view bg | `base09` #c27259 | `base00` #1f1f28 | 4.55:1 |
| ✅ | function on view bg | `base0D` #7e9cd8 | `base00` #1f1f28 | 5.94:1 |
| ✅ | function on current_line | `base0D` #7e9cd8 | `base01` #323137 | 4.68:1 |
| ✅ | keyword on view bg | `base0E` #957fb8 | `base00` #1f1f28 | 4.67:1 |
| ❌ | keyword on current_line | `base0E` #957fb8 | `base01` #323137 | 3.68:1 |
| ✅ | type on view bg | `base0B` #76946a | `base00` #1f1f28 | 4.84:1 |
| ✅ | builtin on view bg | `base0C` #6a9589 | `base00` #1f1f28 | 4.88:1 |
| ✅ | constant/special on view bg | `base0A` #c0a36e | `base00` #1f1f28 | 6.78:1 |
| ✅ | operator on view bg | `base0B` #76946a | `base00` #1f1f28 | 4.84:1 |
| ✅ | preprocessor on view bg | `base0E` #957fb8 | `base00` #1f1f28 | 4.67:1 |
| ✅ | heading on view bg | `base0D` #7e9cd8 | `base00` #1f1f28 | 5.94:1 |
| ✅ | list-marker on view bg | `base09` #c27259 | `base00` #1f1f28 | 4.55:1 |
| ✅ | link-text on view bg | `base0E` #957fb8 | `base00` #1f1f28 | 4.67:1 |
| ✅ | link-destination on view bg | `base0C` #6a9589 | `base00` #1f1f28 | 4.88:1 |
| ✅ | search match on match bg | `base05` #dcd7ba | `base02` #454445 | 6.68:1 |
| ✅ | bracket match on view bg | `base0E` #957fb8 | `base00` #1f1f28 | 4.67:1 |
| ❌ | bracket-mismatch fg on bg | `base08` #c34043 | `base01` #323137 | 2.54:1 |
| ✅ | right_margin text on margin bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | line-numbers on gutter | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |

### intra — gtk3 (9/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | fg on bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | text on base (content area) | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | disabled text on bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | selected text on selected bg | `base05` #dcd7ba | `base02` #454445 | 6.68:1 |
| ✅ | tooltip text on tooltip bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | link on bg | `base0D` #7e9cd8 | `base00` #1f1f28 | 5.94:1 |
| ✅ | link on base (content) | `base0D` #7e9cd8 | `base01` #323137 | 4.68:1 |
| ✅ | success on bg | `base0B` #76946a | `base00` #1f1f28 | 4.84:1 |
| ✅ | warning on bg | `base0A` #c0a36e | `base00` #1f1f28 | 6.78:1 |
| ❌ | error on bg | `base08` #c34043 | `base00` #1f1f28 | 3.22:1 |
| ❌ | border (visual UI) on bg | `base02` #454445 | `base00` #1f1f28 | 1.69:1 |

### intra — gtk4 (10/19 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | window_fg on window_bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | view_fg on view_bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | headerbar_fg on headerbar_bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | sidebar_fg on sidebar_bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | card_fg on card_bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | dialog_fg on dialog_bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | popover_fg on popover_bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ❌ | accent_color on window_bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ❌ | accent_color on view_bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ❌ | accent_color on headerbar_bg | `accent` #e82424 | `base01` #323137 | 2.89:1 |
| ❌ | accent_color on sidebar_bg | `accent` #e82424 | `base01` #323137 | 2.89:1 |
| ❌ | accent_color on card_bg | `accent` #e82424 | `base01` #323137 | 2.89:1 |
| ❌ | accent_fg_color on accent_bg_color | `base00` #1f1f28 | `accent` #e82424 | 3.66:1 |
| ❌ | destructive_fg on destructive_bg | `base00` #1f1f28 | `base08` #c34043 | 3.22:1 |
| ✅ | success_fg on success_bg | `base00` #1f1f28 | `base0B` #76946a | 4.84:1 |
| ✅ | warning_fg on warning_bg | `base00` #1f1f28 | `base0A` #c0a36e | 6.78:1 |
| ❌ | error_fg on error_bg | `base00` #1f1f28 | `base08` #c34043 | 3.22:1 |
| ✅ | blue_3 link on view_bg | `base0D` #7e9cd8 | `base00` #1f1f28 | 5.94:1 |
| ❌ | scrollbar outline on view_bg | `base02` #454445 | `base00` #1f1f28 | 1.69:1 |

### intra — kitty (11/15 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | default text on background | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | selection text on selection bg | `base02` #454445 | `base05` #dcd7ba | 6.68:1 |
| ✅ | cursor on background | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | cursor_text on cursor | `base00` #1f1f28 | `base05` #dcd7ba | 11.26:1 |
| ❌ | url_color on background | `base04` #6b6962 | `base00` #1f1f28 | 2.97:1 |
| ❌ | active border on bg | `base03` #727169 | `base00` #1f1f28 | 3.33:1 |
| ✅ | active tab fg on active tab bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | inactive tab fg on inactive tab bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ❌ | ANSI red on bg (color_01 logs) | `base08` #c34043 | `base00` #1f1f28 | 3.22:1 |
| ✅ | ANSI green on bg | `base0B` #76946a | `base00` #1f1f28 | 4.84:1 |
| ✅ | ANSI yellow on bg | `base0A` #c0a36e | `base00` #1f1f28 | 6.78:1 |
| ✅ | ANSI blue on bg | `base0D` #7e9cd8 | `base00` #1f1f28 | 5.94:1 |
| ✅ | ANSI magenta on bg | `base0E` #957fb8 | `base00` #1f1f28 | 4.67:1 |
| ✅ | ANSI cyan on bg | `base0C` #6a9589 | `base00` #1f1f28 | 4.88:1 |
| ❌ | bright black (comments) on bg | `base03` #727169 | `base00` #1f1f28 | 3.33:1 |

### intra — mako (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ❌ | border on bg | `accent` #e82424 | `base01` #323137 | 2.89:1 |

### intra — openbox (8/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active title text on active title bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | inactive title text on inactive title bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | active button icon on button bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | active button icon hover | `base07` #dcd7ba | `base02` #454445 | 6.68:1 |
| ❌ | inactive button icon on inactive btn bg | `base03` #727169 | `base00` #1f1f28 | 3.33:1 |
| ❌ | active border (outline) on active title | `accent` #e82424 | `base01` #323137 | 2.89:1 |
| ✅ | menu items on menu bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | menu disabled on menu bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | menu active item on highlight | `base05` #dcd7ba | `base02` #454445 | 6.68:1 |
| ✅ | OSD label on osd bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ❌ | close button (red) on active title | `base08` #c34043 | `base01` #323137 | 2.54:1 |

### intra — sidebery (4/7 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active tab fg on active tab bg | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ❌ | inactive tab fg on frame bg | `base04` #6b6962 | `base00` #1f1f28 | 2.97:1 |
| ✅ | toolbar fg on toolbar bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | tab hover fg on hover bg | `base05` #dcd7ba | `base02` #454445 | 6.68:1 |
| ✅ | popup fg on popup bg | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ❌ | scroll progress (accent) on frame bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ❌ | active-tab outline (accent) on tab | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |

### stack — code-in-gtk4 (14/18 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | syntect HEADING on GTK-4 view_bg | `tm:heading` #7e9cd8 | `base00` #1f1f28 | 5.94:1 |
| ✅ | syntect HEADING on GTK-4 popover_bg | `tm:heading` #7e9cd8 | `base01` #323137 | 4.68:1 |
| ✅ | syntect HEADING on GTK-4 sidebar_bg | `tm:heading` #7e9cd8 | `base01` #323137 | 4.68:1 |
| ❌ | syntect HEADING on GTK-4 selection_bg | `tm:heading` #7e9cd8 | `base02` #454445 | 3.52:1 |
| ✅ | syntect comment on GTK-4 view_bg | `tm:comment` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | syntect comment on GTK-4 popover_bg | `tm:comment` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | syntect string on GTK-4 view_bg | `tm:string` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | syntect string on GTK-4 popover_bg | `tm:string` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | syntect keyword on GTK-4 view_bg | `tm:keyword` #957fb8 | `base00` #1f1f28 | 4.67:1 |
| ❌ | syntect keyword on GTK-4 popover_bg | `tm:keyword` #957fb8 | `base01` #323137 | 3.68:1 |
| ✅ | syntect function on GTK-4 view_bg | `tm:function` #7e9cd8 | `base00` #1f1f28 | 5.94:1 |
| ✅ | syntect function on GTK-4 popover_bg | `tm:function` #7e9cd8 | `base01` #323137 | 4.68:1 |
| ✅ | syntect type on GTK-4 view_bg | `tm:type` #c0a36e | `base00` #1f1f28 | 6.78:1 |
| ✅ | syntect number on GTK-4 view_bg | `tm:number` #c47658 | `base00` #1f1f28 | 4.73:1 |
| ❌ | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #1f1f28 | NaN:1 |
| ✅ | syntect default fg on GTK-4 view_bg | `tm:__fg` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ✅ | syntect default fg on GTK-4 popover_bg | `tm:__fg` #dcd7ba | `base01` #323137 | 8.88:1 |
| ❌ | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #1f1f28 | `base00` #1f1f28 | 1.00:1 |

### stack — ff-in-gtk4 (2/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar_text on GTK-4 headerbar | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ✅ | toolbar_field_text on toolbar_field | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ❌ | toolbar_field_border_focus on field bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ❌ | tab_line (accent) on GTK-4 view_bg below | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |

### stack — ff-private (2/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #e82424 | `FF private toolbar_field #42414d` #42414d | 2.25:1 |
| ❌ | accent / focus_border (scheme accent) on chrome bg (FF forced) | `accent` #e82424 | `FF private chrome bg #1c1b22` #1c1b22 | 3.83:1 |
| ✅ | toolbar_field_text (dark scheme base05) on toolbar_field (FF forced) — control | `base05` #dcd7ba | `FF private toolbar_field #42414d` #42414d | 6.91:1 |
| ✅ | toolbar_text (dark scheme base05) on chrome bg (FF forced) — control | `base05` #dcd7ba | `FF private chrome bg #1c1b22` #1c1b22 | 11.77:1 |

### stack — fuzzel-on-gtk4 (2/8 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | selected_fg (base05) on selected_bg (base02) | `base05` #dcd7ba | `base02` #454445 | 6.68:1 |
| ❌ | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #454445 | `base01` #323137 | 1.33:1 |
| ❌ | selection_match (accent) on selected_bg (base02) | `accent` #e82424 | `base02` #454445 | 2.17:1 |
| ❌ | match (accent) on fuzzel bg (base01) | `accent` #e82424 | `base01` #323137 | 2.89:1 |
| ❌ | border (accent) vs GTK-4 window_bg behind | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ✅ | text (base05) on fuzzel bg, GTK-4 window behind | `base05` #dcd7ba | `base01` #323137 | 8.88:1 |
| ❌ | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #454445 | `base00` #1f1f28 | 1.69:1 |
| ❌ | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #454445 | `base01` #323137 | 1.33:1 |

### stack — mako-on-gtk4 (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | border (accent) vs GTK-4 window_bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ✅ | text (base05) vs GTK-4 window_bg seam | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |

### stack — openbox-on-gtk4 (1/3 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | active border (accent) on GTK-4 view_bg | `accent` #e82424 | `base00` #1f1f28 | 3.66:1 |
| ✅ | active title text on GTK-4 view_bg seam | `base05` #dcd7ba | `base00` #1f1f28 | 11.26:1 |
| ❌ | inactive border (base02) on GTK-4 view_bg | `base02` #454445 | `base00` #1f1f28 | 1.69:1 |

