package helpers

import (
	"html/template"
)

var templates = map[string]*template.Template{}

func LoadTemplates() {
	templates = map[string]*template.Template{
		"home": template.Must(
			template.ParseFiles(
				Config.GetString("directories.templates", "./views/")+"home.html",
				Config.GetString("directories.templates", "./views/")+"base.html",
			),
		),
	}
}
