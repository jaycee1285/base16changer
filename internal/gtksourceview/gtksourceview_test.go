package gtksourceview

import (
	"strings"
	"testing"

	"github.com/jaycee1285/base16changer/internal/scheme"
)

func TestRenderV5Current(t *testing.T) {
	s := &scheme.Base16{
		Name:   "Test Scheme",
		Author: "Tester",
		Palette: scheme.Colors{
			Base00: "000000",
			Base01: "111111",
			Base02: "222222",
			Base03: "333333",
			Base04: "444444",
			Base05: "555555",
			Base06: "666666",
			Base07: "777777",
			Base08: "888888",
			Base09: "999999",
			Base0A: "aaaaaa",
			Base0B: "bbbbbb",
			Base0C: "cccccc",
			Base0D: "dddddd",
			Base0E: "eeeeee",
			Base0F: "ffffff",
		},
	}

	got, err := RenderV5Current(s)
	if err != nil {
		t.Fatalf("RenderV5Current returned error: %v", err)
	}

	checks := []string{
		`<style-scheme id="base16-test-scheme" _name="Base16 Test Scheme" version="1.0">`,
		`<color name="base00" value="#000000" />`,
		`<color name="base0A" value="#aaaaaa" />`,
		`<style name="current-line-number"           foreground="base05" background="#rgba(0,0,0,0)" bold="true" />`,
		`<style name="def:keyword"                   foreground="#eeeeee" />`,
	}
	for _, check := range checks {
		if !strings.Contains(got, check) {
			t.Fatalf("rendered XML missing %q\n%s", check, got)
		}
	}
}
