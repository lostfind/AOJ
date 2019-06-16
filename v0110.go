package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var left, right, awnser string
	var splits []string

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			return
		}

		splits = strings.Split(input, "+")
		left = splits[0]
		right = splits[1]

		splits = strings.Split(right, "=")
		right = splits[0]
		awnser = splits[1]

		startIndex := 0
		if len(left) > 1 && string(left[0]) == "X" {
			startIndex = 1
		}
		if len(right) > 1 && string(right[0]) == "X" {
			startIndex = 1
		}
		if len(awnser) > 1 && string(awnser[0]) == "X" {
			startIndex = 1
		}

		isNa := true

		for i := startIndex; i < 10; i++ {
			rpLeft := strings.Replace(left, "X", strconv.Itoa(i), -1)
			rpRight := strings.Replace(right, "X", strconv.Itoa(i), -1)
			rpAwnser := strings.Replace(awnser, "X", strconv.Itoa(i), -1)

			intLeft, _ := strconv.Atoi(rpLeft)
			intRight, _ := strconv.Atoi(rpRight)
			intAwnser, _ := strconv.Atoi(rpAwnser)

			// fmt.Println(intLeft, intRight, intAwnser)
			if (intLeft + intRight) == intAwnser {
				fmt.Println(i)
				isNa = false
				break
			}
		}

		if isNa == true {
			fmt.Println("NA")
		}
	}
}

// 123+4X6=X79
// 12X+4X6=X79
// XX22+89=X2XX
