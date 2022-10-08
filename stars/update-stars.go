package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type GitHubStars []struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Archived    bool   `json:"archived"`
}

// Remove unnecessary fields from API output and write back to file
func main() {
	starsJson := "stars/stars.json"
	file, readFileError := os.ReadFile(starsJson)
	check(readFileError)

	starData := GitHubStars{}

	unmarshalError := json.Unmarshal([]byte(file), &starData)
	check(unmarshalError)

	markdownContents := []string{"## GitHub stars"}
	for _, star := range starData {
		markdownContents = append(markdownContents, fmt.Sprintf("### [%v](%v)\r\n%v", star.FullName, star.HTMLURL, star.Description))
	}

	writeFileError := os.WriteFile("stars/stars.md", []byte(strings.Join(markdownContents, "\r\n\r\n")), 0644)
	check(writeFileError)

	removeFileError := os.Remove(starsJson)
	check(removeFileError)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
