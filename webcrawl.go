package main

import (
	"fmt"
	"encoding/json"
	"io"
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)

type Comp struct {
	Name string
	Year int
	Url string
}
func count(name, url string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	fmt.Printf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}
func do(f func(Comp)) {
	d1 := []byte("{\"Name\": \"Google\", \"Year\": 1998, \"URL\": \"http://google.com/\"}\n{\"Name\": \"Facebook\", \"Year\": 2004, \"URL\": \"http://facebook.com/\"}\n{\"Name\": \"Apple\", \"Year\": 1976, \"URL\": \"http://apple.com/\"}\n{\"Name\": \"Microsoft\", \"Year\": 1975, \"URL\": \"http://microsoft.com/\"}")//, 
	err := ioutil.WriteFile("test.json", d1, 0644)
	input, err := os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(input)
	for {
		var comp Comp
		err := dec.Decode(&comp)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(comp)
	}
}


// START OMIT
func main() {
	start := time.Now()
	do(func(comp Comp){
		count(comp.Name, comp.Url)
	})
	fmt.Printf("%.2fs total \n", time.Since(start).Seconds())
}

// END OMIT