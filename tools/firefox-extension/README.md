# base16-accent (Firefox / LibreWolf)

Themes private windows with the current base16changer accent. Two paths:

## Path 1 — userChrome only (try this first)

base16changer already writes `~/.config/base16changer/librewolf/colors.css`,
which is your `userChrome.css`. As of T5, that file includes a
`:root[privatebrowsingmode="temporary"]` block that re-asserts the scheme's
accent over Firefox's forced private-window styling, with `!important`.

If your `toolkit.legacyUserProfileCustomizations.stylesheets` is `true` (it
already is if userChrome is being applied to normal windows), private windows
just work after the next theme apply. **Open a private window after a
`base16changer -path <scheme>` and check the URL bar.** If the text is
legible against the dark gray field, you're done; ignore the rest of this
file.

If LibreWolf's internal incognito CSS is winning the cascade fight (URL bar
still rendering forced-dark text), continue to Path 2.

## Path 2 — install the .xpi

Pre-built: `tools/firefox-extension/base16-accent.xpi` (1.5K).

```sh
# 1. Install the extension.
#    LibreWolf: about:addons → ⚙ → Install Add-on From File →
#               pick tools/firefox-extension/base16-accent.xpi
#    Firefox stock: requires xpinstall.signatures.required=false in
#                   about:config (Developer/Nightly only) OR self-signing.
#    LibreWolf permits unsigned .xpi installs out of the box.

# 2. Native messaging host — one-time symlink (same pattern as VaultAdd).
mkdir -p ~/.librewolf/native-messaging-hosts
ln -sf "$PWD/tools/firefox-extension/native-host/base16accent.json" \
       ~/.librewolf/native-messaging-hosts/base16accent.json
chmod +x tools/firefox-extension/native-host/base16-accent-host.sh
# For Firefox proper:
# ln -sf .../base16accent.json ~/.mozilla/native-messaging-hosts/base16accent.json

# 3. Verify.
ls ~/.config/base16changer/accent.json   # written by base16changer
nix develop -c base16changer -path "/home/john/.local/share/themes/Catppuccin Latte.yml"
# Open a private window in LibreWolf — the URL bar / toolbar text should
# render in the scheme's accent.
```

That's it. Nothing else to maintain. base16changer writes `accent.json` on
every theme change; the extension reads it fresh on every private-window
open. No extension reload, no browser restart.

## Rebuilding the .xpi

The `.xpi` is committed in the repo. If you change `manifest.json` or
`background.js`, regenerate it:

```sh
nix develop -c scripts/build-firefox-extension.sh
```

The script uses the system's 7zip (`7z a -tzip`) — `.xpi` is just a renamed
zip.

## Why not a pure-extension solution?

Reading a file outside the WebExtension sandbox (any local path) requires
**either** native messaging **or** the `file://` URL scheme, and Firefox
WebExtensions can't fetch `file://` URLs. So if you want the extension to
update without re-install on every theme change, native messaging is the
only path. The host setup is one symlink + one chmod — same pattern your
VaultAdd extension uses.

The userChrome path (Path 1) avoids the extension entirely. Whether it works
depends on Firefox version: if the cascade fight goes our way, you don't
need this extension at all.

## Files

- `manifest.json`, `background.js` — the WebExtension source.
- `base16-accent.xpi` — the built artifact you install.
- `native-host/base16accent.json` — host manifest (registered via the
  symlink above).
- `native-host/base16-accent-host.sh` — host script that reads accent.json
  and frames the response.

## payload contract

`~/.config/base16changer/accent.json` is written by base16changer on every
theme apply:

```json
{
  "accent": "#d20f39",
  "private_accent": "#df8e1d",
  "variant": "light",
  "scheme": "Catppuccin Latte"
}
```

`private_accent` is the same as `accent` unless `accent` fails WCAG AA
against Firefox's forced private-window field bg `#42414d`, in which case
base16changer pre-falls-back to the slot with the best worst-case ratio
across the two forced surfaces (typically a yellow / orange / lavender,
whichever the palette has).
