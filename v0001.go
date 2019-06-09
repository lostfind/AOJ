package main

import (
	"fmt"
	"sort"
)

func main() {
	mountains := make([]int, 10, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&mountains[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(mountains)))

	for i := 0; i < 3; i++ {
		fmt.Println(mountains[i])
	}
}
