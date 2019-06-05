package main

import "fmt"

func main() {
	var result []int
	for {
		var n, a, maxSum int
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		maxSum = 0
		sum := 0
		for i := 0; i < n; i++ {
			fmt.Scan(&a)
			if i == 0 {
				sum = a
				maxSum = sum
			} else if sum+a < a {
				sum = a
				if maxSum < sum {
					maxSum = sum
				}
			} else {
				sum = sum + a
				if maxSum < sum {
					maxSum = sum
				}
			}
		}

		result = append(result, maxSum)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
