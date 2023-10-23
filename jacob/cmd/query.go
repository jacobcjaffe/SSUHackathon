package main

import (
	"log"
	"context"
	"os"
	"strings"
	openai "github.com/sashabaranov/go-openai"
)

func recipeQuery(ingredients []string) string {
	// get the openai key
	data, err := os.ReadFile("/home/jacobjaffe/keys/oai.txt")
	if err != nil {
		log.Fatalf("unable to use openai key")
	}
	key := string(data[:])
	key = strings.TrimSuffix(key, "\n")
	log.Print(key)

	// create OpenAI client
	client := openai.NewClient(key)

	// base query to prompt the recipes
	query := "what are three individual meals I can make that use all or some of the " +
		"following ingredients: "
	for _, ingredient := range ingredients {
		query += ingredient + " "
	}
	log.Printf("here is the query: %v \n", query)
	// send the query to the OpenAI API 
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest {
			Model: openai.GPT3Dot5Turbo0613,
			Messages: []openai.ChatCompletionMessage {
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "you are a skilled chef that recommends dishes based on available ingredients",
				},
				{
					Role: openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("couldn't prompt for recipes: %v\n", err)
	} else {
		log.Println("query complete")
	}

	log.Println(resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content
}
