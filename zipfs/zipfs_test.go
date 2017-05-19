package zipfs

import "fmt"
import "net/http"
func ExampleZipFileSystem() {

	fmt.Println("ZIP File System")
	fmt.Println("(c) johnllao.com 2017")

	router := http.NewServeMux()
	router.Handle("/public/", http.FileServer(NewZipFileSystem("public.zip")))

	server := http.Server{
		Addr   : ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}