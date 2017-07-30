package main

import "encoding/json"
import "fmt"
import "net/http"
import "strconv"
import "sync"
import "github.com/gorilla/mux"

// Contact contains the contact details 
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var contacts = make([]*Contact, 0)
var contactsMap = make(map[int]*Contact)

var lock sync.Mutex

func init() {
	var c *Contact

	c = &Contact{ ID:1, Name:"Bill Gates", Email:"bill@microsoft.com" }
	contacts = append(contacts, c)
	contactsMap[c.ID] = c

	c = &Contact{ ID:2, Name:"Elon Musk", Email:"elon@tesla.com" }
	contacts = append(contacts, c)
	contactsMap[c.ID] = c

}

func main() {
	
	fs := http.StripPrefix("/contacts/", http.FileServer(http.Dir("./public"))) 
	router := mux.NewRouter()
	router.PathPrefix("/contacts/").Handler(fs)
	router.HandleFunc("/api/contacts/{id}", getcontact).Methods(http.MethodGet)
	router.HandleFunc("/api/contacts", getcontacts).Methods(http.MethodGet)
	router.HandleFunc("/api/contacts", addcontact).Methods(http.MethodPost)

	var server http.Server
	server.Addr = ":8080"
	server.Handler = router

	fmt.Println("started...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("%s \n", err.Error())
	}
}

func getcontacts(w http.ResponseWriter, r *http.Request) {

	// marshal the struct into json
	result, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		w.Header().Set("Content-Type:", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	// write the json result into the response
	if _, err = w.Write(result); err != nil {
		w.Header().Set("Content-Type:", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}
}

func getcontact(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.Header().Set("Content-Type:", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
	}

	c, ok := contactsMap[id]
	if !ok {
		w.Header().Set("Content-Type:", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Contact ID %d not found \n", id)
		return
	}

	// marshal the struct into json
	result, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		w.Header().Set("Content-Type:", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	// write the json result into the response
	if _, err = w.Write(result); err != nil {
		w.Header().Set("Content-Type:", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}
}

func addcontact(w http.ResponseWriter, r *http.Request) {
	
}