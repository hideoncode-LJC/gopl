package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler echoes the HTTP request.



func main() {

	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	log.SetPrefix("\033[34m[info]\033[0m]")

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8888", nil) ; err != nil {
		log.Fatal(err)
	} else {
		log.Println("服务器已经启动")
	}





}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm() ; err != nil {
		log.Println(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}