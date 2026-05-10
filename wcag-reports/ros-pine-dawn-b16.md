# WCAG AA report — Rosé Pine Dawn

- **source**: `/home/john/.local/share/themes/rose-pine-dawn.yaml`
- **format**: base16
- **variant**: light
- **accent**: #ea9d34 (explicit)
- **threshold**: AA strict (≥ 4.5:1 for all roles)
- **tmTheme**: `/home/john/.local/share/themes/tmThemes/Rosé Pine Dawn.tmTheme` (bg=#faf4ed, fg=#575279)
- **summary**: 80 pass / 76 fail / 156 total

## Palette

| Slot | Hex |
|------|-----|
| base00 | `#faf4ed` |
| base01 | `#fffaf3` |
| base02 | `#f2e9de` |
| base03 | `#9893a5` |
| base04 | `#797593` |
| base05 | `#575279` |
| base06 | `#575279` |
| base07 | `#cecacd` |
| base08 | `#b4637a` |
| base09 | `#ea9d34` |
| base0A | `#d7827e` |
| base0B | `#286983` |
| base0C | `#56949f` |
| base0D | `#907aa9` |
| base0E | `#ea9d34` |
| base0F | `#cecacd` |
| accent | `#ea9d34` |

## Failures

| Group | Target | Role | fg | bg | Ratio | Suggested swap | Notes |
|-------|--------|------|----|----|-------|----------------|-------|
| intra | kitty | url_color on background | `base04` #797593 | `base00` #faf4ed | **4.02:1** | `base05` (#575279, 6.66:1) |  |
| intra | kitty | active border on bg | `base03` #9893a5 | `base00` #faf4ed | **2.73:1** | `base05` (#575279, 6.66:1) |  |
| intra | kitty | ANSI red on bg (color_01 logs) | `base08` #b4637a | `base00` #faf4ed | **3.84:1** | `base0B` (#286983, 5.59:1) |  |
| intra | kitty | ANSI yellow on bg | `base0A` #d7827e | `base00` #faf4ed | **2.60:1** | `base0B` (#286983, 5.59:1) |  |
| intra | kitty | ANSI blue on bg | `base0D` #907aa9 | `base00` #faf4ed | **3.47:1** | `base0B` (#286983, 5.59:1) |  |
| intra | kitty | ANSI magenta on bg | `base0E` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | kitty | ANSI cyan on bg | `base0C` #56949f | `base00` #faf4ed | **3.14:1** | `base0B` (#286983, 5.59:1) |  |
| intra | kitty | bright black (comments) on bg | `base03` #9893a5 | `base00` #faf4ed | **2.73:1** | `base05` (#575279, 6.66:1) |  |
| intra | fuzzel | match on background | `accent` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) | user-flagged: accent often disappears when accent==base0D and palette has unsaturated blue |
| intra | fuzzel | selection-match on selection | `accent` #ea9d34 | `base02` #f2e9de | **1.87:1** | `base0B` (#286983, 5.09:1) |  |
| intra | fuzzel | border on bg (visual indicator) | `accent` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) |  |
| intra | mako | border on bg | `accent` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) |  |
| intra | openbox | active button icon hover | `base07` #cecacd | `base02` #f2e9de | **1.35:1** | `base05` (#575279, 6.05:1) |  |
| intra | openbox | inactive button icon on inactive btn bg | `base03` #9893a5 | `base00` #faf4ed | **2.73:1** | `base05` (#575279, 6.66:1) |  |
| intra | openbox | active border (outline) on active title | `accent` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) |  |
| intra | openbox | close button (red) on active title | `base08` #b4637a | `base01` #fffaf3 | **4.04:1** | `base0B` (#286983, 5.88:1) |  |
| intra | gtk3 | link on bg | `base0D` #907aa9 | `base00` #faf4ed | **3.47:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gtk3 | link on base (content) | `base0D` #907aa9 | `base01` #fffaf3 | **3.65:1** | `base0B` (#286983, 5.88:1) |  |
| intra | gtk3 | warning on bg | `base0A` #d7827e | `base00` #faf4ed | **2.60:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gtk3 | error on bg | `base08` #b4637a | `base00` #faf4ed | **3.84:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gtk3 | border (visual UI) on bg | `base02` #f2e9de | `base00` #faf4ed | **1.10:1** | — |  |
| intra | gtk4 | accent_color on window_bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) | this is the pair the accent: field is meant to fix |
| intra | gtk4 | accent_color on view_bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gtk4 | accent_color on headerbar_bg | `accent` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) |  |
| intra | gtk4 | accent_color on sidebar_bg | `accent` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) |  |
| intra | gtk4 | accent_color on card_bg | `accent` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) |  |
| intra | gtk4 | accent_fg_color on accent_bg_color | `base00` #faf4ed | `accent` #ea9d34 | **2.05:1** | — | text rendered on solid accent button |
| intra | gtk4 | destructive_fg on destructive_bg | `base00` #faf4ed | `base08` #b4637a | **3.84:1** | — |  |
| intra | gtk4 | warning_fg on warning_bg | `base00` #faf4ed | `base0A` #d7827e | **2.60:1** | — |  |
| intra | gtk4 | error_fg on error_bg | `base00` #faf4ed | `base08` #b4637a | **3.84:1** | — |  |
| intra | gtk4 | blue_3 link on view_bg | `base0D` #907aa9 | `base00` #faf4ed | **3.47:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gtk4 | scrollbar outline on view_bg | `base02` #f2e9de | `base00` #faf4ed | **1.10:1** | — |  |
| intra | gsv | number on view bg | `base09` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | function on view bg | `base0D` #907aa9 | `base00` #faf4ed | **3.47:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | function on current_line | `base0D` #907aa9 | `base01` #fffaf3 | **3.65:1** | `base0B` (#286983, 5.88:1) |  |
| intra | gsv | keyword on view bg | `base0E` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | keyword on current_line | `base0E` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) |  |
| intra | gsv | builtin on view bg | `base0C` #56949f | `base00` #faf4ed | **3.14:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | constant/special on view bg | `base0A` #d7827e | `base00` #faf4ed | **2.60:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | preprocessor on view bg | `base0E` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | heading on view bg | `base0D` #907aa9 | `base00` #faf4ed | **3.47:1** | `base0B` (#286983, 5.59:1) | user-flagged: ayu-light heading reads as white on gray |
| intra | gsv | list-marker on view bg | `base09` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | link-text on view bg | `base0E` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | link-destination on view bg | `base0C` #56949f | `base00` #faf4ed | **3.14:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | bracket match on view bg | `base0E` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | gsv | bracket-mismatch fg on bg | `base08` #b4637a | `base01` #fffaf3 | **4.04:1** | `base0B` (#286983, 5.88:1) |  |
| intra | firefox | field focus border on field bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | firefox | field highlight on field bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | firefox | field highlight text on highlight | `base00` #faf4ed | `accent` #ea9d34 | **2.05:1** | — |  |
| intra | firefox | selected suggestion url on hover bg | `accent` #ea9d34 | `base02` #f2e9de | **1.87:1** | `base0B` (#286983, 5.09:1) |  |
| intra | firefox | tab line (accent underline) on tab bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | firefox | danger (.close-icon hover) on toolbar | `base08` #b4637a | `base01` #fffaf3 | **4.04:1** | `base0B` (#286983, 5.88:1) |  |
| intra | firefox | warning (attention) on toolbar | `base0A` #d7827e | `base01` #fffaf3 | **2.74:1** | `base0B` (#286983, 5.88:1) |  |
| intra | sidebery | inactive tab fg on frame bg | `base04` #797593 | `base00` #faf4ed | **4.02:1** | `base05` (#575279, 6.66:1) | user-flagged equivalent: muted inactive tab text |
| intra | sidebery | scroll progress (accent) on frame bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| intra | sidebery | active-tab outline (accent) on tab | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| stack | code-in-gtk4 | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #faf4ed | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #faf4ed | `base00` #faf4ed | **1.00:1** | — | if tmTheme bg differs from GTK-4 view bg, the gutter strip seam fails |
| stack | ff-in-gtk4 | toolbar_field_border_focus on field bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| stack | ff-in-gtk4 | tab_line (accent) on GTK-4 view_bg below | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| stack | ff-private | toolbar_field_text (scheme base05) on toolbar_field (FF forced #42414d) | `base05` #575279 | `FF private toolbar_field #42414d` #42414d | **1.38:1** | `base07` (#cecacd, 6.18:1) | USER-FLAGGED #1: light theme's dark text lands on Firefox's forced dark gray field bg |
| stack | ff-private | toolbar_field_text (scheme base04 muted) on toolbar_field (FF forced) | `base04` #797593 | `FF private toolbar_field #42414d` #42414d | **2.28:1** | — | muted variant of the same failure (placeholder / disabled URL bar text) |
| stack | ff-private | toolbar_field_text (FF forced #cfcfd8) on toolbar_field (scheme base00) | `FF private toolbar_field_text #cfcfd8` #cfcfd8 | `base00` #faf4ed | **1.42:1** | — |  |
| stack | ff-private | toolbar_field_text (FF forced #fbfbfe chrome fg) on toolbar_field (scheme base00) | `FF private chrome fg #fbfbfe` #fbfbfe | `base00` #faf4ed | **1.06:1** | — | inverse of #1: scheme's light field bg with FF's near-white text washes out |
| stack | ff-private | toolbar_text (scheme base05) on chrome bg (FF forced #1c1b22) | `base05` #575279 | `FF private chrome bg #1c1b22` #1c1b22 | **2.35:1** | `base07` (#cecacd, 10.54:1) |  |
| stack | ff-private | tab_text (scheme base05) on chrome bg (FF forced #1c1b22) | `base05` #575279 | `FF private chrome bg #1c1b22` #1c1b22 | **2.35:1** | `base07` (#cecacd, 10.54:1) |  |
| stack | ff-private | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #ea9d34 | `FF private toolbar_field #42414d` #42414d | **4.47:1** | `base0F` (#cecacd, 6.18:1) |  |
| stack | fuzzel-on-gtk4 | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #f2e9de | `base01` #fffaf3 | **1.16:1** | — | USER-FLAGGED #3b (post-T3): stripe is now base02 vs base01 |
| stack | fuzzel-on-gtk4 | selection_match (accent) on selected_bg (base02) | `accent` #ea9d34 | `base02` #f2e9de | **1.87:1** | `base0B` (#286983, 5.09:1) | USER-FLAGGED #3c (post-T3): typed-letter highlight on the new base02 selection |
| stack | fuzzel-on-gtk4 | match (accent) on fuzzel bg (base01) | `accent` #ea9d34 | `base01` #fffaf3 | **2.16:1** | `base0B` (#286983, 5.88:1) |  |
| stack | fuzzel-on-gtk4 | border (accent) vs GTK-4 window_bg behind | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| stack | fuzzel-on-gtk4 | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #f2e9de | `base00` #faf4ed | **1.10:1** | — | GTK-4 analog of fuzzel #3b |
| stack | fuzzel-on-gtk4 | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #f2e9de | `base01` #fffaf3 | **1.16:1** | — | GTK-3 analog of fuzzel #3b |
| stack | openbox-on-gtk4 | active border (accent) on GTK-4 view_bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |
| stack | openbox-on-gtk4 | inactive border (base02) on GTK-4 view_bg | `base02` #f2e9de | `base00` #faf4ed | **1.10:1** | — |  |
| stack | mako-on-gtk4 | border (accent) vs GTK-4 window_bg | `accent` #ea9d34 | `base00` #faf4ed | **2.05:1** | `base0B` (#286983, 5.59:1) |  |

## All pairs

### intra — firefox (10/17 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar text on toolbar bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | field (URL bar) text on field bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ❌ | field focus border on field bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | field highlight on field bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | field highlight text on highlight | `base00` #faf4ed | `accent` #ea9d34 | 2.05:1 |
| ✅ | popup text on popup bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | popup border on popup bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | selected suggestion url on hover bg | `accent` #ea9d34 | `base02` #f2e9de | 1.87:1 |
| ✅ | sidebar text on sidebar bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | sidebar border on sidebar bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | muted text on toolbar bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | tab selected fg on tab selected bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ❌ | tab line (accent underline) on tab bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ✅ | tab hover fg on tab hover bg | `base05` #575279 | `base02` #f2e9de | 6.05:1 |
| ❌ | danger (.close-icon hover) on toolbar | `base08` #b4637a | `base01` #fffaf3 | 4.04:1 |
| ❌ | warning (attention) on toolbar | `base0A` #d7827e | `base01` #fffaf3 | 2.74:1 |
| ✅ | success (download badge) on toolbar | `base0B` #286983 | `base01` #fffaf3 | 5.88:1 |

### intra — fuzzel (3/6 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | match on background | `accent` #ea9d34 | `base01` #fffaf3 | 2.16:1 |
| ✅ | selection-text on selection | `base05` #575279 | `base02` #f2e9de | 6.05:1 |
| ❌ | selection-match on selection | `accent` #ea9d34 | `base02` #f2e9de | 1.87:1 |
| ❌ | border on bg (visual indicator) | `accent` #ea9d34 | `base01` #fffaf3 | 2.16:1 |
| ✅ | prompt/text on bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |

### intra — gsv (10/24 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on view bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | text on current_line bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | comment on view bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | string on view bg | `base0B` #286983 | `base00` #faf4ed | 5.59:1 |
| ✅ | string on current_line | `base0B` #286983 | `base01` #fffaf3 | 5.88:1 |
| ❌ | number on view bg | `base09` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | function on view bg | `base0D` #907aa9 | `base00` #faf4ed | 3.47:1 |
| ❌ | function on current_line | `base0D` #907aa9 | `base01` #fffaf3 | 3.65:1 |
| ❌ | keyword on view bg | `base0E` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | keyword on current_line | `base0E` #ea9d34 | `base01` #fffaf3 | 2.16:1 |
| ✅ | type on view bg | `base0B` #286983 | `base00` #faf4ed | 5.59:1 |
| ❌ | builtin on view bg | `base0C` #56949f | `base00` #faf4ed | 3.14:1 |
| ❌ | constant/special on view bg | `base0A` #d7827e | `base00` #faf4ed | 2.60:1 |
| ✅ | operator on view bg | `base0B` #286983 | `base00` #faf4ed | 5.59:1 |
| ❌ | preprocessor on view bg | `base0E` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | heading on view bg | `base0D` #907aa9 | `base00` #faf4ed | 3.47:1 |
| ❌ | list-marker on view bg | `base09` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | link-text on view bg | `base0E` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | link-destination on view bg | `base0C` #56949f | `base00` #faf4ed | 3.14:1 |
| ✅ | search match on match bg | `base05` #575279 | `base02` #f2e9de | 6.05:1 |
| ❌ | bracket match on view bg | `base0E` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | bracket-mismatch fg on bg | `base08` #b4637a | `base01` #fffaf3 | 4.04:1 |
| ✅ | right_margin text on margin bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | line-numbers on gutter | `base05` #575279 | `base00` #faf4ed | 6.66:1 |

### intra — gtk3 (6/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | fg on bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | text on base (content area) | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | disabled text on bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | selected text on selected bg | `base05` #575279 | `base02` #f2e9de | 6.05:1 |
| ✅ | tooltip text on tooltip bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ❌ | link on bg | `base0D` #907aa9 | `base00` #faf4ed | 3.47:1 |
| ❌ | link on base (content) | `base0D` #907aa9 | `base01` #fffaf3 | 3.65:1 |
| ✅ | success on bg | `base0B` #286983 | `base00` #faf4ed | 5.59:1 |
| ❌ | warning on bg | `base0A` #d7827e | `base00` #faf4ed | 2.60:1 |
| ❌ | error on bg | `base08` #b4637a | `base00` #faf4ed | 3.84:1 |
| ❌ | border (visual UI) on bg | `base02` #f2e9de | `base00` #faf4ed | 1.10:1 |

### intra — gtk4 (8/19 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | window_fg on window_bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | view_fg on view_bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | headerbar_fg on headerbar_bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | sidebar_fg on sidebar_bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | card_fg on card_bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | dialog_fg on dialog_bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | popover_fg on popover_bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | accent_color on window_bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | accent_color on view_bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | accent_color on headerbar_bg | `accent` #ea9d34 | `base01` #fffaf3 | 2.16:1 |
| ❌ | accent_color on sidebar_bg | `accent` #ea9d34 | `base01` #fffaf3 | 2.16:1 |
| ❌ | accent_color on card_bg | `accent` #ea9d34 | `base01` #fffaf3 | 2.16:1 |
| ❌ | accent_fg_color on accent_bg_color | `base00` #faf4ed | `accent` #ea9d34 | 2.05:1 |
| ❌ | destructive_fg on destructive_bg | `base00` #faf4ed | `base08` #b4637a | 3.84:1 |
| ✅ | success_fg on success_bg | `base00` #faf4ed | `base0B` #286983 | 5.59:1 |
| ❌ | warning_fg on warning_bg | `base00` #faf4ed | `base0A` #d7827e | 2.60:1 |
| ❌ | error_fg on error_bg | `base00` #faf4ed | `base08` #b4637a | 3.84:1 |
| ❌ | blue_3 link on view_bg | `base0D` #907aa9 | `base00` #faf4ed | 3.47:1 |
| ❌ | scrollbar outline on view_bg | `base02` #f2e9de | `base00` #faf4ed | 1.10:1 |

### intra — kitty (7/15 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | default text on background | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | selection text on selection bg | `base02` #f2e9de | `base05` #575279 | 6.05:1 |
| ✅ | cursor on background | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | cursor_text on cursor | `base00` #faf4ed | `base05` #575279 | 6.66:1 |
| ❌ | url_color on background | `base04` #797593 | `base00` #faf4ed | 4.02:1 |
| ❌ | active border on bg | `base03` #9893a5 | `base00` #faf4ed | 2.73:1 |
| ✅ | active tab fg on active tab bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | inactive tab fg on inactive tab bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | ANSI red on bg (color_01 logs) | `base08` #b4637a | `base00` #faf4ed | 3.84:1 |
| ✅ | ANSI green on bg | `base0B` #286983 | `base00` #faf4ed | 5.59:1 |
| ❌ | ANSI yellow on bg | `base0A` #d7827e | `base00` #faf4ed | 2.60:1 |
| ❌ | ANSI blue on bg | `base0D` #907aa9 | `base00` #faf4ed | 3.47:1 |
| ❌ | ANSI magenta on bg | `base0E` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | ANSI cyan on bg | `base0C` #56949f | `base00` #faf4ed | 3.14:1 |
| ❌ | bright black (comments) on bg | `base03` #9893a5 | `base00` #faf4ed | 2.73:1 |

### intra — mako (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | text on background | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | border on bg | `accent` #ea9d34 | `base01` #fffaf3 | 2.16:1 |

### intra — openbox (7/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active title text on active title bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | inactive title text on inactive title bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | active button icon on button bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | active button icon hover | `base07` #cecacd | `base02` #f2e9de | 1.35:1 |
| ❌ | inactive button icon on inactive btn bg | `base03` #9893a5 | `base00` #faf4ed | 2.73:1 |
| ❌ | active border (outline) on active title | `accent` #ea9d34 | `base01` #fffaf3 | 2.16:1 |
| ✅ | menu items on menu bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | menu disabled on menu bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | menu active item on highlight | `base05` #575279 | `base02` #f2e9de | 6.05:1 |
| ✅ | OSD label on osd bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ❌ | close button (red) on active title | `base08` #b4637a | `base01` #fffaf3 | 4.04:1 |

### intra — sidebery (4/7 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active tab fg on active tab bg | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ❌ | inactive tab fg on frame bg | `base04` #797593 | `base00` #faf4ed | 4.02:1 |
| ✅ | toolbar fg on toolbar bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | tab hover fg on hover bg | `base05` #575279 | `base02` #f2e9de | 6.05:1 |
| ✅ | popup fg on popup bg | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | scroll progress (accent) on frame bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | active-tab outline (accent) on tab | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |

### stack — code-in-gtk4 (16/18 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | syntect HEADING on GTK-4 view_bg | `tm:heading` #286983 | `base00` #faf4ed | 5.59:1 |
| ✅ | syntect HEADING on GTK-4 popover_bg | `tm:heading` #286983 | `base01` #fffaf3 | 5.88:1 |
| ✅ | syntect HEADING on GTK-4 sidebar_bg | `tm:heading` #286983 | `base01` #fffaf3 | 5.88:1 |
| ✅ | syntect HEADING on GTK-4 selection_bg | `tm:heading` #286983 | `base02` #f2e9de | 5.09:1 |
| ✅ | syntect comment on GTK-4 view_bg | `tm:comment` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | syntect comment on GTK-4 popover_bg | `tm:comment` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | syntect string on GTK-4 view_bg | `tm:string` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | syntect string on GTK-4 popover_bg | `tm:string` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | syntect keyword on GTK-4 view_bg | `tm:keyword` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | syntect keyword on GTK-4 popover_bg | `tm:keyword` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | syntect function on GTK-4 view_bg | `tm:function` #286983 | `base00` #faf4ed | 5.59:1 |
| ✅ | syntect function on GTK-4 popover_bg | `tm:function` #286983 | `base01` #fffaf3 | 5.88:1 |
| ✅ | syntect type on GTK-4 view_bg | `tm:type` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | syntect number on GTK-4 view_bg | `tm:number` #575279 | `base00` #faf4ed | 6.66:1 |
| ❌ | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #faf4ed | NaN:1 |
| ✅ | syntect default fg on GTK-4 view_bg | `tm:__fg` #575279 | `base00` #faf4ed | 6.66:1 |
| ✅ | syntect default fg on GTK-4 popover_bg | `tm:__fg` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg` #faf4ed | `base00` #faf4ed | 1.00:1 |

### stack — ff-in-gtk4 (2/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | toolbar_text on GTK-4 headerbar | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ✅ | toolbar_field_text on toolbar_field | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ❌ | toolbar_field_border_focus on field bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ❌ | tab_line (accent) on GTK-4 view_bg below | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |

### stack — ff-private (2/9 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | toolbar_field_text (scheme base05) on toolbar_field (FF forced #42414d) | `base05` #575279 | `FF private toolbar_field #42414d` #42414d | 1.38:1 |
| ❌ | toolbar_field_text (scheme base04 muted) on toolbar_field (FF forced) | `base04` #797593 | `FF private toolbar_field #42414d` #42414d | 2.28:1 |
| ✅ | toolbar_field (scheme base00) overrun by FF forced toolbar_field bg | `base00` #faf4ed | `FF private toolbar_field #42414d` #42414d | 9.18:1 |
| ❌ | toolbar_field_text (FF forced #cfcfd8) on toolbar_field (scheme base00) | `FF private toolbar_field_text #cfcfd8` #cfcfd8 | `base00` #faf4ed | 1.42:1 |
| ❌ | toolbar_field_text (FF forced #fbfbfe chrome fg) on toolbar_field (scheme base00) | `FF private chrome fg #fbfbfe` #fbfbfe | `base00` #faf4ed | 1.06:1 |
| ❌ | toolbar_text (scheme base05) on chrome bg (FF forced #1c1b22) | `base05` #575279 | `FF private chrome bg #1c1b22` #1c1b22 | 2.35:1 |
| ❌ | tab_text (scheme base05) on chrome bg (FF forced #1c1b22) | `base05` #575279 | `FF private chrome bg #1c1b22` #1c1b22 | 2.35:1 |
| ❌ | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #ea9d34 | `FF private toolbar_field #42414d` #42414d | 4.47:1 |
| ✅ | accent / focus_border (scheme accent) on chrome bg (FF forced) | `accent` #ea9d34 | `FF private chrome bg #1c1b22` #1c1b22 | 7.62:1 |

### stack — fuzzel-on-gtk4 (2/8 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | selected_fg (base05) on selected_bg (base02) | `base05` #575279 | `base02` #f2e9de | 6.05:1 |
| ❌ | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #f2e9de | `base01` #fffaf3 | 1.16:1 |
| ❌ | selection_match (accent) on selected_bg (base02) | `accent` #ea9d34 | `base02` #f2e9de | 1.87:1 |
| ❌ | match (accent) on fuzzel bg (base01) | `accent` #ea9d34 | `base01` #fffaf3 | 2.16:1 |
| ❌ | border (accent) vs GTK-4 window_bg behind | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ✅ | text (base05) on fuzzel bg, GTK-4 window behind | `base05` #575279 | `base01` #fffaf3 | 7.00:1 |
| ❌ | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #f2e9de | `base00` #faf4ed | 1.10:1 |
| ❌ | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #f2e9de | `base01` #fffaf3 | 1.16:1 |

### stack — mako-on-gtk4 (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | border (accent) vs GTK-4 window_bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ✅ | text (base05) vs GTK-4 window_bg seam | `base05` #575279 | `base00` #faf4ed | 6.66:1 |

### stack — openbox-on-gtk4 (1/3 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | active border (accent) on GTK-4 view_bg | `accent` #ea9d34 | `base00` #faf4ed | 2.05:1 |
| ✅ | active title text on GTK-4 view_bg seam | `base05` #575279 | `base00` #faf4ed | 6.66:1 |
| ❌ | inactive border (base02) on GTK-4 view_bg | `base02` #f2e9de | `base00` #faf4ed | 1.10:1 |

