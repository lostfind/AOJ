package main

import "fmt"

func main() {
	for {
		var a, b, n int
		fmt.Scanf("%d", &n)

		if n == 0 {
			return
		}

		for i := 0; i < n; i++ {
			var cardA, cardB int
			fmt.Scanf("%d %d", &cardA, &cardB)

			if cardA > cardB {
				a += cardA + cardB
			} else if cardB > cardA {
				b += cardA + cardB
			} else {
				a += cardA
				b += cardB
			}
		}

		fmt.Println(a, b)
	}
}
