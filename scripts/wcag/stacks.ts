// Cross-target "literal stacks" — what the user actually sees when one app's
// content is rendered inside another app's chrome.
//
// The three concrete failures driving this file (from the user):
//   1) Firefox/LibreWolf private window: `toolbar_field` ends up dark gray,
//      `toolbar_field_text` ends up black/dark — Firefox forces a dark UI in
//      private windows that the user CSS can't override.
//   2) Syntect heading rendered inside a GTK-4 app: the heading takes the
//      tmTheme's heading color (or, after contrast normalization, the
//      tmTheme's default fg), which doesn't match the GTK-4 surface bg.
//   3) Fuzzel `selected_bg` + `selected_fg` reading as dark-on-dark, AND the
//      selection stripe being invisible against the menu bg.
//
// Each stack expands those three into the broader permutation set.
//
// `Color` is either a slot key (resolved per-scheme from the palette) or a
// literal hex (for forced overlays Firefox/Mozilla ship in private windows
// and for actual tmTheme scope colors looked up from the .tmTheme file).

import type { Pair } from "./targets.ts";
import type { SlotKey, Palette } from "./palette.ts";

export type Color =
  | { kind: "slot"; key: SlotKey }
  | { kind: "literal"; hex: string; label: string }
  | { kind: "tmtheme"; scope: string; label: string };  // resolved per-scheme

export interface StackPair {
  stack: string;
  role: string;
  fg: Color;
  bg: Color;
  variant?: "light" | "dark";
  notes?: string;
}

const slot = (k: SlotKey): Color => ({ kind: "slot", key: k });
const lit  = (hex: string, label: string): Color => ({ kind: "literal", hex, label });
const tm   = (scope: string, label?: string): Color =>
  ({ kind: "tmtheme", scope, label: label ?? scope });

// ---------------------------------------------------------------------------
// Stack — Syntect-rendered code inside a GTK-4 / libadwaita editor
// (user failure 2: ayu-light syntect heading reading "white" / off-theme on
// GTK-4 gray background).
//
// Each pair pulls the REAL color from the matched .tmTheme file (with
// base16changer's normalizeTmThemeContrast already applied) and tests it
// against base16changer's GTK-4 surfaces.
// ---------------------------------------------------------------------------
const codeInGtk4: StackPair[] = [
  // tmTheme scope colors (the actual rendered text) over GTK-4 surfaces.
  // GTK-4 view_bg and the tmTheme bg are NOT the same hex — that mismatch is
  // exactly the failure mode the user reported.
  { stack: "code-in-gtk4", role: "syntect HEADING on GTK-4 view_bg",       fg: tm("heading"),  bg: slot("base00"),
    notes: "user-flagged: heading scope foreground from .tmTheme, against GTK-4 view bg" },
  { stack: "code-in-gtk4", role: "syntect HEADING on GTK-4 popover_bg",    fg: tm("heading"),  bg: slot("base01"),
    notes: "user-flagged: same color but rendered inside a tooltip/popover" },
  { stack: "code-in-gtk4", role: "syntect HEADING on GTK-4 sidebar_bg",    fg: tm("heading"),  bg: slot("base01") },
  { stack: "code-in-gtk4", role: "syntect HEADING on GTK-4 selection_bg", fg: tm("heading"),  bg: slot("base02") },
  { stack: "code-in-gtk4", role: "syntect comment on GTK-4 view_bg",       fg: tm("comment"),  bg: slot("base00") },
  { stack: "code-in-gtk4", role: "syntect comment on GTK-4 popover_bg",    fg: tm("comment"),  bg: slot("base01") },
  { stack: "code-in-gtk4", role: "syntect string on GTK-4 view_bg",        fg: tm("string"),   bg: slot("base00") },
  { stack: "code-in-gtk4", role: "syntect string on GTK-4 popover_bg",     fg: tm("string"),   bg: slot("base01") },
  { stack: "code-in-gtk4", role: "syntect keyword on GTK-4 view_bg",       fg: tm("keyword"),  bg: slot("base00") },
  { stack: "code-in-gtk4", role: "syntect keyword on GTK-4 popover_bg",    fg: tm("keyword"),  bg: slot("base01") },
  { stack: "code-in-gtk4", role: "syntect function on GTK-4 view_bg",      fg: tm("function"), bg: slot("base00") },
  { stack: "code-in-gtk4", role: "syntect function on GTK-4 popover_bg",   fg: tm("function"), bg: slot("base01") },
  { stack: "code-in-gtk4", role: "syntect type on GTK-4 view_bg",          fg: tm("type"),     bg: slot("base00") },
  { stack: "code-in-gtk4", role: "syntect number on GTK-4 view_bg",        fg: tm("number"),   bg: slot("base00") },
  { stack: "code-in-gtk4", role: "syntect link on GTK-4 view_bg",          fg: tm("link"),     bg: slot("base00") },
  { stack: "code-in-gtk4", role: "syntect default fg on GTK-4 view_bg",    fg: tm("__fg"),     bg: slot("base00"),
    notes: "tmTheme global default fg — what bracket/punctuation/everything-else falls back to" },
  { stack: "code-in-gtk4", role: "syntect default fg on GTK-4 popover_bg", fg: tm("__fg"),     bg: slot("base01") },
  { stack: "code-in-gtk4", role: "tmTheme default bg leaks past GTK-4 view_bg", fg: tm("__bg"), bg: slot("base00"),
    notes: "if tmTheme bg differs from GTK-4 view bg, the gutter strip seam fails" },
];

