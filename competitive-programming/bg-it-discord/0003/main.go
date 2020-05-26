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

	input := parseList(scanner.Text())
	sums := make([]int, len(input))
	for i, num := range input {
		sums[i] = sum(input) - num
	}

	for _, sum := range sums {
		fmt.Printf("%d ", sum)
	}
}

func sum(nums []int) int {
	acc := 0
	for _, num := range nums {
		acc += num
	}
	return acc
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
