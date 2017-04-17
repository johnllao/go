package main

import "fmt"
import "net/http"

func css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csshtml")
	fmt.Fprint(w, appCSS)
}

var appCSS = `body { font-family: Arial; font-size: 10pt; }`