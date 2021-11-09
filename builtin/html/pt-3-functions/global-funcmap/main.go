package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

type ViewData struct {
	User User
}

type User struct {
	ID    int
	Email string
}

func main() {
	var err error
	testTemplate, err = template.
		New("hello.gohtml").
		Funcs(
			// PAY ATTENTION HERE
			template.FuncMap{
				"hasPermission": func(feature string) bool {
					return false
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

	user := User{
		ID:    1,
		Email: "jon@calhoun.io",
	}
	vd := ViewData{user}

	// PAY ATTENTION HERE
	err := template.
		/* 	WARNING

		Potential race condition!

		It should be noted here that if you don’t clone the template before calling Funcs that you can
		potentially run into a race condition where multiple web requests are all trying to set different
		FuncMaps for the template. The final result could be that a user gets access to something they
		shouldn’t have access to. This is possible for two reasons:
			- Web requests are handled in goroutines by default, so your server will automatically be processing
			  multiple requests at the same time.
			- We are adding a FuncMap with a closure that uses the user variable. In previous examples we passed
			  the user into the function so this race condition wasn’t possible.
		This is pretty easy to fix with a Clone, but it might be worth noting in your code not to remove the
		call to Clone.

		For this reason We need to clone the template before setting a user-specific FuncMap to avoid any potential race conditions.
		*/
		Must(testTemplate.Clone()).
		Funcs(
			template.FuncMap{
				"hasPermission": func(feature string) bool {
					if user.ID == 1 && feature == "feature-a" {
						return true
					}
					return false
				},
			},
		).Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
