# WCAG AA report — Tokyo Night Terminal Storm

- **source**: `/home/john/.local/share/themes/tokyo-night-terminal-storm.yaml`
- **format**: base16
- **variant**: dark
- **accent**: #ff9e64 (explicit)
- **threshold**: AA strict (≥ 4.5:1 for all roles)
- **tmTheme**: _not found — code-in-gtk4 stack pairs will be skipped_
- **summary**: 66 pass / 85 fail / 151 total

## Palette

| Slot | Hex |
|------|-----|
| base00 | `#24283b` |
| base01 | `#1a1b26` |
| base02 | `#343a52` |
| base03 | `#444b6a` |
| base04 | `#787c99` |
| base05 | `#787c99` |
| base06 | `#cbccd1` |
| base07 | `#d5d6db` |
| base08 | `#f7768e` |
| base09 | `#ff9e64` |
| base0A | `#e0af68` |
| base0B | `#41a6b5` |
| base0C | `#7dcfff` |
| base0D | `#7aa2f7` |
| base0E | `#bb9af7` |
| base0F | `#d18616` |
| accent | `#ff9e64` |

## Failures

| Group | Target | Role | fg | bg | Ratio | Suggested swap | Notes |
|-------|--------|------|----|----|-------|----------------|-------|
| intra | kitty | default text on background | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | kitty | selection text on selection bg | `base02` #343a52 | `base05` #787c99 | **2.75:1** | — | kitty inverts: selection_fg=base02, selection_bg=base05 |
| intra | kitty | cursor on background | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | kitty | cursor_text on cursor | `base00` #24283b | `base05` #787c99 | **3.57:1** | — |  |
| intra | kitty | url_color on background | `base04` #787c99 | `base00` #24283b | **3.57:1** | `base06` (#cbccd1, 9.08:1) |  |
| intra | kitty | active border on bg | `base03` #444b6a | `base00` #24283b | **1.71:1** | `base06` (#cbccd1, 9.08:1) |  |
| intra | kitty | active tab fg on active tab bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | kitty | inactive tab fg on inactive tab bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | kitty | bright black (comments) on bg | `base03` #444b6a | `base00` #24283b | **1.71:1** | `base06` (#cbccd1, 9.08:1) |  |
| intra | fuzzel | text on background | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | fuzzel | selection-text on selection | `base05` #787c99 | `base02` #343a52 | **2.75:1** | `base07` (#d5d6db, 7.73:1) | post-T3: selection now base02 (was base03) for stripe visibility |
| intra | fuzzel | prompt/text on bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | mako | text on background | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | openbox | active title text on active title bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | openbox | inactive title text on inactive title bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) | post-T3: inactive title text now base05 (was base03) |
| intra | openbox | active button icon on button bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | openbox | inactive button icon on inactive btn bg | `base03` #444b6a | `base00` #24283b | **1.71:1** | `base06` (#cbccd1, 9.08:1) |  |
| intra | openbox | menu items on menu bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | openbox | menu disabled on menu bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | openbox | menu active item on highlight | `base05` #787c99 | `base02` #343a52 | **2.75:1** | `base07` (#d5d6db, 7.73:1) |  |
| intra | openbox | OSD label on osd bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | gtk3 | fg on bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | gtk3 | text on base (content area) | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | gtk3 | disabled text on bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | gtk3 | selected text on selected bg | `base05` #787c99 | `base02` #343a52 | **2.75:1** | `base07` (#d5d6db, 7.73:1) |  |
| intra | gtk3 | tooltip text on tooltip bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | gtk3 | border (visual UI) on bg | `base02` #343a52 | `base00` #24283b | **1.30:1** | — |  |
| intra | gtk4 | window_fg on window_bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | gtk4 | view_fg on view_bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | gtk4 | headerbar_fg on headerbar_bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | gtk4 | sidebar_fg on sidebar_bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | gtk4 | card_fg on card_bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | gtk4 | dialog_fg on dialog_bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | gtk4 | popover_fg on popover_bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | gtk4 | scrollbar outline on view_bg | `base02` #343a52 | `base00` #24283b | **1.30:1** | — |  |
| intra | gsv | text on view bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | gsv | text on current_line bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | gsv | comment on view bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | gsv | search match on match bg | `base05` #787c99 | `base02` #343a52 | **2.75:1** | `base07` (#d5d6db, 7.73:1) |  |
| intra | gsv | right_margin text on margin bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | gsv | line-numbers on gutter | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | firefox | toolbar text on toolbar bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | firefox | field (URL bar) text on field bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | firefox | popup text on popup bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | firefox | popup border on popup bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | firefox | sidebar text on sidebar bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | firefox | sidebar border on sidebar bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | firefox | muted text on toolbar bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | firefox | tab selected fg on tab selected bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | firefox | tab hover fg on tab hover bg | `base05` #787c99 | `base02` #343a52 | **2.75:1** | `base07` (#d5d6db, 7.73:1) |  |
| intra | sidebery | active tab fg on active tab bg | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| intra | sidebery | inactive tab fg on frame bg | `base04` #787c99 | `base00` #24283b | **3.57:1** | `base06` (#cbccd1, 9.08:1) | user-flagged equivalent: muted inactive tab text |
| intra | sidebery | toolbar fg on toolbar bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| intra | sidebery | tab hover fg on hover bg | `base05` #787c99 | `base02` #343a52 | **2.75:1** | `base07` (#d5d6db, 7.73:1) |  |
| intra | sidebery | popup fg on popup bg | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| stack | code-in-gtk4 | syntect HEADING on GTK-4 view_bg | `tm:heading (unresolved)` — | `base00` #24283b | **NaN:1** | — | user-flagged: heading scope foreground from .tmTheme, against GTK-4 view bg — skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect HEADING on GTK-4 popover_bg | `tm:heading (unresolved)` — | `base01` #1a1b26 | **NaN:1** | — | user-flagged: same color but rendered inside a tooltip/popover — skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect HEADING on GTK-4 sidebar_bg | `tm:heading (unresolved)` — | `base01` #1a1b26 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect HEADING on GTK-4 selection_bg | `tm:heading (unresolved)` — | `base02` #343a52 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect comment on GTK-4 view_bg | `tm:comment (unresolved)` — | `base00` #24283b | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect comment on GTK-4 popover_bg | `tm:comment (unresolved)` — | `base01` #1a1b26 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect string on GTK-4 view_bg | `tm:string (unresolved)` — | `base00` #24283b | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect string on GTK-4 popover_bg | `tm:string (unresolved)` — | `base01` #1a1b26 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect keyword on GTK-4 view_bg | `tm:keyword (unresolved)` — | `base00` #24283b | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect keyword on GTK-4 popover_bg | `tm:keyword (unresolved)` — | `base01` #1a1b26 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect function on GTK-4 view_bg | `tm:function (unresolved)` — | `base00` #24283b | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect function on GTK-4 popover_bg | `tm:function (unresolved)` — | `base01` #1a1b26 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect type on GTK-4 view_bg | `tm:type (unresolved)` — | `base00` #24283b | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect number on GTK-4 view_bg | `tm:number (unresolved)` — | `base00` #24283b | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #24283b | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect default fg on GTK-4 view_bg | `tm:__fg (unresolved)` — | `base00` #24283b | **NaN:1** | — | tmTheme global default fg — what bracket/punctuation/everything-else falls back to — skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | syntect default fg on GTK-4 popover_bg | `tm:__fg (unresolved)` — | `base01` #1a1b26 | **NaN:1** | — | skipped: fg unresolved (no tmTheme or missing scope) |
| stack | code-in-gtk4 | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg (unresolved)` — | `base00` #24283b | **NaN:1** | — | if tmTheme bg differs from GTK-4 view bg, the gutter strip seam fails — skipped: fg unresolved (no tmTheme or missing scope) |
| stack | ff-in-gtk4 | toolbar_text on GTK-4 headerbar | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| stack | ff-in-gtk4 | toolbar_field_text on toolbar_field | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| stack | ff-private | toolbar_field_text (dark scheme base05) on toolbar_field (FF forced) — control | `base05` #787c99 | `FF private toolbar_field #42414d` #42414d | **2.45:1** | `base07` (#d5d6db, 6.91:1) |  |
| stack | ff-private | toolbar_text (dark scheme base05) on chrome bg (FF forced) — control | `base05` #787c99 | `FF private chrome bg #1c1b22` #1c1b22 | **4.18:1** | `base07` (#d5d6db, 11.77:1) |  |
| stack | fuzzel-on-gtk4 | selected_fg (base05) on selected_bg (base02) | `base05` #787c99 | `base02` #343a52 | **2.75:1** | `base07` (#d5d6db, 7.73:1) | USER-FLAGGED #3a (post-T3): selection bg moved from base03 → base02 |
| stack | fuzzel-on-gtk4 | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #343a52 | `base01` #1a1b26 | **1.52:1** | — | USER-FLAGGED #3b (post-T3): stripe is now base02 vs base01 |
| stack | fuzzel-on-gtk4 | text (base05) on fuzzel bg, GTK-4 window behind | `base05` #787c99 | `base01` #1a1b26 | **4.18:1** | `base07` (#d5d6db, 11.78:1) |  |
| stack | fuzzel-on-gtk4 | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #343a52 | `base00` #24283b | **1.30:1** | — | GTK-4 analog of fuzzel #3b |
| stack | fuzzel-on-gtk4 | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #343a52 | `base01` #1a1b26 | **1.52:1** | — | GTK-3 analog of fuzzel #3b |
| stack | openbox-on-gtk4 | active title text on GTK-4 view_bg seam | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |
| stack | openbox-on-gtk4 | inactive border (base02) on GTK-4 view_bg | `base02` #343a52 | `base00` #24283b | **1.30:1** | — |  |
| stack | mako-on-gtk4 | text (base05) vs GTK-4 window_bg seam | `base05` #787c99 | `base00` #24283b | **3.57:1** | `base07` (#d5d6db, 10.04:1) |  |

