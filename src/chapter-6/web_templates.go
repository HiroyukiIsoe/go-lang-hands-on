package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

// Template for no-template.
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

// get target Template
func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html", "templates/head.html", "templates/foot.html")
	return tmps
}

// index handler.
func index(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Templste string
		Title string
		Message string
	}{
		Templste: "index",
		Title: "Index",
		Message: "This ia Top page.",
	}
	er := page("index").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

// hello handler
func hello(w http.ResponseWriter, rq *http.Request) {
	data := []string{
		"One", "Two", "Three",
	}

	item := struct {
		Title string
		Data []string
	}{
		Title: "Hello",
		Data: data,
	}
	er := page("hello").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func main() {
	// index handling.
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq)
	})
	http.ListenAndServe("", nil)
}
