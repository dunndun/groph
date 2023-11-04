package graphiql

import (
	_ "embed"
	"html/template"
	"net/http"
)

var (
	//go:embed graphiql.html.tmpl
	html    string
	tmpl, _ = template.New("GraphiQL").Parse(html)
)

type Handler struct {
	Url     string         `json:"url"`
	Headers map[string]any `json:"headers,omitempty"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, h)
}
