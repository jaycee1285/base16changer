// Parse base16 (.yaml) and gogh (.yml) theme files into a unified Palette.
// Mirrors internal/scheme/parse.go + gogh.go closely enough for contrast tests.
// Lab interpolation (go-colorful) is approximated with linear sRGB blending —
// see contrast.ts blendHex.

import { readFileSync } from "node:fs";
import { blendHex } from "./contrast.ts";

export type SlotKey =
  | "base00" | "base01" | "base02" | "base03"
  | "base04" | "base05" | "base06" | "base07"
  | "base08" | "base09" | "base0A" | "base0B"
  | "base0C" | "base0D" | "base0E" | "base0F"
  | "accent";

export interface Palette {
  // metadata
  name: string;
  variant: "light" | "dark";
  source: string;            // absolute path
  format: "base16" | "gogh";
  // 16 slots + accent (all hex with leading '#')
  base00: string; base01: string; base02: string; base03: string;
  base04: string; base05: string; base06: string; base07: string;
  base08: string; base09: string; base0A: string; base0B: string;
  base0C: string; base0D: string; base0E: string; base0F: string;
  accent: string;            // = explicit accent if set, else base0D
  accentExplicit: boolean;
}

export const SLOT_KEYS: SlotKey[] = [
  "base00","base01","base02","base03","base04","base05","base06","base07",
  "base08","base09","base0A","base0B","base0C","base0D","base0E","base0F",
  "accent",
];

// --- Tiny YAML extraction (no external dep) --------------------------------

function extractTopKey(content: string, key: string): string | null {
  const re = new RegExp(`(^|\\n)${key}\\s*:\\s*['"]?#?([0-9a-fA-F]{6})['"]?`, "i");
  const m = content.match(re);
  if (m) return "#" + m[2].toLowerCase();
  // also accept non-hex string values for name/variant. `m` flag → $ = EOL.
  const re2 = new RegExp(`^${key}\\s*:\\s*['"]?([^'"\\n#]+?)['"]?\\s*(?:#.*)?$`, "im");
  const m2 = content.match(re2);
  return m2 ? m2[1].trim() : null;
}

function extractIndentedHex(content: string, parent: string, key: string): string | null {
  const lines = content.split(/\r?\n/);
  let inParent = false;
  for (const line of lines) {
    if (!inParent) { if (new RegExp(`^${parent}\\s*:`).test(line)) inParent = true; continue; }
    if (/^\S/.test(line) && !line.startsWith("#")) break;
    const m = line.match(new RegExp(`^\\s+${key}\\s*:\\s*['"]?#?([0-9a-fA-F]{6})['"]?`));
    if (m) return "#" + m[1].toLowerCase();
  }
  return null;
}

// --- Variant inference (matches isLightVariant in parse.go) ----------------

