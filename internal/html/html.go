package html

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func LoadHtmlPage(w http.ResponseWriter, r *http.Request, name string) {
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("frontend/%s", name)))
	tmpl.Execute(w, nil)
}

func LoadImg(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "imgs/")[1]
	http.ServeFile(w, r, fmt.Sprintf("frontend/imgs/%s", filename))
}

func LoadStaticFile(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "static/")[1]
	http.ServeFile(w, r, fmt.Sprintf("frontend/static/%s", filename))
}
