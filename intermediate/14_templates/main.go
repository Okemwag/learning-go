package main

import (
	"bytes"
	"fmt"
	htmpl "html/template"
	ttmpl "text/template"
)

type PageData struct {
	Title string
	Items []string
	Note  string
}

func main() {
	data := PageData{
		Title: "Template Demo",
		Items: []string{"Go", "Templates", "Stdlib"},
		Note:  "<b>escaped in html/template</b>",
	}

	// text/template is for plain text generation.
	textTmpl := ttmpl.Must(ttmpl.New("report").Parse(
		`{{.Title}}
{{range .Items}}- {{.}}
{{end}}Note: {{.Note}}`,
	))

	var textOut bytes.Buffer
	if err := textTmpl.Execute(&textOut, data); err != nil {
		fmt.Println("text template error:", err)
		return
	}
	fmt.Println("text/template output:")
	fmt.Println(textOut.String())

	// html/template is for HTML output and escapes unsafe data automatically.
	htmlTmpl := htmpl.Must(htmpl.New("page").Parse(
		`<h1>{{.Title}}</h1><p>{{.Note}}</p>`,
	))

	var htmlOut bytes.Buffer
	if err := htmlTmpl.Execute(&htmlOut, data); err != nil {
		fmt.Println("html template error:", err)
		return
	}
	fmt.Println("html/template output:")
	fmt.Println(htmlOut.String())
}
