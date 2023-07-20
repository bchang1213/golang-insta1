package TemplateUtility

import (
	"go/build"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = ParseTemplates()
}

func RenderTemplate(file string, data interface{}, w http.ResponseWriter) {
	err := tpl.ExecuteTemplate(w, file, data)
	if err != nil {
		log.Println("LOGGED", err.Error())
		http.Error(w, "Internal server error: File Execution", http.StatusInternalServerError)
	}
}

/* Helper method to get all templates in templates folder.*/
func ParseTemplates() *template.Template {
	templ := template.New("")
	//templ.Funcs(template.FuncMap{"config": func(i string) string { return ConfigUtility.Get(i) }})
	err := filepath.Walk(build.Default.GOPATH+"/src/dep-ex/internal/client/templates",
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, ".gohtml") {
				_, err = templ.ParseFiles(path)
				if err != nil {
					log.Println(err)
				}
			}
			return err
		})
	if err != nil {
		panic(err)
	}
	return templ
}
