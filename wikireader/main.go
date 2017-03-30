package main

import "bufio"
import "encoding/xml"
import "flag"
import "fmt"
import "log"
import "os"
import "strings"
import "time"

// SubLink struct
type SubLink struct {
	Type 	string		`xml:"linktype,attr"`
	Anchor	string		`xml:"anchor"`
	Link	string		`xml:"link"`
}

var filename = flag.String("file", "", "input file name")
var find = flag.String("find", "none", "text to find")

func main() {

	flag.Parse()

	fmt.Println("Wikipedia Reader")
	fmt.Println("version 1.0.0")
	fmt.Println()

	start := time.Now()

	log.Printf("Opening file...")
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("ERR Failed to open xml file. %v", err.Error())
	}
	defer file.Close()

	buff := bufio.NewReader(file)
	decoder := xml.NewDecoder(buff)

	log.Printf("Reading records...")
	count := 0
	for {

		token, err := decoder.Token()
		if err != nil && err.Error() == "EOF" {
			break
		}

		node, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		if node.Name.Local == "sublink" {

			link := new(SubLink)
			err = decoder.DecodeElement(link, &node)
			if err != nil {
				log.Printf("WARN Failed to decode element to SubLink %v", err.Error())
				continue
			}

			if strings.Contains(link.Anchor, *find) {
				count++
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("Total topics that contains '%v': %v", *find, count)
	log.Printf("Elapsed: %v", elapsed)
	fmt.Println("Bye!")
}