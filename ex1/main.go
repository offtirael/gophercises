package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readQuizFile(quizFilePath string) [][]string {
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

	return records
}

func runQuiz(questions [][]string) {
	for _, question := range questions {
		fmt.Println(question[0])
	}
}

func main() {
	records := readQuizFile("problems.csv")
	runQuiz(records)
}
