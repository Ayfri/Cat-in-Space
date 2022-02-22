package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	Client              *http.Client
	DefaultHandler      http.HandlerFunc
	ResourceDirectories map[string]string
	Templates           map[string]*template.Template
	Routes              []string
	Port                string
	tree                *template.Template
}

func (h *Handler) ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) {
	err := h.tree.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) HandleResourcesDir(dir string, dest string) {
	if h.ResourceDirectories == nil {
		h.ResourceDirectories = make(map[string]string)
	}
	h.ResourceDirectories[dir] = dir
	fileServer := http.FileServer(http.Dir(dir))
	http.Handle(dest, http.StripPrefix(dest, fileServer))
}

func (h *Handler) HandleTemplate(filename string) *template.Template {
	t := template.Must(template.ParseFiles(filename))
	if h.Templates == nil {
		h.Templates = make(map[string]*template.Template)
		h.tree = t
	}
	h.Templates[filename] = t
	return t
}

func (h *Handler) HandleTemplates(directory string) {
	templates, err := template.ParseGlob(directory + "/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	if h.Templates == nil {
		h.Templates = make(map[string]*template.Template)
		h.tree = templates
	}
	for _, t := range templates.Templates() {
		h.Templates[t.Name()] = t
	}
}

func (h *Handler) HandleRoute(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.Routes = append(h.Routes, route)
	log.Println(strings.Join(h.Routes, ","))
	http.HandleFunc(route, handler)
}

func (h *Handler) Start(port string) error {
	h.Port = port
	if h.DefaultHandler == nil {
		return http.ListenAndServe(port, nil)
	}
	return http.ListenAndServe(port, h.DefaultHandler)
}
