package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	lists := make([][]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		lists[i] = parseList(scanner.Text())
	}

	spew.Dump(lists)
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
