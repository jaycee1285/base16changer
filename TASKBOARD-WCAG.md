# TASKBOARD-WCAG.md

WCAG-driven legibility hardening for `base16changer`. This board is the
hand-off doc for the work surfaced by the suite at `scripts/wcag/run.ts` and
its reports under `wcag-reports/`. Treat it as the source of truth: any
agent picking this up mid-stream should read this top-to-bottom before
touching code.

---

## Decision log (binding)

These are John's corrections / confirmations, not suggestions. Don't relitigate.

- **Window controls**: GTK/Gnome only. Orchis macOS-style window-button
  hardcoded colours (`#fd5f51` / `#fdbe04` / `#38c76a`) are **out of scope**.
- **Orchis paths in use**: `~/.local/share/themes/orchis-theme-light` and
  `~/.local/share/themes/orchis-theme-dark`. The bundled `openbox-3/`
  sub-theme inside each is the openbox/labWC target. Internal references in
  `internal/orchis/` and `internal/targets/targets.go` to `Orchis-Light-Compact`
  / `Orchis-Dark-Compact` should already track the actual paths; if anything
  still points at the old names, fix as a side step.
- **Accent universality**: every `.yml` and `.yaml` should end up with an
  explicit `accent:` field, including the 15 currently-skipped
  "base0D-already-non-blue" files (write `accent = base0D` verbatim) and the
  2 outliers (`Lunaria Dark.yml`, `grayscale-light.yaml`) that fall back to
  base0D for lack of a better candidate.
- **base05, not base04**, as the swap target for failing `base03`-as-text
  roles. The fallback engine in `scripts/wcag/fallback.ts` already suggests
  base05 in nearly every case — earlier write-up was a typo.
- **Firefox theme.json / base16changer-theme.xpi**: should NOT be generated.
  Normal windows continue to use the system theme. Only private windows are
  themed, via a separate one-time-built `base16-accent` extension that reads
  the current accent through native messaging. See task §5 for mechanics.
- **Threshold**: AA strict (≥ 4.5:1) for every role. No carve-outs for
  large text or UI components.

---

## Status snapshot (when this board was written)

- Suite (`scripts/wcag/run.ts`) runs against 8 schemes (2 light + 2 dark per
  format). Per-scheme reports in `wcag-reports/<slug>.md`, aggregate at
  `wcag-reports/_summary.md`.
- Baseline fail counts:
  | scheme | variant | format | fails / total |
  |---|---|---|---|
  | Catppuccin Latte (.yml)        | light | gogh   | 74 / 156 |
  | Ayu Light                      | light | gogh   | 89 / 156 |
  | Tokyo Night Storm              | dark  | gogh   | 34 / 151 |
  | Kanagawa Wave                  | dark  | gogh   | 59 / 151 |
  | Catppuccin Latte (.yaml)       | light | base16 | 82 / 156 |
  | Rosé Pine Dawn                 | light | base16 | 83 / 156 |
  | Tokyo Night Terminal Storm     | dark  | base16 | 87 / 151 |
  | Everforest Dark Hard           | dark  | base16 | 25 / 151 |
