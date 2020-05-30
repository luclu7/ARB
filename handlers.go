package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

/*type pkg struct {
	ID          int
	Name        string `gorm:"size:255"`
	Status int    `gorm:"type:int"`
} */

type pkg struct {
	ID     int
	Name   string
	Status int
	URL    string
}

type requestResponse struct {
	Type int
	Text string
}

var db *gorm.DB

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, it works !")
}

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

func handlerMarkBuildAsFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var currentPkg pkg
	db.Debug().First(&currentPkg, "Name = ?", name)
	db.Debug().Model(&currentPkg).Update("Status", 1)
	log.Println("/build/complete/mark/" + name)

	a, err := json.Marshal(requestResponse{Type: 200, Text: "Package was successfully marked as built."}) //get json byte array
	if err != nil {
		log.Panic(err)
	}
	n := len(a)        //Find the length of the byte array
	s := string(a[:n]) //convert to string
	fmt.Fprint(w, s)
}

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
