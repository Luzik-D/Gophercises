package main

import (
	"fmt"

	"github.com/Luzik-D/Gophercises/cmd/quizz_game/quizz"
)

func main() {
	fmt.Println("hello")
	g1 := quizz.Game_v1()
	g2 := quizz.Game_v2()
	fmt.Println(g1)
	fmt.Println(g2)
}