## All pairs

### intra — firefox (8/17 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | toolbar text on toolbar bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | field (URL bar) text on field bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ✅ | field focus border on field bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ✅ | field highlight on field bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ✅ | field highlight text on highlight | `base00` #24283b | `accent` #ff9e64 | 7.16:1 |
| ❌ | popup text on popup bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | popup border on popup bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ✅ | selected suggestion url on hover bg | `accent` #ff9e64 | `base02` #343a52 | 5.51:1 |
| ❌ | sidebar text on sidebar bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | sidebar border on sidebar bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | muted text on toolbar bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | tab selected fg on tab selected bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ✅ | tab line (accent underline) on tab bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ❌ | tab hover fg on tab hover bg | `base05` #787c99 | `base02` #343a52 | 2.75:1 |
| ✅ | danger (.close-icon hover) on toolbar | `base08` #f7768e | `base01` #1a1b26 | 6.46:1 |
| ✅ | warning (attention) on toolbar | `base0A` #e0af68 | `base01` #1a1b26 | 8.55:1 |
| ✅ | success (download badge) on toolbar | `base0B` #41a6b5 | `base01` #1a1b26 | 5.98:1 |

### intra — fuzzel (3/6 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | text on background | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ✅ | match on background | `accent` #ff9e64 | `base01` #1a1b26 | 8.40:1 |
| ❌ | selection-text on selection | `base05` #787c99 | `base02` #343a52 | 2.75:1 |
| ✅ | selection-match on selection | `accent` #ff9e64 | `base02` #343a52 | 5.51:1 |
| ✅ | border on bg (visual indicator) | `accent` #ff9e64 | `base01` #1a1b26 | 8.40:1 |
| ❌ | prompt/text on bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |

