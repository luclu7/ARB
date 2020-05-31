package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

/*type pkg struct {
	ID          int
	Name        string `gorm:"size:255"`
	Status int    `gorm:"type:int"`
} */

type pkg struct {
	ID     int
	UUID   uuid.UUID `gorm:"type:uuid`
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

func handlerMarkBuildAsFinished(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("login")
	login := r.FormValue("login")
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
