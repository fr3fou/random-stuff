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

	fmt.Println(Solve(words))
}

func Solve(words []string) bool {
	for i := range words {
		visited := make([]bool, len(words))
		if solve(words, i, visited) {
			return true
		}
	}

	return false
}

func solve(words []string, index int, visited []bool) bool {
	ok := true
	for _, node := range visited {
		if !node {
			ok = false
			break
		}
	}

	if ok {
		return ok
	}

	// Go over the possible candidates
	for _, i := range Next(words, words[index]) {
		if visited[i] {
			continue
		}

		visited[i] = true
		if solve(words, i, visited) {
			return true
		}
		visited[i] = false
	}

	return false
}

// Next gives you an array of possible next words.
// It computes the distances and gives you the words with distance < 2
func Next(words []string, target string) []int {
	next := []int{}

	for i := 0; i < len(words); i++ {
		if Lev(words[i], target) == 1 {
			next = append(next, i)
		}
	}

	return next
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
