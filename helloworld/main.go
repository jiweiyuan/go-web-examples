package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main()  {
	indexFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You are request in Path: %s", r.URL.Path)
	}
	helloFunc := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!")
	}

	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/hello", helloFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
