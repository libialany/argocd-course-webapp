package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Henry 3!")
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
