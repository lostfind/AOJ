package main

import "fmt"

func main() {
	const MAX = 1000000
	prime := make([]int, MAX, MAX)
	check := make([]int, MAX, MAX)
	count := make([]int, MAX, MAX)

	prime[0] = 1
	prime[1] = 1

	// prime init
	for i := 2; i < MAX; i++ {
		if check[i] == 1 {
			continue
		}

		if prime[i] == 0 {
			for j := i * 2; j < MAX; j += i {
				check[j] = 1
				prime[j] = 1
			}
		}
	}

	for i := 2; i < MAX; i++ {
		cnt := 0
		if prime[i] == 0 {
			cnt = 1
		}
		count[i] = count[i-1] + cnt
	}

	for {
		var input int
		inputCount, _ := fmt.Scanf("%d", &input)
		if inputCount != 1 {
			break
		}
		fmt.Println(count[input])
	}
}
