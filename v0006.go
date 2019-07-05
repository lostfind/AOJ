package main

import "fmt"

func main() {
	var str string
	var result string

	fmt.Scanf("%s", &str)

	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}

	fmt.Println(result)
}
