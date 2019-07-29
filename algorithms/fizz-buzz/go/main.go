package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i < 100; i++ {
		answer := fizBuzz(i)
		if answer != "" {
			fmt.Println(i, answer)
		} else {
			fmt.Println(i)
		}
	}
}

func fizBuzz(n int) string {
	var answer strings.Builder
	answer.Grow(4)

	if n%3 == 0 {
		answer.WriteString("Fiz")
	}

	if n%5 == 0 {
		answer.WriteString("Buzz")
	}

	return answer.String()
}
