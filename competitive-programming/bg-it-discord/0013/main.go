package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	words := strings.Split(input, " ")

	for _, word := range words {
		if Solve(words, word) {
			fmt.Println(true)
			return
		}
	}

	fmt.Println(false)
}

func Solve(words []string, target string) bool {
	ds := dists(words, target)

	for j := range ds {
		ignored := Filter(words, target)
		if Solve(ignored, words[j]) {
			return true
		}
	}

	return false
}

// Filter removes the word provided
func Filter(words []string, word string) []string {
	filtered := make([]string, 0, len(words)-1)

	for _, w := range words {
		if w != word {
			filtered = append(filtered, w)
		}
	}

	return filtered
}

// dists compares against every other word for distance
// and find the 1st one that has a distance < 1
func dists(words []string, target string) []int {
	ds := []int{}

	for j := 0; j < len(words); j++ {
		// prevent comparision against self
		word := words[j]

		if word == target {
			continue
		}

		if Lev(word, target) < 2 {
			ds = append(ds, j)
		}
	}

	return ds
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
