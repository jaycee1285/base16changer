package targets

// Embedded templates for each target
// These are based on tinted-theming and Stylix templates

const kittyTemplate = `# Base16 {{scheme-name}}
# Scheme author: {{scheme-author}}
# Template: base16changer

background #{{base00-hex}}
foreground #{{base05-hex}}
selection_background #{{base05-hex}}
selection_foreground #{{base02-hex}}

cursor #{{base05-hex}}
cursor_text_color #{{base00-hex}}

url_color #{{base04-hex}}

active_border_color #{{base03-hex}}
inactive_border_color #{{base01-hex}}

wayland_titlebar_color #{{base00-hex}}

active_tab_background #{{base00-hex}}
active_tab_foreground #{{base05-hex}}
inactive_tab_background #{{base01-hex}}
inactive_tab_foreground #{{base04-hex}}
tab_bar_background #{{base01-hex}}

# normal
color0 #{{base00-hex}}
color1 #{{base08-hex}}
color2 #{{base0B-hex}}
color3 #{{base0A-hex}}
color4 #{{base0D-hex}}
color5 #{{base0E-hex}}
color6 #{{base0C-hex}}
color7 #{{base05-hex}}

# bright
color8 #{{base03-hex}}
color9 #{{base08-hex}}
color10 #{{base0B-hex}}
color11 #{{base0A-hex}}
color12 #{{base0D-hex}}
color13 #{{base0E-hex}}
color14 #{{base0C-hex}}
color15 #{{base07-hex}}

# extended base16
color16 #{{base09-hex}}
color17 #{{base0F-hex}}
color18 #{{base01-hex}}
color19 #{{base02-hex}}
color20 #{{base04-hex}}
color21 #{{base06-hex}}
`

const fuzzelTemplate = `# Base16 {{scheme-name}}
# Scheme author: {{scheme-author}}

[colors]
background={{base01-hex}}f2
text={{base05-hex}}ff
match={{base0D-hex}}ff
selection={{base03-hex}}ff
selection-text={{base05-hex}}ff
selection-match={{base0D-hex}}ff
border={{base0D-hex}}ff
`

// waybarTemplate intentionally empty for v1
const waybarTemplate = ``

// GTK-4 template based on Stylix's comprehensive libadwaita support
const gtk4Template = `/* Base16 {{scheme-name}} */
/* Scheme author: {{scheme-author}} */

@define-color accent_color #{{base0D-hex}};
@define-color accent_bg_color #{{base0D-hex}};
@define-color accent_fg_color #{{base00-hex}};

@define-color destructive_color #{{base08-hex}};
@define-color destructive_bg_color #{{base08-hex}};
@define-color destructive_fg_color #{{base00-hex}};

@define-color success_color #{{base0B-hex}};
@define-color success_bg_color #{{base0B-hex}};
@define-color success_fg_color #{{base00-hex}};

@define-color warning_color #{{base0A-hex}};
@define-color warning_bg_color #{{base0A-hex}};
@define-color warning_fg_color #{{base00-hex}};

@define-color error_color #{{base08-hex}};
@define-color error_bg_color #{{base08-hex}};
@define-color error_fg_color #{{base00-hex}};

@define-color window_bg_color #{{base00-hex}};
@define-color window_fg_color #{{base05-hex}};

@define-color view_bg_color #{{base00-hex}};
@define-color view_fg_color #{{base05-hex}};

@define-color headerbar_bg_color #{{base01-hex}};
@define-color headerbar_fg_color #{{base05-hex}};
@define-color headerbar_backdrop_color @window_bg_color;
@define-color headerbar_shade_color rgba(0, 0, 0, 0.07);

@define-color sidebar_bg_color #{{base01-hex}};
@define-color sidebar_fg_color #{{base05-hex}};
@define-color sidebar_backdrop_color @window_bg_color;
@define-color sidebar_shade_color rgba(0, 0, 0, 0.07);

@define-color card_bg_color #{{base01-hex}};
@define-color card_fg_color #{{base05-hex}};
@define-color card_shade_color rgba(0, 0, 0, 0.07);

@define-color dialog_bg_color #{{base01-hex}};
@define-color dialog_fg_color #{{base05-hex}};

@define-color popover_bg_color #{{base01-hex}};
@define-color popover_fg_color #{{base05-hex}};
@define-color popover_shade_color rgba(0, 0, 0, 0.07);

@define-color shade_color rgba(0, 0, 0, 0.07);
@define-color scrollbar_outline_color #{{base02-hex}};

@define-color blue_1 #{{base0D-hex}};
@define-color green_1 #{{base0B-hex}};
@define-color yellow_1 #{{base0A-hex}};
@define-color orange_1 #{{base09-hex}};
@define-color red_1 #{{base08-hex}};
@define-color purple_1 #{{base0E-hex}};
@define-color brown_1 #{{base0F-hex}};
`

