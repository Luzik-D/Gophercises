package story

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

var CustomHTMLTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Story}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`

var StoryHTMLTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Story}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`

type HandlerOpts func(h *handler)

type handler struct {
	s      StoryMap
	t      *template.Template
	pathFn func(r *http.Request) string
}

func DefaultPathFn(r *http.Request) string {
	path := r.URL.Path

	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

func CustomPathFn(r *http.Request) string {
	path := r.URL.Path

	if path == "" || path == "/" {
		path = "/story/intro"
	}

	if ok := strings.Contains(path, "/story/"); !ok {
		log.Fatal("Unsupported path")
	}
	return path[7:]
}

func SetCustomPathFn(fn func(r *http.Request) string) HandlerOpts {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func SetTemplate(t *template.Template) HandlerOpts {
	return func(h *handler) {
		h.t = t
	}
}

func SetDefaultTemplate() HandlerOpts {
	return func(h *handler) {
		h.t = template.Must(template.New("").Parse(StoryHTMLTemplate))
	}
}

func NewHandler(s StoryMap, opts ...HandlerOpts) http.Handler {
	h := handler{s, nil, DefaultPathFn}
	SetDefaultTemplate()(&h)

	for _, opt := range opts {
		opt(&h)
	}

	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		http.Error(w, "Chapter not found...", http.StatusBadRequest)
	}
}
