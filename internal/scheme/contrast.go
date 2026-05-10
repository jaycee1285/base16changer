package scheme

// WCAG 2.1 sRGB relative-luminance + contrast-ratio math, plus a
// category-aware "fallback slot" picker. Mirrors scripts/wcag/contrast.ts and
// scripts/wcag/fallback.ts so behaviour stays in lockstep with the test
// harness under wcag-reports/.

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// AAThreshold is the WCAG 2.1 AA minimum contrast for normal text. base16changer
// applies it to every role (no large-text or UI-component carve-outs — see
// TASKBOARD-WCAG.md decision log).
const AAThreshold = 4.5

func srgbChannelToLinear(c float64) float64 {
	cs := c / 255.0
	if cs <= 0.03928 {
		return cs / 12.92
	}
	return math.Pow((cs+0.055)/1.055, 2.4)
}

// luminance returns the WCAG relative luminance of a 6-char hex (no '#').
// Returns NaN on a bad input.
func luminance(hex string) float64 {
	h := strings.TrimPrefix(strings.ToLower(hex), "#")
	if len(h) != 6 {
		return math.NaN()
	}
	r, err1 := strconv.ParseUint(h[0:2], 16, 8)
	g, err2 := strconv.ParseUint(h[2:4], 16, 8)
	b, err3 := strconv.ParseUint(h[4:6], 16, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		return math.NaN()
	}
	return 0.2126*srgbChannelToLinear(float64(r)) +
		0.7152*srgbChannelToLinear(float64(g)) +
		0.0722*srgbChannelToLinear(float64(b))
}

// ContrastRatio returns the WCAG 2.1 contrast ratio between two hex colours.
// Either argument may have a leading '#'.
func ContrastRatio(a, b string) float64 {
	la := luminance(a)
	lb := luminance(b)
	if math.IsNaN(la) || math.IsNaN(lb) {
		return math.NaN()
	}
	hi, lo := la, lb
	if lb > la {
		hi, lo = lb, la
	}
	return (hi + 0.05) / (lo + 0.05)
}

// passesAA returns true iff a and b clear the AA threshold against each other.
func passesAA(a, b string) bool {
	r := ContrastRatio(a, b)
	return !math.IsNaN(r) && r >= AAThreshold
}

// promoteAccent picks a replacement accent when the supplied accent fails AA
// on either base00 or base01. It scans chromatic slots and picks the one with
// the highest worst-case ratio across both surfaces; if no candidate clears
// AA on both, it picks the candidate that most improves the worst case
// relative to the original. Returns the original unchanged when no candidate
// strictly improves on it (so we never make legibility worse).
//
// Returns (newAccent, oldAccent, replaced).
func promoteAccent(accent string, p Colors) (string, string, bool) {
	if accent == "" {
		return accent, accent, false
	}
	origMin := minRatio(accent, p.Base00, p.Base01)
	// Already fine — both surfaces clear AA.
	if origMin >= AAThreshold {
		return accent, accent, false
	}
	// Priority order: warm/saturated colour slots first, blue last (since
	// the failing accent is usually base0D blue itself). Within the loop,
	// pick the candidate with the best min(r00, r01) that strictly beats
	// the original.
	type cand struct{ name, hex string }
	candidates := []cand{
		{"base08", p.Base08}, {"base09", p.Base09}, {"base0E", p.Base0E},
		{"base0F", p.Base0F}, {"base0B", p.Base0B}, {"base0C", p.Base0C},
		{"base0A", p.Base0A}, {"base0D", p.Base0D},
	}
	bestHex := accent
	bestMin := origMin
	for _, c := range candidates {
		if c.hex == "" || c.hex == accent {
			continue
		}
		m := minRatio(c.hex, p.Base00, p.Base01)
		if m > bestMin {
			bestMin = m
			bestHex = c.hex
		}
	}
	if bestHex != accent {
		return bestHex, accent, true
	}
	return accent, accent, false
}

// MinRatio returns the smaller of ContrastRatio(fg, a) and ContrastRatio(fg, b),
// or NaN if either input is invalid. Useful when a foreground must clear AA on
// two background surfaces (e.g. view bg + current_line bg).
func MinRatio(fg, a, b string) float64 {
	return minRatio(fg, a, b)
}

func minRatio(fg, a, b string) float64 {
	ra := ContrastRatio(fg, a)
	rb := ContrastRatio(fg, b)
	if math.IsNaN(ra) || math.IsNaN(rb) {
		return math.NaN()
	}
	if ra < rb {
		return ra
	}
	return rb
}

// warnAccentPromoted writes a one-liner to stderr explaining what got
// changed. The format is greppable for log aggregators.
func warnAccentPromoted(scheme, from, to string) {
	fmt.Fprintf(os.Stderr,
		"base16changer: accent auto-promoted scheme=%q from=#%s to=#%s reason=AA-fail-on-base00/01\n",
		scheme, from, to,
	)
}
