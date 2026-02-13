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
@define-color headerbar_border_color rgba({{base01-dec-r}}, {{base01-dec-g}}, {{base01-dec-b}}, 0.7);
@define-color headerbar_backdrop_color @window_bg_color;
@define-color headerbar_shade_color rgba(0, 0, 0, 0.07);
@define-color headerbar_darker_shade_color rgba(0, 0, 0, 0.07);

@define-color sidebar_bg_color #{{base01-hex}};
@define-color sidebar_fg_color #{{base05-hex}};
@define-color sidebar_backdrop_color @window_bg_color;
@define-color sidebar_shade_color rgba(0, 0, 0, 0.07);

@define-color secondary_sidebar_bg_color @sidebar_bg_color;
@define-color secondary_sidebar_fg_color @sidebar_fg_color;
@define-color secondary_sidebar_backdrop_color @sidebar_backdrop_color;
@define-color secondary_sidebar_shade_color @sidebar_shade_color;

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
@define-color blue_2 #{{base0D-hex}};
@define-color blue_3 #{{base0D-hex}};
@define-color blue_4 #{{base0D-hex}};
@define-color blue_5 #{{base0D-hex}};
@define-color green_1 #{{base0B-hex}};
@define-color green_2 #{{base0B-hex}};
@define-color green_3 #{{base0B-hex}};
@define-color green_4 #{{base0B-hex}};
@define-color green_5 #{{base0B-hex}};
@define-color yellow_1 #{{base0A-hex}};
@define-color yellow_2 #{{base0A-hex}};
@define-color yellow_3 #{{base0A-hex}};
@define-color yellow_4 #{{base0A-hex}};
@define-color yellow_5 #{{base0A-hex}};
@define-color orange_1 #{{base09-hex}};
@define-color orange_2 #{{base09-hex}};
@define-color orange_3 #{{base09-hex}};
@define-color orange_4 #{{base09-hex}};
@define-color orange_5 #{{base09-hex}};
@define-color red_1 #{{base08-hex}};
@define-color red_2 #{{base08-hex}};
@define-color red_3 #{{base08-hex}};
@define-color red_4 #{{base08-hex}};
@define-color red_5 #{{base08-hex}};
@define-color purple_1 #{{base0E-hex}};
@define-color purple_2 #{{base0E-hex}};
@define-color purple_3 #{{base0E-hex}};
@define-color purple_4 #{{base0E-hex}};
@define-color purple_5 #{{base0E-hex}};
@define-color brown_1 #{{base0F-hex}};
@define-color brown_2 #{{base0F-hex}};
@define-color brown_3 #{{base0F-hex}};
@define-color brown_4 #{{base0F-hex}};
@define-color brown_5 #{{base0F-hex}};
@define-color light_1 #{{base01-hex}};
@define-color light_2 #{{base01-hex}};
@define-color light_3 #{{base01-hex}};
@define-color light_4 #{{base01-hex}};
@define-color light_5 #{{base01-hex}};
@define-color dark_1 #{{base01-hex}};
@define-color dark_2 #{{base01-hex}};
@define-color dark_3 #{{base01-hex}};
@define-color dark_4 #{{base01-hex}};
@define-color dark_5 #{{base01-hex}};
`

// GTK-3 template — base16 colors + FlatColor widget styling
const gtk3Template = `/* Base16 {{scheme-name}} */
/* Scheme author: {{scheme-author}} */

/* Base16 color scheme */
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
@define-color theme_tooltip_bg_color @tooltip_bg_color;
@define-color theme_tooltip_fg_color @tooltip_fg_color;

@define-color shadow alpha(@theme_fg_color, 0.1);
@define-color info_fg_color @fg_color;
@define-color info_bg_color @base_color;
@define-color warning_fg_color @fg_color;
@define-color warning_bg_color @base_color;
@define-color question_fg_color @fg_color;
@define-color question_bg_color @base_color;
@define-color error_fg_color @fg_color;
@define-color error_bg_color @base_color;
@define-color link_color #{{base0D-hex}};
@define-color success_color #{{base0B-hex}};
@define-color warning_color #{{base0A-hex}};
@define-color error_color #{{base08-hex}};

