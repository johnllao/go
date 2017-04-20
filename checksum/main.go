package main

import "crypto/sha256"
import "flag"
import "fmt"
import "io"
import "os"

var file = flag.String("file", "", "path and filename")

func main() {

	flag.Parse()

	if *file == "" {
		fmt.Println("ERR Path and filename not set")
		os.Exit(1)
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Printf("ERR %s \n", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		fmt.Printf("ERR %s \n", err.Error())
		os.Exit(1)
	}
	
	fmt.Printf("%x\n", hash.Sum(nil))
}