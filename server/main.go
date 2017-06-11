package main

import "fmt"
import "net/http"
import "github.com/gorilla/mux"

// Contact contains the contact details 
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var contacts = map[int]Contact{
	1: Contact{ID: 1, Name: "John Lao"     , Email: "john.lao@gmail.com"},
	2: Contact{ID: 2, Name: "Elmie Antonio", Email: "elmie_0320@yahoo.com"},
}

func main() {
	
	fs := http.StripPrefix("/contacts/", http.FileServer(http.Dir("./public"))) 
	router := mux.NewRouter()
	router.Handle("/contacts/", fs)
	router.HandleFunc("/api/contacts", getcontacts)

	var server http.Server
	server.Addr = ":8080"
	server.Handler = router

	fmt.Println("started...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("%s \n", err.Error())
	}
}

func getcontacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "GET method expected")
		return
	}
}