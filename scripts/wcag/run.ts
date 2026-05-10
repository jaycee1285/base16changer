#!/usr/bin/env bun
// WCAG AA contrast suite for base16changer themes.
// Usage:
//   nix develop -c bun run scripts/wcag/run.ts <theme.yaml|.yml> [more …] [--out wcag-reports]
//
// One markdown report per input scheme is written to <out>/<slug>.md. An
// aggregate `_summary.md` lists every failing pair across every scheme so you
// can see baseline weaknesses at a glance.
//
// AA threshold is 4.5:1 for ALL roles (per project decision: "AA strict").

import { mkdirSync, writeFileSync, existsSync } from "node:fs";
import { resolve, basename, extname, dirname } from "node:path";
import { contrastRatio, AA_THRESHOLD } from "./contrast.ts";
import { loadPalette, slot, type Palette, type SlotKey } from "./palette.ts";
import { allPairs, type Pair } from "./targets.ts";
import { allStackPairs, resolveColor, type StackPair } from "./stacks.ts";
import { suggestFallback, type Suggestion } from "./fallback.ts";
import { loadTmTheme, findTmThemeFor, type TmThemeColors } from "./tmtheme.ts";

interface Result {
  group: "intra" | "stack";
  target: string;            // for intra: target name; for stack: stack name
  role: string;
  fgLabel: string;
  bgLabel: string;
  fgHex: string;
  bgHex: string;
  ratio: number;
  pass: boolean;
  notes?: string;
  suggestion?: Suggestion | null;
}

function runIntra(pal: Palette): Result[] {
  const out: Result[] = [];
  for (const p of allPairs() as Pair[]) {
    if (p.variant && p.variant !== pal.variant) continue;
    const fgHex = slot(pal, p.fg), bgHex = slot(pal, p.bg);
    if (!fgHex || !bgHex) continue;
    const r = contrastRatio(fgHex, bgHex);
    const pass = r >= AA_THRESHOLD;
    const sug = pass ? null : suggestFallback(pal, p.fg, p.bg, bgHex);
    out.push({
      group: "intra", target: p.target, role: p.role,
      fgLabel: p.fg, bgLabel: p.bg, fgHex, bgHex, ratio: r, pass,
      notes: p.notes, suggestion: sug,
    });
  }
  return out;
}

function runStacks(pal: Palette, tm: TmThemeColors | null): Result[] {
  const out: Result[] = [];
  for (const sp of allStackPairs() as StackPair[]) {
    if (sp.variant && sp.variant !== pal.variant) continue;
    const fg = resolveColor(sp.fg, pal, tm);
    const bg = resolveColor(sp.bg, pal, tm);
    if (!fg || !bg) {
      // tmTheme requested but unavailable / scope missing — record skip so
      // the report explains why.
      const which = !fg ? "fg" : "bg";
      const which2 = !fg ? sp.fg : sp.bg;
      const label = which2.kind === "tmtheme"
        ? `tm:${(which2 as any).label} (unresolved)`
        : "?";
      out.push({
        group: "stack", target: sp.stack, role: sp.role,
        fgLabel: fg ? fg.label : label,
        bgLabel: bg ? bg.label : label,
        fgHex: fg ? fg.hex : "—",
        bgHex: bg ? bg.hex : "—",
        ratio: NaN, pass: false,
        notes: (sp.notes ? sp.notes + " — " : "") + `skipped: ${which} unresolved (no tmTheme or missing scope)`,
        suggestion: null,
      });
      continue;
    }
    const r = contrastRatio(fg.hex, bg.hex);
    const pass = r >= AA_THRESHOLD;
    let sug: Suggestion | null = null;
    if (!pass && fg.slotKey) {
      sug = suggestFallback(pal, fg.slotKey, bg.slotKey, bg.hex);
    }
    out.push({
      group: "stack", target: sp.stack, role: sp.role,
      fgLabel: fg.label, bgLabel: bg.label,
      fgHex: fg.hex, bgHex: bg.hex,
      ratio: r, pass, notes: sp.notes, suggestion: sug,
    });
  }
  return out;
}

// --- Reporting -------------------------------------------------------------

function fmt(n: number): string { return Number.isFinite(n) ? n.toFixed(2) : "NaN"; }

function suggestionCell(s: Suggestion | null | undefined): string {
  if (!s) return "—";
  return `\`${s.swap}\` (${s.hex}, ${fmt(s.ratio)}:1)`;
}

