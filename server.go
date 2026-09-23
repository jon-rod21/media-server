package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"strings"
	"database/sql"
	_ "modernc.org/sqlite"
) 

func main() {

	db, err := sql.Open("sqlite", "./media.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	syncMediaFolder(db)


	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs);

	http.HandleFunc("/api/library", func (w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT media_id, title, filename FROM media")
		
		if err != nil {

			log.Println("query failed for media", )
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

// populate the db with new files
func syncMediaFolder(db *sql.DB) {
	entries, err := os.ReadDir("./media")
	if err != nil {
		log.Println("could not read media folder:", err)
		return
	}

	for _, e := range entries {
		filename := e.Name()

		_, err := db.Exec(
			'INSERT OR IGNORE INTO media (title, filename) VALUES (?, ?)',
			filename, filename,
		)
		if err != nil {
			log.Println("insert failed for", filename, ":", err)
		}
	}
}
