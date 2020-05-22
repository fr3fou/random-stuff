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

	for _, list := range lists {

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
