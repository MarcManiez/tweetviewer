package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)

	log.Println("listening...")
	http.ListenAndServe(":3000", nil)
}
