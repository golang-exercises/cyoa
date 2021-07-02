package main

import (
	"bufio"
	mod "cyoa/internal"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type ChapterOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string           `json:"title"`
	Story   []string         `json:"story"`
	Options []ChapterOptions `json:"options"`
}

func main() {
	var story map[string]Chapter
	storyRawData, cliMode := mod.ParseFlags()

	jsonParseErr := json.Unmarshal(storyRawData, &story)

	if jsonParseErr != nil {
		panic(jsonParseErr)
	}

	if !cliMode {
		server := ChapterRouteHandler(story)

		fmt.Println("Story is up!")
		http.ListenAndServe(":8080", server)
	} else {

		fmt.Println("Press [Enter] to start the story line.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ChapterCLIHandler(story, scanner, "intro")
	}

}

func ChapterCLIHandler(s map[string]Chapter, scanner *bufio.Scanner, key string) string {
	optionsLength := GetOptionCount(s[key])

	PrintCLITemplate(s[key])

	if optionsLength == 0 {
		fmt.Println("THE END.")
		os.Exit(1)
	}

	input := HandleInput(scanner)
	idx, err := strconv.Atoi(input)
	if err != nil {
		return ChapterCLIHandler(s, scanner, key)
	} else {
		if optionsLength < idx-1 || idx-1 < 0 {
			fmt.Print("\n\nNo valid input. Please try again.\n\n")
			return ChapterCLIHandler(s, scanner, key)
		}
		key = s[key].Options[idx-1].Arc
		return ChapterCLIHandler(s, scanner, key)
	}
}

func HandleInput(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

func GetOptionCount(c Chapter) int {
	return len(c.Options)
}

func PrintCLITemplate(c Chapter) {
	fmt.Printf("%s\n\n", c.Title)
	for _, val := range c.Story {
		fmt.Printf("%s\n", val)
	}
	for i, val := range c.Options {
		fmt.Printf("\n[%d] %s\n", i+1, val.Text)
	}
}

func ChapterRouteHandler(s map[string]Chapter) http.Handler {
	mux := http.NewServeMux()
	tmpl, err := template.ParseFiles("assets/story.html")
	if err != nil {
		panic(err)
	}

	for key, value := range s {
		path := key
		contents := value
		url := "/"

		if path != "intro" {
			url += path
		}

		mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
			tmpl.Execute(w, contents)
		})
	}

	return mux
}
