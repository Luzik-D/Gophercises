package main

import (
	"fmt"
	"log"

	"github.com/Luzik-D/Gophercises/cmd/quizz_game/quizz"
)

func main() {
	fmt.Println("hello")

	res, err := quizz.RunQuizz("problems.csv")
	if err != nil {
		log.Fatal("Game crushed")
	}

	fmt.Println("Number of correct answers: ", res.CorrectAns)
	fmt.Println("Number of wrong answers: ", res.WrongAns)
}
