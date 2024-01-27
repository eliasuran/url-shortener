package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getURL)
	http.HandleFunc("/set", setURL)
	log.Fatal(http.ListenAndServe((":8080"), nil))
}

func getURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println(r.URL)
		fmt.Fprintf(w, "Hello World!")
		return
	}
	fmt.Println("Only GET is allowed")
	return
}

func setURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println(r.URL)
		fmt.Fprintf(w, "Hello World!")
		return
	}
	fmt.Println("Only POST is allowed")
	return
}
