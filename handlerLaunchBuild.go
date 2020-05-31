package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["package"]
	var currentPkg pkg
	db.First(&currentPkg)
	if !db.Debug().First(&currentPkg, "Name = ?", name).RecordNotFound() {
		fmt.Println("pkg: ", currentPkg)
		if currentPkg.Status == 0 {
			log.Println("Package", name, "is already being built.")
			a, err := json.Marshal(requestResponse{Type: 200, Text: "Your package is already being built!"}) //get json byte array
			if err != nil {
				log.Panic(err)
			}
			n := len(a)        //Find the length of the byte array
			s := string(a[:n]) //convert to string
			fmt.Fprint(w, s)
		} else {
			a, err := json.Marshal(requestResponse{Type: 200, Text: "Your package is being built!"}) //get json byte array
			if err != nil {
				log.Panic(err)
			}
			n := len(a)        //Find the length of the byte array
			s := string(a[:n]) //convert to string
			fmt.Fprint(w, s)
			go buildPackage(name)
			db.Create(&pkg{Name: name, Status: 0})
		}
	} else {
		a, err := json.Marshal(requestResponse{Type: 200, Text: "Your package is being built!"}) //get json byte array
		if err != nil {
			log.Panic(err)
		}
		n := len(a)        //Find the length of the byte array
		s := string(a[:n]) //convert to string
		fmt.Fprint(w, s)
		go buildPackage(name)
		db.Create(&pkg{Name: name, Status: 0})
	}
}
