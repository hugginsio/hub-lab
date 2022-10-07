package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type GitHubStars []struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
	Archived    bool   `json:"archived"`
}

// Remove unnecessary fields from API output and write back to file
func main() {
	filePath := "stars/stars.json"
	file, readFileError := os.ReadFile(filePath)
	check(readFileError)

	starData := GitHubStars{}

	unmarshalError := json.Unmarshal([]byte(file), &starData)
	check(unmarshalError)

	json, marshalError := json.MarshalIndent(starData, "", "  ")
	check(marshalError)

	writeFileError := os.WriteFile(filePath, json, 0644)
	check(writeFileError)

	fmt.Println("stars.json updated successfully")
}

func check(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
