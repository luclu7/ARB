package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerCheckIfBuildFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UUID := vars["UUID"]
	var currentPkg pkg
	if !db.First(&currentPkg, "UUID = ?", UUID).RecordNotFound() {

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
	log.WithFields(log.Fields{
		"UUID": UUID,
	}).Info("Checked status of a container" + UUID)

}
