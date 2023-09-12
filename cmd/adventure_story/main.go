package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Luzik-D/Gophercises/cmd/adventure_story/story"
)

func main() {
	filename := flag.String("file", "gopher.json", "A JSON file with story")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal("Failed to open json file with story: ", err)
	}
	defer file.Close()

	storyMap, err := story.ParseJson(file)
	if err != nil {
		log.Fatal("Failed to parse JSON file with story ", err)
	}

	//hdlr := story.NewHandler(storyMap)

	// below handler with custom opts
	hdlr := story.NewHandler(storyMap, story.SetTemplate(template.Must(template.New("").Parse(story.CustomHTMLTemplate))), story.SetCustomPathFn(story.CustomPathFn))

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", hdlr))
}