### intra — gsv (18/24 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | text on view bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | text on current_line bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | comment on view bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ✅ | string on view bg | `base0B` #41a6b5 | `base00` #24283b | 5.10:1 |
| ✅ | string on current_line | `base0B` #41a6b5 | `base01` #1a1b26 | 5.98:1 |
| ✅ | number on view bg | `base09` #ff9e64 | `base00` #24283b | 7.16:1 |
| ✅ | function on view bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | function on current_line | `base0D` #7aa2f7 | `base01` #1a1b26 | 6.79:1 |
| ✅ | keyword on view bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | keyword on current_line | `base0E` #bb9af7 | `base01` #1a1b26 | 7.39:1 |
| ✅ | type on view bg | `base0B` #41a6b5 | `base00` #24283b | 5.10:1 |
| ✅ | builtin on view bg | `base0C` #7dcfff | `base00` #24283b | 8.49:1 |
| ✅ | constant/special on view bg | `base0A` #e0af68 | `base00` #24283b | 7.28:1 |
| ✅ | operator on view bg | `base0B` #41a6b5 | `base00` #24283b | 5.10:1 |
| ✅ | preprocessor on view bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | heading on view bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | list-marker on view bg | `base09` #ff9e64 | `base00` #24283b | 7.16:1 |
| ✅ | link-text on view bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | link-destination on view bg | `base0C` #7dcfff | `base00` #24283b | 8.49:1 |
| ❌ | search match on match bg | `base05` #787c99 | `base02` #343a52 | 2.75:1 |
| ✅ | bracket match on view bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | bracket-mismatch fg on bg | `base08` #f7768e | `base01` #1a1b26 | 6.46:1 |
| ❌ | right_margin text on margin bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | line-numbers on gutter | `base05` #787c99 | `base00` #24283b | 3.57:1 |

