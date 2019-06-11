package main

import "fmt"

func main() {
	for {
		var n, e, q1, q2 int
		fmt.Scanf("%d", &n)

		if n == 0 {
			break
		}

		result := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Scan(&e, &q1, &q2)
			result[e] += q1 * q2
		}

		emptyFlag := true
		for i, v := range result {
			if v >= 1000000 {
				emptyFlag = false
				fmt.Println(i)
			}
		}

		if emptyFlag {
			fmt.Println("NA")
		}
	}
}
