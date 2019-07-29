package main

import "fmt"

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
	answer := ""

	if n%3 == 0 {
		answer += "Fiz"
	}

	if n%5 == 0 {
		answer += "Buzz"
	}

	return answer
}
