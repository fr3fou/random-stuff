package main

import (
	"bufio"
	"fmt"
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

	if length == 1 {
		printList(lists[0])
	} else {
		printList(MergeK(lists))
	}
}

// O(n * log(k))
func MergeK(lists [][]int) []int {
	length := len(lists)

	if length <= 1 {
		return lists[0]
	}

	pivot := length / 2

	return mergeTwo(
		MergeK(lists[:pivot]),
		MergeK(lists[pivot:]),
	)
}

// O(n * log(k))
func MergeKBad(list [][]int) []int {
	length := len(list)

	switch {
	case length == 2:
		return mergeTwo(list[0], list[1])
	case length == 1:
		return list[0]
	case length == 0:
		return []int{}
	}

	merged := mergeTwo(list[0], list[1])

	for i := 2; i < length; i++ {
		merged = mergeTwo(merged, list[i])
	}

	return merged
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

func printList(list []int) {
	for _, val := range list {
		fmt.Printf("%d ", val)
	}
}
