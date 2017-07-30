package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Tomato Sauce")

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))
	router.HandleFunc("/api/help", helpHandle)

	server := http.Server{
		Addr    : ":8080",
		Handler : router,
	}

	fmt.Println("starting at port 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("ERR: %s \n", err.Error())
	}
}

func helpHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]interface{})
	data["appName"] = "Go (golang) AngularJS Template"
	data["version"] = "1.0.0"
	
	resp, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, err.Error())
		return
	}
	w.Write(resp)
}