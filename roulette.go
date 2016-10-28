// TODO
// * add error output
//

package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func count_bits(number uint64) uint64 {
	var count uint64
	for number > 0 {
		count += number & 1
		number >>= 1
	}
	return count
}

func main() {
	var cells, chips, start, end, i, result, count uint64
	var err error

	if len(os.Args) != 3 {
		fmt.Println("Please type 2 args... For example:")
		fmt.Println("  go run roulette.go 5 3 | head")
		return // fixme
	}

	cells, err = strconv.ParseUint(os.Args[1], 10, 64)
	chips, err = strconv.ParseUint(os.Args[2], 10, 64)

	if err != nil {
		fmt.Println("Input error...")
		return
	}

	if cells < 3 || cells > 63 {
		fmt.Println("Please input cells from 3 to 63")
		return // fixme
	}

	if chips < 3 || chips > 63 {
		fmt.Println("Please input chips from 3 to 36")
		return // fixme
	}

	if chips >= cells {
		fmt.Println("Please input chips < cells")
		return // fixme
	}

	count = count_combinations(cells, chips)
	fmt.Printf("Combinations: %d\n", count)

	if count < 10 {
		fmt.Println("Less 10 results... error")
		return // fixme
	}

	start = 1<<chips - 1
	end = (1<<cells - 1) &^ (1<<(cells-chips) - 1)

	for i = start; i <= end; i++ {
		if count_bits(i) == chips {
			fmt.Printf("%0[2]*[1]b\n", i, cells)
			result++
		}
	}

	// fmt.Println(result)

}

// Very very big numbers
func count_combinations(n, m uint64) (result uint64) {
	x := big.NewInt(1)
	y := big.NewInt(1)
	z := big.NewInt(1)
	r := big.NewInt(1)

	x.MulRange(1, int64(n))
	y.MulRange(1, int64(n-m))
	z.MulRange(1, int64(m))

	r = r.Div(r.Div(x, y), z)

	return r.Uint64()
}