- Three user-flagged failures verified numerically:
  1. Firefox private window `toolbar_field_text on toolbar_field` — 4/4
     light schemes fail at ~1.06–1.54 : 1 against FF's forced `#42414d`.
  2. Syntect heading on GTK-4 surfaces — fails 4/8 on `selection_bg`,
     1/8 on `view_bg/popover_bg/sidebar_bg` (Tokyo Night Terminal Storm
     has no `markup.heading` scope in its tmTheme — that's the NaN).
  3. Fuzzel `selected_fg on selected_bg` — 7/8 fail; stripe visibility
     `selected_bg vs fuzzel bg` — 8/8 fail.

---

## Tasks (numbered for scope; do in order)

Each task should reduce per-scheme fail counts. After each task, re-run the
suite (command at the bottom of this file) and confirm the numbers move in
the right direction. If they don't, stop and diagnose before proceeding.

### Task 1 — Accent universality sweep

**Files:**
- `scripts/convert-with-accents.ts` (modify)

**What:**
Add a `--force` flag that, when set, writes an explicit `accent:` field for
files currently logged under `base0d_already_non_blue_skip` and
`no_accent_found` (in those cases, accent = base0D verbatim, since there's
no better candidate). Update help text + summary to reflect the new mode.

**Acceptance:**
Re-running `convert-with-accents.ts /home/john/.local/share/themes --force`
produces 0 entries in the failure log under `base0d_already_non_blue_skip`
and `no_accent_found`. Every `.yml` / `.yaml` in the dir has an explicit
`accent:` line on disk.

**Risk:** Low. In-place edit, idempotent (`already_has_accent` catches
re-runs).

---

### Task 2 — Parse-time accent validation in `internal/scheme/parse.go`

**Files:**
- `internal/scheme/parse.go` (modify)
- new: `internal/scheme/contrast.go` (a Go port of `scripts/wcag/contrast.ts`
  + `scripts/wcag/fallback.ts`'s `suggestFallback`)

**What:**
After unmarshal, before returning from `Parse()`:
1. Compute `accent` value (the existing fallback to `base0D` if accent
   is empty stays as-is).
2. Compute `contrastRatio(accent, base00)` and `contrastRatio(accent, base01)`.
3. If either is < 4.5:1, scan `{base08, base09, base0E, base0F, base0B, base0C}`
   in that priority order; pick the first slot whose ratio passes both
   surfaces. Replace `s.Accent` with that slot's hex.
4. Emit a stderr line: `accent auto-promoted: scheme=<name> from=<old> to=<new> reason=AA-fail-on-base00/01`.

Don't refuse to load — auto-promote silently except for the warning.

**Acceptance:**
- Suite re-run: `accent_color on view_bg`, `…on window_bg`, `…on headerbar_bg`,
  `…on sidebar_bg`, `…on card_bg`, `tab_line on toolbar`, `mako border`,
  `openbox active border`, `fuzzel match`, `fuzzel border` all drop their
  fail counts substantially (expect roughly 4–6/8 → 0–2/8 each).
- Existing tests still pass.

**Risk:** Medium. `parse.go` is the hot path. The auto-promotion changes
the actual value rendered into templates, so the visual identity of some
themes shifts — that's the point, but verify by re-rendering one theme by
hand.

---

### Task 3 — Role reassignments in `internal/targets/templates.go`

**Files:**
- `internal/targets/templates.go` (modify)

**What — swap `base03 → base05` in these specific roles only:**
- `firefoxTemplate` userChrome / theme JSON: `popup_border`, `sidebar_border`,
  `field_border`, `toolbar_field_border`. Use `base05-hex`.
- `gtk3Template`: `text_color_disabled`. Use `base05-hex`.
- `openboxTemplate`: `window.inactive.label.text.color`,
  `menu.items.disabled.text.color`. Use `base05-hex`.

**What — `fuzzel selected_bg = base02` (separate change):**
- `fuzzelTemplate`: `selection` value should source from `base02-hex`,
  not `base03-hex`. This is the stripe-visibility fix.

**Do NOT change:**
- Kitty `color8` (bright black) — keep at `base03`. ANSI standard slot;
  apps like tmux expect it to be dimmer than fg.
- Any `base02 → base*` decorative borders (`gtk4 scrollbar_outline`,
  `gtk3 border`). The "no candidate" suggestion-engine result for those
  means they're inherently decorative; accept the low contrast.

**Acceptance:**
Suite re-run drops these summary entries from 8/8 to ≤ 1/8:
- `firefox popup_border on popup_bg`
- `firefox sidebar_border on sidebar_bg`
- `gtk3 disabled text on bg`
- `openbox inactive title text`, `openbox menu disabled`
- `fuzzel selected_bg vs fuzzel bg — stripe visibility`

**Risk:** Low. Cosmetic shifts. Visual diff in two test schemes (one light,
one dark) before merging.

---

### Task 4 — gtk-sourceview5 contrast normalization

**Files:**
- `internal/targets/templates.go` (modify) OR
- new: `internal/template/normalize.go` if a shared helper makes sense

**What:**
The tmTheme target already does this: any scope foreground with ratio < 4.5:1
against the tmTheme background gets replaced with the tmTheme's default fg.
Apply the same logic at gsv template render time: for each syntax color
about to be substituted (def:keyword, def:function, def:string, def:type,
def:number, def:operator, def:heading, def:link-text, def:link-destination,
def:builtin, def:constant), check ratio against `base00` AND `base01`
(current_line bg). If either fails, swap to `base05`.

The substitution should happen as a post-processing pass on the rendered
template string, OR by computing a `<role>-aa-fallback-hex` map key and
having the template reference it. Either is fine; the post-pass is simpler.

**Acceptance:**
Suite re-run drops `gsv keyword on current_line`, `gsv function on current_line`,
`gsv string on current_line` failures from 4–5/8 to ≤ 1/8.

**Risk:** Medium. Changes the actual rendered colours in code editors —
visually verifiable by opening a code file in a libadwaita IDE / gnome-text-editor.

---

### Task 5 — Firefox private-window via `base16-accent` extension

**Files (new artefacts):**
- `~/.config/base16changer/accent.json` — a tiny JSON file base16changer
  writes on every render. Contents: `{"accent": "#xxxxxx", "variant": "light|dark"}`
- A native-messaging host script + manifest (one-time install).
- The `base16-accent` WebExtension itself (one-time install).

**Files (base16changer-side changes):**
- `internal/targets/targets.go` (or wherever Firefox targets live):
  - **Remove** generation of `base16changer-theme.xpi` and
    `~/.config/base16changer/firefox/theme.json`. Normal windows go back to
    the system theme.
  - **Add** writing `~/.config/base16changer/accent.json` with the current
    accent value (the same value that ends up in `m["accent-hex"]`) plus
    the variant.
- Document the one-time setup steps in a README or `docs/firefox-private.md`.

**The extension (one-time build, lives outside the base16changer repo
or in a `tools/firefox-extension/` subdir):**

```js
// background.js
async function applyAccent(windowId) {
  const { accent, variant } = await browser.runtime.sendNativeMessage(
    "base16accent", {}
  );
  const fg = accent;
  // FF's forced private surfaces — use them as the bg, accent as the fg.
  await browser.theme.update(windowId, {
    colors: {
      frame:               "#1c1b22",
      tab_background_text: fg,
      toolbar:             "#2b2a33",
      toolbar_text:        fg,
      toolbar_field:       "#42414d",
      toolbar_field_text:  fg,
      popup:               "#1c1b22",
      popup_text:          fg,
      tab_line:            fg,
    },
  });
}

browser.windows.onCreated.addListener(async (w) => {
  if (w.incognito) await applyAccent(w.id);
});

// On install, theme any already-open private windows.
browser.windows.getAll().then((wins) =>
  wins.filter((w) => w.incognito).forEach((w) => applyAccent(w.id))
);
```

```json
// manifest.json
{
  "manifest_version": 2,
  "name": "base16-accent",
  "version": "0.1.0",
  "permissions": ["nativeMessaging", "windows", "theme"],
  "background": { "scripts": ["background.js"], "persistent": false },
  "browser_specific_settings": {
    "gecko": { "id": "base16-accent@base16changer.local" }
  }
}
```

**The native messaging host (one-time install):**

`~/.librewolf/native-messaging-hosts/base16accent.json`:
```json
{
  "name": "base16accent",
  "description": "Reads base16changer accent.json",
  "path": "/home/john/.local/bin/base16-accent-host.sh",
  "type": "stdio",
  "allowed_extensions": ["base16-accent@base16changer.local"]
}
```

`/home/john/.local/bin/base16-accent-host.sh`:
```sh
#!/usr/bin/env bash
# Native messaging protocol: 4-byte little-endian length prefix, then payload.
exec >&2
payload=$(cat ~/.config/base16changer/accent.json)
len=${#payload}
printf '%b' "$(printf '\\x%02x' $((len & 0xff)) $(((len>>8) & 0xff)) $(((len>>16) & 0xff)) $(((len>>24) & 0xff)))"
printf '%s' "$payload"
```

**Acceptance:**
- `base16changer-theme.xpi` is no longer generated by the tool.
- `~/.config/base16changer/accent.json` is written every theme change.
- A LibreWolf private window, opened *after* a theme change, has its
  toolbar_text / toolbar_field_text / tab_line / popup_text reading in the
  current accent color against FF's forced dark surfaces.
- Switching themes in base16changer and opening a NEW private window
  reflects the new accent without reloading the extension or restarting
  LibreWolf.

**Pre-flight contrast check (add to Task 2's accent validation):**
After accent auto-promotion, also verify the chosen accent passes 4.5:1
against `#42414d` AND `#1c1b22`. If not, fall back to `base07` for the
extension's purposes (this fallback only affects the `accent.json` written
for the extension; it doesn't override the GTK/openbox/fuzzel accent).
Practically: write `accent.json` as
`{"accent": "...", "private_accent": "...", "variant": "..."}` so the
extension can consume the private-safe variant separately if needed.

**Risk:** Medium-high. New artefacts, new install steps. The native
messaging path needs to be set up on John's system once. Test the
extension build with `web-ext run` before committing.

---

### Task 6 — Suite as regression harness

**Files:**
- (none — process step)

**What:**
After each of tasks 1–5, run the suite and diff `wcag-reports/_summary.md`
against the prior run to confirm the expected fail count drops. The expected
deltas:

| After task | Expected drop in summed fail count across 8 schemes |
|---|---|
| 1 | minimal (just adds explicit accent fields; doesn't change the rendering) |
| 2 | ~30 fails removed (all the accent-on-light-surface cases) |
| 3 | ~20 fails removed (the seven base03→base05 swaps + fuzzel stripe) |
| 4 | ~15 fails removed (gsv current_line cases) |
| 5 | 4 ff-private fails replaced with passes for light schemes; the dark-scheme controls stay as they were |

If a delta is significantly off, stop and inspect.

---

## Re-run command (regression harness)

```sh
nix develop -c bun run scripts/wcag/run.ts \
  "/home/john/.local/share/themes/Catppuccin Latte.yml" \
  "/home/john/.local/share/themes/Ayu Light.yml" \
  "/home/john/.local/share/themes/Tokyo Night Storm.yml" \
  "/home/john/.local/share/themes/Kanagawa Wave.yml" \
  "/home/john/.local/share/themes/catppuccin-latte.yaml" \
  "/home/john/.local/share/themes/rose-pine-dawn.yaml" \
  "/home/john/.local/share/themes/tokyo-night-terminal-storm.yaml" \
  "/home/john/.local/share/themes/everforest-dark-hard.yaml" \
  --out wcag-reports
```

---

## Open questions / things not decided yet

- Whether the Go port of `suggestFallback` should also live in
  `internal/scheme/contrast.go` or somewhere more shared. Cheapest is to
  put it next to `parse.go`.
- Whether the `base16-accent` extension lives in `tools/firefox-extension/`
  inside this repo or in its own repo. Lean toward `tools/` here for now;
  it's tightly coupled.
- Whether to ship the native-messaging host script alongside the extension
  or in `scripts/install-base16-accent.sh`. The latter is cleaner.
- What to do with `internal/scheme/gogh.go`'s Lab interpolation (uses
  `go-colorful`) when the suite uses sRGB linear. Not a blocker for any
  task here, but the suite's reported ratios for derived gogh slots
  (base01/02/04/06) are ~±0.1 off from base16changer's runtime values.
  If precision matters, port `colorful.BlendLab` to TS for the suite.

---

## Files this work touches (cheat sheet)

| File | Task | Nature |
|---|---|---|
| `scripts/convert-with-accents.ts` | 1 | add `--force` flag |
| `internal/scheme/parse.go` | 2 | accent auto-promotion |
| `internal/scheme/contrast.go` (new) | 2 | Go port of WCAG math + fallback |
| `internal/targets/templates.go` | 3, 4 | role reassignments + gsv normalization |
| `internal/targets/targets.go` | 5 | drop ff theme/xpi gen, write `accent.json` |
| `tools/firefox-extension/` (new) | 5 | extension source + native host |
| `scripts/install-base16-accent.sh` (new) | 5 | one-time installer |
| `wcag-reports/*.md` | 6 | regenerated after every task |

---

## DO NOT

- Don't relitigate the decision log without flagging to John first.
- Don't generate the Firefox theme.json or .xpi.
- Don't swap base03 to base04 anywhere — base05 is the right target.
- Don't touch Orchis macOS button colors.
- Don't change kitty's `color8` (ANSI bright black). Leave at base03.
- Don't add fallbacks for purely-decorative borders (gtk3 border, gtk4
  scrollbar_outline). They're decorative.