function renderReport(pal: Palette, results: Result[], tm: TmThemeColors | null): string {
  const fails = results.filter(r => !r.pass);
  const passes = results.length - fails.length;
  const byTarget = new Map<string, Result[]>();
  for (const r of results) {
    const key = `${r.group}:${r.target}`;
    if (!byTarget.has(key)) byTarget.set(key, []);
    byTarget.get(key)!.push(r);
  }
  const sortedKeys = [...byTarget.keys()].sort();

  const lines: string[] = [];
  lines.push(`# WCAG AA report — ${pal.name}`);
  lines.push("");
  lines.push(`- **source**: \`${pal.source}\``);
  lines.push(`- **format**: ${pal.format}`);
  lines.push(`- **variant**: ${pal.variant}`);
  lines.push(`- **accent**: ${pal.accent}${pal.accentExplicit ? " (explicit)" : " (= base0D)"}`);
  lines.push(`- **threshold**: AA strict (≥ ${AA_THRESHOLD}:1 for all roles)`);
  lines.push(`- **tmTheme**: ${tm ? `\`${tm.path}\` (bg=${tm.background}, fg=${tm.foreground})` : "_not found — code-in-gtk4 stack pairs will be skipped_"}`);
  lines.push(`- **summary**: ${passes} pass / ${fails.length} fail / ${results.length} total`);
  lines.push("");

  // Palette dump for cross-reference.
  lines.push("## Palette");
  lines.push("");
  lines.push("| Slot | Hex |");
  lines.push("|------|-----|");
  for (const k of ["base00","base01","base02","base03","base04","base05","base06","base07",
                    "base08","base09","base0A","base0B","base0C","base0D","base0E","base0F","accent"] as SlotKey[]) {
    lines.push(`| ${k} | \`${slot(pal, k)}\` |`);
  }
  lines.push("");

  // Failures first — that's what the user actually iterates on.
  if (fails.length) {
    lines.push("## Failures");
    lines.push("");
    lines.push("| Group | Target | Role | fg | bg | Ratio | Suggested swap | Notes |");
    lines.push("|-------|--------|------|----|----|-------|----------------|-------|");
    for (const r of fails) {
      const note = r.notes ? r.notes.replace(/\|/g, "\\|") : "";
      lines.push(
        `| ${r.group} | ${r.target} | ${r.role.replace(/\|/g, "\\|")} | ` +
        `\`${r.fgLabel}\` ${r.fgHex} | \`${r.bgLabel}\` ${r.bgHex} | ` +
        `**${fmt(r.ratio)}:1** | ${suggestionCell(r.suggestion)} | ${note} |`
      );
    }
    lines.push("");
  } else {
    lines.push("## Failures");
    lines.push("");
    lines.push("_None — every tested pair meets AA._");
    lines.push("");
  }

  // Full result table per target/stack for forensic review.
  lines.push("## All pairs");
  lines.push("");
  for (const key of sortedKeys) {
    const [group, target] = key.split(":");
    const rs = byTarget.get(key)!;
    const tFails = rs.filter(r => !r.pass).length;
    lines.push(`### ${group} — ${target} (${rs.length - tFails}/${rs.length} pass)`);
    lines.push("");
    lines.push("| ✓ | Role | fg | bg | Ratio |");
    lines.push("|---|------|----|----|-------|");
    for (const r of rs) {
      lines.push(
        `| ${r.pass ? "✅" : "❌"} | ${r.role.replace(/\|/g, "\\|")} | ` +
        `\`${r.fgLabel}\` ${r.fgHex} | \`${r.bgLabel}\` ${r.bgHex} | ${fmt(r.ratio)}:1 |`
      );
    }
    lines.push("");
  }
  return lines.join("\n");
}

