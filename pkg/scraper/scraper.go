package scraper

import (
	"fmt"
	"net/http"
	"net/url"

	"strings"

	"github.com/PuerkitoBio/goquery"
	embed "github.com/pedrosouza458/go-open-graph-scraper/utils"
)

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
	websites, err := embed.GetWebsites()
	if err != nil {
		fmt.Println("Error reading embedded JSON:", err)
		return "", fmt.Errorf("failed to read embedded websites.json: %v", err)
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
