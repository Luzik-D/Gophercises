package story

import (
	"html/template"
	"log"
	"net/http"
)

var StoryHTMLTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title> Adventure Story </title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Story}}
			<p>{{.}}</p>
		{{end}}
		<ul>
		{{range .Options}}
			<li><a href="/{{.Arc}}">{{.Text}}</a></li>
		{{end}}
		</ul>
	</body>
</html>`

type handler struct {
	s StoryMap
}

func NewHandler(s StoryMap) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Parse(StoryHTMLTemplate))

	err := tmpl.Execute(w, h.s["intro"])
	if err != nil {
		log.Fatal("Failed to execute html template: ", err)
	}
}
