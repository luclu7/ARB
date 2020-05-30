package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&pkg{})

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/build/launch/{package}", handlerBuild)
	r.HandleFunc("/build/complete/mark/{name}", handlerMarkBuildAsFinished)
	r.HandleFunc("/build/complete/check/{name}", handlerCheckIfBuildFinished)
	r.HandleFunc("/upload", uploadFile)
	http.Handle("/", r)
	fmt.Println("Starting on http://0.0.0.0:8081...")
	http.ListenAndServe(":8081", r)
}
