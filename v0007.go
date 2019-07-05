package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	hellGate := 100.000

	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		hellGate *= 1.05
		hellGate = math.Ceil(hellGate)
	}

	fmt.Println(int(hellGate * 1000))
}
