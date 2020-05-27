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
	sum := 0

	// O(2n) is still O(n) 😎

	for _, num := range input {
		sum += num
	}

	for _, num := range input {
		fmt.Printf("%d ", sum-num)
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
