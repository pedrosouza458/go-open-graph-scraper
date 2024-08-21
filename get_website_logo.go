package main

import (
	"encoding/json"
	"fmt"
	"io"

	"os"
	"strings"
)

type Website struct {
	Website string `json:"website"`
	Logo    string `json:"logo"`
}

func getWebsiteLogo(url string) (string, error) {

	// Open the JSON file
	jsonFile, err := os.Open("websites.json")
	if err != nil {
		return "", fmt.Errorf("failed to open websites.json: %v", err)
	}
	defer jsonFile.Close()

	// Read the JSON file
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return "", fmt.Errorf("failed to read websites.json: %v", err)
	}

	// Parse the JSON file into a slice of Website structs
	var websites []Website
	if err := json.Unmarshal(byteValue, &websites); err != nil {
		return "", fmt.Errorf("failed to unmarshal websites.json: %v", err)
	}

	// Search for the website and return the logo if found
	for _, website := range websites {
		if strings.Contains(url, website.Website) {
			return website.Logo, nil
		}
	}

	// Return an error if no match is found
	return "", fmt.Errorf(url)
}
