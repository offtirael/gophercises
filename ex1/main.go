package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Question struct {
	text   string
	answer string
}

func readQuizFile(quizFilePath string) []Question {
	fl, err := os.Open(quizFilePath)
	if err != nil {
		log.Fatal("Unable to read Quiz file "+quizFilePath, err)
	}
	defer fl.Close()

	csvReader := csv.NewReader(fl)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse Quiz file "+quizFilePath, err)
	}

	var questions []Question
	for _, v := range records {
		questions = append(questions, Question{text: v[0], answer: v[1]})
	}

	return questions
}

func runQuiz(questions []Question) (int, int) {
	correctCount, incorrectCount := 0, 0
	var usersAnswer string
	for _, question := range questions {
		fmt.Printf("%s: ", question.text)
		fmt.Scanf("%s", &usersAnswer)
		if usersAnswer == question.answer {
			correctCount++
		} else {
			incorrectCount++
		}
	}
	return correctCount, incorrectCount
}

func main() {
	quizFile := flag.String("file", "problems.csv", "Quiz file name.")
	flag.Parse()

	records := readQuizFile(*quizFile)
	correctCount, incorrectCount := runQuiz(records)
	fmt.Println("Your results!")
	fmt.Printf("Correct answer count: %d\n", correctCount)
	fmt.Printf("Incorrect answer count: %d\n", incorrectCount)
}
