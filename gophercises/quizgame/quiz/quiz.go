package quiz

import (
	"fmt"
	"strings"
	"time"
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
	fmt.Println("Press enter to start quiz...")
	fmt.Scanf("%v") // waits until enter is pressed

	timer := *time.NewTimer(time.Duration(quiz.timeout) * time.Second)
	// <-timer.C // block until we receive a message

	fmt.Printf("Starting quiz! There will be %d questions...\n", len(quiz.questionsAndAnswers))

questionLoop:
	for q, a := range quiz.questionsAndAnswers {
		fmt.Printf("What is %s?\n", q)

		answerCh := make(chan string)
		go func(answerCh chan string) {
			var answer string
			fmt.Scan(&answer)

			answerCh <- answer
		}(answerCh)

		// wait either for a message from the timer telling us time has expired OR a message from the answer channel with the user input
		select {
		case <-timer.C:
			// fmt.Printf("\nTime limit reached. Results -> Correct: %d, Total: %d\n", quiz.correct, len(quiz.questionsAndAnswers))
			// os.Exit(0)
			fmt.Println("\nTimeout reached!!")
			break questionLoop // break out of the question loop and continue to the end of the program
		case answer := <-answerCh:
			if cleanAnswer := strings.ToLower(strings.TrimSpace(answer)); cleanAnswer == a {
				quiz.correct++
			} else {
				quiz.incorrect++
			}
			// no default case exists, so it blocks until we get a message from one of the channels
		}
	}

}

// return total Qs and also amount of correct and incorrect answers
func (sk scoreKeeper) Results() (correct int, incorrect int) {
	return sk.correct, sk.incorrect
}
