package main

import "flag"
import "net/http"
import "log"
import "regexp"
import "sync"
import "fmt"

var symURL = "http://www.batstrading.com/market_data/symbol_listing/csv"

type store struct {
	sync.Mutex
	m map[string]Company
}

var s = store{m:make(map[string]Company)}

func init() {
	flag.Parse()
}

func main() {

	download()
	start()
	
	log.Println("Bye!")
}

func download() {

	list := NewCompanyList().Download()

	w := new(sync.WaitGroup)
	q := make(chan int, 100)
	for c := range list {
		go save(q, w, c)
	}
	w.Wait()
}

func save(q chan int, w *sync.WaitGroup, c Company) {

	q <- 1
	w.Add(1)

	s.Lock()
	s.m[c.Symbol] = c
	s.Unlock()

	w.Done()
	<-q
}

func start() {

	router := http.NewServeMux()
	router.HandleFunc("/symbols", symbols)
	router.HandleFunc("/companies/", company)
	server := http.Server{
		Addr : ":8080",
		Handler : router,
	}
	log.Println("http started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ERR", err.Error())
	}
}

func company(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")

	regex, err := regexp.Compile("^/companies/(\\w+)$")
	if err != nil {
		fmt.Fprintf(w, "ERR: %s \n", err.Error())
		return
	}
	if !regex.MatchString(r.URL.Path) {
		fmt.Fprintf(w, "ERR: Invalid URL \n")
		return
	}
	comp := s.m[regex.FindStringSubmatch(r.URL.Path)[1]]
	fmt.Fprintf(w, "%s \n"               , comp.Symbol)
	fmt.Fprintf(w, "Name       : %s   \n", comp.Name)
	fmt.Fprintf(w, "Sector     : %s   \n", comp.Sector)
	fmt.Fprintf(w, "Industry   : %s   \n", comp.Industry)
	fmt.Fprintf(w, "Market Cap : %.2f \n", comp.MarketCap)
}

func symbols(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	for sym, company := range s.m {
		fmt.Fprintln(w, sym, company.Name)
	}
}
