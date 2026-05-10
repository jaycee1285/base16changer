// Fallback engine. For a failing (fg over bg) pair, search the palette for an
// alternative slot in the SAME category whose ratio passes. Used to inform
// future base16changer auto-fallback logic. Keeps stylistic intent (a failing
// "accent" suggests another color slot, not base05; a failing muted "comment"
// suggests another muted slot, not a hot pink).

import { contrastRatio, AA_THRESHOLD } from "./contrast.ts";
import { type Palette, type SlotKey, slot, SLOT_CATEGORY } from "./palette.ts";

export interface Suggestion {
  swap: SlotKey;
  hex: string;
  ratio: number;
}

const ORDER: Record<string, SlotKey[]> = {
  text:    ["base05", "base06", "base07", "base04", "base03"],
  muted:   ["base04", "base03", "base05", "base06"],
  surface: ["base01", "base02", "base00", "base03"],
  // accents/colors: try every chromatic slot then the explicit accent
  color:   ["base08","base09","base0A","base0B","base0C","base0D","base0E","base0F","accent"],
  accent:  ["accent","base0D","base0E","base08","base09","base0A","base0B","base0C","base0F"],
};

export function suggestFallback(
  pal: Palette,
  failingFg: SlotKey,
  bgKey: SlotKey | null,
  bgHex: string,
  threshold = AA_THRESHOLD,
): Suggestion | null {
  const cat = SLOT_CATEGORY[failingFg];
  const candidates = (ORDER[cat] ?? []).filter(k => k !== failingFg);
  let best: Suggestion | null = null;
  for (const k of candidates) {
    const hex = slot(pal, k);
    if (!hex) continue;
    if (bgKey && k === bgKey) continue; // don't suggest same-as-bg
    const r = contrastRatio(hex, bgHex);
    if (!Number.isFinite(r)) continue;
    if (r < threshold) continue;
    if (!best || r > best.ratio) best = { swap: k, hex, ratio: r };
  }
  return best;
}
