package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerCheckIfBuildFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var currentPkg pkg
	if !db.Debug().First(&currentPkg, "Name = ?", name).RecordNotFound() {

		a, err := json.Marshal(currentPkg) //get json byte array
		if err != nil {
			log.Panic(err)
		}
		n := len(a)        //Find the length of the byte array
		s := string(a[:n]) //convert to string
		fmt.Fprint(w, s)   //write to response
	} else {
		w.WriteHeader(http.StatusNotFound)
		a, err := json.Marshal(requestResponse{Type: 404, Text: "Package not found !"}) //get json byte array
		if err != nil {
			log.Panic(err)
		}
		n := len(a)        //Find the length of the byte array
		s := string(a[:n]) //convert to string
		fmt.Fprint(w, s)
	}
	log.Println("/build/complete/check/" + name)

}
