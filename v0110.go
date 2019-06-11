package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	var left, right, awnser string
	var splits []string

	for {
		i, _ := fmt.Scanf("%s", &input)

		if i != 1 {
			break
		}

		splits = strings.Split(input, "+")
		left = splits[0]
		right = splits[1]

		splits = strings.Split(right, "=")
		right = splits[0]
		awnser = splits[1]

		var rpLeft string
		for i := 0; i < 10; i++ {

			rpLeft = strings.Replace(left, "X", strconv.Itoa(i), -1)
			// var l, r, a int
			fmt.Println(rpLeft)

		}

	}
	fmt.Println("AAAA", left, right, awnser)
}
