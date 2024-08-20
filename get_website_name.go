package main

import (
	"net/url"
	"strings"
)

func getWebsiteName(rawurl string) (string, error) {
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
