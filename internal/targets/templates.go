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
match={{accent-hex}}ff
selection={{base03-hex}}ff
selection-text={{base05-hex}}ff
selection-match={{accent-hex}}ff
border={{accent-hex}}ff
`

// waybarTemplate intentionally empty for v1
const waybarTemplate = ``

// GTK-4 template based on Stylix's comprehensive libadwaita support
const gtk4Template = `/* Base16 {{scheme-name}} */
/* Scheme author: {{scheme-author}} */

@define-color accent_color #{{accent-hex}};
@define-color accent_bg_color #{{accent-hex}};
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

// Mako notification daemon config
const makoTemplate = `# Base16 {{scheme-name}}
# Scheme author: {{scheme-author}}

background-color=#{{base01-hex}}FF
text-color=#{{base05-hex}}FF
border-color=#{{accent-hex}}FF
`

// LibreWolf/Firefox colors.css — base16 chrome theming
const librewolfTemplate = `/* Generated by base16changer. Do not edit by hand.
   Purpose: imported from main userChrome.css for Firefox/LibreWolf chrome theming.
   Scope: colors and light UI surface overrides only.
*/

:root {
  /* Base16-derived internal tokens */
  --mm-bg: #{{base00-hex}};
  --mm-fg: #{{base05-hex}};
  --mm-surface: #{{base01-hex}};
  --mm-surface-alt: #{{base02-hex}};
  --mm-surface-hover: #{{base06-hex}};
  --mm-border: #{{base03-hex}};
  --mm-accent: #{{accent-hex}};
  --mm-accent-2: #{{base0E-hex}};
  --mm-muted: #{{base04-hex}};
  --mm-danger: #{{base08-hex}};
  --mm-warning: #{{base0A-hex}};
  --mm-success: #{{base0B-hex}};

  --mm-popup-bg: #{{base01-hex}};
  --mm-popup-fg: #{{base05-hex}};
  --mm-popup-border: #{{base03-hex}};

  --mm-field-bg: #{{base00-hex}};
  --mm-field-fg: #{{base05-hex}};
  --mm-field-border: #{{base03-hex}};
  --mm-field-focus: #{{accent-hex}};

  --mm-sidebar-bg: #{{base00-hex}};
  --mm-sidebar-fg: #{{base05-hex}};
  --mm-sidebar-border: #{{base03-hex}};

  --mm-tab-selected-bg: #{{base00-hex}};
  --mm-tab-selected-fg: #{{base05-hex}};
  --mm-tab-hover-bg: #{{base02-hex}};
  --mm-tab-hover-fg: #{{base05-hex}};
  --mm-tab-line: #{{accent-hex}};

  --mm-toolbar-bg: #{{base01-hex}};
  --mm-toolbar-fg: #{{base05-hex}};
  --mm-toolbar-border: #{{base03-hex}};

  /* Firefox-facing variables */
  --lwt-accent-color: var(--mm-bg) !important;
  --lwt-text-color: var(--mm-fg) !important;

  --toolbar-bgcolor: var(--mm-toolbar-bg) !important;
  --toolbar-color: var(--mm-toolbar-fg) !important;
  --toolbar-field-background-color: var(--mm-field-bg) !important;
  --toolbar-field-color: var(--mm-field-fg) !important;
  --toolbar-field-border-color: var(--mm-field-border) !important;
  --toolbar-field-focus-border-color: var(--mm-field-focus) !important;

  --arrowpanel-background: var(--mm-popup-bg) !important;
  --arrowpanel-color: var(--mm-popup-fg) !important;
  --arrowpanel-border-color: var(--mm-popup-border) !important;

  --sidebar-background-color: var(--mm-sidebar-bg) !important;
  --sidebar-text-color: var(--mm-sidebar-fg) !important;
  --sidebar-border-color: var(--mm-sidebar-border) !important;

  --tabs-navbar-separator-color: var(--mm-toolbar-border) !important;
  --chrome-content-separator-color: var(--mm-toolbar-border) !important;
}

/* Main top chrome */
#navigator-toolbox {
  background: var(--mm-toolbar-bg) !important;
  color: var(--mm-toolbar-fg) !important;
  border-color: var(--mm-toolbar-border) !important;
}

#nav-bar,
#TabsToolbar,
#PersonalToolbar,
#toolbar-menubar {
  background: var(--mm-toolbar-bg) !important;
  color: var(--mm-toolbar-fg) !important;
  border-color: var(--mm-toolbar-border) !important;
}

/* Toolbar buttons */
.toolbarbutton-1,
#PanelUI-menu-button,
#nav-bar toolbarbutton {
  color: var(--mm-toolbar-fg) !important;
}

.toolbarbutton-1:hover,
#PanelUI-menu-button:hover,
#nav-bar toolbarbutton:hover {
  background: var(--mm-surface-hover) !important;
}

.toolbarbutton-1[open="true"],
.toolbarbutton-1[checked="true"] {
  background: var(--mm-surface-alt) !important;
}

/* Tabs */
.tabbrowser-tab .tab-background {
  background: var(--mm-surface-alt) !important;
  border-color: transparent !important;
}

.tabbrowser-tab:hover .tab-background {
  background: var(--mm-tab-hover-bg) !important;
}

.tabbrowser-tab .tab-label {
  color: var(--mm-fg) !important;
}

.tabbrowser-tab[selected="true"] .tab-background {
  background: var(--mm-tab-selected-bg) !important;
  border-color: var(--mm-border) !important;
}

.tabbrowser-tab[selected="true"] .tab-label {
  color: var(--mm-tab-selected-fg) !important;
}

.tabbrowser-tab[usercontextid] .tab-context-line {
  background-color: var(--mm-tab-line) !important;
}

/* URL bar and other input fields */
#urlbar-background,
#searchbar,
.findbar-textbox,
#searchTextbox {
  background: var(--mm-field-bg) !important;
  color: var(--mm-field-fg) !important;
  border-color: var(--mm-field-border) !important;
}

#urlbar-input,
#searchbar input,
.findbar-textbox {
  color: var(--mm-field-fg) !important;
}

#urlbar[open] #urlbar-background,
#urlbar[focused="true"] #urlbar-background,
#searchbar:focus-within,
.findbar-textbox:focus {
  border-color: var(--mm-field-focus) !important;
  box-shadow: none !important;
}

/* URL bar suggestions/results */
.urlbarView {
  background: var(--mm-popup-bg) !important;
  color: var(--mm-popup-fg) !important;
  border-color: var(--mm-popup-border) !important;
}

.urlbarView-row {
  color: var(--mm-popup-fg) !important;
}

.urlbarView-row[selected] {
  background: var(--mm-surface-hover) !important;
}

.urlbarView-url {
  color: var(--mm-accent) !important;
}

/* Popups, menus, panels */
menupopup,
panel,
.panel-arrowcontent,
.panel-subview-body,
panelview {
  background: var(--mm-popup-bg) !important;
  color: var(--mm-popup-fg) !important;
  border-color: var(--mm-popup-border) !important;
}

menu,
menuitem,
menucaption {
  color: var(--mm-popup-fg) !important;
}

menu:hover,
menuitem:hover,
menu[_moz-menuactive="true"],
menuitem[_moz-menuactive="true"] {
  background: var(--mm-surface-hover) !important;
  color: var(--mm-popup-fg) !important;
}

/* Sidebar */
#sidebar-box,
#sidebar,
#sidebar-header,
#sidebar-main,
#webext-panels-browser {
  background: var(--mm-sidebar-bg) !important;
  color: var(--mm-sidebar-fg) !important;
  border-color: var(--mm-sidebar-border) !important;
}

/* Bookmarks/history/library panels often inherit weirdly */
#bookmarksPanel,
#PlacesToolbar,
#history-panel,
#placeContent {
  background: var(--mm-popup-bg) !important;
  color: var(--mm-popup-fg) !important;
}

/* Find bar */
findbar {
  background: var(--mm-toolbar-bg) !important;
  color: var(--mm-toolbar-fg) !important;
  border-color: var(--mm-toolbar-border) !important;
}

/* Notification bars / infobars */
.notificationbox-stack notification,
notification-message {
  background: var(--mm-surface) !important;
  color: var(--mm-fg) !important;
  border-color: var(--mm-border) !important;
}

/* Downloads / panels / subviews */
#appMenu-popup,
#widget-overflow,
#customizationui-widget-panel,
#downloadsPanel {
  --panel-background: var(--mm-popup-bg) !important;
  --panel-color: var(--mm-popup-fg) !important;
  --panel-border-color: var(--mm-popup-border) !important;
}

/* Separators */
menuseparator,
toolbarseparator,
.separator {
  border-color: var(--mm-border) !important;
}

/* Disabled or secondary text */
.description,
.toolbarbutton-text,
#statuspanel-label,
.identity-color-label {
  color: var(--mm-muted) !important;
}

/* Optional semantic accents */
.findbar-find-next,
.findbar-find-previous {
  color: var(--mm-accent) !important;
}

toolbarbutton[attention="true"] {
  color: var(--mm-warning) !important;
}

toolbarbutton[badge-status="downloading"] {
  color: var(--mm-success) !important;
}

/* Close buttons / destructive UI can inherit danger color if desired */
.close-icon:hover,
toolbarbutton.close-icon:hover {
  color: var(--mm-danger) !important;
}
`

