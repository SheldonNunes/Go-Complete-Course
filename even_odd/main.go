package main

import (
	"fmt"
)

func main() {
	nums := []int { 0,1,2,3,4,5,6,7,8,9,10 }

	for _, num := range nums {
		fmt.Printf("%v is %v\n", num, getNumberEvenOrOdd(num))
	}
}

func getNumberEvenOrOdd(num int) string {
	if(num % 2 == 0) {
		return "even"
	}
	return "odd"
}