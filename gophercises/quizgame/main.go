package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	problemsCsvFilename := flag.String("problemsCsvFilename", "problems.csv", "The filename of the csv holding the problems for the quiz")
	flag.Parse()

	probs, err := readProblems(problemsCsvFilename)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("read %d bytes: %q\n", count, (*probs)[:count])
	fmt.Println(probs)
	// for i := 0; i < count; i++ {
	// 	fmt.Printf("%q", (*probs)[i])
	// }
}

func readProblems(filename *string) (string, error) {
	file, err := os.Open(*filename)
	if err != nil {
		return "", errors.New(err.Error())
	}

	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		return "", errors.New(err.Error())
	}

	probsString := fmt.Sprintf("%q", data[:count])

	return probsString, nil
}
