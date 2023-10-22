package main

import (
	"net/http"
	"log"
	"io"
)

func ScrapeRecipes(ingredients []string) []string {
	dummy := make([]string, 10)

	baseURL := "https://www.budgetbytes.com"
	resp, err := http.Get(baseURL)
	if err != nil {
		log.Fatalf("failed to query for recipe %v\n", err)
	}
	defer resp.Body.Close()
	
	// why does the body need to be read?
	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		log.Fatalf("failed to read body of query response %v\n", err)
	}

	// debug output body
	stringBody := string(body)
	log.Printf(stringBody)
	return dummy
}
