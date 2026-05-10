// Intra-target role pairs. Each Pair is a foreground role rendered ON a
// background role within the same target. Slot keys reference the unified
// palette (see palette.ts). Sourced from internal/targets/templates.go +
// targets.go (see role inventory in commit notes).
//
// Variant-specific pairs use `variant` to opt-in. Pairs without `variant`
// run on both light and dark schemes.

import type { SlotKey } from "./palette.ts";

export type Variant = "light" | "dark";

export interface Pair {
  target: string;
  role: string;        // human-readable (used in report)
  fg: SlotKey;
  bg: SlotKey;
  variant?: Variant;   // restrict to this variant if set
  notes?: string;      // shown in report when failing
}

// Helper for table compactness.
const P = (target: string, role: string, fg: SlotKey, bg: SlotKey, extra: Partial<Pair> = {}): Pair =>
  ({ target, role, fg, bg, ...extra });

// ---------------------------------------------------------------------------
// Kitty (terminal). ANSI color slots tested against the terminal background
// because that's what users actually read.
// ---------------------------------------------------------------------------
const kitty: Pair[] = [
  P("kitty", "default text on background",        "base05", "base00"),
  P("kitty", "selection text on selection bg",    "base02", "base05",
    { notes: "kitty inverts: selection_fg=base02, selection_bg=base05" }),
  P("kitty", "cursor on background",              "base05", "base00"),
  P("kitty", "cursor_text on cursor",             "base00", "base05"),
  P("kitty", "url_color on background",           "base04", "base00"),
  P("kitty", "active border on bg",               "base03", "base00"),
  P("kitty", "active tab fg on active tab bg",    "base05", "base00"),
  P("kitty", "inactive tab fg on inactive tab bg","base05", "base01"),
  P("kitty", "ANSI red on bg (color_01 logs)",    "base08", "base00"),
  P("kitty", "ANSI green on bg",                  "base0B", "base00"),
  P("kitty", "ANSI yellow on bg",                 "base0A", "base00"),
  P("kitty", "ANSI blue on bg",                   "base0D", "base00"),
  P("kitty", "ANSI magenta on bg",                "base0E", "base00"),
  P("kitty", "ANSI cyan on bg",                   "base0C", "base00"),
  P("kitty", "bright black (comments) on bg",     "base03", "base00"),
];

// ---------------------------------------------------------------------------
// Fuzzel (Wayland launcher). Background carries 0xf2 alpha; we test against
// the solid base01 — alpha cases are also tested in stacks.ts.
// ---------------------------------------------------------------------------
const fuzzel: Pair[] = [
  P("fuzzel", "text on background",                "base05", "base01"),
  P("fuzzel", "match on background",               "accent", "base01",
    { notes: "user-flagged: accent often disappears when accent==base0D and palette has unsaturated blue" }),
  P("fuzzel", "selection-text on selection",       "base05", "base02",
    { notes: "post-T3: selection now base02 (was base03) for stripe visibility" }),
  P("fuzzel", "selection-match on selection",      "accent", "base02"),
  P("fuzzel", "border on bg (visual indicator)",   "accent", "base01"),
  P("fuzzel", "prompt/text on bg",                 "base05", "base01"),
];

// ---------------------------------------------------------------------------
// Mako (notification daemon).
// ---------------------------------------------------------------------------
const mako: Pair[] = [
  P("mako", "text on background",  "base05", "base01"),
  P("mako", "border on bg",        "accent", "base01"),
];

// ---------------------------------------------------------------------------
// Openbox / LabWC.
// ---------------------------------------------------------------------------
const openbox: Pair[] = [
  P("openbox", "active title text on active title bg",   "base05", "base01"),
  P("openbox", "inactive title text on inactive title bg","base05", "base00",
    { notes: "post-T3: inactive title text now base05 (was base03)" }),
  P("openbox", "active button icon on button bg",        "base05", "base01"),
  P("openbox", "active button icon hover",               "base07", "base02"),
  P("openbox", "inactive button icon on inactive btn bg","base03", "base00"),
  P("openbox", "active border (outline) on active title","accent", "base01"),
  P("openbox", "menu items on menu bg",                  "base05", "base00"),
  P("openbox", "menu disabled on menu bg",               "base05", "base00"),
  P("openbox", "menu active item on highlight",          "base05", "base02"),
  P("openbox", "OSD label on osd bg",                    "base05", "base00"),
  P("openbox", "close button (red) on active title",     "base08", "base01"),
];

