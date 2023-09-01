package quizz

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type QAnswer struct {
	CorrectAns int
	WrongAns   int
}

func Game_v1() string {
	return "Init part 1"
}

func askQuestions(records [][]string, s *QAnswer) {
	var userAns string
	for _, rec := range records {
		if len(rec) != 2 {
			log.Fatal("Incorrect record type")
		}
		q := rec[0]
		ans := rec[1]

		fmt.Println(q)
		fmt.Scan(&userAns)

		if userAns != ans {
			s.WrongAns++
		} else {
			s.CorrectAns++
		}
	}
}

func RunQuizz(file *os.File) (QAnswer, error) {
	correctAnsCnt := 0
	wrongAnsCnt := 0
	ans := QAnswer{correctAnsCnt, wrongAnsCnt}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read csv")
	}

	//start quizz
	askQuestions(records, &ans)

	return ans, nil
}
