// Parse the .tmTheme file base16changer actually copies for syntect, so the
// suite can test the real heading / comment / keyword colors (not a base-slot
// proxy). Mirrors the contrast-normalization step described in the inventory:
// any scope foreground with ratio < 4.5:1 against the tmTheme's declared bg
// is replaced with the tmTheme's default fg.

import { readFileSync, existsSync } from "node:fs";
import { contrastRatio, AA_THRESHOLD } from "./contrast.ts";

const TM_DIR_DEFAULT = "/home/john/.local/share/themes/tmThemes";

export interface TmThemeColors {
  path: string;
  background: string;        // global bg
  foreground: string;        // global fg
  scopes: Record<string, string>; // scope → foreground (post-normalization)
}

const SCOPES_OF_INTEREST: { key: string; pattern: RegExp }[] = [
  { key: "heading",        pattern: /markup\.heading|entity\.name\.section/i },
  { key: "comment",        pattern: /\bcomment\b/i },
  { key: "string",         pattern: /\bstring\b/i },
  { key: "keyword",        pattern: /\bkeyword(?!\.)/i },
  { key: "function",       pattern: /entity\.name\.function|support\.function/i },
  { key: "type",           pattern: /entity\.name\.type|support\.type|storage\.type/i },
  { key: "constant",       pattern: /constant(?!\.character)/i },
  { key: "variable",       pattern: /variable(?!\.parameter)/i },
  { key: "number",         pattern: /constant\.numeric/i },
  { key: "operator",       pattern: /keyword\.operator/i },
  { key: "link",           pattern: /markup\.underline\.link|markup\.link/i },
  { key: "punctuation",    pattern: /\bpunctuation\b/i },
];

// tmTheme files have a nested structure where a scope rule's `<key>scope</key>`
// sits in an outer dict and its `<key>foreground</key>` sits in an inner
// settings dict. Rather than walk the XML tree, scan all foreground/background
// tags and the most-recent scope tag preceding each. The first foreground
// found before any scope tag is the global default.
function parseDicts(xml: string): { scope: string | null; foreground: string | null; background: string | null }[] {
  type Tag = { kind: "scope" | "fg" | "bg"; value: string; pos: number };
  const tags: Tag[] = [];
  const reScope = /<key>scope<\/key>\s*<string>([\s\S]*?)<\/string>/g;
  const reFg    = /<key>foreground<\/key>\s*<string>\s*(#[0-9a-fA-F]{6,8})\s*<\/string>/g;
  const reBg    = /<key>background<\/key>\s*<string>\s*(#[0-9a-fA-F]{6,8})\s*<\/string>/g;
  let m: RegExpExecArray | null;
  while ((m = reScope.exec(xml)) !== null) tags.push({ kind: "scope", value: m[1], pos: m.index });
  while ((m = reFg.exec(xml))    !== null) tags.push({ kind: "fg",    value: m[1].slice(0, 7).toLowerCase(), pos: m.index });
  while ((m = reBg.exec(xml))    !== null) tags.push({ kind: "bg",    value: m[1].slice(0, 7).toLowerCase(), pos: m.index });
  tags.sort((a, b) => a.pos - b.pos);

  const out: { scope: string | null; foreground: string | null; background: string | null }[] = [];

  // Globals: the first contiguous run of fg/bg tags before the first scope.
  let i = 0;
  let gFg: string | null = null, gBg: string | null = null;
  while (i < tags.length && tags[i].kind !== "scope") {
    if (tags[i].kind === "fg" && !gFg) gFg = tags[i].value;
    if (tags[i].kind === "bg" && !gBg) gBg = tags[i].value;
    i++;
  }
  out.push({ scope: null, foreground: gFg, background: gBg });

  // Each scope tag from here on adopts the NEXT foreground (and bg if any)
  // before the following scope tag.
  while (i < tags.length) {
    if (tags[i].kind !== "scope") { i++; continue; }
    const scope = tags[i].value;
    let fg: string | null = null, bg: string | null = null;
    let j = i + 1;
    while (j < tags.length && tags[j].kind !== "scope") {
      if (tags[j].kind === "fg" && !fg) fg = tags[j].value;
      if (tags[j].kind === "bg" && !bg) bg = tags[j].value;
      j++;
    }
    out.push({ scope, foreground: fg, background: bg });
    i = j;
  }
  return out;
}

export function findTmThemeFor(schemeName: string, dir = TM_DIR_DEFAULT): string | null {
  const tries = [
    `${dir}/${schemeName}.tmTheme`,
    `${dir}/${schemeName.replace(/\s+/g, "")}.tmTheme`,
    `${dir}/${schemeName.toLowerCase().replace(/\s+/g, "-")}.tmTheme`,
  ];
  for (const p of tries) if (existsSync(p)) return p;
  return null;
}

export function loadTmTheme(path: string): TmThemeColors | null {
  if (!existsSync(path)) return null;
  const xml = readFileSync(path, "utf8");
  const dicts = parseDicts(xml);

  // Global bg/fg: first dict that has both (the "settings" root).
  const globals = dicts.find(d => d.background && d.foreground);
  if (!globals || !globals.background || !globals.foreground) return null;
  const bg = globals.background!;
  const fg = globals.foreground!;

  const scopes: Record<string, string> = {};
  for (const d of dicts) {
    if (!d.scope || !d.foreground) continue;
    for (const s of SCOPES_OF_INTEREST) {
      if (scopes[s.key]) continue;       // first match wins
      if (s.pattern.test(d.scope)) scopes[s.key] = d.foreground;
    }
  }

  // Apply base16changer's normalizeTmThemeContrast: any scope below AA against
  // the tmTheme bg gets replaced with the global fg.
  for (const k of Object.keys(scopes)) {
    if (contrastRatio(scopes[k], bg) < AA_THRESHOLD) scopes[k] = fg;
  }

  return { path, background: bg, foreground: fg, scopes };
}
