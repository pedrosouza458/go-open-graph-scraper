package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var websitesJSON []byte

type Website struct {
	Website string `json:"website"`
	Logo    string `json:"logo"`
}

func GetWebsiteName(rawurl string) (string, error) {
	parsedUrl, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	hostname := parsedUrl.Hostname()

	// Split the hostname into parts
	parts := strings.Split(hostname, ".")

	// Handle common cases: www.example.com or example.com
	// If there are 3 parts, "www.example.com", take the second part ("example")
	// If there are 2 parts, "example.com", take the first part ("example")
	if len(parts) > 2 {
		return parts[1], nil
	} else if len(parts) == 2 {
		return parts[0], nil
	}

	return hostname, nil
}

func GetWebsiteLogo(url string) (string, error) {
	byteValue, err := os.ReadFile("websites.json")
	if err != nil {
		fmt.Println("Error reading embedded JSON:", err)
		return "", fmt.Errorf("failed to read embedded websites.json: %v", err)
	}

	// Parse the JSON file into a slice of Website structs
	var websites []Website
	if err := json.Unmarshal(byteValue, &websites); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return "", fmt.Errorf("failed to unmarshal websites.json: %v", err)
	}

	// Search for the website and return the logo if found
	for _, website := range websites {
		if strings.Contains(url, website.Website) {
			return website.Logo, nil
		}
	}

	// Return an error if no match is found
	return "", fmt.Errorf("no logo found for URL: %s", url)
}

func GetWebsitePageTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	title, _ := doc.Find(`meta[property="og:title"]`).Attr("content")
	return title, nil
}

func GetWebsiteImg(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	imgURL, _ := doc.Find(`meta[property="og:image"]`).Attr("content")
	return imgURL, nil
}

func GetWebsiteDescription(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	description, _ := doc.Find(`meta[property="og:description"]`).Attr("content")
	return description, nil
}