### intra — gtk3 (5/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | fg on bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | text on base (content area) | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | disabled text on bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | selected text on selected bg | `base05` #787c99 | `base02` #343a52 | 2.75:1 |
| ❌ | tooltip text on tooltip bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ✅ | link on bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | link on base (content) | `base0D` #7aa2f7 | `base01` #1a1b26 | 6.79:1 |
| ✅ | success on bg | `base0B` #41a6b5 | `base00` #24283b | 5.10:1 |
| ✅ | warning on bg | `base0A` #e0af68 | `base00` #24283b | 7.28:1 |
| ✅ | error on bg | `base08` #f7768e | `base00` #24283b | 5.51:1 |
| ❌ | border (visual UI) on bg | `base02` #343a52 | `base00` #24283b | 1.30:1 |

### intra — gtk4 (11/19 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | window_fg on window_bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | view_fg on view_bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | headerbar_fg on headerbar_bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | sidebar_fg on sidebar_bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | card_fg on card_bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | dialog_fg on dialog_bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | popover_fg on popover_bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ✅ | accent_color on window_bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ✅ | accent_color on view_bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ✅ | accent_color on headerbar_bg | `accent` #ff9e64 | `base01` #1a1b26 | 8.40:1 |
| ✅ | accent_color on sidebar_bg | `accent` #ff9e64 | `base01` #1a1b26 | 8.40:1 |
| ✅ | accent_color on card_bg | `accent` #ff9e64 | `base01` #1a1b26 | 8.40:1 |
| ✅ | accent_fg_color on accent_bg_color | `base00` #24283b | `accent` #ff9e64 | 7.16:1 |
| ✅ | destructive_fg on destructive_bg | `base00` #24283b | `base08` #f7768e | 5.51:1 |
| ✅ | success_fg on success_bg | `base00` #24283b | `base0B` #41a6b5 | 5.10:1 |
| ✅ | warning_fg on warning_bg | `base00` #24283b | `base0A` #e0af68 | 7.28:1 |
| ✅ | error_fg on error_bg | `base00` #24283b | `base08` #f7768e | 5.51:1 |
| ✅ | blue_3 link on view_bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ❌ | scrollbar outline on view_bg | `base02` #343a52 | `base00` #24283b | 1.30:1 |