// Firefox theme payload for Firefox Color/theme API consumers.
const firefoxThemeTemplate = `{
  "title": "Base16 {{scheme-name}}",
  "colors": {
    "frame": {"r": {{base00-dec-r}}, "g": {{base00-dec-g}}, "b": {{base00-dec-b}}},
    "frame_inactive": {"r": {{base01-dec-r}}, "g": {{base01-dec-g}}, "b": {{base01-dec-b}}, "a": 1},
    "tab_background_text": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}},
    "toolbar": {"r": {{base01-dec-r}}, "g": {{base01-dec-g}}, "b": {{base01-dec-b}}, "a": 1},
    "toolbar_text": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}, "a": 1},
    "toolbar_top_separator": {"r": {{base03-dec-r}}, "g": {{base03-dec-g}}, "b": {{base03-dec-b}}, "a": 1},
    "toolbar_bottom_separator": {"r": {{base03-dec-r}}, "g": {{base03-dec-g}}, "b": {{base03-dec-b}}, "a": 1},
    "toolbar_vertical_separator": {"r": {{base03-dec-r}}, "g": {{base03-dec-g}}, "b": {{base03-dec-b}}, "a": 1},
    "toolbar_field": {"r": {{base00-dec-r}}, "g": {{base00-dec-g}}, "b": {{base00-dec-b}}, "a": 1},
    "toolbar_field_text": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}, "a": 1},
    "toolbar_field_border": {"r": {{base03-dec-r}}, "g": {{base03-dec-g}}, "b": {{base03-dec-b}}, "a": 1},
    "toolbar_field_border_focus": {"r": {{accent-dec-r}}, "g": {{accent-dec-g}}, "b": {{accent-dec-b}}, "a": 1},
    "toolbar_field_focus": {"r": {{base00-dec-r}}, "g": {{base00-dec-g}}, "b": {{base00-dec-b}}, "a": 1},
    "toolbar_field_highlight": {"r": {{accent-dec-r}}, "g": {{accent-dec-g}}, "b": {{accent-dec-b}}, "a": 1},
    "toolbar_field_highlight_text": {"r": {{base00-dec-r}}, "g": {{base00-dec-g}}, "b": {{base00-dec-b}}, "a": 1},
    "icons": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}, "a": 1},
    "icons_attention": {"r": {{base0A-dec-r}}, "g": {{base0A-dec-g}}, "b": {{base0A-dec-b}}, "a": 1},
    "button_background_hover": {"r": {{base02-dec-r}}, "g": {{base02-dec-g}}, "b": {{base02-dec-b}}, "a": 1},
    "button_background_active": {"r": {{base02-dec-r}}, "g": {{base02-dec-g}}, "b": {{base02-dec-b}}, "a": 1},
    "tab_selected": {"r": {{base00-dec-r}}, "g": {{base00-dec-g}}, "b": {{base00-dec-b}}, "a": 1},
    "tab_text": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}, "a": 1},
    "tab_line": {"r": {{accent-dec-r}}, "g": {{accent-dec-g}}, "b": {{accent-dec-b}}, "a": 1},
    "tab_loading": {"r": {{accent-dec-r}}, "g": {{accent-dec-g}}, "b": {{accent-dec-b}}, "a": 1},
    "tab_background_separator": {"r": {{base03-dec-r}}, "g": {{base03-dec-g}}, "b": {{base03-dec-b}}, "a": 1},
    "popup": {"r": {{base01-dec-r}}, "g": {{base01-dec-g}}, "b": {{base01-dec-b}}, "a": 1},
    "popup_text": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}, "a": 1},
    "popup_border": {"r": {{base03-dec-r}}, "g": {{base03-dec-g}}, "b": {{base03-dec-b}}, "a": 1},
    "popup_highlight": {"r": {{base02-dec-r}}, "g": {{base02-dec-g}}, "b": {{base02-dec-b}}, "a": 1},
    "popup_highlight_text": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}, "a": 1},
    "sidebar": {"r": {{base00-dec-r}}, "g": {{base00-dec-g}}, "b": {{base00-dec-b}}},
    "sidebar_text": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}, "a": 1},
    "sidebar_border": {"r": {{base03-dec-r}}, "g": {{base03-dec-g}}, "b": {{base03-dec-b}}, "a": 1},
    "sidebar_highlight": {"r": {{base02-dec-r}}, "g": {{base02-dec-g}}, "b": {{base02-dec-b}}, "a": 1},
    "sidebar_highlight_text": {"r": {{base05-dec-r}}, "g": {{base05-dec-g}}, "b": {{base05-dec-b}}, "a": 1}
  }
}
`

