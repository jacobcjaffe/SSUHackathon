package main

import (
	"net/http"
)

// home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("found home"))
}

// getting image from the frontend, saves image to disk, saves 
func GetImage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getting image binary from frontend"))
}

// retrieves the recipes
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getting recipes from chatgpt"))
}
