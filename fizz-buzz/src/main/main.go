package main

import (
	"fmt"
	"strconv"
)

func isDivisible(i int, divider int) bool {
	return (i % divider) == 0
}

func fizzbuzz(i int) string {
	if isDivisible(i, 3) && isDivisible(i, 5) {
		return "fizzbuzz"
	} else if isDivisible(i, 3) {
		return "fizz"
	} else if isDivisible(i, 5) {
		return "buzz"
	}
	return strconv.Itoa(i)
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