### intra — kitty (6/15 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | default text on background | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | selection text on selection bg | `base02` #343a52 | `base05` #787c99 | 2.75:1 |
| ❌ | cursor on background | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | cursor_text on cursor | `base00` #24283b | `base05` #787c99 | 3.57:1 |
| ❌ | url_color on background | `base04` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | active border on bg | `base03` #444b6a | `base00` #24283b | 1.71:1 |
| ❌ | active tab fg on active tab bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | inactive tab fg on inactive tab bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ✅ | ANSI red on bg (color_01 logs) | `base08` #f7768e | `base00` #24283b | 5.51:1 |
| ✅ | ANSI green on bg | `base0B` #41a6b5 | `base00` #24283b | 5.10:1 |
| ✅ | ANSI yellow on bg | `base0A` #e0af68 | `base00` #24283b | 7.28:1 |
| ✅ | ANSI blue on bg | `base0D` #7aa2f7 | `base00` #24283b | 5.78:1 |
| ✅ | ANSI magenta on bg | `base0E` #bb9af7 | `base00` #24283b | 6.30:1 |
| ✅ | ANSI cyan on bg | `base0C` #7dcfff | `base00` #24283b | 8.49:1 |
| ❌ | bright black (comments) on bg | `base03` #444b6a | `base00` #24283b | 1.71:1 |

### intra — mako (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | text on background | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ✅ | border on bg | `accent` #ff9e64 | `base01` #1a1b26 | 8.40:1 |

### intra — openbox (3/11 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | active title text on active title bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | inactive title text on inactive title bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | active button icon on button bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ✅ | active button icon hover | `base07` #d5d6db | `base02` #343a52 | 7.73:1 |
| ❌ | inactive button icon on inactive btn bg | `base03` #444b6a | `base00` #24283b | 1.71:1 |
| ✅ | active border (outline) on active title | `accent` #ff9e64 | `base01` #1a1b26 | 8.40:1 |
| ❌ | menu items on menu bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | menu disabled on menu bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | menu active item on highlight | `base05` #787c99 | `base02` #343a52 | 2.75:1 |
| ❌ | OSD label on osd bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ✅ | close button (red) on active title | `base08` #f7768e | `base01` #1a1b26 | 6.46:1 |

### intra — sidebery (2/7 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | active tab fg on active tab bg | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | inactive tab fg on frame bg | `base04` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | toolbar fg on toolbar bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | tab hover fg on hover bg | `base05` #787c99 | `base02` #343a52 | 2.75:1 |
| ❌ | popup fg on popup bg | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ✅ | scroll progress (accent) on frame bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ✅ | active-tab outline (accent) on tab | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |

