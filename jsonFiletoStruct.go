package main

import (
	"fmt"
	"encoding/json"
	"io"
	"os"
	"log"
	"io/ioutil"
	// "bufio"
)
type Comp struct {
	Name string
	Year int
	URL string
}
func main() {
	
	d1 := []byte("{\"Name\": \"Google\", \"Year\": 1998, \"URL\": \"http://google.com/\"}\n{\"Name\": \"Facebook\", \"Year\": 2004, \"URL\": \"http://facebook.com/\"}\n{\"Name\": \"Apple\", \"Year\": 1976, \"URL\": \"http://apple.com/\"}\n{\"Name\": \"Microsoft\", \"Year\": 1975, \"URL\": \"http://microsoft.com/\"}")//, 
	err := ioutil.WriteFile("test.json", d1, 0644)
// START OMIT
	input, err := os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(input)
	for {
		var comp Comp
		err := decoder.Decode(&comp)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("%v\n", comp)
	}
}
// END OMIT