// ---------------------------------------------------------------------------
// GTK-3.
// ---------------------------------------------------------------------------
const gtk3: Pair[] = [
  P("gtk3", "fg on bg",                       "base05", "base00"),
  P("gtk3", "text on base (content area)",    "base05", "base01"),
  P("gtk3", "disabled text on bg",            "base05", "base00"),
  P("gtk3", "selected text on selected bg",   "base05", "base02"),
  P("gtk3", "tooltip text on tooltip bg",     "base05", "base00"),
  P("gtk3", "link on bg",                     "base0D", "base00"),
  P("gtk3", "link on base (content)",         "base0D", "base01"),
  P("gtk3", "success on bg",                  "base0B", "base00"),
  P("gtk3", "warning on bg",                  "base0A", "base00"),
  P("gtk3", "error on bg",                    "base08", "base00"),
  P("gtk3", "border (visual UI) on bg",       "base02", "base00"),
];

// ---------------------------------------------------------------------------
// GTK-4 / libadwaita. Most-load-bearing target — same role names users see in
// adwaita-1 tools. accent_fg = base00 (text on accent) often fails for low-
// saturation accents; flag every accent-on-bg pair.
// ---------------------------------------------------------------------------
const gtk4: Pair[] = [
  P("gtk4", "window_fg on window_bg",                 "base05", "base00"),
  P("gtk4", "view_fg on view_bg",                     "base05", "base00"),
  P("gtk4", "headerbar_fg on headerbar_bg",           "base05", "base01"),
  P("gtk4", "sidebar_fg on sidebar_bg",               "base05", "base01"),
  P("gtk4", "card_fg on card_bg",                     "base05", "base01"),
  P("gtk4", "dialog_fg on dialog_bg",                 "base05", "base01"),
  P("gtk4", "popover_fg on popover_bg",               "base05", "base01"),
  P("gtk4", "accent_color on window_bg",              "accent", "base00",
    { notes: "this is the pair the accent: field is meant to fix" }),
  P("gtk4", "accent_color on view_bg",                "accent", "base00"),
  P("gtk4", "accent_color on headerbar_bg",           "accent", "base01"),
  P("gtk4", "accent_color on sidebar_bg",             "accent", "base01"),
  P("gtk4", "accent_color on card_bg",                "accent", "base01"),
  P("gtk4", "accent_fg_color on accent_bg_color",     "base00", "accent",
    { notes: "text rendered on solid accent button" }),
  P("gtk4", "destructive_fg on destructive_bg",       "base00", "base08"),
  P("gtk4", "success_fg on success_bg",               "base00", "base0B"),
  P("gtk4", "warning_fg on warning_bg",               "base00", "base0A"),
  P("gtk4", "error_fg on error_bg",                   "base00", "base08"),
  P("gtk4", "blue_3 link on view_bg",                 "base0D", "base00"),
  P("gtk4", "scrollbar outline on view_bg",           "base02", "base00"),
];

