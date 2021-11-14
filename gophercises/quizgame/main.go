package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tzarick/go-learning/gophercises/quizgame/quiz"
)

func main() {
	problemsCsvFilename := flag.String("csv", "problems.csv", "The filename of the csv holding the problems for the quiz")
	flag.Parse()

	probs, err := readProblems(problemsCsvFilename)
	if err != nil {
		log.Fatal(err)
	}

	questionAnswerMap := buildQuestionAnswerMap(probs)

	quiz := quiz.NewQuiz(&questionAnswerMap, 0)
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

func buildQuestionAnswerMap(probs []string) map[string]string {
	questionAnswer := map[string]string{}
	for _, v := range probs {
		if len(v) > 0 {
			questionElements := strings.Split(v, ",")
			elementLen := len(questionElements)
			answer := questionElements[elementLen-1]
			question := strings.Join(questionElements[:(elementLen-1)], ",") // account for questions that have commas in them
			questionAnswer[question] = answer
		}
	}

	return questionAnswer
}

func readProblems(filename *string) ([]string, error) {
	file, err := os.Open(*filename)
	if err != nil {
		return []string{}, errors.New(err.Error())
	}

	defer file.Close()

	data := make([]byte, 200)
	count, err := file.Read(data)
	if err != nil {
		return []string{}, errors.New(err.Error())
	}

	probsString := string(data[:count])
	// println("p string:", probsString)
	problemElements := strings.Split(probsString, "\n")

	return problemElements, nil
}
