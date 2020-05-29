package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, it works !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)
	fmt.Println("Starting on localhost:8081...")
	http.ListenAndServe(":8081", r)
}
