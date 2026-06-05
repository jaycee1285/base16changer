# Openbox XBM Notes

LabWC loads Openbox-style decoration assets from the selected theme directory.
For this app, that means the active theme name in `rc.xml` points here:

```text
/home/john/.local/share/themes/Orchis-Light-Compact/openbox-3/
/home/john/.local/share/themes/Orchis-Dark-Compact/openbox-3/
```

The `themerc` file controls geometry and colors. XBM files provide the button
masks/icons that those colors are painted onto.

## Put XBM Files Here

Light variant:

```text
/home/john/.local/share/themes/Orchis-Light-Compact/openbox-3/*.xbm
```

Dark variant:

```text
/home/john/.local/share/themes/Orchis-Dark-Compact/openbox-3/*.xbm
```

If a button icon looks wrong, missing, too OS-X-like, or semantically confused,
this is the first place to check.

## Useful Button Names

The common Openbox button masks are:

```text
close.xbm
iconify.xbm
max.xbm
max_toggled.xbm
shade.xbm
shade_toggled.xbm
desk.xbm
desk_toggled.xbm
```

Themes may also include extras such as `bullet.xbm` for menus.

## Theme Goal

Do not use the old red/yellow/green macOS window-control idea as the design
model. It fights Base16 palettes and breaks careful schemes like Gruvbox,
Solarized, and Flexoki.

The intended roles are:

```text
normal button:      foreground on normal title/button surface
hover button:       foreground on hover surface
selected/toggled:   foreground on selected/accent surface
close/destructive:  foreground on destructive red surface
```

So the XBM shape should be semantic and neutral:

```text
close     = destructive close shape
iconify   = minimize shape
max       = maximize shape
max_toggle = restore/unmaximize shape
```

The color comes from `themerc`; the XBM should not try to encode red/yellow/green
meaning by shape or by theme assumptions.

## After Changing XBM Files

Run a normal theme apply from the TUI:

```text
./base16changer
```

Then pick a scheme/wallpaper and apply. The app updates `rc.xml`, runs
`labwc -r`, and restarts ferritebar. If LabWC still shows old button masks,
restart LabWC or log out/in; some decoration assets may be cached more strongly
than plain colors.
