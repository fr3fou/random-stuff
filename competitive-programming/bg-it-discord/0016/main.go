package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	lists := make([][]int, n)
	length := 0

	for i := 0; i < n; i++ {
		scanner.Scan()
		list := parseList(scanner.Text())
		length += list[0]   // first num is length
		lists[i] = list[1:] // rest is the list
	}

	merged := make([]int, length)

}

func mergeTwo(a []int, b []int) []int {
	length := len(a) + len(b)
	sorted := make([]int, 0, length-2)

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}

	// drain the last elements
	for i < len(a) {
		sorted = append(sorted, a[i])
		i++
	}

	// drain the last elements
	for j < len(b) {
		sorted = append(sorted, b[j])
		j++
	}

	return sorted
}

// O(n * log(k))
func Merge(list [][]int) []int {
	return merge(list)
}

func merge(list [][]int) []int {
	if low < high {

	}
}

func parseList(line string) []int {
	list := []int{}
	nums := strings.Split(line, " ")
	for _, n := range nums {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		list = append(list, num)
	}

	return list
}
