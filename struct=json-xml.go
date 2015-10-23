package main

import (
	"fmt"
	"encoding/json"
	"log"
)
type Lang struct {
	Names string
	PresentationTopic string
	Date int
	URL string
}
// START OMIT
func main() {
	lang := Lang{"Victor Oza, Corey Hom, Michael McInerney", "Go",102315, "http://google.org/"}
	data,err := json.Marshal(lang)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("%s\n", data)
}
// END OMIT