// GTK-3 template (simpler, uses FlatColor-style injection)
const gtk3Template = `/* Base16 {{scheme-name}} */
/* Scheme author: {{scheme-author}} */

@define-color bg_color #{{base00-hex}};
@define-color fg_color #{{base05-hex}};
@define-color base_color #{{base01-hex}};
@define-color text_color #{{base05-hex}};
@define-color text_color_disabled #{{base03-hex}};
@define-color selected_bg_color #{{base02-hex}};
@define-color selected_fg_color #{{base05-hex}};
@define-color tooltip_bg_color #{{base00-hex}};
@define-color tooltip_fg_color #{{base05-hex}};

@define-color theme_bg_color @bg_color;
@define-color theme_fg_color @fg_color;
@define-color theme_base_color @base_color;
@define-color theme_text_color @text_color;
@define-color theme_selected_bg_color @selected_bg_color;
@define-color theme_selected_fg_color @selected_fg_color;
`

// Openbox themerc for labwc
const openboxTemplate = `# Base16 {{scheme-name}}
# Scheme author: {{scheme-author}}

# Window geometry
border.width: 1
padding.width: 4
padding.height: 4
window.handle.width: 0
window.client.padding.width: 0
window.client.padding.height: 0

# Menu geometry
menu.overlap.x: 0
menu.overlap.y: 0

# Border colors
window.active.border.color: #{{base0D-hex}}
window.inactive.border.color: #{{base02-hex}}
menu.border.color: #{{base02-hex}}

# Title bar
window.active.title.bg: flat solid
window.active.title.bg.color: #{{base01-hex}}
window.inactive.title.bg: flat solid
window.inactive.title.bg.color: #{{base00-hex}}

# Title text
window.active.label.text.color: #{{base05-hex}}
window.inactive.label.text.color: #{{base03-hex}}
window.label.text.justify: center

# Buttons
window.active.button.unpressed.bg: flat solid
window.active.button.unpressed.bg.color: #{{base01-hex}}
window.active.button.unpressed.image.color: #{{base05-hex}}

window.active.button.pressed.bg: flat solid
window.active.button.pressed.bg.color: #{{base02-hex}}
window.active.button.pressed.image.color: #{{base05-hex}}

window.active.button.hover.bg: flat solid
window.active.button.hover.bg.color: #{{base02-hex}}
window.active.button.hover.image.color: #{{base07-hex}}

window.inactive.button.unpressed.bg: flat solid
window.inactive.button.unpressed.bg.color: #{{base00-hex}}
window.inactive.button.unpressed.image.color: #{{base03-hex}}

window.inactive.button.pressed.bg: flat solid
window.inactive.button.pressed.bg.color: #{{base01-hex}}
window.inactive.button.pressed.image.color: #{{base03-hex}}

window.inactive.button.hover.bg: flat solid
window.inactive.button.hover.bg.color: #{{base01-hex}}
window.inactive.button.hover.image.color: #{{base05-hex}}

# Close button
window.active.button.close.unpressed.image.color: #{{base08-hex}}
window.active.button.close.hover.image.color: #{{base08-hex}}
window.active.button.close.pressed.image.color: #{{base08-hex}}

# Menu
menu.title.bg: flat solid
menu.title.bg.color: #{{base01-hex}}
menu.title.text.color: #{{base05-hex}}
menu.title.text.justify: center

menu.items.bg: flat solid
menu.items.bg.color: #{{base00-hex}}
menu.items.text.color: #{{base05-hex}}
menu.items.disabled.text.color: #{{base03-hex}}

menu.items.active.bg: flat solid
menu.items.active.bg.color: #{{base02-hex}}
menu.items.active.text.color: #{{base05-hex}}

# OSD (on-screen display)
osd.bg: flat solid
osd.bg.color: #{{base00-hex}}
osd.border.color: #{{base02-hex}}
osd.label.text.color: #{{base05-hex}}
`
