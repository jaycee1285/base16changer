// base16-accent — themes LibreWolf private windows with the current
// base16changer accent. Reads the accent fresh via native messaging on every
// private-window open, so theme changes apply immediately without an
// extension reload.
//
// Native messaging host: `base16accent` (see ../native-host/).
// Payload contract: { accent, private_accent, variant, scheme }.
//
// FF forces a dark UI in private windows; the accent is the only readable
// scheme-derived colour against those forced surfaces. `private_accent` is
// the same as `accent` unless the accent fails AA on FF's forced field bg
// (#42414d), in which case base16changer pre-falls-back to base07.

const FF_PRIV_BG       = "#1c1b22";
const FF_PRIV_TOOLBAR  = "#2b2a33";
const FF_PRIV_FIELD_BG = "#42414d";

async function readAccent() {
  try {
    const reply = await browser.runtime.sendNativeMessage("base16accent", {});
    if (reply && typeof reply === "object" && reply.private_accent) {
      return reply.private_accent;
    }
    if (reply && typeof reply === "object" && reply.accent) {
      return reply.accent;
    }
  } catch (e) {
    console.error("base16-accent: native messaging failed:", e);
  }
  return null;
}

async function applyToWindow(windowId) {
  const fg = await readAccent();
  if (!fg) return;
  await browser.theme.update(windowId, {
    colors: {
      frame:                FF_PRIV_BG,
      tab_background_text:  fg,
      toolbar:              FF_PRIV_TOOLBAR,
      toolbar_text:         fg,
      toolbar_field:        FF_PRIV_FIELD_BG,
      toolbar_field_text:   fg,
      toolbar_field_focus:  FF_PRIV_FIELD_BG,
      popup:                FF_PRIV_BG,
      popup_text:           fg,
      tab_line:             fg,
      tab_loading:          fg,
      icons:                fg,
    },
  });
}

browser.windows.onCreated.addListener(async (w) => {
  if (w.incognito) await applyToWindow(w.id);
});

// On install / browser startup, theme any already-open private windows.
browser.runtime.onStartup.addListener(async () => {
  const wins = await browser.windows.getAll();
  for (const w of wins) if (w.incognito) await applyToWindow(w.id);
});
browser.runtime.onInstalled.addListener(async () => {
  const wins = await browser.windows.getAll();
  for (const w of wins) if (w.incognito) await applyToWindow(w.id);
});
