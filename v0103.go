package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var inning int

	fmt.Scanf("%d", &inning)

	score := make([]int, inning)

	for i := 0; i < inning; i++ {
		var outCount, runnerCount int

		for scanner.Scan() {
			event := scanner.Text()

			// HIT, HOMERUN, OUT
			switch event {
			case "HIT":
				runnerCount++
				if runnerCount > 3 {
					runnerCount--
					score[i]++
				}
			case "HOMERUN":
				score[i] += runnerCount + 1
				runnerCount = 0
			case "OUT":
				outCount++
			}

			if outCount == 3 {
				break
			}
		}
	}

	for _, v := range score {
		fmt.Println(v)
	}
}
