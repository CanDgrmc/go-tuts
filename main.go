package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// "" => string
// '' => character
var templates *template.Template
var router *mux.Router

type Profile struct {
	Name  string
	Value string
}

func createRouter() {
	router = mux.NewRouter()
}

func createEndpoints() {
	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/hello", helloHandler).Methods("GET")
}

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	createRouter()
	createEndpoints()
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "index.html", nil)

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	arr := []interface{}{Profile{"Can", "Test"}}
	p, err := json.MarshalIndent(arr, "", "	")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(p)
}
