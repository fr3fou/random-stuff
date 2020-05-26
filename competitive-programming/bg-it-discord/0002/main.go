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
	nums := map[int]int{}
	for _, num := range input {
		nums[num]++
	}
	for num, occ := range nums {
		if occ == 1 {
			fmt.Println(num)
			return
		}
	}

	fmt.Println(-1)
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
