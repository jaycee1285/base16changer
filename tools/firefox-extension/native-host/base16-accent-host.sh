#!/usr/bin/env bash
# base16-accent native messaging host. Reads the current accent that
# base16changer wrote to ~/.config/base16changer/accent.json and replies on
# stdout in the WebExtension native-messaging protocol:
#
#   <4-byte little-endian length><JSON payload>
#
# Stderr is ignored by the browser; we route any errors there. Stdout MUST
# contain only the framed payload.
#
# Install (one-time):
#   1. Copy this file to ~/.local/bin/base16-accent-host.sh and chmod +x.
#   2. Copy ../base16accent.json to ~/.librewolf/native-messaging-hosts/
#      (Firefox uses ~/.mozilla/native-messaging-hosts/).
#   3. Edit the copied JSON's "path" field if your install location differs.
#   4. Install the extension at ../ (load as temporary or sign + install).

set -euo pipefail

CONFIG="${BASE16_ACCENT_JSON:-$HOME/.config/base16changer/accent.json}"

# Read the request frame from the browser. We don't actually use the request
# body — the extension just sends `{}` — but the protocol requires us to
# consume the 4-byte length prefix and the body before responding.
if ! head -c 4 >/dev/null; then
  echo "base16-accent-host: failed to read request length" >&2
  exit 1
fi

# We can ignore the request body length because we always reply the same way.
# But we still need to drain whatever the browser sent so it doesn't block.
# (Browsers send small JSON; absorbing extras is harmless.)
cat >/dev/null &
DRAIN_PID=$!

if [[ ! -r "$CONFIG" ]]; then
  payload='{"error":"accent.json not found","path":"'"$CONFIG"'"}'
else
  # Compact the JSON onto a single line — the framing wants a byte length, and
  # jq is overkill here since base16changer writes well-formed JSON. tr is
  # enough to strip newlines for length accuracy.
  payload="$(tr -d '\r\n' <"$CONFIG")"
fi

len=${#payload}

# Emit 4-byte little-endian length prefix.
printf '\\x%02x\\x%02x\\x%02x\\x%02x' \
  $(( len        & 0xff )) \
  $(((len >>  8) & 0xff )) \
  $(((len >> 16) & 0xff )) \
  $(((len >> 24) & 0xff )) | xargs -0 printf '%b'

printf '%s' "$payload"

# Reap the drain background process if still alive.
kill "$DRAIN_PID" 2>/dev/null || true
wait "$DRAIN_PID" 2>/dev/null || true
