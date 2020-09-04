package main

import (
	"fmt"
)

func main() {
	fmt.Println(split("foo bar baz", " "))
	fmt.Println(split("foo-bar-baz", "-"))
}

func split(s string, str string) []string {
	val := []string{}
	temp := ""

	for _, elem := range s {
		if string(elem) == str {
			val = append(val, temp)
			temp = ""
		} else {
			temp += string(elem)
		}
	}

	// append the last elem
	if temp != "" {
		val = append(val, temp)
	}

	return val
}
