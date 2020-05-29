package main

import (
	"fmt"
	"github.com/gorilla/mux"
	docker "github.com/luclu7/ARB/docker"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, it works !")
}

func handlerBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(request)
	pkg := vars["package"]
	docker.buildPackage(pkg)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/build/{package}", handlerBuild)
	http.Handle("/", r)
	fmt.Println("Starting on localhost:8081...")
	http.ListenAndServe(":8081", r)
}
