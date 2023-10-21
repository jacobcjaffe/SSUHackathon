package main

import (
	"net/http"
	"log"
)

// currently: main is working, talking to the google vision api
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Print("Listening...\n")
	VisionTest("static/images/test1.jpg")
	http.ListenAndServe(":3000", mux)
}
