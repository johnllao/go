package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// DebugHandler custom handler
type DebugHandler struct {}

// NewDebugHandler returns instance of the DebugHandler
func NewDebugHandler() *DebugHandler {
	return &DebugHandler{}
}

func (h *DebugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	buffer := bytes.NewBufferString("")
	buffer.WriteString(fmt.Sprintf("r.URL.Path = '%s' \n ", r.URL.Path))
	
	w.Header().Set("Content-Type", "text/plain")
	w.Write(buffer.Bytes())
}

func main() {

	router := http.NewServeMux()
	router.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	//router.Handle("/public/", NewDebugHandler())

	server := http.Server{
		Addr : ":8080",
		Handler : router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("ERR %s", err.Error())
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {

	buffer := bytes.NewBufferString(indexHTML)

	w.Header().Set("Content-Type", "text/html")
	w.Write(buffer.Bytes())
}

var indexHTML = `<!DOCTYPE html>
<html>
<head>
<title>Title of the document</title>
</head>

<body>
The content of the document......
</body>

</html>`