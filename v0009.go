package main

import (
	"fmt"
)

func main() {
	var primes []int
	maxPrime := 2
	primes = append(primes, 2)

	// input
	for i := 0; i < 10; i++ {
		var input int
		inputCount, _ := fmt.Scanf("%d", &input)
		if inputCount != 1 {
			break
		} else {
			primeCount := 0

			if input > maxPrime {
				for x := maxPrime + 1; x <= input; x++ {
					isPrime := isPrime(x)
					// Primeであれば追加
					if isPrime {
						primes = append(primes, x)
						maxPrime = x
					}
				}
			}

			for _, primeValue := range primes {
				if primeValue <= input {
					primeCount++
				} else {
					break
				}
			}
			fmt.Println(primeCount)
		}
	}
}

func isPrime(value int) bool {
	result := true
	for i := 2; i < value; i++ {
		if value%i == 0 {
			result = false
			break
		}
	}

	return result
}
