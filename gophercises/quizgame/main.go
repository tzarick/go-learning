package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tzarick/go-learning/gophercises/quizgame/quiz"
)

func main() {
	problemsCsvFilename := flag.String("csv", "problems.csv", "The filename of the csv holding the problems for the quiz")
	userTimeout := flag.String("timeout", "30", "Time limit for the quiz")
	flag.Parse()

	timeout, _ := strconv.Atoi(*userTimeout)

	probs, err := readProblems(problemsCsvFilename)
	if err != nil {
		log.Fatal(err) // Fatal is equivalent to print + os.Exit(1)
	}

	questionAnswerMap := buildQuestionAnswerMap(probs)

	quiz := quiz.NewQuiz(&questionAnswerMap, timeout)

	quiz.Administer()
	correct, incorrect := quiz.Results()

	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf(`
		Results:

		Total Qs: %v
		Correct: %v
		Incorrect: %v
	`, correct+incorrect, correct, incorrect)
}

func buildQuestionAnswerMap(lines [][]string) map[string]string {
	questionAnswer := map[string]string{}
	for _, line := range lines {
		elementLen := len(line)
		answer := strings.TrimSpace(line[elementLen-1])
		question := strings.Join(line[:(elementLen-1)], " ") // account for multi word questions
		questionAnswer[question] = answer
	}

	return questionAnswer
}

// func parseData(lines [][]string) []problem {
// 	clean := make([]problem, len(lines))
// 	for i, line := range lines {
// 		lineLen := len(line)
// 		clean[i] = problem{
// 			q: strings.Join(line[:(lineLen-1)], " "),
// 			a: strings.TrimSpace([lineLen-1]),
// 		}
// 	}

// 	return clean
// }

// type problem struct {
// 	q string
// 	a string
// }

func readProblems(filename *string) (probs [][]string, err error) {
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *filename)
		return [][]string{}, errors.New(err.Error())
	}

	defer file.Close()

	r := csv.NewReader(file)
	data, err := r.ReadAll() // this returns a 2D string slice
	if err != nil {
		fmt.Printf("Error reading CSV file: %s\n", *filename)
		return [][]string{}, errors.New(err.Error())
	}

	return data, nil
}
