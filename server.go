package main

import (
	"fmt"
	"log"
	"net/http"
) 

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
