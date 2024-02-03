package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sliceLength int

	if len(os.Args) > 1 {
		s, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		sliceLength = s
	}

	intSlice := initNumSlice(uint(sliceLength))
	for _, num := range intSlice {
		if isEven(num) {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}

func isEven(num int) bool {
	if num % 2 == 0 {
		return true
	} else {
		return false
	}
}

// Create a slice with a numeral range of 0 to provided max
func initNumSlice(max uint) []int {
	if max == 0 {
		max = 10
	}

	s := make([]int, max)

	for i := range s {
		s[i] = i + 1
	}

	return s
}