@define-color border_color #{{base02-hex}};
@define-color button_normal_color @base_color;
@define-color button_default_active_color shade(@theme_selected_bg_color, 0.857);
@define-color entry_border_color shade(@theme_base_color, 0.9);
@define-color sel_color @selected_bg_color;
@define-color switch_bg_color @base_color;
@define-color panel_bg_color @bg_color;
@define-color panel_fg_color @fg_color;
@define-color borders @border_color;
@define-color scrollbar_trough shade(@theme_base_color, 0.9);
@define-color scrollbar_slider_prelight mix(@scrollbar_trough, @theme_fg_color, 0.5);

@define-color osd_separator #{{base02-hex}};
@define-color osd_fg @fg_color;
@define-color osd_bg @bg_color;

@define-color wm_bg @theme_bg_color;
@define-color wm_title_focused @theme_fg_color;
@define-color wm_title_unfocused @theme_text_color;
@define-color wm_border_focused @border_color;
@define-color wm_border_unfocused @border_color;

/* FlatColor widget styling */
@import url("../../FlatColor/gtk-3.0/gtk-widgets.css");
@import url("../../FlatColor/gtk-3.0/gtk-widgets-assets.css");
@import url("../../FlatColor/gtk-3.0/widgets/button.css");
@import url("../../FlatColor/gtk-3.0/widgets/cell-row.css");
@import url("../../FlatColor/gtk-3.0/widgets/check-radio.css");
@import url("../../FlatColor/gtk-3.0/widgets/column-header.css");
@import url("../../FlatColor/gtk-3.0/widgets/calendar.css");
@import url("../../FlatColor/gtk-3.0/widgets/entry.css");
@import url("../../FlatColor/gtk-3.0/widgets/infobar.css");
@import url("../../FlatColor/gtk-3.0/widgets/menu.css");
@import url("../../FlatColor/gtk-3.0/widgets/notebook.css");
@import url("../../FlatColor/gtk-3.0/widgets/progress-scale.css");
@import url("../../FlatColor/gtk-3.0/widgets/scrollbar.css");
@import url("../../FlatColor/gtk-3.0/widgets/separator.css");
@import url("../../FlatColor/gtk-3.0/widgets/sidebar.css");
@import url("../../FlatColor/gtk-3.0/widgets/spinbutton.css");
@import url("../../FlatColor/gtk-3.0/widgets/spinner.css");
@import url("../../FlatColor/gtk-3.0/widgets/switch.css");
@import url("../../FlatColor/gtk-3.0/widgets/color-chooser.css");
@import url("../../FlatColor/gtk-3.0/widgets/toolbar.css");
@import url("../../FlatColor/gtk-3.0/widgets/header-bar.css");
@import url("../../FlatColor/gtk-3.0/widgets/osd.css");
@import url("../../FlatColor/gtk-3.0/widgets/csd.css");
@import url("../../FlatColor/gtk-3.0/widgets/combobox.css");
@import url("../../FlatColor/gtk-3.0/widgets/selection-mode.css");
`

// GTK-2 template — base16 color scheme + FlatColor widget styling
const gtk2Template = `# Base16 {{scheme-name}}
# Scheme author: {{scheme-author}}

gtk-color-scheme = "bg_color:#{{base00-hex}}
color0:#{{base00-hex}}
text_color:#{{base05-hex}}
selected_bg_color:#{{base02-hex}}
selected_fg_color:#{{base05-hex}}
tooltip_bg_color:#{{base00-hex}}
tooltip_fg_color:#{{base05-hex}}
titlebar_bg_color:#{{base01-hex}}
titlebar_fg_color:#{{base05-hex}}
menu_bg_color:#{{base01-hex}}
menu_fg_color:#{{base05-hex}}
link_color:#{{base0D-hex}}"

include "../../FlatColor/gtk-2.0/gtkrc"
`

// index.theme for ~/.themes/Base16/
const indexThemeTemplate = `[Desktop Entry]
Type=X-GNOME-Metatheme
Name=Base16
Comment=Base16 color scheme
Encoding=UTF-8

[X-GNOME-Metatheme]
GtkTheme=Base16
MetacityTheme=Base16
IconTheme=Adwaita
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
