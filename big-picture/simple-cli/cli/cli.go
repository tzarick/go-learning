package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Run_cli(path, level *string) {
	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n')

		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
