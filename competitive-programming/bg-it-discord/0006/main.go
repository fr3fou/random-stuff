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

	year, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	if year%100 == 0 {
		if year%400 == 0 {
			fmt.Println(true)
			return
		}
	} else if year%4 == 0 {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}
