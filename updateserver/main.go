package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "newversion.exe")
	})

	err := http.ListenAndServe(":2233", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
