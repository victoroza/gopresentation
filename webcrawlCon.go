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
func count(name, url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: %s\n", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, dt)
}
func do(f func(Comp)) {
	input, err := os.Open("companies.article")
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
	c := make(chan string)
	n := 0
	do(func(comp Comp){
		n++
		go count(comp.Name, comp.Url, c)
	})
	for i := 0; i < n; i++ {
		fmt.Print(<-c)
	}
	fmt.Printf("%.2fs total \n", time.Since(start).Seconds())
}

// END OMIT