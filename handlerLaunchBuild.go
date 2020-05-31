package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

func handlerBuild(w http.ResponseWriter, r *http.Request) {
	var request buildRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Error(err)
	}
	name := request.PackageName
	var currentPkg pkg
	db.First(&currentPkg)
	UUID := UUIDToString(uuid.New())
	a, err := json.Marshal(requestResponse{Type: 200, UUID: UUID, Text: "The build is being launched."})
	if err != nil {
		log.Panic(err)
	}
	n := len(a)        //Find the length of the byte array
	s := string(a[:n]) //convert to string
	fmt.Fprint(w, s)
	secret := randstr.Hex(16)
	go buildPackage(name, UUID, secret)
	db.Create(&pkg{Name: name, Status: 0, UUID: UUID, Secret: secret})
	log.WithFields(log.Fields{
		"UUID":    UUID,
		"package": name,
	}).Info("Build created")

}
