package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	problemsCsvFilename := flag.String("problemsCsvFilename", "problems.csv", "The filename of the csv holding the problems for the quiz")
	flag.Parse()

	probs, err := readProblems(problemsCsvFilename)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("read %d bytes: %q\n", count, (*probs)[:count])
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

	fmt.Println(questionAnswer)

	// for i := 0; i < count; i++ {
	// 	fmt.Printf("%q", (*probs)[i])
	// }
}

func readProblems(filename *string) ([]string, error) {
	file, err := os.Open(*filename)
	if err != nil {
		return []string{}, errors.New(err.Error())
	}

	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		return []string{}, errors.New(err.Error())
	}

	probsString := string(data[:count])
	problemElements := strings.Split(probsString, "\n")

	return problemElements, nil
}
