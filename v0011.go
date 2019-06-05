package main

import "fmt"

func main() {
	var w, n int
	fmt.Scan(&w)
	numbers := make([]int, w)

	fmt.Scan(&n)

	intersect := make([][2]int, n)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scanf("%d,%d", &x, &y)
		intersect[i] = [2]int{x, y}
	}

	for i := 0; i < w; i++ {
		numbers[i] = i + 1
	}

	for _, v := range intersect {
		swap(numbers, v[0], v[1])
	}

	for _, v := range numbers {
		fmt.Println(v)
	}
}

func swap(arr []int, x, y int) {
	tmp := arr[x-1]
	arr[x-1] = arr[y-1]
	arr[y-1] = tmp
}
