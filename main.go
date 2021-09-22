package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Super Sweet Website</h1>")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h2>FAQ</h2>"+
		"<ul><li>why</li></ul>")
}
func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "putting our best team together to find that for you")
}

func main() {
	r := httprouter.New()

	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)

	r.NotFound = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", r)
}