### stack — code-in-gtk4 (0/18 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | syntect HEADING on GTK-4 view_bg | `tm:heading (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect HEADING on GTK-4 popover_bg | `tm:heading (unresolved)` — | `base01` #1a1b26 | NaN:1 |
| ❌ | syntect HEADING on GTK-4 sidebar_bg | `tm:heading (unresolved)` — | `base01` #1a1b26 | NaN:1 |
| ❌ | syntect HEADING on GTK-4 selection_bg | `tm:heading (unresolved)` — | `base02` #343a52 | NaN:1 |
| ❌ | syntect comment on GTK-4 view_bg | `tm:comment (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect comment on GTK-4 popover_bg | `tm:comment (unresolved)` — | `base01` #1a1b26 | NaN:1 |
| ❌ | syntect string on GTK-4 view_bg | `tm:string (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect string on GTK-4 popover_bg | `tm:string (unresolved)` — | `base01` #1a1b26 | NaN:1 |
| ❌ | syntect keyword on GTK-4 view_bg | `tm:keyword (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect keyword on GTK-4 popover_bg | `tm:keyword (unresolved)` — | `base01` #1a1b26 | NaN:1 |
| ❌ | syntect function on GTK-4 view_bg | `tm:function (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect function on GTK-4 popover_bg | `tm:function (unresolved)` — | `base01` #1a1b26 | NaN:1 |
| ❌ | syntect type on GTK-4 view_bg | `tm:type (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect number on GTK-4 view_bg | `tm:number (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect link on GTK-4 view_bg | `tm:link (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect default fg on GTK-4 view_bg | `tm:__fg (unresolved)` — | `base00` #24283b | NaN:1 |
| ❌ | syntect default fg on GTK-4 popover_bg | `tm:__fg (unresolved)` — | `base01` #1a1b26 | NaN:1 |
| ❌ | tmTheme default bg leaks past GTK-4 view_bg | `tm:__bg (unresolved)` — | `base00` #24283b | NaN:1 |

### stack — ff-in-gtk4 (2/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | toolbar_text on GTK-4 headerbar | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | toolbar_field_text on toolbar_field | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ✅ | toolbar_field_border_focus on field bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ✅ | tab_line (accent) on GTK-4 view_bg below | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |

### stack — ff-private (2/4 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | accent / tab_line (scheme accent) on toolbar_field (FF forced) | `accent` #ff9e64 | `FF private toolbar_field #42414d` #42414d | 4.93:1 |
| ✅ | accent / focus_border (scheme accent) on chrome bg (FF forced) | `accent` #ff9e64 | `FF private chrome bg #1c1b22` #1c1b22 | 8.40:1 |
| ❌ | toolbar_field_text (dark scheme base05) on toolbar_field (FF forced) — control | `base05` #787c99 | `FF private toolbar_field #42414d` #42414d | 2.45:1 |
| ❌ | toolbar_text (dark scheme base05) on chrome bg (FF forced) — control | `base05` #787c99 | `FF private chrome bg #1c1b22` #1c1b22 | 4.18:1 |

### stack — fuzzel-on-gtk4 (3/8 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ❌ | selected_fg (base05) on selected_bg (base02) | `base05` #787c99 | `base02` #343a52 | 2.75:1 |
| ❌ | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | `base02` #343a52 | `base01` #1a1b26 | 1.52:1 |
| ✅ | selection_match (accent) on selected_bg (base02) | `accent` #ff9e64 | `base02` #343a52 | 5.51:1 |
| ✅ | match (accent) on fuzzel bg (base01) | `accent` #ff9e64 | `base01` #1a1b26 | 8.40:1 |
| ✅ | border (accent) vs GTK-4 window_bg behind | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ❌ | text (base05) on fuzzel bg, GTK-4 window behind | `base05` #787c99 | `base01` #1a1b26 | 4.18:1 |
| ❌ | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | `base02` #343a52 | `base00` #24283b | 1.30:1 |
| ❌ | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | `base02` #343a52 | `base01` #1a1b26 | 1.52:1 |

### stack — mako-on-gtk4 (1/2 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | border (accent) vs GTK-4 window_bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ❌ | text (base05) vs GTK-4 window_bg seam | `base05` #787c99 | `base00` #24283b | 3.57:1 |

### stack — openbox-on-gtk4 (1/3 pass)

| ✓ | Role | fg | bg | Ratio |
|---|------|----|----|-------|
| ✅ | active border (accent) on GTK-4 view_bg | `accent` #ff9e64 | `base00` #24283b | 7.16:1 |
| ❌ | active title text on GTK-4 view_bg seam | `base05` #787c99 | `base00` #24283b | 3.57:1 |
| ❌ | inactive border (base02) on GTK-4 view_bg | `base02` #343a52 | `base00` #24283b | 1.30:1 |

