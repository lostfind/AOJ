package main

import "fmt"

func main() {
	var a, b int
	var gcd, lcm int

	for {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			return
		}

		var big, small int
		// GCD
		if a > b {
			big = a
			small = b
		} else {
			big = b
			small = a
		}

		for {
			if big%small == 0 {
				gcd = small
				break
			} else {
				small, big = big%small, small
			}
		}

		// LCM
		lcm = a * b / gcd

		fmt.Println(gcd, lcm)
	}
}
