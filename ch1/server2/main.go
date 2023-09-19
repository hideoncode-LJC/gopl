package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Locker
var count = 0

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counthandler)

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counthandler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
	
}