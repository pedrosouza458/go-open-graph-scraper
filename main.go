package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	var baseURL string
	fmt.Scanln(&baseURL)

	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := goquery.NewDocumentFromResponse(resp)

	logo, _ := getWebsiteLogo(baseURL)
	name, _ := getWebsiteName(baseURL)
	imgURL, _ := doc.Find(`meta[property="og:image"]`).Attr("content")
	pageName, _ := doc.Find(`meta[property="og:title"]`).Attr("content")
	description, _ := doc.Find(`meta[property="og:description"]`).Attr("content")

	fmt.Println("Logo: " + logo)
	fmt.Println("Website name: " + name)
	fmt.Println("Image:" + imgURL)
	fmt.Println("Page Name:" + pageName)
	fmt.Println("Page Description: " + description)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
