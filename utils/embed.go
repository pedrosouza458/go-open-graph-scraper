package embed

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed websites.json
var dataJSON []byte

type WebsiteData struct {
	Website string `json:"website"`
	Logo    string `json:"logo"`
}

// GetWebsites returns the parsed website data from the embedded JSON.
func GetWebsites() ([]WebsiteData, error) {
	var websites []WebsiteData
	err := json.Unmarshal(dataJSON, &websites)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return websites, nil
}
