package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

}