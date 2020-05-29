package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type pkg struct {
	ID          int
	Name        string `gorm:"size:255"`
	BuildStatus int    `gorm:"type:int"`
}

type Build struct {
	ID     int
	Name   string
	Status int
	URL    string
}

var db *gorm.DB

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, it works !")
}

func handlerBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkgName := vars["package"]
	var currentPkg pkg
	db.First(&currentPkg)
	if !db.Debug().First(&currentPkg, "Name = ?", pkgName).RecordNotFound() {
		log.Println("Package", pkgName, "is already being built.")
		fmt.Fprintf(w, "Your package was already built.\n")

	} else {

		buildPackage(pkgName)
		db.Create(&pkg{Name: pkgName, BuildStatus: 0})
		fmt.Fprintf(w, "Your package is building! Please recheck later.\n")
	}
}

func handlerMarkBuildAsFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["name"]
	var currentPkg pkg
	db.First(&currentPkg)
	db.Model(&currentPkg).Update("BuildStatus", 1)
	log.Println("/build/complete/mark/" + id)
}

func handlerCheckIfBuildFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var currentPkg pkg
	db.Debug().First(&currentPkg, "Name = ?", name)
	fmt.Fprintf(w, "status: "+strconv.Itoa(currentPkg.BuildStatus))
	fmt.Println(currentPkg)
	log.Println("/build/complete/check/" + name)

}
