package main

import "fmt"

func main() {
	numbers := make([]int, 10, 10)
	// sum : 0 ~ 330
	// var counts [10][331]int

	for {
		var n, s int
		scanCount, _ := fmt.Scanf("%d %d", &n, &s)

		if scanCount == 2 {
			count := combination(numbers, s, n, 1, 0)
			fmt.Println(count)
		} else {
			break
		}
	}
}

func combination(numbers []int, target, maxDepth, depth, total int) int {
	var count int

	for i, v := range numbers {
		if v == 1 {
			continue
		}

		sumNumber := total + (i * depth)
		if sumNumber == target && depth == maxDepth {
			count++
			break
		} else if sumNumber > target {
			break
		}

		if depth < maxDepth && sumNumber <= target {
			numbers[i] = 1
			count += combination(numbers, target, maxDepth, depth+1, sumNumber)
			numbers[i] = 0
		}
	}

	return count
}
