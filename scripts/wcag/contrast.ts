// WCAG 2.1 sRGB relative luminance + contrast ratio.
// Formula: https://www.w3.org/TR/WCAG21/#dfn-relative-luminance
//          https://www.w3.org/TR/WCAG21/#dfn-contrast-ratio

export const AA_THRESHOLD = 4.5; // user requested AA-strict for ALL roles

export type RGB = readonly [number, number, number];

export function hexToRgb(hex: string): RGB | null {
  const s = hex.replace(/^#/, "").toLowerCase();
  if (!/^[0-9a-f]{6}$/.test(s)) return null;
  const v = parseInt(s, 16);
  return [(v >> 16) & 0xff, (v >> 8) & 0xff, v & 0xff];
}

export function rgbToHex([r, g, b]: RGB): string {
  const h = (n: number) => n.toString(16).padStart(2, "0");
  return `#${h(r)}${h(g)}${h(b)}`;
}

function srgbChannelToLinear(c: number): number {
  const cs = c / 255;
  return cs <= 0.03928 ? cs / 12.92 : Math.pow((cs + 0.055) / 1.055, 2.4);
}

export function luminance(rgb: RGB): number {
  const [r, g, b] = rgb;
  return (
    0.2126 * srgbChannelToLinear(r) +
    0.7152 * srgbChannelToLinear(g) +
    0.0722 * srgbChannelToLinear(b)
  );
}

export function contrastRatio(a: string, b: string): number {
  const ra = hexToRgb(a), rb = hexToRgb(b);
  if (!ra || !rb) return NaN;
  const la = luminance(ra), lb = luminance(rb);
  const [hi, lo] = la > lb ? [la, lb] : [lb, la];
  return (hi + 0.05) / (lo + 0.05);
}

export function passes(a: string, b: string, threshold = AA_THRESHOLD): boolean {
  const r = contrastRatio(a, b);
  return Number.isFinite(r) && r >= threshold;
}

// Simple linear sRGB blend (gogh interpolation approximation — base16changer
// itself uses Lab via go-colorful; we use sRGB linear which is close enough
// for contrast testing but won't match base16changer's runtime values exactly).
export function blendHex(a: string, b: string, t: number): string {
  const ra = hexToRgb(a)!, rb = hexToRgb(b)!;
  const mix = (x: number, y: number) =>
    Math.round(x + (y - x) * Math.max(0, Math.min(1, t)));
  return rgbToHex([mix(ra[0], rb[0]), mix(ra[1], rb[1]), mix(ra[2], rb[2])] as RGB);
}