// ---------------------------------------------------------------------------
// Stack — Firefox / LibreWolf in private window (user failure 1)
//
// Mozilla source: toolkit/themes/shared/in-content/info-pages.css and
// browser/themes/shared/incognito/* hard-pin a dark UI in private mode.
// Observed forced colors in LibreWolf 121:
//   chrome bg:           #1c1b22
//   chrome fg:           #fbfbfe
//   toolbar_field bg:    #42414d   (URL bar background)
//   toolbar_field_text:  #cfcfd8   (URL bar text — but this gets overridden by
//                                   the user's userChrome --mm-field-fg = base05
//                                   in some Firefox/LibreWolf versions, which
//                                   is the user-flagged failure mode)
//
// The pair names match the literal Firefox theme-API keys so they grep cleanly.
// ---------------------------------------------------------------------------
const FF_PRIV_BG       = "#1c1b22";
const FF_PRIV_FG       = "#fbfbfe";
const FF_PRIV_FIELD_BG = "#42414d";
const FF_PRIV_FIELD_FG = "#cfcfd8";

const ffPrivate: StackPair[] = [
  // The user's exact failure: scheme's `toolbar_field_text` (= --mm-field-fg = base05)
  // dragged into Firefox's forced `toolbar_field` (#42414d).
  { stack: "ff-private", role: "toolbar_field_text (scheme base05) on toolbar_field (FF forced #42414d)",
    fg: slot("base05"), bg: lit(FF_PRIV_FIELD_BG, "FF private toolbar_field #42414d"),
    variant: "light",
    notes: "USER-FLAGGED #1: light theme's dark text lands on Firefox's forced dark gray field bg" },
  { stack: "ff-private", role: "toolbar_field_text (scheme base04 muted) on toolbar_field (FF forced)",
    fg: slot("base04"), bg: lit(FF_PRIV_FIELD_BG, "FF private toolbar_field #42414d"),
    variant: "light",
    notes: "muted variant of the same failure (placeholder / disabled URL bar text)" },
  { stack: "ff-private", role: "toolbar_field (scheme base00) overrun by FF forced toolbar_field bg",
    fg: slot("base00"), bg: lit(FF_PRIV_FIELD_BG, "FF private toolbar_field #42414d"),
    variant: "light",
    notes: "what the URL bar would have been vs. what FF actually paints — purely diagnostic" },

  // Inverse: FF's forced light text sometimes leaks onto a scheme-colored bg
  // (when userChrome partially wins).
  { stack: "ff-private", role: "toolbar_field_text (FF forced #cfcfd8) on toolbar_field (scheme base00)",
    fg: lit(FF_PRIV_FIELD_FG, "FF private toolbar_field_text #cfcfd8"), bg: slot("base00"),
    variant: "light" },
  { stack: "ff-private", role: "toolbar_field_text (FF forced #fbfbfe chrome fg) on toolbar_field (scheme base00)",
    fg: lit(FF_PRIV_FG, "FF private chrome fg #fbfbfe"), bg: slot("base00"),
    variant: "light",
    notes: "inverse of #1: scheme's light field bg with FF's near-white text washes out" },

  // Surrounding chrome — scheme text on FF's forced chrome bg.
  { stack: "ff-private", role: "toolbar_text (scheme base05) on chrome bg (FF forced #1c1b22)",
    fg: slot("base05"), bg: lit(FF_PRIV_BG, "FF private chrome bg #1c1b22"),
    variant: "light" },
  { stack: "ff-private", role: "tab_text (scheme base05) on chrome bg (FF forced #1c1b22)",
    fg: slot("base05"), bg: lit(FF_PRIV_BG, "FF private chrome bg #1c1b22"),
    variant: "light" },
  { stack: "ff-private", role: "accent / tab_line (scheme accent) on toolbar_field (FF forced)",
    fg: slot("accent"), bg: lit(FF_PRIV_FIELD_BG, "FF private toolbar_field #42414d") },
  { stack: "ff-private", role: "accent / focus_border (scheme accent) on chrome bg (FF forced)",
    fg: slot("accent"), bg: lit(FF_PRIV_BG, "FF private chrome bg #1c1b22") },

  // Dark-scheme controls — should mostly pass; included to confirm the
  // failure is a light-theme phenomenon.
  { stack: "ff-private", role: "toolbar_field_text (dark scheme base05) on toolbar_field (FF forced) — control",
    fg: slot("base05"), bg: lit(FF_PRIV_FIELD_BG, "FF private toolbar_field #42414d"),
    variant: "dark" },
  { stack: "ff-private", role: "toolbar_text (dark scheme base05) on chrome bg (FF forced) — control",
    fg: slot("base05"), bg: lit(FF_PRIV_BG, "FF private chrome bg #1c1b22"),
    variant: "dark" },
];

