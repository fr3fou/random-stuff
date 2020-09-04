package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(split("foo bar baz", " "))
	fmt.Println(split("foo-bar-baz", "-"))
}

func split(s string, str string) []string {
	val := []string{}
	temp := &strings.Builder{}

	for _, elem := range s {
		if string(elem) == str {
			val = append(val, temp.String())
			temp.Reset()
		} else {
			temp.WriteString(string(elem))
		}
	}

	// append the last elem
	if temp.String() != "" {
		val = append(val, temp.String())
	}

	return val
}
