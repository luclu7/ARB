package main

import (
	log "github.com/sirupsen/logrus"
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
	db.AutoMigrate(&File{})

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/build/launch", handlerBuild).Methods("POST")
	r.HandleFunc("/build/complete", handlerMarkBuildAsFinished).Methods("POST")
	r.HandleFunc("/build/addURL", handlerRegisterURLs).Methods("POST")
	r.HandleFunc("/build/getURL/{UUID}", handlerGetURLs)
	r.HandleFunc("/build/check/{UUID}", handlerCheckIfBuildFinished)
	http.Handle("/", r)
	log.Info("Starting on http://0.0.0.0:8081...")
	http.ListenAndServe(":8081", r)
}
