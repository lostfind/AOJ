package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		w, _ := strconv.Atoi(input[0])
		h, _ := strconv.Atoi(input[1])

		if w == 0 && h == 0 {
			return
		}

		isArrive := make([][]bool, w)
		tiles := make([][]string, w)

		for i := 0; i < w; i++ {
			rowArrive := make([]bool, h)
			rowTiles := make([]string, h)
			scanner.Scan()
			row := scanner.Text()
			for j := 0; j < h; j++ {
				rowTiles[j] = string(row[j])
			}
			tiles[i] = rowTiles
			isArrive[i] = rowArrive
		}

		var i, j int
		isEnd := false
		for {
			if isArrive[i][j] == true {
				fmt.Println("LOOP")
				break
			}
			isArrive[i][j] = true

			switch tiles[i][j] {
			case ">":
				j++
			case "^":
				i--
			case "v":
				i++
			case "<":
				j--
			case ".":
				isEnd = true
			}

			if isEnd == true {
				fmt.Println(j, i)
				break
			}
		}
	}
}
