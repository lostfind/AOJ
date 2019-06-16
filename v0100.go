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
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}

		sales := make([]int, 4000)
		members := make([]int, 0)
		memberCheck := make(map[int]bool)

		for i := 0; i < n; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			e, _ := strconv.Atoi(input[0])
			p, _ := strconv.Atoi(input[1])
			q, _ := strconv.Atoi(input[2])

			if e >= 1 && e <= 4000 && memberCheck[e] == false {
				members = append(members, e)
				memberCheck[e] = true
			}
			sales[e-1] += p * q
		}

		emptyFlag := true
		for _, v := range members {
			if sales[v-1] >= 1000000 {
				fmt.Println(v)
				emptyFlag = false
			}
		}

		if emptyFlag == true {
			fmt.Println("NA")
		}
	}
}
