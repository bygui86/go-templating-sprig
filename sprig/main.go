package main

import (
	"os"
	// "html/template"
	"text/template"

	"github.com/Masterminds/sprig"
)

func main() {
	template1 := `{{- $s := "a b c d" -}}{{- $l := splitList " " $s -}}{{ index $l 1 }}`
	sprigParsing(template1)

	template2 := `
{{- $noValues := dict -}}
{{- $overrides := dict "foo" (dict "a" "b") -}}
{{- $common := dict "foo" (dict "c" "d") "bar" (dict "e" "f") -}}

initial state:
$noValues: {{ $noValues }}
$overrides: {{ $overrides }}
$common: {{ $common}}

{{- with merge $noValues $overrides $common }}

after merge:
$noValues: {{ $noValues }}
$overrides: {{ $overrides }}
$common: {{ $common}}

{{- $_ := set .bar "e" "qwerty" -}}
{{- end }}

after set:
$noValues: {{ $noValues }}
$overrides: {{ $overrides }}
$common: {{ $common}}
`
	sprigParsing(template2)
}

func sprigParsing(strTemplate string) {
	type Data struct{}
	data := Data{}

	tmpl, err := template.
		New("sprig").
		// Funcs(sprig.FuncMap()). // aka sprig.HtmlFuncMap()
		Funcs(sprig.TxtFuncMap()).
		Parse(strTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