// ---------------------------------------------------------------------------
// Stack — Firefox in normal window inside GTK-4 (user CSS effective)
// ---------------------------------------------------------------------------
const ffInGtk4: StackPair[] = [
  { stack: "ff-in-gtk4", role: "toolbar_text on GTK-4 headerbar",         fg: slot("base05"), bg: slot("base01") },
  { stack: "ff-in-gtk4", role: "toolbar_field_text on toolbar_field",     fg: slot("base05"), bg: slot("base00") },
  { stack: "ff-in-gtk4", role: "toolbar_field_border_focus on field bg",  fg: slot("accent"), bg: slot("base00") },
  { stack: "ff-in-gtk4", role: "tab_line (accent) on GTK-4 view_bg below",fg: slot("accent"), bg: slot("base00") },
];

// ---------------------------------------------------------------------------
// Stack — Fuzzel floating over a GTK-4 desktop (user failure 3)
//
// Three failure modes in the user's report:
//   a) `selected_fg on selected_bg` — text on the highlight stripe
//   b) `selected_bg vs background` — the highlight stripe being visible
//      against the rest of the menu
//   c) Same as (a) but for the match-highlight on the selection
// All three live here so they're co-located with the user's mental model.
// ---------------------------------------------------------------------------
const fuzzelOnGtk4: StackPair[] = [
  // (a) selected_fg on selected_bg — post-T3 selected_bg is base02.
  { stack: "fuzzel-on-gtk4", role: "selected_fg (base05) on selected_bg (base02)",
    fg: slot("base05"), bg: slot("base02"),
    notes: "USER-FLAGGED #3a (post-T3): selection bg moved from base03 → base02" },
  // (b) selection stripe visibility — selected_bg vs the menu bg.
  { stack: "fuzzel-on-gtk4", role: "selected_bg (base02) vs fuzzel bg (base01) — stripe visibility",
    fg: slot("base02"), bg: slot("base01"),
    notes: "USER-FLAGGED #3b (post-T3): stripe is now base02 vs base01" },
  // (c) match-highlight inside the selection.
  { stack: "fuzzel-on-gtk4", role: "selection_match (accent) on selected_bg (base02)",
    fg: slot("accent"), bg: slot("base02"),
    notes: "USER-FLAGGED #3c (post-T3): typed-letter highlight on the new base02 selection" },
  // Match-highlight on the unselected menu bg.
  { stack: "fuzzel-on-gtk4", role: "match (accent) on fuzzel bg (base01)",
    fg: slot("accent"), bg: slot("base01") },
  // Fuzzel content vs. the desktop / GTK-4 window behind it (alpha bleed).
  { stack: "fuzzel-on-gtk4", role: "border (accent) vs GTK-4 window_bg behind",
    fg: slot("accent"), bg: slot("base00") },
  { stack: "fuzzel-on-gtk4", role: "text (base05) on fuzzel bg, GTK-4 window behind",
    fg: slot("base05"), bg: slot("base01") },
  // Symmetric checks of failures (b) for other targets that use the same
  // base02/base03 highlight pattern — adjacent to the user's reported case.
  { stack: "fuzzel-on-gtk4", role: "GTK-4 selection_bg (base02) vs view_bg (base00) — selection visibility",
    fg: slot("base02"), bg: slot("base00"),
    notes: "GTK-4 analog of fuzzel #3b" },
  { stack: "fuzzel-on-gtk4", role: "GTK-3 selected_bg (base02) vs base_bg (base01) — selection visibility",
    fg: slot("base02"), bg: slot("base01"),
    notes: "GTK-3 analog of fuzzel #3b" },
];

