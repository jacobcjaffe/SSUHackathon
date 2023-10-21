package main

import (
	"context"
	"log"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

func PopulateBannedMap() map[string]bool {
	mp := make(map[string]bool)
	mp["food"] = true;
	mp["ingredient"] = true;
	mp["nutrition"] = true;
	mp["fruit"] = true;
	mp["staple"] = true;
	mp["plant"] = true;
	mp["produce"] = true;
	mp["wood"] = true;
	return mp
}

/// method to send the image in binary format to the Google Vision client, 
/// receives labels of the image contents in JSON format
func VisionTest(fileName string) (map[string]bool) {
	// TODO: should i copy the results into a string array, or not?
	ctx := context.Background()

	// creates a client
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("failed to connect to client: %v", err)
	}
	// defer is like a finally statement
	defer client.Close()

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// generate new image with format compatability with the client
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("failed to generate labeled image: %v", err)
	}

	// generate labels for the image
	labels, err := client.DetectLabels(ctx, image, nil, 20) 
	if err != nil {
		log.Fatalf("failed to generate labels: %v", err)
	}


	// copy the results into a string
	labelMap := make(map[string]bool, 20)
	for _, label := range labels {
		labelMap[strings.ToLower(label.GetDescription())] = true
	}
	label[

	// debug print to log
	log.Println("original labels: ")
	for str := range labelMap {
		log.Println(str)
	}
	log.Println()

	bannedMap := PopulateBannedMap()
	narrowedMap := NarrowVisionObjects(labelMap, bannedMap)
	// debugging after
	log.Println("narrowed labels: ")
	for str := range narrowedMap {
		log.Println(str)
	}
	return narrowedMap
}

/// function to narrow down the labels by removing popular key words that are irrelevant
func NarrowVisionObjects(labels map[string]bool, banned map[string]bool) map[string]bool{
	for label := range labels {
		for bannedWord := range banned {
			if strings.Contains(label, bannedWord) {
				delete(labels, label)
				break
			}
		}
	}
	return labels
}
