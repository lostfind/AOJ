package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	result := make([]string, n, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()

		result[i] = strings.Replace(input, "Hoshino", "Hoshina", -1)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