function inferVariant(declared: string | null, base00: string, base05: string): "light" | "dark" {
  const d = (declared ?? "").trim().toLowerCase();
  if (d === "light") return "light";
  if (d === "dark") return "dark";
  const sumHex = (h: string) => {
    const n = parseInt(h.replace(/^#/, ""), 16);
    return ((n >> 16) & 0xff) + ((n >> 8) & 0xff) + (n & 0xff);
  };
  return sumHex(base00) > sumHex(base05) ? "light" : "dark";
}

// --- Base16 (.yaml) --------------------------------------------------------

function parseBase16(path: string, content: string): Palette {
  const get = (k: string) => extractIndentedHex(content, "palette", k);
  const required: Record<string, string | null> = {};
  for (const k of ["base00","base01","base02","base03","base04","base05",
                   "base06","base07","base08","base09","base0A","base0B",
                   "base0C","base0D","base0E","base0F"]) {
    required[k] = get(k);
  }
  for (const [k, v] of Object.entries(required)) {
    if (!v) throw new Error(`base16 palette missing ${k} in ${path}`);
  }
  const accentRaw = extractTopKey(content, "accent");
  const accentHex = accentRaw && /^#?[0-9a-fA-F]{6}$/.test(accentRaw.replace(/^#/, ""))
    ? "#" + accentRaw.replace(/^#/, "").toLowerCase() : null;
  const name = extractTopKey(content, "name") ?? path.split("/").pop() ?? "unknown";
  const declaredVariant = extractTopKey(content, "variant");
  const variant = inferVariant(declaredVariant, required.base00!, required.base05!);

  return {
    name, variant, source: path, format: "base16",
    base00: required.base00!, base01: required.base01!, base02: required.base02!,
    base03: required.base03!, base04: required.base04!, base05: required.base05!,
    base06: required.base06!, base07: required.base07!, base08: required.base08!,
    base09: required.base09!, base0A: required.base0A!, base0B: required.base0B!,
    base0C: required.base0C!, base0D: required.base0D!, base0E: required.base0E!,
    base0F: required.base0F!,
    accent: accentHex ?? required.base0D!,
    accentExplicit: !!accentHex,
  };
}

// --- Gogh (.yml) ----------------------------------------------------------

interface Gogh {
  name: string; variant: string;
  c01: string; c02: string; c03: string; c04: string;
  c05: string; c06: string; c07: string; c08: string;
  c09: string; c10: string; c11: string; c12: string;
  c13: string; c14: string; c15: string; c16: string;
  bg: string; fg: string; cursor: string;
}

function parseGogh(path: string, content: string): Palette {
  const need = (k: string): string => {
    const v = extractTopKey(content, k);
    if (!v || !/^#?[0-9a-fA-F]{6}$/.test(v.replace(/^#/, ""))) {
      throw new Error(`gogh missing ${k} in ${path}`);
    }
    return "#" + v.replace(/^#/, "").toLowerCase();
  };
  const opt = (k: string): string => {
    try { return need(k); } catch { return ""; }
  };
  const g: Gogh = {
    name: extractTopKey(content, "name") ?? path.split("/").pop() ?? "unknown",
    variant: extractTopKey(content, "variant") ?? "",
    c01: need("color_01"), c02: need("color_02"), c03: need("color_03"), c04: need("color_04"),
    c05: need("color_05"), c06: need("color_06"), c07: need("color_07"), c08: need("color_08"),
    c09: need("color_09"), c10: need("color_10"), c11: need("color_11"), c12: need("color_12"),
    c13: need("color_13"), c14: need("color_14"), c15: need("color_15"), c16: need("color_16"),
    bg: need("background"), fg: need("foreground"), cursor: opt("cursor"),
  };
  // Mirror gogh.go ToBase16 (see internal/scheme/gogh.go).
  const orange = blendHex(g.c02, g.c04, 0.5);          // red ↔ yellow
  const brown  = blendHex(orange, g.bg, 0.4);
  const base00 = g.bg;
  const base05 = g.fg;
  const base01 = blendHex(base00, base05, 0.1);
  const base02 = blendHex(base00, base05, 0.2);
  const base04 = blendHex(base00, base05, 0.4);
  const base06 = blendHex(base00, base05, 0.8);

  const accentRaw = extractTopKey(content, "accent");
  const accentHex = accentRaw && /^#?[0-9a-fA-F]{6}$/.test(accentRaw.replace(/^#/, ""))
    ? "#" + accentRaw.replace(/^#/, "").toLowerCase() : null;
  const variant = inferVariant(g.variant || null, base00, base05);

  return {
    name: g.name, variant, source: path, format: "gogh",
    base00, base01, base02,
    base03: g.c09,           // bright black
    base04, base05, base06,
    base07: g.c16,           // bright white
    base08: g.c02,           // red
    base09: orange,          // derived
    base0A: g.c04,           // yellow
    base0B: g.c03,           // green
    base0C: g.c07,           // cyan
    base0D: g.c05,           // blue
    base0E: g.c06,           // magenta
    base0F: brown,            // derived
    accent: accentHex ?? g.c05,
    accentExplicit: !!accentHex,
  };
}

// --- Public ----------------------------------------------------------------

export function loadPalette(path: string): Palette {
  const content = readFileSync(path, "utf8");
  const isGogh = /\bcolor_01\s*:/i.test(content);
  const isBase16 = /\bpalette\s*:/i.test(content) && /\bbase00\s*:/i.test(content);
  if (isBase16) return parseBase16(path, content);
  if (isGogh) return parseGogh(path, content);
  throw new Error(`unrecognized scheme format: ${path}`);
}

export function slot(p: Palette, key: SlotKey): string {
  return (p as any)[key];
}

// Slot category → used by the fallback engine to keep stylistic intent.
export type SlotCategory = "surface" | "muted" | "text" | "color" | "accent";
export const SLOT_CATEGORY: Record<SlotKey, SlotCategory> = {
  base00: "surface", base01: "surface", base02: "surface", base03: "muted",
  base04: "muted",   base05: "text",    base06: "text",    base07: "text",
  base08: "color",   base09: "color",   base0A: "color",   base0B: "color",
  base0C: "color",   base0D: "color",   base0E: "color",   base0F: "color",
  accent: "accent",
};
