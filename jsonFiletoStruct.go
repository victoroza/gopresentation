package main

import (
	"fmt"
	"encoding/json"
	"io"
	"os"
	"log"
)
type Comp struct {
	Name string
	Year int
	URL string
}
// START OMIT
func main() {
	input, err := os.Open("companies.article")
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