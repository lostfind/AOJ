package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	dictionary := make(map[string][]int)

	for {
		var word string
		var page int
		fmt.Scanf("%s %d", &word, &page)

		if word == "" || page == 0 {
			break
		}

		pages := dictionary[word]
		pages = append(pages, page)
		dictionary[word] = pages
	}

	words := make([]string, 0)
	for i := range dictionary {
		words = append(words, i)
	}

	sort.Strings(words)

	for _, v := range words {
		sort.Ints(dictionary[v])
		fmt.Println(v)
		for i, page := range dictionary[v] {
			if i == len(dictionary[v])-1 {
				fmt.Printf("%d\n", page)
			} else {
				fmt.Printf("%d ", page)
			}
		}
	}
}
