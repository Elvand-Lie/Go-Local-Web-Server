package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() error: %v", err)
		return
	}
	fmt.Fprintf(w, "Post successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name : %s", name)
	fmt.Fprintf(w, "address : %s", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "This page didn't exist.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "This method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello there.")
}

func main() {
	fmt.Printf("We're serving on port 8000.")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
