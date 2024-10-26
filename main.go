package main

import (
	"fmt"
)

func main() {
	tess := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tessTwo := []int{100, 101, 102, 103}
	testResult := append(tess, tessTwo...)

	fmt.Println(testResult)
	sumAll := 0
	for _, value := range testResult {
		sumAll += value
		fmt.Print(sumAll, " --> ")
	}
	fmt.Println("\nИтого: ", sumAll)

}
