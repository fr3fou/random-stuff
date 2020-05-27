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

	product := 1
	moreZeroes := false
	zeroIndex := -1

	for i, num := range input {
		if num == 0 {
			if zeroIndex == -1 {
				zeroIndex = i
			} else {
				moreZeroes = true
			}
		} else {
			product *= num
		}
	}

	for i, num := range input {
		if moreZeroes {
			fmt.Printf("0 ")
			continue
		}
		if zeroIndex != -1 {
			if i == zeroIndex {
				fmt.Printf("%d ", product)
			} else {
				fmt.Printf("0 ")
			}
		} else {
			fmt.Printf("%d ", product/num)
		}
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
