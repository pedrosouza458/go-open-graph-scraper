package scraper

import (
	"encoding/json"
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

// Base struct for the JSON response
type OpenGraphResponse struct {
	Type    string      `json:"type"`
	Details interface{} `json:"details,omitempty"`
}

// Music type details
type MusicDetails struct {
	Duration    int      `json:"duration,omitempty"`
	Album       string   `json:"album,omitempty"`
	Musician    []string `json:"musician,omitempty"`
	ReleaseDate string   `json:"release_date,omitempty"`
}

// Video type details
type VideoDetails struct {
	Duration    int      `json:"duration,omitempty"`
	ReleaseDate string   `json:"release_date,omitempty"`
	Actor       []string `json:"actor,omitempty"`
	Director    []string `json:"director,omitempty"`
	Writer      []string `json:"writer,omitempty"`
	Tag         []string `json:"tag,omitempty"`
}

// No specific metadata details
type GenericDetails struct{}

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
		return "", nil
	}

	// Search for the website and return the logo if found
	for _, website := range websites {
		if strings.Contains(url, website.Website) {
			return website.Logo, nil
		}
	}

	// Return an error if no match is found
	return "", nil
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

func GetWebsiteType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", nil
	}

	// Find the `og:type` meta tag content
	ogType, _ := doc.Find(`meta[property="og:type"]`).Attr("content")

	// Create response object based on type
	response := OpenGraphResponse{
		Type: ogType,
	}

	// Determine details based on type
	switch {
	case ogType == "music.song":
		response.Details = MusicDetails{
			// Populate with relevant details if available
		}
	case ogType == "video.movie":
		response.Details = VideoDetails{
			// Populate with relevant details if available
		}
	default:
		response.Details = GenericDetails{}
	}

	// Marshal response to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(jsonData), nil
}

func GetWebsiteLocale(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	locale, _ := doc.Find(`meta[property="og:locale"]`).Attr("content")
	return locale, nil

}

func GetWebsiteVideo(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	video, _ := doc.Find(`meta[property="og:video"]`).Attr("content")
	return video, nil

}
