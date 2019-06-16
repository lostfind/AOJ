package main

import "fmt"

func main() {
	maxGram := 5000
	// store := map[int]int{
	// 	1500: 2244, // 500 * 3
	// 	1200: 1870, // 300 * 4
	// 	1000: 1520, // 200 * 5
	// 	500:  850,
	// 	300:  550,
	// 	200:  380,
	// }

	for {
		var gram int
		fmt.Scanf("%d", &gram)

		if gram == 0 {
			return
		}

		var minPrice int

		// (x * 200) + (y * 300) + (z * 500) == gram
		for x := 0; (x * 200) <= maxGram; x++ {
			for y := 0; (y * 300) <= maxGram; y++ {
				for z := 0; (z * 500) <= maxGram; z++ {
					if (x*200)+(y*300)+(z*500) == gram {
						price := 0
						price += (z / 3) * 2244
						price += (z % 3) * 850
						price += (y / 4) * 1870
						price += (y % 4) * 550
						price += (x / 5) * 1520
						price += (x % 5) * 380

						if price < minPrice || minPrice == 0 {
							minPrice = price
						}
					}
				}
			}
		}

		fmt.Println(minPrice)
	}
}

// 1 袋あたりの量			200g	300g	500g
// 1 袋の単価(定価)		380円	550円	850円
// 割引が適用される単位		5袋毎	4袋毎	3袋毎
// 割引率							20%引き	15%引き	12%引き
