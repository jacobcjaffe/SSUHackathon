package main

import (
	"context"
	"log"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

/// populate a list of banned words to narrow down the labels
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

/// populate a list of allowed words, only these labels can be used
func PopulateAllowedStringsMap() map[string]bool {
	mp := make(map[string]bool)

	// fruits
	mp["apple"] = true
	mp["banana"] = true
	mp["cherry"] = true
	mp["pear"] = true
	mp["peach"] = true

	// 
	mp["bread"] = true
	mp["flour"] = true
	mp["bean"] = true

	// dairy
	mp["milk"] = true
	mp["yogurt"] = true
	mp["butter"] = true

	// vegetables
	mp["lettuce"] = true
	mp["broccoli"] = true
	mp["peanut"] = true
	mp["onion"] = true
	mp["avacado"] = true
	mp["green onion"] = true
	mp["artichoke"] = true
	mp["carrot"] = true
	mp["chard"] = true
	mp["pepper"] = true
	mp["eggplang"] = true
	mp["bear"] = true
	
	// meat
	mp["steak"] = true
	mp["beef"] = true
	mp["chicken"] = true
	mp["sardine"] = true
	mp["fish"] = true

	// beverages
	mp["wine"] = true
	return mp
}

/// method to send the image in binary format to the Google Vision client, 
/// receives labels of the image contents in JSON format
func VisionTest(fileName string) ([]string) {
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


	var margin float32 = .75
	// copy the results into a string
	labelMap := make(map[string]bool, 20)
	for _, label := range labels {
		if (label.GetScore() >= margin) {
			labelMap[label.GetDescription()] = true
		}
		log.Printf("label: %v, Topicality: %f, " +
			"Score: %f", label.GetDescription(), label.GetTopicality(), label.GetScore())
	}

	// debug print to log
	log.Println("original labels: ")
	for str := range labelMap {
		log.Println(str)
	}
	log.Println()

	allowed := PopulateAllowedStringsMap()
	narrowedMap := NarrowVisionObjects(labelMap, allowed)
	// debugging after
	log.Println("narrowed labels: ")
	for str := range narrowedMap {
		log.Println(str)
	}

	strArray := make([]string, len(narrowedMap))
	idx := 0;
	for key, _ := range narrowedMap {
		strArray[idx] = key
		idx++;
	}
	return strArray
}

/// function to narrow down the labels by removing popular key words that are irrelevant
func NarrowVisionObjects(labels map[string]bool, allowed map[string]bool) map[string]bool{
	// TODO make this more efficient
	contains := false
	for label := range labels {
		contains = false
		for word := range allowed {
			if strings.Contains(strings.ToLower(label), word) {
				contains = true
				break
			}
		}
		if contains == false {
			delete(labels, label)
		}
	}
	return labels
}
