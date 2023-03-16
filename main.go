package main

import (
	"fmt"
	"sort"
)

func main() {
	var z, e int
	caseNumber := 1

	for fmt.Scan(&z); z > 0; fmt.Scan(&z) {
		profits := make([]int, z)
		orders := make([][]int, z)

		for i := 0; i < z; i++ {
			fmt.Scan(&e)
			trunks := make([]int, e)
			for j := 0; j < e; j++ {
				fmt.Scan(&trunks[j])
			}
			profits[i], orders[i] = getMaxProfit(trunks)
		}

		for i, profit := range profits {
			fmt.Printf("Case %d\n", caseNumber)
			fmt.Println("Max profit:", profit)
			fmt.Print("Order:")
			for _, order := range orders[i] {
				fmt.Print(" [", order, "]")
			}
			fmt.Println("\n")
			caseNumber++
		}
	}
}

func getMaxProfit(trunks []int) (int, []int) {
	sort.Ints(trunks)
	profit := 0
	order := make([]int, 0)

	counts := []int{0, 0, 0}

	for _, length := range trunks {
		counts[length%3]++
	}

	profit += counts[0]*2 + counts[1]*(-1) + counts[2]*3

	for i := 0; i < counts[2]; i++ {
		order = append(order, 2)
	}
	for i := 0; i < counts[1]; i++ {
		order = append(order, 1)
	}
	for i := 0; i < counts[0]; i++ {
		order = append(order, 3)
	}

	return profit, order
}
