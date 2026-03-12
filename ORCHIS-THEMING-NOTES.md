# Orchis Theming Notes

This repo uses Orchis as the base GTK theme because a well-designed grey/white theme is a better starting point than the old flat Base16 GTK templates. The point is not to force every surface into pure scheme colors. The point is to keep the app/workspace coherent while still benefiting from Orchis's better widget design.

That said, the session that produced this note surfaced a few important constraints and fixes.

## What We Found

### 1. `GTK_THEME=Base16` was overriding everything

The main runtime bug was not in Orchis at first. It was the session environment forcing:

```sh
GTK_THEME=Base16
```

That caused GTK apps like Thunar to ignore the normal GTK theme selection flow and keep using the old `~/.themes/Base16` theme assets.

That explained why:

- Thunar still showed stale Solarized-style colors
- some browser-adjacent surfaces looked like they were inheriting old GTK colors
- changing the generated Orchis theme did not seem to matter

Once that override was removed, the live behavior started making sense again.

### 2. The old `Base16` GTK theme was stale

The older generated theme in `~/.themes/Base16/` still contained old colors, including Solarized Light values like:

- `#FDF6E3`

That stale theme was a real source of confusion during debugging because it was still reachable through the GTK environment override.

### 3. The Orchis generator was only partly Base16-aware

The original Orchis integration in this repo only fed Base16 colors into:

- accent/semantic colors
- some named palette entries

but it left the full neutral grey ramp hard-coded to the stock Orchis values.

That meant many GTK surfaces still looked like the original grey Orchis family even after changing schemes.

In practice, the theme looked "kind of themed" but not actually driven by the scheme's main background/base neutrals.

### 4. Orchis also has many literal white surfaces upstream

Even after fixing the neutral ramp, many compiled GTK surfaces still resolved to literal white or near-white because Orchis upstream uses:

- `#FFFFFF`
- `white`
- `rgba(255,255,255,...)`

in a lot of places.

So the issue was not only the generator. Some surfaces were hard-coded deeper in Orchis itself.

## Changes Made

### 1. Neutral ramp is now derived from Base16

The generator now exports a scheme-derived neutral ramp instead of using the old fixed grey ladder.

This means light and dark themes both pull more of their core surfaces from:

- `base00` through `base07`

instead of looking like the same stock grey Orchis variant with only accent changes.

### 2. Core Orchis surfaces were rebalanced

The following shared surface relationships were adjusted so they no longer collapse into the same bright value in light mode:

- `background`
- `surface`
- `base`
- `titlebar`

That matters because titlebar/base/background separation is visually important in file managers, dialogs, lists, and browsers.

### 3. Shared "white-like" surface is no longer literal white

The generated Orchis `$white` token was changed to use the top neutral from the active scheme rather than hard-coded `#FFFFFF`.

This reduces the number of places where a light theme forces pure white even when the selected scheme is cream, parchment, blue-grey, or something else.

### 4. Some shared libadwaita and GNOME app snippets were patched

The highest-impact shared bindings and a few obvious literal-white snippets were replaced with scheme-derived values.

This is not a full upstream Orchis rewrite. It is a pragmatic first pass aimed at:

- GTK3/GTK4 core surfaces
- titlebar/base/content separation
- removing the most obvious hard-coded white previews/snippets

## Current Model

The intended model for this repo is now:

1. Orchis provides the structural/widget design language.
2. Base16 provides the accent colors and the neutral family.
3. Light and dark remain separate variants.
4. The theme should feel like "Orchis designed around this scheme", not "flat Base16" and not "stock grey Orchis plus a new accent".

## Tradeoff

This is still a tradeoff, not a pure Base16 port.

That is deliberate.

For day-to-day use, a thoughtfully designed grey/white theme often makes a better work surface than the older flat Base16 GTK templates. The goal is to preserve that design quality while letting the chosen scheme matter in a real way.

So the right question is not:

- "Is every pixel directly mapped from Base16?"

It is:

- "Does the resulting GTK theme still feel intentionally designed, while honestly following the scheme?"

That is the standard this Orchis path should aim for.

## Remaining Work

There are still more literal light/dark values upstream in Orchis that could be normalized later.

Likely future cleanup:

- audit remaining `#FFFFFF` and `white` backgrounds in the Orchis SCSS
- audit hard-coded dark literals like `#000000` / `#202020`
- continue replacing local literals with scheme-aware surface tokens
- decide whether the old `~/.themes/Base16` output should still exist at all

## Practical Lesson

The debugging path here took time because all the individual pieces used to make sense:

- old Base16 theme output
- Orchis output
- GTK settings
- `GTK_THEME`
- browser chrome CSS
- browser theme API
- Sidebery custom/theme state

What broke things was not one absurd choice. It was too many theme authorities active at once.

Sometimes it just takes three hours.
