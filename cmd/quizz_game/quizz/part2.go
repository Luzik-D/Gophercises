package quizz

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Question struct {
	q string
	a string
}

func Game_v2() string {
	return "Init p2"
}

// / Parse q-a exprs from csv file to array of struct Question
func getQuestions(lines [][]string) []Question {
	q := make([]Question, len(lines))

	for i, l := range lines {
		q[i] = Question{
			q: l[0],
			a: l[1],
		}
	}

	return q
}

// / Print questions in stdout and wait for user answer.
func askQuestionsv2(q_data []Question, num_c, num_w *int, ch chan (bool)) {
	num_q := len(q_data)
	var ans string

	for i, q := range q_data {
		fmt.Printf("Question number %d/%d: %s\n", i+1, num_q, q.q)
		fmt.Scan(&ans)
		if ans != q.a {
			(*num_w)++
		} else {
			(*num_c)++
		}
	}

	ch <- true
}

func startQuizz(q_data []Question, timeAmnt int) {
	timer := time.NewTimer(time.Duration(timeAmnt) * time.Second)
	done := make(chan bool)

	num_corr := 0
	num_wrong := 0

	// run timer
	go func() {
		<-timer.C
	}()

	// run questions
	go askQuestionsv2(q_data, &num_corr, &num_wrong, done)

	for {
		select {
		case <-timer.C:
			fmt.Printf("\nWhoops, end of time. You answered %d questions.\n", num_corr+num_wrong)
			fmt.Printf("Your result is: %d correct and %d wrong answers\n", num_corr, num_wrong)
			return
		case <-done:
			fmt.Println("\nYou answered all questions")
			fmt.Printf("Your result is: %d correct and %d wrong answers\n", num_corr, num_wrong)
			return
		}
	}
}

func QuizzGame(file *os.File, timeAmnt int) {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Failed to parse the csv file: " + file.Name())
	}

	quizz_data := getQuestions(records)

	fmt.Printf("You have %d seconds to pass the quizz. Type enter to run.\n", timeAmnt)
	dummy := new(int)
	fmt.Scanln(dummy) // wait user's input

	startQuizz(quizz_data, timeAmnt)
}
