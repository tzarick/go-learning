package quiz

import (
	"fmt"
	"strings"
)

type Quiz struct {
	questionsAndAnswers map[string]string
	timeout             int
	scoreKeeper
}

type scoreKeeper struct {
	correct   int
	incorrect int
}

func NewQuiz(questionMap *map[string]string, timeout int) *Quiz {
	return &Quiz{
		questionsAndAnswers: *questionMap,
		timeout:             timeout,
	}
}

func (quiz *Quiz) Administer() {
	fmt.Printf("Starting quiz! There will be %d questions...\n", len(quiz.questionsAndAnswers))
	for q, a := range quiz.questionsAndAnswers {
		quiz.askQuestion(q, a)
	}
}

func (quiz *Quiz) askQuestion(q, a string) {
	fmt.Printf("What is %s?\n", q)
	var answer string
	_, err := fmt.Scan(&answer)

	if err != nil {
		quiz.incorrect++
	} else {
		if cleanAnswer := strings.ToLower(strings.TrimSpace(answer)); cleanAnswer == a {
			quiz.correct++
		} else {
			quiz.incorrect++
		}
	}
}

// return total Qs and also amount of correct and incorrect answers
func (sk scoreKeeper) Results() (int, int) {
	return sk.correct, sk.incorrect
}
