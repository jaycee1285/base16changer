# WCAG AA report — Catppuccin Latte

- **source**: `/home/john/.local/share/themes/catppuccin-latte.yaml`
- **format**: base16
- **variant**: light
- **accent**: #fe640b (explicit)
- **threshold**: AA strict (≥ 4.5:1 for all roles)
- **tmTheme**: `/home/john/.local/share/themes/tmThemes/catppuccin-latte.tmTheme` (bg=#eff1f5, fg=#4c4f69)
- **summary**: 81 pass / 75 fail / 156 total

## Palette

| Slot | Hex |
|------|-----|
| base00 | `#eff1f5` |
| base01 | `#e6e9ef` |
| base02 | `#ccd0da` |
| base03 | `#bcc0cc` |
| base04 | `#acb0be` |
| base05 | `#4c4f69` |
| base06 | `#dc8a78` |
| base07 | `#7287fd` |
| base08 | `#d20f39` |
| base09 | `#fe640b` |
| base0A | `#df8e1d` |
| base0B | `#40a02b` |
| base0C | `#179299` |
| base0D | `#1e66f5` |
| base0E | `#8839ef` |
| base0F | `#dd7878` |
| accent | `#fe640b` |

## Failures

| Group | Target | Role | fg | bg | Ratio | Suggested swap | Notes |
|-------|--------|------|----|----|-------|----------------|-------|
| intra | kitty | url_color on background | `base04` #acb0be | `base00` #eff1f5 | **1.91:1** | `base05` (#4c4f69, 7.06:1) |  |
| intra | kitty | active border on bg | `base03` #bcc0cc | `base00` #eff1f5 | **1.61:1** | `base05` (#4c4f69, 7.06:1) |  |
| intra | kitty | ANSI green on bg | `base0B` #40a02b | `base00` #eff1f5 | **2.96:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | kitty | ANSI yellow on bg | `base0A` #df8e1d | `base00` #eff1f5 | **2.31:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | kitty | ANSI blue on bg | `base0D` #1e66f5 | `base00` #eff1f5 | **4.34:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | kitty | ANSI cyan on bg | `base0C` #179299 | `base00` #eff1f5 | **3.31:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | kitty | bright black (comments) on bg | `base03` #bcc0cc | `base00` #eff1f5 | **1.61:1** | `base05` (#4c4f69, 7.06:1) |  |
| intra | fuzzel | match on background | `accent` #fe640b | `base01` #e6e9ef | **2.45:1** | — | user-flagged: accent often disappears when accent==base0D and palette has unsaturated blue |
| intra | fuzzel | selection-match on selection | `accent` #fe640b | `base02` #ccd0da | **1.93:1** | — |  |
| intra | fuzzel | border on bg (visual indicator) | `accent` #fe640b | `base01` #e6e9ef | **2.45:1** | — |  |
| intra | mako | border on bg | `accent` #fe640b | `base01` #e6e9ef | **2.45:1** | — |  |
| intra | openbox | active button icon hover | `base07` #7287fd | `base02` #ccd0da | **2.06:1** | `base05` (#4c4f69, 5.17:1) |  |
| intra | openbox | inactive button icon on inactive btn bg | `base03` #bcc0cc | `base00` #eff1f5 | **1.61:1** | `base05` (#4c4f69, 7.06:1) |  |
| intra | openbox | active border (outline) on active title | `accent` #fe640b | `base01` #e6e9ef | **2.45:1** | — |  |
| intra | openbox | close button (red) on active title | `base08` #d20f39 | `base01` #e6e9ef | **4.46:1** | — |  |
| intra | gtk3 | link on bg | `base0D` #1e66f5 | `base00` #eff1f5 | **4.34:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gtk3 | link on base (content) | `base0D` #1e66f5 | `base01` #e6e9ef | **4.04:1** | — |  |
| intra | gtk3 | success on bg | `base0B` #40a02b | `base00` #eff1f5 | **2.96:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gtk3 | warning on bg | `base0A` #df8e1d | `base00` #eff1f5 | **2.31:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gtk3 | border (visual UI) on bg | `base02` #ccd0da | `base00` #eff1f5 | **1.37:1** | — |  |
| intra | gtk4 | accent_color on window_bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) | this is the pair the accent: field is meant to fix |
| intra | gtk4 | accent_color on view_bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gtk4 | accent_color on headerbar_bg | `accent` #fe640b | `base01` #e6e9ef | **2.45:1** | — |  |
| intra | gtk4 | accent_color on sidebar_bg | `accent` #fe640b | `base01` #e6e9ef | **2.45:1** | — |  |
| intra | gtk4 | accent_color on card_bg | `accent` #fe640b | `base01` #e6e9ef | **2.45:1** | — |  |
| intra | gtk4 | accent_fg_color on accent_bg_color | `base00` #eff1f5 | `accent` #fe640b | **2.64:1** | — | text rendered on solid accent button |
| intra | gtk4 | success_fg on success_bg | `base00` #eff1f5 | `base0B` #40a02b | **2.96:1** | — |  |
| intra | gtk4 | warning_fg on warning_bg | `base00` #eff1f5 | `base0A` #df8e1d | **2.31:1** | — |  |
| intra | gtk4 | blue_3 link on view_bg | `base0D` #1e66f5 | `base00` #eff1f5 | **4.34:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gtk4 | scrollbar outline on view_bg | `base02` #ccd0da | `base00` #eff1f5 | **1.37:1** | — |  |
| intra | gsv | string on view bg | `base0B` #40a02b | `base00` #eff1f5 | **2.96:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | string on current_line | `base0B` #40a02b | `base01` #e6e9ef | **2.75:1** | — |  |
| intra | gsv | number on view bg | `base09` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | function on view bg | `base0D` #1e66f5 | `base00` #eff1f5 | **4.34:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | function on current_line | `base0D` #1e66f5 | `base01` #e6e9ef | **4.04:1** | — |  |
| intra | gsv | keyword on current_line | `base0E` #8839ef | `base01` #e6e9ef | **4.45:1** | — |  |
| intra | gsv | type on view bg | `base0B` #40a02b | `base00` #eff1f5 | **2.96:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | builtin on view bg | `base0C` #179299 | `base00` #eff1f5 | **3.31:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | constant/special on view bg | `base0A` #df8e1d | `base00` #eff1f5 | **2.31:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | operator on view bg | `base0B` #40a02b | `base00` #eff1f5 | **2.96:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | heading on view bg | `base0D` #1e66f5 | `base00` #eff1f5 | **4.34:1** | `base08` (#d20f39, 4.80:1) | user-flagged: ayu-light heading reads as white on gray |
| intra | gsv | list-marker on view bg | `base09` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | link-destination on view bg | `base0C` #179299 | `base00` #eff1f5 | **3.31:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | gsv | bracket-mismatch fg on bg | `base08` #d20f39 | `base01` #e6e9ef | **4.46:1** | — |  |
| intra | firefox | field focus border on field bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | firefox | field highlight on field bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | firefox | field highlight text on highlight | `base00` #eff1f5 | `accent` #fe640b | **2.64:1** | — |  |
| intra | firefox | selected suggestion url on hover bg | `accent` #fe640b | `base02` #ccd0da | **1.93:1** | — |  |
| intra | firefox | tab line (accent underline) on tab bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | firefox | danger (.close-icon hover) on toolbar | `base08` #d20f39 | `base01` #e6e9ef | **4.46:1** | — |  |
| intra | firefox | warning (attention) on toolbar | `base0A` #df8e1d | `base01` #e6e9ef | **2.15:1** | — |  |
| intra | firefox | success (download badge) on toolbar | `base0B` #40a02b | `base01` #e6e9ef | **2.75:1** | — |  |
| intra | sidebery | inactive tab fg on frame bg | `base04` #acb0be | `base00` #eff1f5 | **1.91:1** | `base05` (#4c4f69, 7.06:1) | user-flagged equivalent: muted inactive tab text |
| intra | sidebery | scroll progress (accent) on frame bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| intra | sidebery | active-tab outline (accent) on tab | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| stack | code-in-gtk4 | syntect keyword on GTK-4 popover_bg | `tm:keyword` #8839ef | `base01` #e6e9ef | **4.45:1** | — |  |
| stack | code-in-gtk4 | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #eff1f5 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #eff1f5 | `base00` #eff1f5 | **1.00:1** | — | if tmTheme bg differs from GTK-4 view bg, the gutter strip seam fails |
| stack | ff-in-gtk4 | toolbar_field_border_focus on field bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| stack | ff-in-gtk4 | tab_line (accent) on GTK-4 view_bg below | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| stack | ff-private | toolbar_field_text (scheme base05) on toolbar_field (FF forced #42414d) | `base05` #4c4f69 | `FF private toolbar_field #42414d` #42414d | **1.26:1** | `base03` (#bcc0cc, 5.51:1) | USER-FLAGGED #1: light theme's dark text lands on Firefox's forced dark gray field bg |
| stack | ff-private | toolbar_field_text (FF forced #cfcfd8) on toolbar_field (scheme base00) | `FF private toolbar_field_text #cfcfd8` #cfcfd8 | `base00` #eff1f5 | **1.37:1** | — |  |
| stack | ff-private | toolbar_field_text (FF forced #fbfbfe chrome fg) on toolbar_field (scheme base00) | `FF private chrome fg #fbfbfe` #fbfbfe | `base00` #eff1f5 | **1.09:1** | — | inverse of #1: scheme's light field bg with FF's near-white text washes out |
| stack | ff-private | toolbar_text (scheme base05) on chrome bg (FF forced #1c1b22) | `base05` #4c4f69 | `FF private chrome bg #1c1b22` #1c1b22 | **2.14:1** | `base03` (#bcc0cc, 9.40:1) |  |
| stack | ff-private | tab_text (scheme base05) on chrome bg (FF forced #1c1b22) | `base05` #4c4f69 | `FF private chrome bg #1c1b22` #1c1b22 | **2.14:1** | `base03` (#bcc0cc, 9.40:1) |  |
| stack | ff-private | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #fe640b | `FF private toolbar_field #42414d` #42414d | **3.36:1** | — |  |
| stack | fuzzel-on-gtk4 | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #ccd0da | `base01` #e6e9ef | **1.27:1** | — | USER-FLAGGED #3b (post-T3): stripe is now base02 vs base01 |
| stack | fuzzel-on-gtk4 | selection_match (accent) on selected_bg (base02) | `accent` #fe640b | `base02` #ccd0da | **1.93:1** | — | USER-FLAGGED #3c (post-T3): typed-letter highlight on the new base02 selection |
| stack | fuzzel-on-gtk4 | match (accent) on fuzzel bg (base01) | `accent` #fe640b | `base01` #e6e9ef | **2.45:1** | — |  |
| stack | fuzzel-on-gtk4 | border (accent) vs GTK-4 window_bg behind | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| stack | fuzzel-on-gtk4 | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #ccd0da | `base00` #eff1f5 | **1.37:1** | — | GTK-4 analog of fuzzel #3b |
| stack | fuzzel-on-gtk4 | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #ccd0da | `base01` #e6e9ef | **1.27:1** | — | GTK-3 analog of fuzzel #3b |
| stack | openbox-on-gtk4 | active border (accent) on GTK-4 view_bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |
| stack | openbox-on-gtk4 | inactive border (base02) on GTK-4 view_bg | `base02` #ccd0da | `base00` #eff1f5 | **1.37:1** | — |  |
| stack | mako-on-gtk4 | border (accent) vs GTK-4 window_bg | `accent` #fe640b | `base00` #eff1f5 | **2.64:1** | `base08` (#d20f39, 4.80:1) |  |

## All pairs

### intra — firefox (9/17 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar text on toolbar bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | field (URL bar) text on field bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | field focus border on field bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ❌ | field highlight on field bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ❌ | field highlight text on highlight | `base00` #eff1f5 | `accent` #fe640b | 2.64:1 |
| ✅ | popup text on popup bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | popup border on popup bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ❌ | selected suggestion url on hover bg | `accent` #fe640b | `base02` #ccd0da | 1.93:1 |
| ✅ | sidebar text on sidebar bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | sidebar border on sidebar bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | muted text on toolbar bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | tab selected fg on tab selected bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | tab line (accent underline) on tab bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ✅ | tab hover fg on tab hover bg | `base05` #4c4f69 | `base02` #ccd0da | 5.17:1 |
| ❌ | danger (.close-icon hover) on toolbar | `base08` #d20f39 | `base01` #e6e9ef | 4.46:1 |
| ❌ | warning (attention) on toolbar | `base0A` #df8e1d | `base01` #e6e9ef | 2.15:1 |
| ❌ | success (download badge) on toolbar | `base0B` #40a02b | `base01` #e6e9ef | 2.75:1 |

### intra — fuzzel (3/6 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ❌ | match on background | `accent` #fe640b | `base01` #e6e9ef | 2.45:1 |
| ✅ | selection-text on selection | `base05` #4c4f69 | `base02` #ccd0da | 5.17:1 |
| ❌ | selection-match on selection | `accent` #fe640b | `base02` #ccd0da | 1.93:1 |
| ❌ | border on bg (visual indicator) | `accent` #fe640b | `base01` #e6e9ef | 2.45:1 |
| ✅ | prompt/text on bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |

### intra — gsv (10/24 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on view bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | text on current_line bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | comment on view bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | string on view bg | `base0B` #40a02b | `base00` #eff1f5 | 2.96:1 |
| ❌ | string on current_line | `base0B` #40a02b | `base01` #e6e9ef | 2.75:1 |
| ❌ | number on view bg | `base09` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ❌ | function on view bg | `base0D` #1e66f5 | `base00` #eff1f5 | 4.34:1 |
| ❌ | function on current_line | `base0D` #1e66f5 | `base01` #e6e9ef | 4.04:1 |
| ✅ | keyword on view bg | `base0E` #8839ef | `base00` #eff1f5 | 4.79:1 |
| ❌ | keyword on current_line | `base0E` #8839ef | `base01` #e6e9ef | 4.45:1 |
| ❌ | type on view bg | `base0B` #40a02b | `base00` #eff1f5 | 2.96:1 |
| ❌ | builtin on view bg | `base0C` #179299 | `base00` #eff1f5 | 3.31:1 |
| ❌ | constant/special on view bg | `base0A` #df8e1d | `base00` #eff1f5 | 2.31:1 |
| ❌ | operator on view bg | `base0B` #40a02b | `base00` #eff1f5 | 2.96:1 |
| ✅ | preprocessor on view bg | `base0E` #8839ef | `base00` #eff1f5 | 4.79:1 |
| ❌ | heading on view bg | `base0D` #1e66f5 | `base00` #eff1f5 | 4.34:1 |
| ❌ | list-marker on view bg | `base09` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ✅ | link-text on view bg | `base0E` #8839ef | `base00` #eff1f5 | 4.79:1 |
| ❌ | link-destination on view bg | `base0C` #179299 | `base00` #eff1f5 | 3.31:1 |
| ✅ | search match on match bg | `base05` #4c4f69 | `base02` #ccd0da | 5.17:1 |
| ✅ | bracket match on view bg | `base0E` #8839ef | `base00` #eff1f5 | 4.79:1 |
| ❌ | bracket-mismatch fg on bg | `base08` #d20f39 | `base01` #e6e9ef | 4.46:1 |
| ✅ | right_margin text on margin bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | line-numbers on gutter | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |

### intra — gtk3 (6/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | fg on bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | text on base (content area) | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | disabled text on bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | selected text on selected bg | `base05` #4c4f69 | `base02` #ccd0da | 5.17:1 |
| ✅ | tooltip text on tooltip bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | link on bg | `base0D` #1e66f5 | `base00` #eff1f5 | 4.34:1 |
| ❌ | link on base (content) | `base0D` #1e66f5 | `base01` #e6e9ef | 4.04:1 |
| ❌ | success on bg | `base0B` #40a02b | `base00` #eff1f5 | 2.96:1 |
| ❌ | warning on bg | `base0A` #df8e1d | `base00` #eff1f5 | 2.31:1 |
| ✅ | error on bg | `base08` #d20f39 | `base00` #eff1f5 | 4.80:1 |
| ❌ | border (visual UI) on bg | `base02` #ccd0da | `base00` #eff1f5 | 1.37:1 |

### intra — gtk4 (9/19 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | window_fg on window_bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | view_fg on view_bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | headerbar_fg on headerbar_bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | sidebar_fg on sidebar_bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | card_fg on card_bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | dialog_fg on dialog_bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | popover_fg on popover_bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ❌ | accent_color on window_bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ❌ | accent_color on view_bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ❌ | accent_color on headerbar_bg | `accent` #fe640b | `base01` #e6e9ef | 2.45:1 |
| ❌ | accent_color on sidebar_bg | `accent` #fe640b | `base01` #e6e9ef | 2.45:1 |
| ❌ | accent_color on card_bg | `accent` #fe640b | `base01` #e6e9ef | 2.45:1 |
| ❌ | accent_fg_color on accent_bg_color | `base00` #eff1f5 | `accent` #fe640b | 2.64:1 |
| ✅ | destructive_fg on destructive_bg | `base00` #eff1f5 | `base08` #d20f39 | 4.80:1 |
| ❌ | success_fg on success_bg | `base00` #eff1f5 | `base0B` #40a02b | 2.96:1 |
| ❌ | warning_fg on warning_bg | `base00` #eff1f5 | `base0A` #df8e1d | 2.31:1 |
| ✅ | error_fg on error_bg | `base00` #eff1f5 | `base08` #d20f39 | 4.80:1 |
| ❌ | blue_3 link on view_bg | `base0D` #1e66f5 | `base00` #eff1f5 | 4.34:1 |
| ❌ | scrollbar outline on view_bg | `base02` #ccd0da | `base00` #eff1f5 | 1.37:1 |

### intra — kitty (8/15 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | default text on background | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | selection text on selection bg | `base02` #ccd0da | `base05` #4c4f69 | 5.17:1 |
| ✅ | cursor on background | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | cursor_text on cursor | `base00` #eff1f5 | `base05` #4c4f69 | 7.06:1 |
| ❌ | url_color on background | `base04` #acb0be | `base00` #eff1f5 | 1.91:1 |
| ❌ | active border on bg | `base03` #bcc0cc | `base00` #eff1f5 | 1.61:1 |
| ✅ | active tab fg on active tab bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | inactive tab fg on inactive tab bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | ANSI red on bg (color_01 logs) | `base08` #d20f39 | `base00` #eff1f5 | 4.80:1 |
| ❌ | ANSI green on bg | `base0B` #40a02b | `base00` #eff1f5 | 2.96:1 |
| ❌ | ANSI yellow on bg | `base0A` #df8e1d | `base00` #eff1f5 | 2.31:1 |
| ❌ | ANSI blue on bg | `base0D` #1e66f5 | `base00` #eff1f5 | 4.34:1 |
| ✅ | ANSI magenta on bg | `base0E` #8839ef | `base00` #eff1f5 | 4.79:1 |
| ❌ | ANSI cyan on bg | `base0C` #179299 | `base00` #eff1f5 | 3.31:1 |
| ❌ | bright black (comments) on bg | `base03` #bcc0cc | `base00` #eff1f5 | 1.61:1 |

### intra — mako (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ❌ | border on bg | `accent` #fe640b | `base01` #e6e9ef | 2.45:1 |

### intra — openbox (7/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active title text on active title bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | inactive title text on inactive title bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | active button icon on button bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ❌ | active button icon hover | `base07` #7287fd | `base02` #ccd0da | 2.06:1 |
| ❌ | inactive button icon on inactive btn bg | `base03` #bcc0cc | `base00` #eff1f5 | 1.61:1 |
| ❌ | active border (outline) on active title | `accent` #fe640b | `base01` #e6e9ef | 2.45:1 |
| ✅ | menu items on menu bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | menu disabled on menu bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | menu active item on highlight | `base05` #4c4f69 | `base02` #ccd0da | 5.17:1 |
| ✅ | OSD label on osd bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | close button (red) on active title | `base08` #d20f39 | `base01` #e6e9ef | 4.46:1 |

### intra — sidebery (4/7 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active tab fg on active tab bg | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | inactive tab fg on frame bg | `base04` #acb0be | `base00` #eff1f5 | 1.91:1 |
| ✅ | toolbar fg on toolbar bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | tab hover fg on hover bg | `base05` #4c4f69 | `base02` #ccd0da | 5.17:1 |
| ✅ | popup fg on popup bg | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ❌ | scroll progress (accent) on frame bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ❌ | active-tab outline (accent) on tab | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |

### stack — code-in-gtk4 (15/18 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | syntect HEADING on GTK-4 view_bg | `tm:heading` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | syntect HEADING on GTK-4 popover_bg | `tm:heading` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | syntect HEADING on GTK-4 sidebar_bg | `tm:heading` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | syntect HEADING on GTK-4 selection_bg | `tm:heading` #4c4f69 | `base02` #ccd0da | 5.17:1 |
| ✅ | syntect comment on GTK-4 view_bg | `tm:comment` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | syntect comment on GTK-4 popover_bg | `tm:comment` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | syntect string on GTK-4 view_bg | `tm:string` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | syntect string on GTK-4 popover_bg | `tm:string` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | syntect keyword on GTK-4 view_bg | `tm:keyword` #8839ef | `base00` #eff1f5 | 4.79:1 |
| ❌ | syntect keyword on GTK-4 popover_bg | `tm:keyword` #8839ef | `base01` #e6e9ef | 4.45:1 |
| ✅ | syntect function on GTK-4 view_bg | `tm:function` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | syntect function on GTK-4 popover_bg | `tm:function` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | syntect type on GTK-4 view_bg | `tm:type` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | syntect number on GTK-4 view_bg | `tm:number` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #eff1f5 | NaN:1 |
| ✅ | syntect default fg on GTK-4 view_bg | `tm:__fg` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ✅ | syntect default fg on GTK-4 popover_bg | `tm:__fg` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ❌ | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #eff1f5 | `base00` #eff1f5 | 1.00:1 |

### stack — ff-in-gtk4 (2/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar_text on GTK-4 headerbar | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ✅ | toolbar_field_text on toolbar_field | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | toolbar_field_border_focus on field bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ❌ | tab_line (accent) on GTK-4 view_bg below | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |

### stack — ff-private (3/9 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | toolbar_field_text (scheme base05) on toolbar_field (FF forced #42414d) | `base05` #4c4f69 | `FF private toolbar_field #42414d` #42414d | 1.26:1 |
| ✅ | toolbar_field_text (scheme base04 muted) on toolbar_field (FF forced) | `base04` #acb0be | `FF private toolbar_field #42414d` #42414d | 4.63:1 |
| ✅ | toolbar_field (scheme base00) overrun by FF forced toolbar_field bg | `base00` #eff1f5 | `FF private toolbar_field #42414d` #42414d | 8.86:1 |
| ❌ | toolbar_field_text (FF forced #cfcfd8) on toolbar_field (scheme base00) | `FF private toolbar_field_text #cfcfd8` #cfcfd8 | `base00` #eff1f5 | 1.37:1 |
| ❌ | toolbar_field_text (FF forced #fbfbfe chrome fg) on toolbar_field (scheme base00) | `FF private chrome fg #fbfbfe` #fbfbfe | `base00` #eff1f5 | 1.09:1 |
| ❌ | toolbar_text (scheme base05) on chrome bg (FF forced #1c1b22) | `base05` #4c4f69 | `FF private chrome bg #1c1b22` #1c1b22 | 2.14:1 |
| ❌ | tab_text (scheme base05) on chrome bg (FF forced #1c1b22) | `base05` #4c4f69 | `FF private chrome bg #1c1b22` #1c1b22 | 2.14:1 |
| ❌ | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #fe640b | `FF private toolbar_field #42414d` #42414d | 3.36:1 |
| ✅ | accent / focus_border (scheme accent) on chrome bg (FF forced) | `accent` #fe640b | `FF private chrome bg #1c1b22` #1c1b22 | 5.73:1 |

### stack — fuzzel-on-gtk4 (2/8 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | selected_fg (base05) on selected_bg (base02) | `base05` #4c4f69 | `base02` #ccd0da | 5.17:1 |
| ❌ | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #ccd0da | `base01` #e6e9ef | 1.27:1 |
| ❌ | selection_match (accent) on selected_bg (base02) | `accent` #fe640b | `base02` #ccd0da | 1.93:1 |
| ❌ | match (accent) on fuzzel bg (base01) | `accent` #fe640b | `base01` #e6e9ef | 2.45:1 |
| ❌ | border (accent) vs GTK-4 window_bg behind | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ✅ | text (base05) on fuzzel bg, GTK-4 window behind | `base05` #4c4f69 | `base01` #e6e9ef | 6.57:1 |
| ❌ | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #ccd0da | `base00` #eff1f5 | 1.37:1 |
| ❌ | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #ccd0da | `base01` #e6e9ef | 1.27:1 |

### stack — mako-on-gtk4 (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | border (accent) vs GTK-4 window_bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ✅ | text (base05) vs GTK-4 window_bg seam | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |

### stack — openbox-on-gtk4 (1/3 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | active border (accent) on GTK-4 view_bg | `accent` #fe640b | `base00` #eff1f5 | 2.64:1 |
| ✅ | active title text on GTK-4 view_bg seam | `base05` #4c4f69 | `base00` #eff1f5 | 7.06:1 |
| ❌ | inactive border (base02) on GTK-4 view_bg | `base02` #ccd0da | `base00` #eff1f5 | 1.37:1 |

