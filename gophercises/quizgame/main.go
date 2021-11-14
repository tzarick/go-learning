package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	probs, count, err := readProblems("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, (*probs)[:count])
}

func readProblems(filename string) (*[]byte, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return &[]byte{}, 0, errors.New(err.Error())
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		return &[]byte{}, 0, errors.New(err.Error())
	}

	return &data, count, nil
}
