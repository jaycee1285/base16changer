package gtksourceview

import (
	"fmt"

	"github.com/jaycee1285/base16changer/internal/scheme"
	"github.com/jaycee1285/base16changer/internal/template"
)

// RenderV5Current renders a GtkSourceView 5 style-scheme XML for the active
// scheme. It is based on tinted-gtksourceview's v4 template, with the v5
// current-line-number behavior made explicit.
func RenderV5Current(s *scheme.Base16) (string, error) {
	data := s.ToMap()
	data["scheme-description"] = fmt.Sprintf("Base16 %s GtkSourceView 5 style scheme", s.Name)

	return template.RenderString(v5Template, data)
}

const v5Template = `<?xml version="1.0" encoding="UTF-8"?>
<!--
  Tinted GtkSourceView Base16 {{scheme-name}} Color Scheme

  {{scheme-name}} Color Scheme by {{scheme-author}}

  Based on tinted-gtksourceview by Rob Loach (https://robloach.net)
  https://github.com/robloach/tinted-gtksourceview
-->

<style-scheme id="base16-{{scheme-slug}}" _name="Base16 {{scheme-name}}" version="1.0">
    <author>{{scheme-author}}</author>
    <_description>{{scheme-description}}</_description>

    <!-- Base16 Palette -->
    <color name="base00" value="#{{base00-hex}}" />
    <color name="base01" value="#{{base01-hex}}" />
    <color name="base02" value="#{{base02-hex}}" />
    <color name="base03" value="#{{base03-hex}}" />
    <color name="base04" value="#{{base04-hex}}" />
    <color name="base05" value="#{{base05-hex}}" />
    <color name="base06" value="#{{base06-hex}}" />
    <color name="base07" value="#{{base07-hex}}" />
    <color name="base08" value="#{{base08-hex}}" />
    <color name="base09" value="#{{base09-hex}}" />
    <color name="base0A" value="#{{base0A-hex}}" />
    <color name="base0B" value="#{{base0B-hex}}" />
    <color name="base0C" value="#{{base0C-hex}}" />
    <color name="base0D" value="#{{base0D-hex}}" />
    <color name="base0E" value="#{{base0E-hex}}" />
    <color name="base0F" value="#{{base0F-hex}}" />

    <!-- Global Settings -->
    <style name="text"                          foreground="base05" background="base00" />
    <style name="selection"                     background="base02" />
    <style name="cursor"                        foreground="base05" />
    <style name="secondary-cursor"              foreground="base06" />
    <style name="current-line"                  background="base01" />
    <style name="line-numbers"                  foreground="base05" background="base00" />
    <style name="current-line-number"           foreground="base05" background="#rgba(0,0,0,0)" bold="true" />
    <style name="background-pattern"            background="base0F" />

    <!-- Bracket Matching -->
    <style name="bracket"                       foreground="#{{base0E-aa-hex}}" />
    <style name="bracket-match"                 foreground="#{{base0E-aa-hex}}" bold="true" />
    <style name="bracket-mismatch"              foreground="base08" background="base01" />

    <!-- Right Margin -->
    <style name="right-margin"                  foreground="base05" background="base01" />

    <!-- Search Matching -->
    <style name="search-match"                  foreground="base05" background="base02" />

    <!-- Comments -->
    <style name="def:comment"                   foreground="base05" />
    <style name="def:shebang"                   foreground="base08" bold="true" />
    <style name="def:doc-comment-element"       italic="true" />
    <style name="def:doc-comment"               italic="true" />

    <!-- Constants -->
    <style name="def:constant"                  foreground="#{{base0A-aa-hex}}" />
    <style name="def:character"                 foreground="#{{base0B-aa-hex}}" />
    <style name="def:string"                    foreground="#{{base0B-aa-hex}}" />
    <style name="def:special-char"              foreground="#{{base0B-aa-hex}}" />
    <style name="def:number"                    foreground="#{{base09-aa-hex}}" />
    <style name="def:floating-point"            foreground="#{{base09-aa-hex}}" />
    <style name="def:decimal"                   foreground="#{{base09-aa-hex}}" />
    <style name="def:base-n-integer"            foreground="#{{base09-aa-hex}}" />
    <style name="def:complex"                   foreground="#{{base09-aa-hex}}" />
    <style name="def:special-constant"          foreground="#{{base09-aa-hex}}" />
    <style name="def:boolean"                   foreground="#{{base09-aa-hex}}" />
    <style name="def:null-value"                foreground="#{{base09-aa-hex}}" />

    <!-- Identifiers -->
    <style name="def:identifier"                foreground="#{{base0E-aa-hex}}" />

    <!-- Functions -->
    <style name="def:function"                  foreground="#{{base0D-aa-hex}}" />
    <style name="def:function-name"             foreground="#{{base0D-aa-hex}}" />

    <!-- Builtin -->
    <style name="def:builtin"                   foreground="#{{base0C-aa-hex}}" />
    <style name="def:built-in-function"         foreground="#{{base0C-aa-hex}}" />

    <!-- Keywords -->
    <style name="def:keyword"                   foreground="#{{base0E-aa-hex}}" />

    <!-- Statements -->
    <style name="def:statement"                 foreground="#{{base09-aa-hex}}" />

    <!-- Types -->
    <style name="def:type"                      foreground="#{{base0B-aa-hex}}" />

    <!-- Markup -->
    <style name="def:emphasis"                  italic="true" />
    <style name="def:strong-emphasis"           foreground="#{{base09-aa-hex}}" />
    <style name="def:inline-code"               foreground="#{{base0D-aa-hex}}" />
    <style name="def:insertion"                 underline="single" />
    <style name="def:deletion"                  strikethrough="true" />
    <style name="def:link-text"                 foreground="#{{base0E-aa-hex}}" />
    <style name="def:link-symbol"               foreground="base01" bold="true" />
    <style name="def:link-destination"          foreground="#{{base0C-aa-hex}}" underline="single" />
    <style name="def:heading"                   foreground="#{{base0D-aa-hex}}" bold="true" />
    <style name="def:thematic-break"            foreground="base0F" />
    <style name="def:preformatted-section"      foreground="#{{base0D-aa-hex}}" />
    <style name="def:list-marker"               foreground="#{{base09-aa-hex}}" />

    <!-- Operators -->
    <style name="def:operator"                  foreground="#{{base0B-aa-hex}}" />

    <!-- Others -->
    <style name="def:preprocessor"              foreground="#{{base0E-aa-hex}}" />
    <style name="def:error"                     foreground="base08" bold="true" />
    <style name="def:note"                      foreground="base0E" bold="true" />
    <style name="def:net-address"               italic="true" underline="single" />
    <style name="def:warning"                   foreground="base0A" />
    <style name="def:reserved"                  foreground="base0E" />
    <style name="def:underlined"                underline="single" />

    <!-- C -->
    <style name="c:preprocessor"                foreground="#{{base0E-aa-hex}}" />
    <style name="c:type-keyword"                foreground="#{{base0E-aa-hex}}" />
    <style name="c:function-name"               foreground="#{{base0D-aa-hex}}" />

    <!-- Python -->
    <style name="python:string-conversion"      foreground="#{{base0C-aa-hex}}" />
    <style name="python:class-name"             foreground="#{{base0A-aa-hex}}" />

    <!-- YAML -->
    <style name="yaml:map-key"                  foreground="#{{base0C-aa-hex}}" />

    <!-- JSON -->
    <style name="json:keyname"                  foreground="#{{base0C-aa-hex}}" />

    <!-- GTK-DOC Overrides -->
    <style name="gtk-doc:type"                  foreground="base0E" />
    <style name="gtk-doc:function"              foreground="base0D" />
    <style name="gtk-doc:function-name"         bold="true" />
    <style name="gtk-doc:property-name"         bold="true" />
    <style name="gtk-doc:signal-name"           bold="true" />
    <style name="gtk-doc:parameter"             bold="true" />
    <style name="gtk-doc:constant"              foreground="base09" />
    <style name="gtk-doc:return"                bold="true" />
    <style name="gtk-doc:since"                 bold="true" />
    <style name="gtk-doc:deprecated"            bold="true" />
</style-scheme>
`