const firefoxManifestTemplate = `{
  "manifest_version": 2,
  "name": "Base16 Changer Theme",
  "version": "{{extension-version}}",
  "description": "Generated theme bridge for LibreWolf/Firefox and Sidebery.",
  "permissions": ["theme"],
  "background": {
    "scripts": ["background.js"]
  },
  "browser_specific_settings": {
    "gecko": {
      "id": "{{extension-id}}"
    }
  }
}
`

const firefoxBackgroundTemplate = `async function applyTheme() {
  const url = browser.runtime.getURL('theme.json');
  const response = await fetch(url);
  const theme = await response.json();
  await browser.theme.update(theme);
}

applyTheme().catch(console.error);
browser.runtime.onInstalled.addListener(() => applyTheme().catch(console.error));
browser.runtime.onStartup.addListener(() => applyTheme().catch(console.error));
`

// Sidebery sidebar tab extension — applied via userContent.css @import
const sideberyTemplate = `/* Generated by base16changer. Do not edit by hand.
   Purpose: Sidebery tab sidebar theming via userContent.css
   Scope: Sidebery CSS custom properties only.
*/

#root.root {
  --tabs-activated-shadow: 0 1px 4px -1px rgba(0,0,0,0.282), inset 0 0 0 1px #{{accent-hex}}, inset 4px 0 0 0 #{{accent-hex}};
  --tabs-border-radius: 5px;
  --tabs-activated-fg: #{{base05-hex}};
  --tabs-activated-bg: #{{base00-hex}};
  --tabs-normal-fg: #{{base04-hex}};
  --tabs-normal-bg: transparent;

  --frame-bg: #{{base00-hex}};
  --frame-fg: #{{base05-hex}};
  --toolbar-bg: #{{base01-hex}};
  --toolbar-fg: #{{base05-hex}};

  --nav-btn-fg: #{{base05-hex}};

  --tabs-hover-bg: #{{base02-hex}};
  --tabs-hover-fg: #{{base05-hex}};

  --popup-bg: #{{base01-hex}};
  --popup-fg: #{{base05-hex}};
  --popup-border: #{{base03-hex}};

  --separator-color: #{{base03-hex}};
  --scroll-progress-h: #{{accent-hex}};
}
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
window.active.border.color: #{{accent-hex}}
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