// ---------------------------------------------------------------------------
// Stack — Openbox titlebar over GTK-4 window content (decorative seam)
// ---------------------------------------------------------------------------
const openboxOnGtk4: StackPair[] = [
  { stack: "openbox-on-gtk4", role: "active border (accent) on GTK-4 view_bg",  fg: slot("accent"), bg: slot("base00") },
  { stack: "openbox-on-gtk4", role: "active title text on GTK-4 view_bg seam", fg: slot("base05"), bg: slot("base00") },
  { stack: "openbox-on-gtk4", role: "inactive border (base02) on GTK-4 view_bg",fg: slot("base02"), bg: slot("base00") },
];

// ---------------------------------------------------------------------------
// Stack — Mako notification floating over desktop
// ---------------------------------------------------------------------------
const makoOnGtk4: StackPair[] = [
  { stack: "mako-on-gtk4", role: "border (accent) vs GTK-4 window_bg",     fg: slot("accent"), bg: slot("base00") },
  { stack: "mako-on-gtk4", role: "text (base05) vs GTK-4 window_bg seam",  fg: slot("base05"), bg: slot("base00") },
];

export const ALL_STACKS: Record<string, StackPair[]> = {
  "code-in-gtk4":    codeInGtk4,
  "ff-in-gtk4":      ffInGtk4,
  "ff-private":      ffPrivate,
  "fuzzel-on-gtk4":  fuzzelOnGtk4,
  "openbox-on-gtk4": openboxOnGtk4,
  "mako-on-gtk4":    makoOnGtk4,
};

export function allStackPairs(): StackPair[] {
  return Object.values(ALL_STACKS).flat();
}

// Resolve a Color against a palette + an optional tmTheme. Returns null if a
// tmTheme color is requested but no tmTheme is loaded or the scope wasn't
// present.
export function resolveColor(
  c: Color,
  pal: Palette,
  tm: { background: string; foreground: string; scopes: Record<string, string> } | null,
): { hex: string; label: string; slotKey: SlotKey | null } | null {
  if (c.kind === "literal") return { hex: c.hex, label: c.label, slotKey: null };
  if (c.kind === "slot")    return { hex: (pal as any)[c.key], label: c.key, slotKey: c.key };
  // tmtheme
  if (!tm) return null;
  let hex: string | undefined;
  if (c.scope === "__bg") hex = tm.background;
  else if (c.scope === "__fg") hex = tm.foreground;
  else hex = tm.scopes[c.scope];
  if (!hex) return null;
  return { hex, label: `tm:${c.label}`, slotKey: null };
}
