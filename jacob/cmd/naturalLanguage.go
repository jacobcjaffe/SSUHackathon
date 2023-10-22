package main

import (
	"context"
	"log"

	language "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/language/apiv1/languagepb"
)

func ParseIntoBuckets() {
	log.Println("parsing into buckets")
	ctx := context.Background()

	// initialize a client
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to connect to natural language client: %v\n", err)
	}
	defer client.Close()

	text := "apple peach"

	// make the request by initializing a langaugepb struct
	classification, err := client.ClassifyText(ctx, &languagepb.ClassifyTextRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content {
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		//EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("failed to classify the text %v\n", err)
	}

	log.Printf("Text: %v\n", text)
	pb := classification.GetCategories()
	for l := range pb {
		log.Println(l)
	}
}
