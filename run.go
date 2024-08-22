package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedrosouza458/go-open-graph-scraper/pkg/scraper"
)

func main() {

	var baseURL string
	fmt.Scanln(&baseURL)

	// Retrieve and display website name
	name, err := scraper.GetWebsiteName(baseURL)
	if err != nil {
		fmt.Println("Error fetching website name:", err)
	} else {
		fmt.Println("Website name: " + name)
	}

	// Retrieve and display website logo
	logo, err := scraper.GetWebsiteLogo(baseURL)
	if err != nil {
		fmt.Println("Error fetching logo:", err)
	} else {
		fmt.Println("Logo: " + logo)
	}

	// Retrieve and display website image
	imgURL, err := scraper.GetWebsiteImg(baseURL)
	if err != nil {
		fmt.Println("Error fetching image URL:", err)
	} else {
		fmt.Println("Image: " + imgURL)
	}

	// Retrieve and display website image
	video, err := scraper.GetWebsiteVideo(baseURL)
	if err != nil {
		fmt.Println("Error fetching video:", err)
	} else {
		fmt.Println("Video: " + video)
	}

	// Retrieve and display page title
	pageTitle, err := scraper.GetWebsitePageTitle(baseURL)
	if err != nil {
		fmt.Println("Error fetching page title:", err)
	} else {
		fmt.Println("Page Name: " + pageTitle)
	}

	// Retrieve and display page description
	description, err := scraper.GetWebsiteDescription(baseURL)
	if err != nil {
		fmt.Println("Error fetching page description:", err)
	} else {
		fmt.Println("Page Description: " + description)
	}

	// Retrieve and display page type
	page_type, err := scraper.GetWebsiteType(baseURL)
	if err != nil {
		fmt.Println("Error fetching page type:", err)
	} else {
		fmt.Println("Page type: " + page_type)
	}

	// Retrieve and display page type
	locale, err := scraper.GetWebsiteLocale(baseURL)
	if err != nil {
		fmt.Println("Error fetching page locale:", err)
	} else {
		fmt.Println("Page locale: " + locale)
	}

	// Keep the server running
	log.Fatal(http.ListenAndServe(":8080", nil))
}
