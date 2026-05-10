# WCAG AA — aggregate summary

Schemes tested: 8
Threshold: AA strict (≥ 4.5:1)

Buckets are sorted by fail count. A pair that fails in many schemes is a strong
candidate for an in-`base16changer` fallback rule (run on parse, before targets).

| Fails / Tested | Group | Target | Role | Failing schemes (ratio → suggested swap) |
|---|---|---|---|---|
| 8/8 | intra | gtk3 | border (visual UI) on bg | Catppuccin Latte (1.35 → no candidate); Ayu Light (1.33 → no candidate); Tokyo Night Storm (1.62 → no candidate); Kanagawa Wave (1.69 → no candidate); Catppuccin Latte (1.37 → no candidate); Rosé Pine Dawn (1.10 → no candidate); Tokyo Night Terminal Storm (1.30 → no candidate); Everforest Dark Hard (1.54 → no candidate) |
| 8/8 | intra | gtk4 | scrollbar outline on view_bg | Catppuccin Latte (1.35 → no candidate); Ayu Light (1.33 → no candidate); Tokyo Night Storm (1.62 → no candidate); Kanagawa Wave (1.69 → no candidate); Catppuccin Latte (1.37 → no candidate); Rosé Pine Dawn (1.10 → no candidate); Tokyo Night Terminal Storm (1.30 → no candidate); Everforest Dark Hard (1.54 → no candidate) |
| 8/8 | intra | kitty | active border on bg | Catppuccin Latte (4.37 → base05@7.06); Ayu Light (3.05 → base05@6.22); Tokyo Night Storm (1.63 → base05@9.02); Kanagawa Wave (3.33 → base05@11.26); Catppuccin Latte (1.61 → base05@7.06); Rosé Pine Dawn (2.73 → base05@6.66); Tokyo Night Terminal Storm (1.71 → base06@9.08); Everforest Dark Hard (4.24 → base06@11.40) |
| 8/8 | intra | kitty | bright black (comments) on bg | Catppuccin Latte (4.37 → base05@7.06); Ayu Light (3.05 → base05@6.22); Tokyo Night Storm (1.63 → base05@9.02); Kanagawa Wave (3.33 → base05@11.26); Catppuccin Latte (1.61 → base05@7.06); Rosé Pine Dawn (2.73 → base05@6.66); Tokyo Night Terminal Storm (1.71 → base06@9.08); Everforest Dark Hard (4.24 → base06@11.40) |
| 8/8 | intra | openbox | inactive button icon on inactive btn bg | Catppuccin Latte (4.37 → base05@7.06); Ayu Light (3.05 → base05@6.22); Tokyo Night Storm (1.63 → base05@9.02); Kanagawa Wave (3.33 → base05@11.26); Catppuccin Latte (1.61 → base05@7.06); Rosé Pine Dawn (2.73 → base05@6.66); Tokyo Night Terminal Storm (1.71 → base06@9.08); Everforest Dark Hard (4.24 → base06@11.40) |
| 8/8 | stack | code-in-gtk4 | syntect link on GTK-4 view_bg | Catppuccin Latte (NaN → no candidate); Ayu Light (NaN → no candidate); Tokyo Night Storm (NaN → no candidate); Kanagawa Wave (NaN → no candidate); Catppuccin Latte (NaN → no candidate); Rosé Pine Dawn (NaN → no candidate); Tokyo Night Terminal Storm (NaN → no candidate); Everforest Dark Hard (NaN → no candidate) |
| 8/8 | stack | code-in-gtk4 | tmTheme default bg leaks past GTK-4 view_bg | Catppuccin Latte (1.00 → no candidate); Ayu Light (1.00 → no candidate); Tokyo Night Storm (1.00 → no candidate); Kanagawa Wave (1.00 → no candidate); Catppuccin Latte (1.00 → no candidate); Rosé Pine Dawn (1.00 → no candidate); Tokyo Night Terminal Storm (NaN → no candidate); Everforest Dark Hard (1.00 → no candidate) |
| 8/8 | stack | fuzzel-on-gtk4 | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | Catppuccin Latte (1.17 → no candidate); Ayu Light (1.16 → no candidate); Tokyo Night Storm (1.28 → no candidate); Kanagawa Wave (1.33 → no candidate); Catppuccin Latte (1.27 → no candidate); Rosé Pine Dawn (1.16 → no candidate); Tokyo Night Terminal Storm (1.52 → no candidate); Everforest Dark Hard (1.34 → no candidate) |
| 8/8 | stack | fuzzel-on-gtk4 | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | Catppuccin Latte (1.35 → no candidate); Ayu Light (1.33 → no candidate); Tokyo Night Storm (1.62 → no candidate); Kanagawa Wave (1.69 → no candidate); Catppuccin Latte (1.37 → no candidate); Rosé Pine Dawn (1.10 → no candidate); Tokyo Night Terminal Storm (1.30 → no candidate); Everforest Dark Hard (1.54 → no candidate) |
| 8/8 | stack | fuzzel-on-gtk4 | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | Catppuccin Latte (1.17 → no candidate); Ayu Light (1.16 → no candidate); Tokyo Night Storm (1.28 → no candidate); Kanagawa Wave (1.33 → no candidate); Catppuccin Latte (1.27 → no candidate); Rosé Pine Dawn (1.16 → no candidate); Tokyo Night Terminal Storm (1.52 → no candidate); Everforest Dark Hard (1.34 → no candidate) |
| 8/8 | stack | openbox-on-gtk4 | inactive border (base02) on GTK-4 view_bg | Catppuccin Latte (1.35 → no candidate); Ayu Light (1.33 → no candidate); Tokyo Night Storm (1.62 → no candidate); Kanagawa Wave (1.69 → no candidate); Catppuccin Latte (1.37 → no candidate); Rosé Pine Dawn (1.10 → no candidate); Tokyo Night Terminal Storm (1.30 → no candidate); Everforest Dark Hard (1.54 → no candidate) |
| 7/8 | intra | firefox | danger (.close-icon hover) on toolbar | Catppuccin Latte (4.15 → no candidate); Ayu Light (3.50 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.54 → base0A@5.34); Catppuccin Latte (4.46 → no candidate); Rosé Pine Dawn (4.04 → base0B@5.88); Everforest Dark Hard (4.38 → base0A@6.59) |
| 7/8 | intra | firefox | selected suggestion url on hover bg | Catppuccin Latte (3.55 → no candidate); Ayu Light (1.53 → no candidate); Tokyo Night Storm (3.40 → base0C@5.24); Kanagawa Wave (2.17 → no candidate); Catppuccin Latte (1.93 → no candidate); Rosé Pine Dawn (1.87 → base0B@5.09); Everforest Dark Hard (4.12 → base0A@4.90) |
| 7/8 | intra | fuzzel | selection-match on selection | Catppuccin Latte (3.55 → no candidate); Ayu Light (1.53 → no candidate); Tokyo Night Storm (3.40 → base0C@5.24); Kanagawa Wave (2.17 → no candidate); Catppuccin Latte (1.93 → no candidate); Rosé Pine Dawn (1.87 → base0B@5.09); Everforest Dark Hard (4.12 → base0A@4.90) |
| 7/8 | intra | gsv | bracket-mismatch fg on bg | Catppuccin Latte (4.15 → no candidate); Ayu Light (3.50 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.54 → base0A@5.34); Catppuccin Latte (4.46 → no candidate); Rosé Pine Dawn (4.04 → base0B@5.88); Everforest Dark Hard (4.38 → base0A@6.59) |
| 7/8 | intra | kitty | url_color on background | Catppuccin Latte (1.91 → base05@7.06); Ayu Light (1.83 → base05@6.22); Tokyo Night Storm (2.69 → base05@9.02); Kanagawa Wave (2.97 → base05@11.26); Catppuccin Latte (1.91 → base05@7.06); Rosé Pine Dawn (4.02 → base05@6.66); Tokyo Night Terminal Storm (3.57 → base06@9.08) |
| 7/8 | intra | openbox | close button (red) on active title | Catppuccin Latte (4.15 → no candidate); Ayu Light (3.50 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.54 → base0A@5.34); Catppuccin Latte (4.46 → no candidate); Rosé Pine Dawn (4.04 → base0B@5.88); Everforest Dark Hard (4.38 → base0A@6.59) |
| 7/8 | intra | sidebery | inactive tab fg on frame bg | Catppuccin Latte (1.91 → base05@7.06); Ayu Light (1.83 → base05@6.22); Tokyo Night Storm (2.69 → base05@9.02); Kanagawa Wave (2.97 → base05@11.26); Catppuccin Latte (1.91 → base05@7.06); Rosé Pine Dawn (4.02 → base05@6.66); Tokyo Night Terminal Storm (3.57 → base06@9.08) |
| 7/8 | stack | fuzzel-on-gtk4 | selection_match (accent) on selected_bg (base02) | Catppuccin Latte (3.55 → no candidate); Ayu Light (1.53 → no candidate); Tokyo Night Storm (3.40 → base0C@5.24); Kanagawa Wave (2.17 → no candidate); Catppuccin Latte (1.93 → no candidate); Rosé Pine Dawn (1.87 → base0B@5.09); Everforest Dark Hard (4.12 → base0A@4.90) |
| 6/8 | intra | fuzzel | border on bg (visual indicator) | Catppuccin Latte (4.15 → no candidate); Ayu Light (1.78 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.89 → base0A@5.34); Catppuccin Latte (2.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 6/8 | intra | fuzzel | match on background | Catppuccin Latte (4.15 → no candidate); Ayu Light (1.78 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.89 → base0A@5.34); Catppuccin Latte (2.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 6/8 | intra | gtk4 | accent_color on card_bg | Catppuccin Latte (4.15 → no candidate); Ayu Light (1.78 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.89 → base0A@5.34); Catppuccin Latte (2.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 6/8 | intra | gtk4 | accent_color on headerbar_bg | Catppuccin Latte (4.15 → no candidate); Ayu Light (1.78 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.89 → base0A@5.34); Catppuccin Latte (2.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 6/8 | intra | gtk4 | accent_color on sidebar_bg | Catppuccin Latte (4.15 → no candidate); Ayu Light (1.78 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.89 → base0A@5.34); Catppuccin Latte (2.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 6/8 | intra | mako | border on bg | Catppuccin Latte (4.15 → no candidate); Ayu Light (1.78 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.89 → base0A@5.34); Catppuccin Latte (2.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 6/8 | intra | openbox | active border (outline) on active title | Catppuccin Latte (4.15 → no candidate); Ayu Light (1.78 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.89 → base0A@5.34); Catppuccin Latte (2.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 6/8 | stack | fuzzel-on-gtk4 | match (accent) on fuzzel bg (base01) | Catppuccin Latte (4.15 → no candidate); Ayu Light (1.78 → no candidate); Tokyo Night Storm (4.36 → base0C@6.72); Kanagawa Wave (2.89 → base0A@5.34); Catppuccin Latte (2.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 5/8 | intra | gsv | keyword on current_line | Catppuccin Latte (2.02 → no candidate); Ayu Light (2.83 → no candidate); Kanagawa Wave (3.68 → base0A@5.34); Catppuccin Latte (4.45 → no candidate); Rosé Pine Dawn (2.16 → base0B@5.88) |
| 5/8 | stack | ff-private | accent / tab_line (scheme accent) on toolbar_field (FF forced) | Catppuccin Latte (1.85 → no candidate); Tokyo Night Storm (3.79 → base0C@5.84); Kanagawa Wave (2.25 → no candidate); Catppuccin Latte (3.36 → no candidate); Rosé Pine Dawn (4.47 → base0F@6.18) |
| 4/8 | intra | firefox | field focus border on field bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | intra | firefox | field highlight on field bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | intra | firefox | field highlight text on highlight | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → no candidate); Catppuccin Latte (2.64 → no candidate); Rosé Pine Dawn (2.05 → no candidate) |
| 4/8 | intra | firefox | success (download badge) on toolbar | Catppuccin Latte (2.56 → no candidate); Ayu Light (2.08 → no candidate); Kanagawa Wave (3.81 → base0A@5.34); Catppuccin Latte (2.75 → no candidate) |
| 4/8 | intra | firefox | tab line (accent underline) on tab bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | intra | firefox | warning (attention) on toolbar | Catppuccin Latte (2.00 → no candidate); Ayu Light (1.61 → no candidate); Catppuccin Latte (2.15 → no candidate); Rosé Pine Dawn (2.74 → base0B@5.88) |
| 4/8 | intra | gsv | builtin on view bg | Catppuccin Latte (3.31 → base08@4.80); Ayu Light (2.18 → no candidate); Catppuccin Latte (3.31 → base08@4.80); Rosé Pine Dawn (3.14 → base0B@5.59) |
| 4/8 | intra | gsv | constant/special on view bg | Catppuccin Latte (2.31 → base08@4.80); Ayu Light (1.84 → no candidate); Catppuccin Latte (2.31 → base08@4.80); Rosé Pine Dawn (2.60 → base0B@5.59) |
| 4/8 | intra | gsv | function on current_line | Catppuccin Latte (3.76 → no candidate); Ayu Light (2.44 → no candidate); Catppuccin Latte (4.04 → no candidate); Rosé Pine Dawn (3.65 → base0B@5.88) |
| 4/8 | intra | gsv | function on view bg | Catppuccin Latte (4.34 → base08@4.80); Ayu Light (2.79 → no candidate); Catppuccin Latte (4.34 → base08@4.80); Rosé Pine Dawn (3.47 → base0B@5.59) |
| 4/8 | intra | gsv | heading on view bg | Catppuccin Latte (4.34 → base08@4.80); Ayu Light (2.79 → no candidate); Catppuccin Latte (4.34 → base08@4.80); Rosé Pine Dawn (3.47 → base0B@5.59) |
| 4/8 | intra | gsv | link-destination on view bg | Catppuccin Latte (3.31 → base08@4.80); Ayu Light (2.18 → no candidate); Catppuccin Latte (3.31 → base08@4.80); Rosé Pine Dawn (3.14 → base0B@5.59) |
| 4/8 | intra | gsv | list-marker on view bg | Catppuccin Latte (3.64 → base08@4.80); Ayu Light (3.01 → no candidate); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | intra | gsv | number on view bg | Catppuccin Latte (3.64 → base08@4.80); Ayu Light (3.01 → no candidate); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | intra | gsv | string on current_line | Catppuccin Latte (2.56 → no candidate); Ayu Light (2.08 → no candidate); Kanagawa Wave (3.81 → base0A@5.34); Catppuccin Latte (2.75 → no candidate) |
| 4/8 | intra | gtk3 | link on base (content) | Catppuccin Latte (3.76 → no candidate); Ayu Light (2.44 → no candidate); Catppuccin Latte (4.04 → no candidate); Rosé Pine Dawn (3.65 → base0B@5.88) |
| 4/8 | intra | gtk3 | link on bg | Catppuccin Latte (4.34 → base08@4.80); Ayu Light (2.79 → no candidate); Catppuccin Latte (4.34 → base08@4.80); Rosé Pine Dawn (3.47 → base0B@5.59) |
| 4/8 | intra | gtk3 | warning on bg | Catppuccin Latte (2.31 → base08@4.80); Ayu Light (1.84 → no candidate); Catppuccin Latte (2.31 → base08@4.80); Rosé Pine Dawn (2.60 → base0B@5.59) |
| 4/8 | intra | gtk4 | accent_color on view_bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | intra | gtk4 | accent_color on window_bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | intra | gtk4 | accent_fg_color on accent_bg_color | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → no candidate); Catppuccin Latte (2.64 → no candidate); Rosé Pine Dawn (2.05 → no candidate) |
| 4/8 | intra | gtk4 | blue_3 link on view_bg | Catppuccin Latte (4.34 → base08@4.80); Ayu Light (2.79 → no candidate); Catppuccin Latte (4.34 → base08@4.80); Rosé Pine Dawn (3.47 → base0B@5.59) |
| 4/8 | intra | gtk4 | warning_fg on warning_bg | Catppuccin Latte (2.31 → no candidate); Ayu Light (1.84 → no candidate); Catppuccin Latte (2.31 → no candidate); Rosé Pine Dawn (2.60 → no candidate) |
| 4/8 | intra | kitty | ANSI blue on bg | Catppuccin Latte (4.34 → base08@4.80); Ayu Light (2.79 → no candidate); Catppuccin Latte (4.34 → base08@4.80); Rosé Pine Dawn (3.47 → base0B@5.59) |
| 4/8 | intra | kitty | ANSI cyan on bg | Catppuccin Latte (3.31 → base08@4.80); Ayu Light (2.18 → no candidate); Catppuccin Latte (3.31 → base08@4.80); Rosé Pine Dawn (3.14 → base0B@5.59) |
| 4/8 | intra | kitty | ANSI yellow on bg | Catppuccin Latte (2.31 → base08@4.80); Ayu Light (1.84 → no candidate); Catppuccin Latte (2.31 → base08@4.80); Rosé Pine Dawn (2.60 → base0B@5.59) |
| 4/8 | intra | openbox | active button icon hover | Catppuccin Latte (1.19 → base05@5.23); Ayu Light (1.33 → base05@4.67); Catppuccin Latte (2.06 → base05@5.17); Rosé Pine Dawn (1.35 → base05@6.05) |
| 4/8 | intra | sidebery | active-tab outline (accent) on tab | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | intra | sidebery | scroll progress (accent) on frame bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | stack | code-in-gtk4 | syntect HEADING on GTK-4 selection_bg | Tokyo Night Storm (3.57 → no candidate); Kanagawa Wave (3.52 → no candidate); Tokyo Night Terminal Storm (NaN → no candidate); Everforest Dark Hard (4.12 → no candidate) |
| 4/8 | stack | code-in-gtk4 | syntect keyword on GTK-4 popover_bg | Catppuccin Latte (4.14 → no candidate); Kanagawa Wave (3.68 → no candidate); Catppuccin Latte (4.45 → no candidate); Tokyo Night Terminal Storm (NaN → no candidate) |
| 4/8 | stack | ff-in-gtk4 | tab_line (accent) on GTK-4 view_bg below | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | stack | ff-in-gtk4 | toolbar_field_border_focus on field bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/4 | stack | ff-private | tab_text (scheme base05) on chrome bg (FF forced #1c1b22) | Catppuccin Latte (2.14 → base07@9.40); Ayu Light (2.63 → base07@16.37); Catppuccin Latte (2.14 → base03@9.40); Rosé Pine Dawn (2.35 → base07@10.54) |
| 4/4 | stack | ff-private | toolbar_field_text (FF forced #cfcfd8) on toolbar_field (scheme base00) | Catppuccin Latte (1.37 → no candidate); Ayu Light (1.48 → no candidate); Catppuccin Latte (1.37 → no candidate); Rosé Pine Dawn (1.42 → no candidate) |
| 4/4 | stack | ff-private | toolbar_field_text (FF forced #fbfbfe chrome fg) on toolbar_field (scheme base00) | Catppuccin Latte (1.09 → no candidate); Ayu Light (1.01 → no candidate); Catppuccin Latte (1.09 → no candidate); Rosé Pine Dawn (1.06 → no candidate) |
| 4/4 | stack | ff-private | toolbar_field_text (scheme base05) on toolbar_field (FF forced #42414d) | Catppuccin Latte (1.26 → base07@5.51); Ayu Light (1.54 → base07@9.60); Catppuccin Latte (1.26 → base03@5.51); Rosé Pine Dawn (1.38 → base07@6.18) |
| 4/4 | stack | ff-private | toolbar_text (scheme base05) on chrome bg (FF forced #1c1b22) | Catppuccin Latte (2.14 → base07@9.40); Ayu Light (2.63 → base07@16.37); Catppuccin Latte (2.14 → base03@9.40); Rosé Pine Dawn (2.35 → base07@10.54) |
| 4/8 | stack | fuzzel-on-gtk4 | border (accent) vs GTK-4 window_bg behind | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | stack | mako-on-gtk4 | border (accent) vs GTK-4 window_bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 4/8 | stack | openbox-on-gtk4 | active border (accent) on GTK-4 view_bg | Ayu Light (2.04 → no candidate); Kanagawa Wave (3.66 → base0A@6.78); Catppuccin Latte (2.64 → base08@4.80); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 3/8 | intra | gsv | bracket match on view bg | Catppuccin Latte (2.34 → base08@4.80); Ayu Light (3.24 → no candidate); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 3/8 | intra | gsv | keyword on view bg | Catppuccin Latte (2.34 → base08@4.80); Ayu Light (3.24 → no candidate); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 3/8 | intra | gsv | link-text on view bg | Catppuccin Latte (2.34 → base08@4.80); Ayu Light (3.24 → no candidate); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 3/8 | intra | gsv | operator on view bg | Catppuccin Latte (2.96 → base08@4.80); Ayu Light (2.38 → no candidate); Catppuccin Latte (2.96 → base08@4.80) |
| 3/8 | intra | gsv | preprocessor on view bg | Catppuccin Latte (2.34 → base08@4.80); Ayu Light (3.24 → no candidate); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 3/8 | intra | gsv | string on view bg | Catppuccin Latte (2.96 → base08@4.80); Ayu Light (2.38 → no candidate); Catppuccin Latte (2.96 → base08@4.80) |
| 3/8 | intra | gsv | type on view bg | Catppuccin Latte (2.96 → base08@4.80); Ayu Light (2.38 → no candidate); Catppuccin Latte (2.96 → base08@4.80) |
| 3/8 | intra | gtk3 | error on bg | Ayu Light (4.00 → no candidate); Kanagawa Wave (3.22 → base0A@6.78); Rosé Pine Dawn (3.84 → base0B@5.59) |
| 3/8 | intra | gtk3 | success on bg | Catppuccin Latte (2.96 → base08@4.80); Ayu Light (2.38 → no candidate); Catppuccin Latte (2.96 → base08@4.80) |
| 3/8 | intra | gtk4 | destructive_fg on destructive_bg | Ayu Light (4.00 → no candidate); Kanagawa Wave (3.22 → no candidate); Rosé Pine Dawn (3.84 → no candidate) |
| 3/8 | intra | gtk4 | error_fg on error_bg | Ayu Light (4.00 → no candidate); Kanagawa Wave (3.22 → no candidate); Rosé Pine Dawn (3.84 → no candidate) |
| 3/8 | intra | gtk4 | success_fg on success_bg | Catppuccin Latte (2.96 → no candidate); Ayu Light (2.38 → no candidate); Catppuccin Latte (2.96 → no candidate) |
| 3/8 | intra | kitty | ANSI green on bg | Catppuccin Latte (2.96 → base08@4.80); Ayu Light (2.38 → no candidate); Catppuccin Latte (2.96 → base08@4.80) |
| 3/8 | intra | kitty | ANSI magenta on bg | Catppuccin Latte (2.34 → base08@4.80); Ayu Light (3.24 → no candidate); Rosé Pine Dawn (2.05 → base0B@5.59) |
| 3/8 | intra | kitty | ANSI red on bg (color_01 logs) | Ayu Light (4.00 → no candidate); Kanagawa Wave (3.22 → base0A@6.78); Rosé Pine Dawn (3.84 → base0B@5.59) |
| 2/8 | stack | ff-private | accent / focus_border (scheme accent) on chrome bg (FF forced) | Catppuccin Latte (3.15 → base0F@6.93); Kanagawa Wave (3.83 → base0A@7.08) |
| 1/8 | intra | firefox | field (URL bar) text on field bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | firefox | muted text on toolbar bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | firefox | popup border on popup bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | firefox | popup text on popup bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | firefox | sidebar border on sidebar bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | firefox | sidebar text on sidebar bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | firefox | tab hover fg on tab hover bg | Tokyo Night Terminal Storm (2.75 → base07@7.73) |
| 1/8 | intra | firefox | tab selected fg on tab selected bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | firefox | toolbar text on toolbar bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | fuzzel | prompt/text on bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | fuzzel | selection-text on selection | Tokyo Night Terminal Storm (2.75 → base07@7.73) |
| 1/8 | intra | fuzzel | text on background | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gsv | comment on view bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | gsv | line-numbers on gutter | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | gsv | right_margin text on margin bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gsv | search match on match bg | Tokyo Night Terminal Storm (2.75 → base07@7.73) |
| 1/8 | intra | gsv | text on current_line bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gsv | text on view bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | gtk3 | disabled text on bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | gtk3 | fg on bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | gtk3 | selected text on selected bg | Tokyo Night Terminal Storm (2.75 → base07@7.73) |
| 1/8 | intra | gtk3 | text on base (content area) | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gtk3 | tooltip text on tooltip bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | gtk4 | card_fg on card_bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gtk4 | dialog_fg on dialog_bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gtk4 | headerbar_fg on headerbar_bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gtk4 | popover_fg on popover_bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gtk4 | sidebar_fg on sidebar_bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | gtk4 | view_fg on view_bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | gtk4 | window_fg on window_bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | kitty | active tab fg on active tab bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | kitty | cursor on background | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | kitty | cursor_text on cursor | Tokyo Night Terminal Storm (3.57 → no candidate) |
| 1/8 | intra | kitty | default text on background | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | kitty | inactive tab fg on inactive tab bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | kitty | selection text on selection bg | Tokyo Night Terminal Storm (2.75 → no candidate) |
| 1/8 | intra | mako | text on background | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | openbox | active button icon on button bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | openbox | active title text on active title bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | openbox | inactive title text on inactive title bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | openbox | menu active item on highlight | Tokyo Night Terminal Storm (2.75 → base07@7.73) |
| 1/8 | intra | openbox | menu disabled on menu bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | openbox | menu items on menu bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | openbox | OSD label on osd bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | sidebery | active tab fg on active tab bg | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | intra | sidebery | popup fg on popup bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | intra | sidebery | tab hover fg on hover bg | Tokyo Night Terminal Storm (2.75 → base07@7.73) |
| 1/8 | intra | sidebery | toolbar fg on toolbar bg | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | stack | code-in-gtk4 | syntect comment on GTK-4 popover_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect comment on GTK-4 view_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect default fg on GTK-4 popover_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect default fg on GTK-4 view_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect function on GTK-4 popover_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect function on GTK-4 view_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect HEADING on GTK-4 popover_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect HEADING on GTK-4 sidebar_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect HEADING on GTK-4 view_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect keyword on GTK-4 view_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect number on GTK-4 view_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect string on GTK-4 popover_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect string on GTK-4 view_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | code-in-gtk4 | syntect type on GTK-4 view_bg | Tokyo Night Terminal Storm (NaN → no candidate) |
| 1/8 | stack | ff-in-gtk4 | toolbar_field_text on toolbar_field | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | stack | ff-in-gtk4 | toolbar_text on GTK-4 headerbar | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/4 | stack | ff-private | toolbar_field_text (dark scheme base05) on toolbar_field (FF forced) — control | Tokyo Night Terminal Storm (2.45 → base07@6.91) |
| 1/4 | stack | ff-private | toolbar_field_text (scheme base04 muted) on toolbar_field (FF forced) | Rosé Pine Dawn (2.28 → no candidate) |
| 1/4 | stack | ff-private | toolbar_text (dark scheme base05) on chrome bg (FF forced) — control | Tokyo Night Terminal Storm (4.18 → base07@11.77) |
| 1/8 | stack | fuzzel-on-gtk4 | selected_fg (base05) on selected_bg (base02) | Tokyo Night Terminal Storm (2.75 → base07@7.73) |
| 1/8 | stack | fuzzel-on-gtk4 | text (base05) on fuzzel bg, GTK-4 window behind | Tokyo Night Terminal Storm (4.18 → base07@11.78) |
| 1/8 | stack | mako-on-gtk4 | text (base05) vs GTK-4 window_bg seam | Tokyo Night Terminal Storm (3.57 → base07@10.04) |
| 1/8 | stack | openbox-on-gtk4 | active title text on GTK-4 view_bg seam | Tokyo Night Terminal Storm (3.57 → base07@10.04) |

## Recommended `base16changer` parse-time fallbacks

Pairs failing in ≥ 50% of tested schemes — these are the candidates for an
auto-fallback rule in `internal/scheme/parse.go` (e.g. `ToMap()` could emit a
derived `<role>-fallback-hex` map key).

| Group | Target | Role | Most-frequent suggested swap |
|-------|--------|------|------------------------------|
| intra | gtk3 | border (visual UI) on bg | no candidate |
| intra | gtk4 | scrollbar outline on view_bg | no candidate |
| intra | kitty | active border on bg | base05@7.06 |
| intra | kitty | bright black (comments) on bg | base05@7.06 |
| intra | openbox | inactive button icon on inactive btn bg | base05@7.06 |
| stack | code-in-gtk4 | syntect link on GTK-4 view_bg | no candidate |
| stack | code-in-gtk4 | tmTheme default bg leaks past GTK-4 view_bg | no candidate |
| stack | fuzzel-on-gtk4 | GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility | no candidate |
| stack | fuzzel-on-gtk4 | GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility | no candidate |
| stack | fuzzel-on-gtk4 | selected_bg (base02) vs fuzzel bg (base01) — stripe visibility | no candidate |
| stack | openbox-on-gtk4 | inactive border (base02) on GTK-4 view_bg | no candidate |
| intra | firefox | danger (.close-icon hover) on toolbar | no candidate |
| intra | firefox | selected suggestion url on hover bg | no candidate |
| intra | fuzzel | selection-match on selection | no candidate |
| intra | gsv | bracket-mismatch fg on bg | no candidate |
| intra | kitty | url_color on background | base05@7.06 |
| intra | openbox | close button (red) on active title | no candidate |
| intra | sidebery | inactive tab fg on frame bg | base05@7.06 |
| stack | fuzzel-on-gtk4 | selection_match (accent) on selected_bg (base02) | no candidate |
| intra | fuzzel | border on bg (visual indicator) | no candidate |
| intra | fuzzel | match on background | no candidate |
| intra | gtk4 | accent_color on card_bg | no candidate |
| intra | gtk4 | accent_color on headerbar_bg | no candidate |
| intra | gtk4 | accent_color on sidebar_bg | no candidate |
| intra | mako | border on bg | no candidate |
| intra | openbox | active border (outline) on active title | no candidate |
| stack | fuzzel-on-gtk4 | match (accent) on fuzzel bg (base01) | no candidate |
| intra | gsv | keyword on current_line | no candidate |
| stack | ff-private | accent / tab_line (scheme accent) on toolbar_field (FF forced) | no candidate |
| intra | firefox | field focus border on field bg | no candidate |
| intra | firefox | field highlight on field bg | no candidate |
| intra | firefox | field highlight text on highlight | no candidate |
| intra | firefox | success (download badge) on toolbar | no candidate |
| intra | firefox | tab line (accent underline) on tab bg | no candidate |
| intra | firefox | warning (attention) on toolbar | no candidate |
| intra | gsv | builtin on view bg | base08@4.80 |
| intra | gsv | constant/special on view bg | base08@4.80 |
| intra | gsv | function on current_line | no candidate |
| intra | gsv | function on view bg | base08@4.80 |
| intra | gsv | heading on view bg | base08@4.80 |
| intra | gsv | link-destination on view bg | base08@4.80 |
| intra | gsv | list-marker on view bg | base08@4.80 |
| intra | gsv | number on view bg | base08@4.80 |
| intra | gsv | string on current_line | no candidate |
| intra | gtk3 | link on base (content) | no candidate |
| intra | gtk3 | link on bg | base08@4.80 |
| intra | gtk3 | warning on bg | base08@4.80 |
| intra | gtk4 | accent_color on view_bg | no candidate |
| intra | gtk4 | accent_color on window_bg | no candidate |
| intra | gtk4 | accent_fg_color on accent_bg_color | no candidate |
| intra | gtk4 | blue_3 link on view_bg | base08@4.80 |
| intra | gtk4 | warning_fg on warning_bg | no candidate |
| intra | kitty | ANSI blue on bg | base08@4.80 |
| intra | kitty | ANSI cyan on bg | base08@4.80 |
| intra | kitty | ANSI yellow on bg | base08@4.80 |
| intra | openbox | active button icon hover | base05@5.23 |
| intra | sidebery | active-tab outline (accent) on tab | no candidate |
| intra | sidebery | scroll progress (accent) on frame bg | no candidate |
| stack | code-in-gtk4 | syntect HEADING on GTK-4 selection_bg | no candidate |
| stack | code-in-gtk4 | syntect keyword on GTK-4 popover_bg | no candidate |
| stack | ff-in-gtk4 | tab_line (accent) on GTK-4 view_bg below | no candidate |
| stack | ff-in-gtk4 | toolbar_field_border_focus on field bg | no candidate |
| stack | ff-private | tab_text (scheme base05) on chrome bg (FF forced #1c1b22) | base07@9.40 |
| stack | ff-private | toolbar_field_text (FF forced #cfcfd8) on toolbar_field (scheme base00) | no candidate |
| stack | ff-private | toolbar_field_text (FF forced #fbfbfe chrome fg) on toolbar_field (scheme base00) | no candidate |
| stack | ff-private | toolbar_field_text (scheme base05) on toolbar_field (FF forced #42414d) | base07@5.51 |
| stack | ff-private | toolbar_text (scheme base05) on chrome bg (FF forced #1c1b22) | base07@9.40 |
| stack | fuzzel-on-gtk4 | border (accent) vs GTK-4 window_bg behind | no candidate |
| stack | mako-on-gtk4 | border (accent) vs GTK-4 window_bg | no candidate |
| stack | openbox-on-gtk4 | active border (accent) on GTK-4 view_bg | no candidate |
