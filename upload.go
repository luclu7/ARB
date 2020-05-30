package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["name"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'name' is missing")
		return
	}

	name := keys[0]

	file, err := os.Create(getEnv("OUT_FOLDER", "./out/") + name)
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}
