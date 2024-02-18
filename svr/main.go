package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	homePage, err := template.ParseFiles("htmx/mainPage/mainPage.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	homePage.Execute(w, nil)
}
func main() {

	//initialize router
	r := mux.NewRouter()
	//server mux
	r.HandleFunc("/", HomePageHandler)

	// Serve static files (CSS, JS, etc.)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)
	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
