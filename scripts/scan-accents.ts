#!/usr/bin/env bun
// Scan Gogh (.yml) and Base16 (.yaml) theme files for non-blue accent colors.
//
// Algorithm (per ACCENT-COLORS.md):
//   Gogh:    cursor -> distinctiveness filter, else most-saturated non-blue ANSI
//            slot, else null (would fall back to color_05 blue).
//   Base16:  if base0D already non-blue+saturated, that IS the accent. Otherwise
//            scan base06,07,08-0F for most-saturated non-blue.
//
// Usage:  nix develop -c bun run scripts/scan-accents.ts <dir>
//   (a project hook blocks bare `bun run`; always invoke through nix develop)

import { readdirSync, readFileSync, statSync } from "node:fs";
import { extname, join } from "node:path";

const SAT_MIN = 35;        // %
const DIST_MIN = 60;       // RGB euclidean
const BLUE_HUE_LO = 200;   // deg
const BLUE_HUE_HI = 260;
const BASE0D_SAT_OK = 25;  // already-good base0D threshold

type RGB = [number, number, number];
type HSL = { h: number; s: number; l: number };

function hexToRgb(hex: string): RGB | null {
  const m = hex.replace(/^#/, "").match(/^([0-9a-fA-F]{6})$/);
  if (!m) return null;
  const v = parseInt(m[1], 16);
  return [(v >> 16) & 0xff, (v >> 8) & 0xff, v & 0xff];
}

function rgbToHsl([r, g, b]: RGB): HSL {
  const R = r / 255, G = g / 255, B = b / 255;
  const max = Math.max(R, G, B), min = Math.min(R, G, B);
  const l = (max + min) / 2;
  let h = 0, s = 0;
  if (max !== min) {
    const d = max - min;
    s = l > 0.5 ? d / (2 - max - min) : d / (max + min);
    switch (max) {
      case R: h = ((G - B) / d + (G < B ? 6 : 0)); break;
      case G: h = ((B - R) / d + 2); break;
      case B: h = ((R - G) / d + 4); break;
    }
    h *= 60;
  }
  return { h, s: s * 100, l: l * 100 };
}

function dist(a: RGB, b: RGB): number {
  const dr = a[0] - b[0], dg = a[1] - b[1], db = a[2] - b[2];
  return Math.sqrt(dr * dr + dg * dg + db * db);
}

function saturation(hex: string): number {
  const rgb = hexToRgb(hex);
  if (!rgb) return 0;
  return rgbToHsl(rgb).s;
}

function isBlue(hex: string): boolean {
  const rgb = hexToRgb(hex);
  if (!rgb) return false;
  const { h, s } = rgbToHsl(rgb);
  if (s < 15) return false; // gray, not blue
  return h >= BLUE_HUE_LO && h <= BLUE_HUE_HI;
}

// --- File parsing -----------------------------------------------------------

function extractKey(content: string, key: string): string | null {
  // matches: `key: '#abcdef'` or `key: "#abcdef"` or `key: abcdef` (with optional comments)
  const re = new RegExp(`(^|\\n)\\s*${key}\\s*:\\s*['"]?#?([0-9a-fA-F]{6})['"]?`, "i");
  const m = content.match(re);
  return m ? "#" + m[2].toLowerCase() : null;
}

function extractIndentedKey(content: string, parent: string, key: string): string | null {
  // crude but works: locate parent line, search subsequent indented region
  const lines = content.split(/\r?\n/);
  let inParent = false;
  for (const line of lines) {
    if (!inParent) {
      if (new RegExp(`^${parent}\\s*:`).test(line)) inParent = true;
      continue;
    }
    // exit parent when a top-level (non-indented, non-empty, non-comment) line appears
    if (/^\S/.test(line) && !line.startsWith("#")) break;
    const m = line.match(new RegExp(`^\\s+${key}\\s*:\\s*['"]?#?([0-9a-fA-F]{6})['"]?`));
    if (m) return "#" + m[1].toLowerCase();
  }
  return null;
}

// --- Detection --------------------------------------------------------------

export type Detect =
  | { kind: "gogh"; accent: string | null; source: string; base0d: string | null }
  | { kind: "base16"; accent: string | null; source: string; base0d: string | null; skipReason?: string }
  | { kind: "unknown" }
  | { kind: "parse_error"; reason: string };

export function detect(content: string): Detect {
  const isGogh = /\bcolor_01\s*:/i.test(content);
  const isBase16 = /\bpalette\s*:/i.test(content) && /\bbase00\s*:/i.test(content);

  if (isGogh) return detectGogh(content);
  if (isBase16) return detectBase16(content);
  return { kind: "unknown" };
}

function pad2(n: number): string { return n < 10 ? `0${n}` : `${n}`; }

function detectGogh(content: string): Detect {
  const cursor = extractKey(content, "cursor");
  const bg = extractKey(content, "background");
  const fg = extractKey(content, "foreground");
  const slots: { key: string; hex: string }[] = [];
  for (let i = 1; i <= 16; i++) {
    const k = `color_${pad2(i)}`;
    const v = extractKey(content, k);
    if (v) slots.push({ key: k, hex: v });
  }
  const base0d = slots.find(s => s.key === "color_05")?.hex ?? null;

  if (!bg || !fg || slots.length === 0) {
    return { kind: "parse_error", reason: "missing background/foreground/color slots" };
  }
  const bgRgb = hexToRgb(bg)!, fgRgb = hexToRgb(fg)!;

  // 1. Cursor candidate
  if (cursor) {
    const c = hexToRgb(cursor);
    if (c && !isBlue(cursor)
        && saturation(cursor) > SAT_MIN
        && dist(c, bgRgb) > DIST_MIN
        && dist(c, fgRgb) > DIST_MIN) {
      return { kind: "gogh", accent: cursor, source: "cursor", base0d };
    }
  }
  // 2. Most saturated non-blue ANSI slot passing distinctiveness
  let best: { key: string; hex: string; sat: number } | null = null;
  for (const s of slots) {
    if (isBlue(s.hex)) continue;
    const sat = saturation(s.hex);
    if (sat < SAT_MIN) continue;
    const c = hexToRgb(s.hex);
    if (!c) continue;
    if (dist(c, bgRgb) < DIST_MIN) continue;
    if (dist(c, fgRgb) < DIST_MIN) continue;
    if (!best || sat > best.sat) best = { ...s, sat };
  }
  if (best) return { kind: "gogh", accent: best.hex, source: best.key, base0d };
  return { kind: "gogh", accent: null, source: "fallback (color_05 blue)", base0d };
}

function detectBase16(content: string): Detect {
  const get = (k: string) => extractIndentedKey(content, "palette", k);
  const palette: Record<string, string | null> = {};
  const keys = ["base00","base01","base02","base03","base04","base05",
                "base06","base07","base08","base09","base0A","base0B",
                "base0C","base0D","base0E","base0F"];
  for (const k of keys) palette[k] = get(k);

  const base0d = palette.base0D;
  if (!base0d) return { kind: "parse_error", reason: "missing base0D" };

  // If base0D is already non-blue + saturated, that IS the accent (no override needed).
  if (!isBlue(base0d) && saturation(base0d) > BASE0D_SAT_OK) {
    return {
      kind: "base16",
      accent: null,
      source: "base0D-already-non-blue",
      base0d,
      skipReason: "base0d_already_non_blue",
    };
  }

  // Check base06,07 (often repurposed) and base08-0F (skip 0D itself since it's blue here).
  const candidateKeys = ["base06","base07","base08","base09","base0A","base0B","base0C","base0E","base0F"];
  let best: { key: string; hex: string; sat: number } | null = null;
  for (const k of candidateKeys) {
    const hex = palette[k];
    if (!hex) continue;
    if (isBlue(hex)) continue;
    const sat = saturation(hex);
    if (sat < SAT_MIN) continue;
    if (!best || sat > best.sat) best = { key: k, hex, sat };
  }
  if (best) return { kind: "base16", accent: best.hex, source: best.key, base0d };
  return { kind: "base16", accent: null, source: "no-non-blue-candidate", base0d };
}

// --- CLI --------------------------------------------------------------------

function walk(dir: string): string[] {
  const out: string[] = [];
  for (const f of readdirSync(dir)) {
    const p = join(dir, f);
    if (statSync(p).isFile() && (extname(f) === ".yml" || extname(f) === ".yaml")) out.push(p);
  }
  return out.sort();
}

if (import.meta.main) {
  const dir = process.argv[2];
  if (!dir) {
    console.error("usage: bun run scripts/scan-accents.ts <dir>");
    process.exit(2);
  }
  for (const path of walk(dir)) {
    let content: string;
    try { content = readFileSync(path, "utf8"); }
    catch (e) { console.log(`READ_ERR\t${path}\t${(e as Error).message}`); continue; }
    const d = detect(content);
    const name = path.split("/").pop();
    switch (d.kind) {
      case "gogh":
      case "base16":
        console.log(
          `${d.kind.toUpperCase().padEnd(7)}  ${(d.accent ?? "—").padEnd(8)}  ` +
          `src=${d.source.padEnd(28)}  base0D=${d.base0d ?? "—"}  ${name}`
        );
        break;
      case "unknown": console.log(`UNKNOWN                            ${name}`); break;
      case "parse_error": console.log(`PARSE_ERR  ${d.reason.padEnd(40)}  ${name}`); break;
    }
  }
}
