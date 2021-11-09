package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir = "layouts"
var flashRotator int = 0

type View struct {
	Template *template.Template
	Layout   string
}

type ViewData struct {
	Flashes map[string]string
	Data    interface{}
}

func NewView(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	vd := ViewData{
		Flashes: flashes(),
		Data:    data,
	}
	return v.Template.ExecuteTemplate(w, v.Layout, vd)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func flashes() map[string]string {
	flashRotator = flashRotator + 1
	if flashRotator%3 == 0 {
		return map[string]string{
			"warning": "You are about to exceed your plan limts!",
		}
	} else {
		return map[string]string{}
	}
}
