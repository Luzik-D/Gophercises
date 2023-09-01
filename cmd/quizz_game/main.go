package main

import (
	"flag"
	"log"
	"os"

	"github.com/Luzik-D/Gophercises/cmd/quizz_game/quizz"
)

const QUIZZ_TIME = 20

func main() {

	file_name := flag.String("csv", "problems.csv", "a csv file in format 'question, answer'")
	flag.Parse()

	file, err := os.Open(*file_name)
	defer file.Close()
	if err != nil {
		log.Fatal("Failed to open csv file")
	}

	// uncomment /*--*/ section to run first version of quizz game
	/*res, err := quizz.RunQuizz(file)
	if err != nil {
		log.Fatal("Game crushed")
	}

	fmt.Println("Number of correct answers: ", res.CorrectAns)
	fmt.Println("Number of wrong answers: ", res.WrongAns)*/

	quizz.QuizzGame(file, QUIZZ_TIME)
}
