package targets

import (
	"strings"
	"testing"
)

func TestNormalizeTmThemeContrast(t *testing.T) {
	input := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<plist version="1.0">
<dict>
  <key>settings</key>
  <array>
    <dict>
      <key>settings</key>
      <dict>
        <key>background</key>
        <string>#202020</string>
        <key>foreground</key>
        <string>#ddc7a1</string>
      </dict>
    </dict>
    <dict>
      <key>name</key>
      <string>Comments</string>
      <key>settings</key>
      <dict>
        <key>foreground</key>
        <string>#5a524c</string>
      </dict>
    </dict>
  </array>
</dict>
</plist>`)

	got := string(normalizeTmThemeContrast(input, 4.5))
	if strings.Contains(got, "#5a524c") {
		t.Fatalf("low-contrast foreground was not normalized:\n%s", got)
	}
	if count := strings.Count(got, "#ddc7a1"); count != 2 {
		t.Fatalf("expected default foreground to appear twice, got %d:\n%s", count, got)
	}
}
