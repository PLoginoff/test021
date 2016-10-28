package main

import (
	"fmt"
)

func main() {
	var first, second, start, max, i, count, now, result uint64

	first = 36
	second = 18
	start = 1<<second - 1
	max = (1<<first - 1) &^ (1<<(first-second) - 1)

	fmt.Printf("%0[2]*[1]b\n", start, first)
	fmt.Printf("%b\n", max)

	for i = start; i <= max; i++ {
		now = i
		count = 0
		for now != 0 {
			if now&1 == 1 {
				count++
			}
			now >>= 1
		}
		if count == second {
			result++
		}
	}

	fmt.Println(result)
}
