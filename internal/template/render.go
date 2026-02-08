package template

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Render processes a mustache-like template with the given data
// Uses Go templates but with {{variable}} syntax compatibility
func Render(templatePath string, data map[string]string) (string, error) {
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("read template: %w", err)
	}

	return RenderString(string(content), data)
}

// RenderString processes an embedded template string
func RenderString(templateContent string, data map[string]string) (string, error) {
	// Convert mustache {{var}} to Go template {{.var}}
	// Also handle {{scheme-name}} -> {{index . "scheme-name"}}
	converted := convertMustacheToGo(templateContent)

	tmpl, err := template.New("base16").Parse(converted)
	if err != nil {
		return "", fmt.Errorf("parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("execute template: %w", err)
	}

	return buf.String(), nil
}

// convertMustacheToGo converts mustache-style {{var}} to Go template syntax
func convertMustacheToGo(content string) string {
	// Simple replacement: {{something}} -> {{index . "something"}}
	// This handles hyphenated keys like scheme-name, base00-hex
	result := content

	// Find all {{...}} patterns and convert them
	for {
		start := strings.Index(result, "{{")
		if start == -1 {
			break
		}
		end := strings.Index(result[start:], "}}")
		if end == -1 {
			break
		}
		end += start

		varName := strings.TrimSpace(result[start+2 : end])

		// Skip if it's already a Go template directive
		if strings.HasPrefix(varName, ".") ||
			strings.HasPrefix(varName, "if") ||
			strings.HasPrefix(varName, "range") ||
			strings.HasPrefix(varName, "end") ||
			strings.HasPrefix(varName, "index") {
			// Move past this match to avoid infinite loop
			result = result[:start] + "<<PROCESSED>>" + varName + "<<END>>" + result[end+2:]
			continue
		}

		// Convert to index syntax for hyphenated names
		replacement := fmt.Sprintf(`{{index . "%s"}}`, varName)
		result = result[:start] + replacement + result[end+2:]
	}

	// Restore processed markers
	result = strings.ReplaceAll(result, "<<PROCESSED>>", "{{")
	result = strings.ReplaceAll(result, "<<END>>", "}}")

	return result
}
