package main

import (
	"fmt"
	"log"
	"net/http"
)

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Query()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusTemporaryRedirect)
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contacts", contactsHandler)
	addr := "localhost:8080"
	fmt.Printf("Starting server at %s\n", addr)
	error := http.ListenAndServe(addr, nil)
	if error != nil {
		log.Fatal(error)
	}
}
