package main

import (
	"fmt"
)

type Goods struct {
	value  int
	weight int
}

func main() {
	var N, W int

	fmt.Scanf("%d %d", &N, &W)

	objects := make([]Goods, N)

	// 品物入力
	for i := 0; i < N; i++ {
		var v, w int
		fmt.Scanf("%d %d", &v, &w)

		goods := Goods{value: v, weight: w}
		objects[i] = goods
	}

	value := Knapsack(W, objects)
	fmt.Println(value)
}

// Knapsack is Knapsack calculate
func Knapsack(bagWeight int, objects []Goods) int {
	preBag := make([]int, bagWeight+1)
	bag := make([]int, bagWeight+1)

	for _, obj := range objects {
		remain := 0
		for i := obj.weight; i <= bagWeight; i++ {
			remainValue := preBag[remain]

			if bag[i] < obj.value+remainValue {
				bag[i] = obj.value + remainValue
			}

			remain++
		}
		copy(preBag, bag)

	}

	return bag[bagWeight]
}
