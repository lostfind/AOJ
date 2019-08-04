package main

import "fmt"

func main() {
	var n int

	for {
		result := 0
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			return
		}

		if n <= 36 {
			for a := 9; a >= 0; a-- {
				for b := 9; b >= 0; b-- {
					for c := 9; c >= 0; c-- {
						for d := 9; d >= 0; d-- {
							if a+b+c+d == n {
								// fmt.Println(a, b, c, d)
								result++
							}
						}
					}
				}
			}
		}
		fmt.Println(result)
	}
}
