package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	"os"
	"io"
	"time"
)

// home page
func home(w http.ResponseWriter, r *http.Request) {
	files := []string {
		"./static/html/index.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalf("couldn't parse template files: %v", err)
	}

	err = ts.ExecuteTemplate(w, "home", nil)
	if err != nil {
		log.Fatalf("server error, %v", err)
	}
}

// getting image from the frontend, saves image to disk, saves 
func GetImage(w http.ResponseWriter, r *http.Request) {
	log.Println("received")
	
	// receive body of request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	/*
	fmt.Println("HERE IS THE BODY")
	fmt.Printf("%s", body)
	*/
	fmt.Println("parsing the form")
	r.ParseForm();

	/*
	for key, value := range r.Form {
		fmt.Printf("%s = %s", key, value)
	}
	*/


	f, err := os.CreateTemp("/home/jacobjaffe/", "tempfile.jpg")
	o, err := os.CreateTemp("/home/jacobjaffe/", "tempfile.jpg")
	outFile, err  := os.Create("tempfile.jpg")
	if err != nil {
		log.Fatalf("couldn't create temp file %v", err)
	}
	defer os.Remove(o.Name())
	defer os.Remove(f.Name())
	f.Write(body)
	os.Chmod(f.Name(), 0777)
	os.Chmod(o.Name(), 0777)

	fmt.Println("temp file name is " + f.Name())
	// use unix commands to remove the unnecessary form data
	
	time.Sleep(100)
	counter := 0
	byteIndex:=0
	for i:= 0; i < len(body); i++ {
		if body[i] == 0x0D {
			counter++
		}
		if counter == 4 {
			byteIndex = i;
			fmt.Printf("found it %d\n", byteIndex)
			break;
		}
	}
	dummy := body[byteIndex + 2: ]
	outFile.Write(dummy)

	//strVec := VisionTest("static/images/test3.webp")
	strVec := VisionTest(outFile.Name())
	recipes := recipeQuery(strVec)
	returnstr := make([]byte, 1)
	for i := 0; i < len(recipes); i++ {
		if recipes[i] == '\n' {
			returnstr = append(returnstr, '<', 'b', 'r', '/', '>')
		} else {
			returnstr = append(returnstr, recipes[i])
		}
	}
	w.Write([]byte(returnstr))

}

// retrieves the recipes
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getting recipes from chatgpt"))
}
