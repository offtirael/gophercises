package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
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

func runQuiz(questions []Question, timer *time.Timer) (score int) {
	correctCount := 0
	var usersAnswer string
	answers := make(chan string)

	for _, question := range questions {
		fmt.Printf("%s: ", question.text)
		go func() {
			fmt.Scanf("%s", &usersAnswer)
			answers <- usersAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTimeout!")
			return correctCount
		case answer := <-answers:
			if answer == question.answer {
				correctCount++
			}
		}
	}

	return correctCount
}

func main() {
	quizFile := flag.String("file", "problems.csv", "Quiz file name.")
	timeout := flag.Int("timeout", 30, "Quiz timeout.")
	shuffle := flag.Bool("shuffle", false, "Shuffle quiz")
	flag.Parse()

	records := readQuizFile(*quizFile)
	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	}

	fmt.Println("Press Enter to start the Quiz")
	fmt.Scanln()

	timer := time.NewTimer(time.Second * time.Duration(*timeout))
	correctCount := runQuiz(records, timer)
	fmt.Println("Your results:")
	fmt.Printf("Correct answer count: %d\n", correctCount)
	fmt.Printf("Incorrect answer count: %d\n", len(records)-correctCount)
}
