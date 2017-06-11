package main

import "fmt"
import "net/http"

func main() {
	
	fs := http.StripPrefix("/contacts/", http.FileServer(http.Dir("./public")))
	router := http.NewServeMux()
	router.Handle("/contacts/", fs)

	var server http.Server
	server.Addr = ":8080"
	server.Handler = router

	fmt.Println("started...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("%s \n", err.Error())
	}
}