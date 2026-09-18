package main

import (
	"fmt"
	"log"
	"net/http"
) 

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs);

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}

