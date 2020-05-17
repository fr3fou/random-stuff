package main

import (
	"fmt"
	"math"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// input := scanner.Text()
	// words := strings.Split(input, " ")
	fmt.Println(Lev("cat", "bot"))
}

func Lev(a, b string) int {
	return lev(a, b, len(a), len(b))
}

func lev(a, b string, i, j int) int {
	if Min(i, j) == 0 {
		return Max(i, j)
	}

	indicator := 1
	if a[i-1] == b[j-1] {
		indicator = 0
	}

	return Min(
		lev(a, b, i-1, j)+1,
		lev(a, b, i, j-1)+1,
		lev(a, b, i-1, j-1)+indicator,
	)
}

func Max(nums ...int) int {
	max := math.MinInt64

	for _, num := range nums {
		if num >= max {
			max = num
		}
	}

	return max
}

func Min(nums ...int) int {
	min := math.MaxInt64

	for _, num := range nums {
		if num <= min {
			min = num
		}
	}

	return min
}
