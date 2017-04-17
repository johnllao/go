package main

import "encoding/json"
import "flag"
import "fmt"
import "io/ioutil"
import "net/http"
import "log"
import "os"
import "runtime"

var port = flag.Int("port", 8080, "http port")

func main() {
	log.Printf("starting fileupload")
	log.Printf("environment")
	log.Printf("process id : %v", os.Getpid())
	log.Printf("args")
	log.Printf("port : %v", *port)

	
	fileHandler := http.StripPrefix("/public/", http.FileServer(http.Dir("../public/")))

	mux := http.NewServeMux()
	mux.Handle("/public/", fileHandler)
	mux.HandleFunc("/stats", stats)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/", index)

	server := http.Server{}
	server.Addr = fmt.Sprintf(":%v", *port)
	server.Handler = mux
	
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

func stats(w http.ResponseWriter, r *http.Request) {

	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)

	result := make(map[string]interface{})
	result["processId"] = os.Getpid()
	
	data, _ := json.MarshalIndent(result, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func upload(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Invalid http method. POST is expected")
		return
	}
	content, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	defer content.Close()

	body, err := ioutil.ReadAll(content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintf(w, "size : %v \n", len(body))
	fmt.Fprintln(w, string(body))
}