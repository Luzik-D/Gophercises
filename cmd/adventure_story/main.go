package main

import (
	"fmt"

	"github.com/Luzik-D/Gophercises/cmd/adventure_story/story"
)

func main() {
	fmt.Println("Choose your own adventure")
	story.ParseJson()
}
