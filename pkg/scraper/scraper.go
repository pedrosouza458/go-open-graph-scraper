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

/*
Get the website name by a url string provided, will return either
the name or empty string if it cannot find.
*/
func GetWebsiteName(rawurl string) (string, error) {
	parsedUrl, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	hostname := parsedUrl.Hostname()
	parts := strings.Split(hostname, ".")
	if len(parts) > 2 {
		return parts[1], nil
	} else if len(parts) == 2 {
		return parts[0], nil
	}

	return hostname, nil
}

/*
Get the website logo by a url string provided, will return either
the logo or empty string if it cannot find; check the full websites
json if logos in: https://github.com/pedrosouza458/go-open-graph-scraper/blob/main/utils/websites.json.
*/
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

/*
Get the website title by a url string provided accesing the og:title
of website url, will return either the title or empty string if it
cannot find.
*/
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

/*
Get the website page image by a url string provided accesing the og:image
of website url, will return either the image or empty string if it
cannot find.
*/
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

/*
Get the website page description by a url string provided accesing the og:description
of website url, will return either the description or empty string if it
cannot find.
*/
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

/*
Get the website page type object by a url string provided accesing
the og:type of website url, will return either the type object or
empty string if it cannot find.
*/
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

/*
Get the website locale by a url string provided accesing the og:locale
of website url, will return either the locale or empty string if it
cannot find.
*/
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

/*
Get the website video by a url string provided accesing the og:video
of website url, will return either the video url or empty string if it
cannot find.
*/
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
