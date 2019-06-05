package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		var result string
		line := scanner.Text()
		key := make(map[string]string)

		var i int32
		for ; i < 52; i++ {
			makeKey(key, i)
			result = decode(key, line)

			if strings.Contains(result, "the") {
				break
			} else if strings.Contains(result, "this") {
				break
			} else if strings.Contains(result, "that") {
				break
			}
		}

		fmt.Println(result)
	}
}

func makeKey(key map[string]string, shift int32) {
	for i := 'a'; i <= 'z'; i++ {
		cipher := (i + shift)
		if cipher > 'z' {
			cipher = cipher%'z' + 'a' - 1
		} else if cipher < 'a' {
			cipher = cipher + ('z' - 'a') + 1
		}
		key[string(i)] = string(cipher)
	}
}

func decode(key map[string]string, cipher string) string {
	var result string
	for _, v := range cipher {
		if key[string(v)] != "" {
			result += key[string(v)]
		} else {
			result += string(v)
		}
	}
	return result
}
