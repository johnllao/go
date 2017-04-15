package main

import "bytes"
import "flag"
import "io/ioutil"
import "net/http"
import "fmt"
import "os"
import "sync"
import "time"

// Stat struct
type Stat struct {
	ID int
	Elapsed time.Duration
	Error string
	Status int
}

var method = flag.String("method", "GET", "http method")
var url = flag.String("url", "", "url")
var concurrency = flag.Int("n", 1, "number of concurrent requests")

var lock = new(sync.Mutex)
var requestStats = make([]Stat, 0)
var wait = sync.WaitGroup{}

func main() {

	flag.Parse()

	if len(*url) == 0 {
		fmt.Printf("ERR Invalid URL")
		os.Exit(1)
	}

	wait.Add(*concurrency)

	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("ERR %v", err.Error())
		os.Exit(1)
	}

	for i := 1; i <= *concurrency; i++ {
		go send(i, body)
	}
	wait.Wait()

	for _, i := range requestStats {
		fmt.Println(i)
	}
}

func send(id int, body []byte) {

	defer wait.Done()

	stat := Stat{ ID: id, Elapsed: 0, Error: "" }

	req, err := http.NewRequest(*method, *url, bytes.NewReader(body))
	if err != nil {
		stat.Error = err.Error()
		lock.Lock()
		requestStats = append(requestStats, stat)
		lock.Unlock()
		return
	}

	client := http.Client{}

	start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		stat.Error = err.Error()
		lock.Lock()
		requestStats = append(requestStats, stat)
		lock.Unlock()
		return
	}
	
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		stat.Error = err.Error()
		lock.Lock()
		requestStats = append(requestStats, stat)
		lock.Unlock()
		return
	}

	stat.Elapsed = time.Since(start)
	stat.Status = res.StatusCode
	lock.Lock()
	requestStats = append(requestStats, stat)
	lock.Unlock()
	return
}