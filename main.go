package main

import "fmt"

// import "fmt"

func main() {
	intSlice := initNumSlice(10)
	for _, num := range intSlice {
		if isEven(num) {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}

	}
}

func isEven(num int) bool {
	remainder := num % 2
	if remainder == 0 {
		return true
	} else {
		return false
	}
}

// Create a slice with a numeral range of 0 to provided max
func initNumSlice(max uint) []int {
	s := make([]int, max)

	for i := range s {
		s[i] = i + 1
	}

	return s
}