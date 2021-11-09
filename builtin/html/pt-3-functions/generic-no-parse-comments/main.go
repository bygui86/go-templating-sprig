package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

func main() {
	var err error
	testTemplate, err = template.
		New("hello.gohtml").
		Funcs(
			// PAY ATTENTION HERE
			template.FuncMap{
				"htmlSafe": func(html string) template.HTML {
					return template.HTML(html)
				},
			},
		).
		ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := testTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
