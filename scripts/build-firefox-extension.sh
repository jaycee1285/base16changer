#!/usr/bin/env bash
# Package tools/firefox-extension/ into a loadable .xpi using 7z (system tool).
# Output: tools/firefox-extension/base16-accent.xpi
#
# Install in LibreWolf:
#   about:addons → cog icon → Install Add-on From File
# (LibreWolf permits unsigned extensions by default. Stock Firefox requires
#  Developer/Nightly or xpinstall.signatures.required=false in about:config.)

set -euo pipefail

repo_root="$(cd "$(dirname "$0")/.." && pwd)"
ext_dir="$repo_root/tools/firefox-extension"
out="$ext_dir/base16-accent.xpi"

if [[ ! -f "$ext_dir/manifest.json" ]]; then
  echo "extension source not found at $ext_dir" >&2
  exit 1
fi

rm -f "$out"

# .xpi is a zip with a renamed extension. Pack manifest.json + background.js
# only — don't include README/native-host (those live outside the .xpi).
( cd "$ext_dir" && 7z a -tzip "$out" manifest.json background.js >/dev/null )

echo "wrote $out"
ls -lh "$out"
