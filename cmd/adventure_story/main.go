package main

import (
	"flag"
	"fmt"
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

	fmt.Println("Server is listening...")
	hdlr := story.NewHandler(storyMap)

	fmt.Println()
	log.Fatal(http.ListenAndServe(":8080", hdlr))
}
