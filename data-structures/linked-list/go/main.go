package main

import (
	"fmt"

	"github.com/fr3fou/reversing-linked-list/llist"
)

func main() {
	l := llist.New(9)
	l.Add(5)
	l.Add(4)
	l.Add(7)
	l.Add(23)
	l.Add(12)
	fmt.Printf("Original: %+v\n", l.ToArray())
	l.Reverse()
	fmt.Printf("Reversed: %+v\n", l.ToArray())
}
