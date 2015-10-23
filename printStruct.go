package main

import "fmt"
type Lang struct {
	Names string
	PresentationTopic string
	Date int
	URL string
}
func main() {
	lang := Lang{"Victor Oza, Corey Hom, Michael McInerney", "Go",102315, "http://google.org/"}
    fmt.Printf("%v\n", lang)
}
