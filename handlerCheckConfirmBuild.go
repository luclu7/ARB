package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
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

// UUIDToString converts uuid.UUIDs to... strings
func UUIDToString(UUID uuid.UUID) string {
	return fmt.Sprintf("%s", UUID)
}

var db *gorm.DB

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, it works !")
}

func handlerRegisterURLs(w http.ResponseWriter, r *http.Request) {
	var request addURLRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Error(err)
	}

	var read pkg
	//if !db.First(&read, "UUID =", uuid).Where("Secret =", secret).RecordNotFound {
	chain := db.Where("UUID = ?", request.UUID)
	chain = chain.Where("secret = ?", request.Secret)
	chain.Find(&read)

	if !chain.Find(&read).RecordNotFound() {
		log.WithFields(log.Fields{
			"UUID": request.UUID,
			"URL":  request.URL,
		}).Info("URL added")

		db.Create(&File{UUID: request.UUID, URL: request.URL})
	}
}

func handlerGetURLs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UUID := vars["UUID"]

	rows, err := db.Model(&File{}).Where("UUID = ?", UUID).Select("ID, UUID, URL").Rows() // (*sql.Rows, error)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var urls []File
	for rows.Next() {
		var id int
		var UUID string
		var URL string
		rows.Scan(&id, &UUID, &URL)
		urls = append(urls, File{ID: id, UUID: UUID, URL: URL})
	}
	a, err := json.Marshal(urls) //get json byte array
	if err != nil {
		log.Panic(err)
	}
	n := len(a)        //Find the length of the byte array
	s := string(a[:n]) //convert to string
	fmt.Fprint(w, s)
}

func handlerMarkBuildAsFinished(w http.ResponseWriter, r *http.Request) {
	UUID := r.FormValue("uuid")
	Secret := r.FormValue("secret")
	var currentPkg pkg
	if !db.First(&currentPkg, "UUID = ?", UUID).RecordNotFound() {
		if currentPkg.Secret == Secret {
			db.Model(&currentPkg).Update("Status", 1)
			log.WithFields(log.Fields{
				"UUID":    UUID,
				"package": currentPkg.Name,
			}).Info("Build finished")

			a, err := json.Marshal(requestResponse{Type: 200, UUID: currentPkg.UUID, Text: "Package was successfully marked as built."}) //get json byte array
			if err != nil {
				log.Panic(err)
			}
			n := len(a)        //Find the length of the byte array
			s := string(a[:n]) //convert to string
			fmt.Fprint(w, s)
		}
	} else {
		a, err := json.Marshal(requestResponse{Type: 403, UUID: currentPkg.UUID, Text: "You don't have the right token to mark this build as finished."}) //get json byte array
		if err != nil {
			log.Panic(err)
		}
		n := len(a)        //Find the length of the byte array
		s := string(a[:n]) //convert to string
		fmt.Fprint(w, s)

	}
}
