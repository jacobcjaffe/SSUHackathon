package main

import (
	"net/http"
	"log"
)

// currently: main is working, talking to the google vision api
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/upload", GetImage)

	log.Print("Listening...\n")
	/*
	strArray := VisionTest("static/images/test1.jpg")
	recipeQuery(strArray)
	*/
	http.ListenAndServe(":3000", mux)
}
