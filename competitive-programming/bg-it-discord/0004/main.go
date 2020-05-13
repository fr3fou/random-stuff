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
	a, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	b, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Println(Path(a, b, make(Operations, 1)))
}

type OP int

const (
	Double OP = iota
	Halve
	AddTwo
)

func (o OP) String() string {
	switch o {
	case Double:
		return "*"
	case Halve:
		return "/"
	case AddTwo:
		return "+"
	default:
		return ""
	}
}

func (o OP) Apply(a int) int {
	switch o {
	case Double:
		return a * 2
	case Halve:
		return a / 2
	case AddTwo:
		return a + 2
	default:
		return -1
	}
}

type Operations []OP

func (ops Operations) String(a, b int) string {
	builder := &strings.Builder{}
	fmt.Fprintf(builder, "%d ", a)
	for _, op := range ops {
		a = op.Apply(a)
		fmt.Fprintf(builder, "%s 2 -> %d ", op, a)
	}

	return builder.String()
}

func (ops Operations) Apply(a int) int {
	for _, op := range ops {
		a = op.Apply(a)
	}

	return a
}

func Path(a, b int, ops Operations) int {
	if a == b {
		return 1
	}

	length := len(ops)
	if force(0, a, b, a, ops) {
		// fmt.Println(ops.String(a, b))
		return length + 1
	}

	return Path(a, b, make(Operations, length+1))
}

func force(index, a, b, current int, ops Operations) bool {
	if index >= len(ops) {
		return false
	}

	for i := Double; i <= AddTwo; i++ {
		if i == Halve {
			if current%2 == 0 {
				ops[index] = i
			} else {
				continue
			}
		} else {
			ops[index] = i
		}

		app := ops.Apply(a)
		if app == b || force(index+1, a, b, app, ops) {
			return true
		}
	}

	return false
}
