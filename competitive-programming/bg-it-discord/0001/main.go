package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()
	a, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	input = scanner.Text()
	b, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(a + b)
}
