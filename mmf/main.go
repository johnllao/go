package main

import "fmt"
import "os"

func main() {

	file, err := os.Open("")
	if err != nil {
		fmt.Println(err.Error())
	}
	_ = file
	
}