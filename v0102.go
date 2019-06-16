package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}

		input := make([][]int, n+1, n+1)
		totalRow := make([]int, n+1, n+1)

		for i := 0; i < n; i++ {
			row := make([]int, n+1, n+1)
			scanner.Scan()

			row = numbers(scanner.Text())

			input[i] = row
			total(totalRow, row)
		}

		input[n] = totalRow

		for xk := range input {
			for yk, yv := range input[xk] {
				if yk == n {
					fmt.Printf("%5d\n", yv)
				} else {
					fmt.Printf("%5d", yv)
				}
			}
		}
	}
}

func total(sum, row []int) {
	for i := 0; i < len(row); i++ {
		sum[i] += row[i]
	}
}

func numbers(s string) []int {
	var n []int
	var sum int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
			sum += i
		}
	}
	n = append(n, sum)
	return n
}
