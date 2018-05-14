package main

import (
	"io"
	"log"
	"net/http"
)

const version string = "2.0.1"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number
func versionHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()

    if err != nil {
        panic(err)
    }

	io.WriteString(w, "Hello " + name)
}

func main() {
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9090", nil)
}
