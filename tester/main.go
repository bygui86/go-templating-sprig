package main

import (
	"html/template"
	"os"

	"github.com/Masterminds/sprig"
)

func main() {
	test := `{{- $s := "a b c d" -}}{{- $l := splitList " " $s -}}{{ index $l 1 }}`

	type Data struct{}
	data := Data{}

	tmpl, err := template.
		New("sprig").
		Funcs(sprig.FuncMap()).
		Parse(test)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
