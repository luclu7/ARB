package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

type pkg struct {
	gorm.Model
	Name        string
	buildStatus bool
}

var db *gorm.DB

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, it works !")
}

func handlerBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkgName := vars["package"]
	db.Create(&pkg{Name: pkgName, buildStatus: false})
	id := 1
	buildPackage(pkgName, id)
	fmt.Fprintf(w, "Your package is building! Please recheck later.")
}

func handlerMarkBuildAsFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var currentPkg pkg
	db.First(&currentPkg, id)
	db.Model(&currentPkg).Update("buildStatus", true)
}
func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&pkg{})

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/build/{package}", handlerBuild)
	r.HandleFunc("/markcomplete/{id}", handlerMarkBuildAsFinished)
	http.Handle("/", r)
	fmt.Println("Starting on http://localhost:8081...")
	http.ListenAndServe(":8081", r)
}
