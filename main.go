package main

import (
	"fmt"
	"net/http"
	"portfolio/views"       
	"portfolio/controllers" 
	"github.com/gorilla/mux"
)
var (
	homeView 	*views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>I'm sorry bro. 404, not found.</h1>")	
}

func main() {	
	
	homeView = views.NewView("bootstrap", "views/home.html")
	contactView = views.NewView("bootstrap", "views/contact.html")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":3000", r)
}

func must(err error){
	if err != nil {
		panic(err)
	}
}