// Aggregate summary across all schemes — surfaces the *baseline* role/stack
// pairs that fail across multiple themes (i.e. things base16changer should
// auto-fix at parse time).
function renderSummary(perScheme: { pal: Palette; results: Result[] }[]): string {
  type Bucket = { fails: number; total: number; samples: { name: string; ratio: number; suggestion: string }[] };
  const buckets = new Map<string, Bucket>();
  for (const { pal, results } of perScheme) {
    for (const r of results) {
      const key = `${r.group}::${r.target}::${r.role}`;
      let b = buckets.get(key);
      if (!b) { b = { fails: 0, total: 0, samples: [] }; buckets.set(key, b); }
      b.total++;
      if (!r.pass) {
        b.fails++;
        b.samples.push({
          name: pal.name,
          ratio: r.ratio,
          suggestion: r.suggestion ? `${r.suggestion.swap}@${fmt(r.suggestion.ratio)}` : "no candidate",
        });
      }
    }
  }
  const ranked = [...buckets.entries()]
    .filter(([, b]) => b.fails > 0)
    .sort((a, b) => b[1].fails - a[1].fails || a[0].localeCompare(b[0]));

  const lines: string[] = [];
  lines.push("# WCAG AA — aggregate summary");
  lines.push("");
  lines.push(`Schemes tested: ${perScheme.length}`);
  lines.push(`Threshold: AA strict (≥ ${AA_THRESHOLD}:1)`);
  lines.push("");
  lines.push("Buckets are sorted by fail count. A pair that fails in many schemes is a strong");
  lines.push("candidate for an in-`base16changer` fallback rule (run on parse, before targets).");
  lines.push("");
  lines.push("| Fails / Tested | Group | Target | Role | Failing schemes (ratio → suggested swap) |");
  lines.push("|---|---|---|---|---|");
  for (const [key, b] of ranked) {
    const [group, target, role] = key.split("::");
    const samples = b.samples
      .map(s => `${s.name} (${fmt(s.ratio)} → ${s.suggestion})`)
      .join("; ");
    lines.push(`| ${b.fails}/${b.total} | ${group} | ${target} | ${role.replace(/\|/g, "\\|")} | ${samples.replace(/\|/g, "\\|")} |`);
  }
  lines.push("");

  // Recommendation block: any role failing in ≥ 50% of schemes deserves a
  // hard-coded fallback in base16changer.
  const halfBar = perScheme.length / 2;
  const universal = ranked.filter(([, b]) => b.fails >= halfBar);
  lines.push("## Recommended `base16changer` parse-time fallbacks");
  lines.push("");
  if (!universal.length) {
    lines.push("_No pair fails in a majority of tested schemes._");
  } else {
    lines.push("Pairs failing in ≥ 50% of tested schemes — these are the candidates for an");
    lines.push("auto-fallback rule in `internal/scheme/parse.go` (e.g. `ToMap()` could emit a");
    lines.push("derived `<role>-fallback-hex` map key).");
    lines.push("");
    lines.push("| Group | Target | Role | Most-frequent suggested swap |");
    lines.push("|-------|--------|------|------------------------------|");
    for (const [key, b] of universal) {
      const [group, target, role] = key.split("::");
      const swapCounts = new Map<string, number>();
      for (const s of b.samples) swapCounts.set(s.suggestion, (swapCounts.get(s.suggestion) ?? 0) + 1);
      const top = [...swapCounts.entries()].sort((a, b2) => b2[1] - a[1])[0]?.[0] ?? "—";
      lines.push(`| ${group} | ${target} | ${role.replace(/\|/g, "\\|")} | ${top} |`);
    }
  }
  return lines.join("\n");
}

function slugify(s: string): string {
  return s.toLowerCase().replace(/[^a-z0-9]+/g, "-").replace(/^-+|-+$/g, "") || "scheme";
}

// --- CLI -------------------------------------------------------------------

function parseArgs(argv: string[]): { inputs: string[]; out: string } {
  const inputs: string[] = [];
  let out = "wcag-reports";
  for (let i = 0; i < argv.length; i++) {
    const a = argv[i];
    if (a === "--out") { out = argv[++i]; continue; }
    inputs.push(a);
  }
  return { inputs, out };
}

if (import.meta.main) {
  const { inputs, out } = parseArgs(process.argv.slice(2));
  if (inputs.length === 0) {
    console.error(
      "usage: nix develop -c bun run scripts/wcag/run.ts <theme.yaml|.yml> [more …] [--out <dir>]\n" +
      "       (run 8 themes — 2 light + 2 dark of each format — for a baseline)"
    );
    process.exit(2);
  }
  const outDir = resolve(out);
  if (!existsSync(outDir)) mkdirSync(outDir, { recursive: true });

  const perScheme: { pal: Palette; results: Result[] }[] = [];
  for (const input of inputs) {
    const path = resolve(input);
    const ext = extname(path);
    if (ext !== ".yaml" && ext !== ".yml") {
      console.error(`SKIP ${path}: not .yaml/.yml`);
      continue;
    }
    let pal: Palette;
    try { pal = loadPalette(path); }
    catch (e) { console.error(`PARSE_FAIL ${path}: ${(e as Error).message}`); continue; }
    const tmPath = findTmThemeFor(pal.name);
    const tm = tmPath ? loadTmTheme(tmPath) : null;
    const results = [...runIntra(pal), ...runStacks(pal, tm)];
    const slug = slugify(pal.name) + (pal.format === "gogh" ? "-gogh" : "-b16");
    const reportPath = resolve(outDir, `${slug}.md`);
    writeFileSync(reportPath, renderReport(pal, results, tm) + "\n");
    const fails = results.filter(r => !r.pass).length;
    console.log(`${fails === 0 ? "PASS" : "FAIL"} ${pal.name} [${pal.variant}/${pal.format}]: ${fails}/${results.length} fail → ${reportPath}`);
    perScheme.push({ pal, results });
  }

  if (perScheme.length > 1) {
    const summaryPath = resolve(outDir, "_summary.md");
    writeFileSync(summaryPath, renderSummary(perScheme) + "\n");
    console.log(`SUMMARY ${perScheme.length} schemes → ${summaryPath}`);
  }
}
