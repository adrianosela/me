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

	// when running on Google App Engine,
	// env FILENAME must be defined in app.yaml
	filename := os.Getenv("FILENAME")
	var h http.Handler
	switch filepath.Ext(filename) {
	case ".md":
		h = mdHandler(filename)
	default: // { ".html" , ".pdf" }
		h = fileHandler(filename)
	}

	log.Printf("serving %s on %s\n", filename, port)
	err := http.ListenAndServe(port, h)
	if err != nil {
		log.Fatal(err)
	}
}
