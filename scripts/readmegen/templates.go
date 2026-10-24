package main

import (
	"embed"
	"strings"
	"text/template"
)

//go:embed templates/*.md.tmpl
var templatesFS embed.FS

var templates = template.Must(template.New("").Funcs(template.FuncMap{
	"trim":  func(s string) string { return strings.TrimRight(s, "\n") },
	"table": func(t table) string { return t.String() },
}).ParseFS(templatesFS, "templates/*.md.tmpl"))

func renderTemplate(name string, data any) (string, error) {
	var b strings.Builder
	if err := templates.ExecuteTemplate(&b, name, data); err != nil {
		return "", err
	}
	return strings.TrimRight(b.String(), "\n"), nil
}
