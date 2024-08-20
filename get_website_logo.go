package main

import "strings"

func getWebsiteLogo(urlImg string) (string, error) {
	if strings.HasPrefix(urlImg, "https://medium.com") {
		return "https://www.svgrepo.com/show/354057/medium-icon.svg", nil
	}
	if strings.HasPrefix(urlImg, "https://google.com") {
		return "https://www.svgrepo.com/show/380993/google-logo-search-new.svg", nil
	}
	if strings.HasPrefix(urlImg, "https://www.tabnews.com.br") {
		return "https://i.imgur.com/O5OOa3s.png", nil
	}
	return "", nil
}
