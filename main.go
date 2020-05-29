package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type pkg struct {
	name        string
	buildStatus string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, it works !")
}

func handlerBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkg := vars["package"]
	buildPackage(pkg)
	fmt.Fprintf(w, "Your package is building! Please recheck later.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/build/{package}", handlerBuild)
	http.Handle("/", r)
	fmt.Println("Starting on localhost:8081...")
	http.ListenAndServe(":8081", r)
}
