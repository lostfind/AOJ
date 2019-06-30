package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	var a, b, c int
	for i := 0; i < n; i++ {
		var isTriangle bool
		fmt.Scanf("%d %d %d", &a, &b, &c)

		if a*a == b*b+c*c || b*b == a*a+c*c || c*c == b*b+a*a {
			isTriangle = true
		}

		if isTriangle == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
