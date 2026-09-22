package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"strings"
) 

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs);

	http.HandleFunc("/api/library", func (w http.ResponseWriter, r *http.Request) {
		entries, err := os.ReadDir("./media")
		if err != nil {
			http.Error(w, "could not read media folder", http.StatusInternalServerError)
			return
		}

		var names []string
		for _, e := range entries {
			names = append(names, e.Name())
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(names)
	})

	http.HandleFunc("/stream/", func (w http.ResponseWriter, r *http.Request) {
		filename := strings.TrimPrefix(r.URL.Path, "/stream/")
		http.ServeFile(w, r, "./media/" + filename)
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