// ---------------------------------------------------------------------------
// GtkSourceView 5. Comments use base05 in the template (same as default text)
// — we test the *effective* contrast against base00 view bg as well as the
// base01 current_line bg (because a line being edited shifts the bg).
// ---------------------------------------------------------------------------
const gsv: Pair[] = [
  P("gsv", "text on view bg",                "base05", "base00"),
  P("gsv", "text on current_line bg",        "base05", "base01"),
  P("gsv", "comment on view bg",             "base05", "base00"),
  P("gsv", "string on view bg",              "base0B", "base00"),
  P("gsv", "string on current_line",         "base0B", "base01"),
  P("gsv", "number on view bg",              "base09", "base00"),
  P("gsv", "function on view bg",            "base0D", "base00"),
  P("gsv", "function on current_line",       "base0D", "base01"),
  P("gsv", "keyword on view bg",             "base0E", "base00"),
  P("gsv", "keyword on current_line",        "base0E", "base01"),
  P("gsv", "type on view bg",                "base0B", "base00"),
  P("gsv", "builtin on view bg",             "base0C", "base00"),
  P("gsv", "constant/special on view bg",    "base0A", "base00"),
  P("gsv", "operator on view bg",            "base0B", "base00"),
  P("gsv", "preprocessor on view bg",        "base0E", "base00"),
  P("gsv", "heading on view bg",             "base0D", "base00",
    { notes: "user-flagged: ayu-light heading reads as white on gray" }),
  P("gsv", "list-marker on view bg",         "base09", "base00"),
  P("gsv", "link-text on view bg",           "base0E", "base00"),
  P("gsv", "link-destination on view bg",    "base0C", "base00"),
  P("gsv", "search match on match bg",       "base05", "base02"),
  P("gsv", "bracket match on view bg",       "base0E", "base00"),
  P("gsv", "bracket-mismatch fg on bg",      "base08", "base01"),
  P("gsv", "right_margin text on margin bg", "base05", "base01"),
  P("gsv", "line-numbers on gutter",         "base05", "base00"),
];

// ---------------------------------------------------------------------------
// Firefox / LibreWolf userChrome.
// Private-window override is simulated in stacks.ts.
// ---------------------------------------------------------------------------
const ff: Pair[] = [
  P("firefox", "toolbar text on toolbar bg",            "base05", "base01"),
  P("firefox", "field (URL bar) text on field bg",      "base05", "base00"),
  P("firefox", "field focus border on field bg",        "accent", "base00"),
  P("firefox", "field highlight on field bg",           "accent", "base00"),
  P("firefox", "field highlight text on highlight",     "base00", "accent"),
  P("firefox", "popup text on popup bg",                "base05", "base01"),
  P("firefox", "popup border on popup bg",              "base05", "base01"),
  P("firefox", "selected suggestion url on hover bg",   "accent", "base02"),
  P("firefox", "sidebar text on sidebar bg",            "base05", "base00"),
  P("firefox", "sidebar border on sidebar bg",          "base05", "base00"),
  P("firefox", "muted text on toolbar bg",              "base05", "base01"),
  P("firefox", "tab selected fg on tab selected bg",    "base05", "base00"),
  P("firefox", "tab line (accent underline) on tab bg", "accent", "base00"),
  P("firefox", "tab hover fg on tab hover bg",          "base05", "base02"),
  P("firefox", "danger (.close-icon hover) on toolbar", "base08", "base01"),
  P("firefox", "warning (attention) on toolbar",        "base0A", "base01"),
  P("firefox", "success (download badge) on toolbar",   "base0B", "base01"),
];

// ---------------------------------------------------------------------------
// Sidebery (sidebar tab extension).
// ---------------------------------------------------------------------------
const sidebery: Pair[] = [
  P("sidebery", "active tab fg on active tab bg",       "base05", "base00"),
  P("sidebery", "inactive tab fg on frame bg",          "base04", "base00",
    { notes: "user-flagged equivalent: muted inactive tab text" }),
  P("sidebery", "toolbar fg on toolbar bg",             "base05", "base01"),
  P("sidebery", "tab hover fg on hover bg",             "base05", "base02"),
  P("sidebery", "popup fg on popup bg",                 "base05", "base01"),
  P("sidebery", "scroll progress (accent) on frame bg", "accent", "base00"),
  P("sidebery", "active-tab outline (accent) on tab",   "accent", "base00"),
];

// ---------------------------------------------------------------------------
// All targets bundled.
// ---------------------------------------------------------------------------
export const ALL_TARGETS: Record<string, Pair[]> = {
  kitty, fuzzel, mako, openbox, gtk3, gtk4, gsv, firefox: ff, sidebery,
};

export function allPairs(): Pair[] {
  return Object.values(ALL_TARGETS).flat();
}
