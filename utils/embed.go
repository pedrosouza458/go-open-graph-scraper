package embed

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed websites.json
var dataJSON []byte

// Define a single website's data structure
type WebsiteData struct {
	Website string `json:"website"`
	Logo    string `json:"logo"`
}

// ParsedData now represents a slice of WebsiteData
type ParsedData []WebsiteData

func GetData() (ParsedData, error) {
	var data ParsedData
	err := json.Unmarshal(dataJSON, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return data, nil
}
