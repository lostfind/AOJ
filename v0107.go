package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var d, w, h, n int
		var p1, p2, p3 float64
		fmt.Scanf("%d %d %d", &d, &w, &h)

		if d == 0 && w == 0 && h == 0 {
			return
		} else {
			p1 = math.Hypot(float64(d), float64(w))
			p2 = math.Hypot(float64(w), float64(h))
			p3 = math.Hypot(float64(d), float64(h))
		}

		fmt.Println("len:", p1, p2, p3)

		fmt.Scanf("%d", &n)
		for i := 0; i < n; i++ {
			var r float64
			fmt.Scanf("%f", &r)

			if p1 < r*2 || p2 < r*2 || p3 < r*2 {
				fmt.Println("OK")
			} else {
				fmt.Println("NA")
			}
		}
	}
}
