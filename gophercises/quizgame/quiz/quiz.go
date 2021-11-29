package quiz

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Quiz struct {
	questionsAndAnswers map[string]string
	timeout             int
	startTime           time.Time
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
	fmt.Scanf("Press enter to start quiz...") // waits until enter is pressed

	wg := &sync.WaitGroup{}

	quiz.startTime = time.Now()

	wg.Add(2)
	go func(startTime time.Time, timeout int, wg *sync.WaitGroup) {
		for {
			if int(time.Since(startTime).Seconds()) >= timeout {
				// terminate quiz somehow
				correct, incorrect := quiz.Results()

				fmt.Println(strings.Repeat("-", 30))
				fmt.Printf(`
					Results:
			
					Total Qs: %v
					Correct: %v
					Incorrect: %v
				`, correct+incorrect, correct, incorrect)

				os.Exit(0)
				wg.Done()
			}
		}
	}(quiz.startTime, quiz.timeout, wg)

	go func(wg *sync.WaitGroup) {
		fmt.Printf("Starting quiz! There will be %d questions...\n", len(quiz.questionsAndAnswers))
		for q, a := range quiz.questionsAndAnswers {
			quiz.askQuestion(q, a)
		}

		wg.Done()
	}(wg)

	wg.Wait()
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
func (sk scoreKeeper) Results() (correct int, incorrect int) {
	return sk.correct, sk.incorrect
}
