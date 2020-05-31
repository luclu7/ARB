package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
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
	UUID   string `gorm:"type:string"`
	Name   string
	Status int
}

type File struct {
	UUID string `gorm:"type:string"`
	URL  string
}

type requestResponse struct {
	Type int
	Text string
}

func UUIDToString(UUID uuid.UUID) string {
	return fmt.Sprintf("%s", UUID)
}

var db *gorm.DB

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, it works !")
}

func handlerRegisterURLs(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("uuid")
	url := r.FormValue("url")

	db.Create(&File{UUID: uuid, URL: url})

	var currentPkg pkg
	db.Debug().First(&currentPkg, "UUID = ?", uuid)
	db.Debug().Model(&currentPkg).Update("Status", 1)
	db.Debug().Model(&currentPkg).Update("URL", url)
}

func handlerGetURLs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UUID := vars["UUID"]
	fmt.Println(db.Where("UUID = ?", UUID))
}

func handlerMarkBuildAsFinished(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("uuid")
	urls := r.FormValue("urls")
	var currentPkg pkg
	db.Debug().First(&currentPkg, "UUID = ?", uuid)
	db.Debug().Model(&currentPkg).Update("Status", 1)
	db.Debug().Model(&currentPkg).Update("URL", urls)
	log.Println("/build/complete/mark/" + uuid)

	a, err := json.Marshal(requestResponse{Type: 200, Text: "Package was successfully marked as built."}) //get json byte array
	if err != nil {
		log.Panic(err)
	}
	n := len(a)        //Find the length of the byte array
	s := string(a[:n]) //convert to string
	fmt.Fprint(w, s)
}
