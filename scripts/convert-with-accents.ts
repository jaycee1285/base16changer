#!/usr/bin/env bun
// Insert an `accent: "xxxxxx"` field into Gogh (.yml) and Base16 (.yaml) theme
// files, using the detection logic in scan-accents.ts.
//
// Logs ONLY non-conversions (anything that doesn't end with a fresh accent
// field on disk), grouped by category, so failures are easy to tail and iterate
// on.
//
// Usage:  nix develop -c bun run scripts/convert-with-accents.ts <dir> [--log <path>] [--force]
//   (a project hook blocks bare `bun run`; always invoke through nix develop)
//
//   --force  Also write `accent: <base0D>` for files that would otherwise be
//            skipped by the detection algorithm:
//              - base16 schemes whose base0D is already non-blue+saturated
//                (the "control group" — accent = base0D verbatim, made
//                explicit so every file is self-describing).
//              - gogh schemes where no non-blue saturated candidate exists
//                (e.g. grayscale themes — accent = base0D verbatim too;
//                no better option, but at least the field is present).
//
// Default log:  ./convert-with-accents.log  (line-buffered, tail -F friendly)

import { readdirSync, readFileSync, statSync, writeFileSync, appendFileSync, openSync, closeSync } from "node:fs";
import { extname, join, resolve } from "node:path";
import { detect } from "./scan-accents.ts";

type FailCat =
  | "parse_failure"
  | "unknown_structure"
  | "no_accent_found"
  | "base0d_already_non_blue_skip"
  | "already_has_accent"
  | "write_failure"
  | "read_failure"
  | "forced_base0d";  // --force path: explicit accent=base0D written

const argv = process.argv.slice(2);
const dirArg = argv.find(a => !a.startsWith("--"));
const logIdx = argv.indexOf("--log");
const logPath = resolve(logIdx >= 0 ? argv[logIdx + 1] : "convert-with-accents.log");
const force = argv.includes("--force");

if (!dirArg) {
  console.error("usage: bun run scripts/convert-with-accents.ts <dir> [--log <path>] [--force]");
  process.exit(2);
}
const dir = resolve(dirArg);

// truncate / create log
closeSync(openSync(logPath, "w"));

const counts: Record<FailCat, number> = {
  parse_failure: 0, unknown_structure: 0, no_accent_found: 0,
  base0d_already_non_blue_skip: 0, already_has_accent: 0,
  write_failure: 0, read_failure: 0, forced_base0d: 0,
};
let converted = 0, total = 0;

function logFail(cat: FailCat, file: string, reason: string) {
  counts[cat]++;
  appendFileSync(logPath, `[${cat}] ${file}\t${reason}\n`);
}

function hasAccentField(content: string): boolean {
  return /(^|\n)\s*accent\s*:\s*['"]?[#]?[0-9a-fA-F]{6}/i.test(content);
}

// Insert `accent: "xxxxxx"` into a file string. Returns the new content.
//   Gogh:    insert after the `name:` line if present, else at top.
//   Base16:  insert immediately before the `palette:` line.
function insertAccent(content: string, kind: "gogh" | "base16", hex: string): string {
  const bare = hex.replace(/^#/, "").toLowerCase();
  const line = `accent: "${bare}"`;
  const lines = content.split(/\r?\n/);
  if (kind === "base16") {
    const idx = lines.findIndex(l => /^palette\s*:/.test(l));
    if (idx >= 0) {
      lines.splice(idx, 0, line);
      return lines.join("\n");
    }
    // fallback: top
    return line + "\n" + content;
  }
  // gogh
  const idx = lines.findIndex(l => /^\s*name\s*:/i.test(l));
  if (idx >= 0) {
    lines.splice(idx + 1, 0, line);
    return lines.join("\n");
  }
  // fallback: after a leading `---` if present, else top
  if (lines[0]?.trim() === "---") {
    lines.splice(1, 0, line);
    return lines.join("\n");
  }
  return line + "\n" + content;
}

function walk(d: string): string[] {
  const out: string[] = [];
  for (const f of readdirSync(d)) {
    const p = join(d, f);
    if (statSync(p).isFile() && (extname(f) === ".yml" || extname(f) === ".yaml")) out.push(p);
  }
  return out.sort();
}

for (const path of walk(dir)) {
  total++;
  let content: string;
  try { content = readFileSync(path, "utf8"); }
  catch (e) { logFail("read_failure", path, (e as Error).message); continue; }

  if (hasAccentField(content)) {
    logFail("already_has_accent", path, "file already contains an accent: field");
    continue;
  }

  const d = detect(content);

  if (d.kind === "unknown") {
    logFail("unknown_structure", path, "neither gogh (color_01) nor base16 (palette.base00) markers found");
    continue;
  }
  if (d.kind === "parse_error") {
    logFail("parse_failure", path, d.reason);
    continue;
  }
  if (d.kind === "base16" && d.skipReason === "base0d_already_non_blue") {
    if (!force) {
      logFail("base0d_already_non_blue_skip", path, `base0D=${d.base0d} already non-blue + saturated`);
      continue;
    }
    // --force: write accent = base0D verbatim so the file is self-describing.
    if (!d.base0d) { logFail("parse_failure", path, "missing base0D under --force"); continue; }
    const next = insertAccent(content, "base16", d.base0d);
    try { writeFileSync(path, next); converted++; counts.forced_base0d++;
          appendFileSync(logPath, `[forced_base0d] ${path}\taccent=${d.base0d} (already non-blue base0D, made explicit)\n`); }
    catch (e) { logFail("write_failure", path, (e as Error).message); }
    continue;
  }
  if (!d.accent) {
    if (!force) {
      logFail("no_accent_found", path,
        `no non-blue saturated candidate (kind=${d.kind}, source=${d.source}, base0D=${d.base0d ?? "—"})`);
      continue;
    }
    // --force: fall back to base0D for grayscale / unsaturated themes.
    if (!d.base0d) { logFail("parse_failure", path, "missing base0D under --force"); continue; }
    const next = insertAccent(content, d.kind, d.base0d);
    try { writeFileSync(path, next); converted++; counts.forced_base0d++;
          appendFileSync(logPath, `[forced_base0d] ${path}\taccent=${d.base0d} (no non-blue candidate; base0D used as fallback)\n`); }
    catch (e) { logFail("write_failure", path, (e as Error).message); }
    continue;
  }

  const next = insertAccent(content, d.kind, d.accent);
  try {
    writeFileSync(path, next);
    converted++;
  } catch (e) {
    logFail("write_failure", path, (e as Error).message);
  }
}

const summary = [
  `# convert-with-accents summary`,
  `# dir=${dir}`,
  `# total=${total}  converted=${converted}`,
  ...Object.entries(counts)
    .filter(([, n]) => n > 0)
    .map(([k, n]) => `# ${k}=${n}`),
].join("\n") + "\n";
appendFileSync(logPath, "\n" + summary);
console.log(summary);
console.log(`log: ${logPath}`);
