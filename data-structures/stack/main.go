package main

import (
	"fmt"

	"github.com/fr3fou/go-stack/stack"
)

func main() {
	s := stack.New()
	s.Push(5)
	s.Push(4)

	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
}
