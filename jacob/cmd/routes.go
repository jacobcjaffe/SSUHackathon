package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("found home"))
}

// getting image from the frontend
func GetImage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getting image binary from frontend"))
}
