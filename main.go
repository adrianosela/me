package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const defaultPort = ":80"

func main() {
	// when running on Google App Engine, the PORT
	// env variable is set by the runtime
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = defaultPort
	}

	// set in app.yaml for Google App Engine,
	// or in ENV when running locally
	filename := os.Getenv("FILENAME")
	var h http.Handler
	switch filepath.Ext(filename) {
	case ".md":
		h = mdHandler(filename)
	case ".pdf":
		h = fileHandler(filename)
	default:
		h = fileHandler(filename)
	}

	err := http.ListenAndServe(port, h)
	if err != nil {
		log.Fatal(err)
	}
}
