package main

import (
	"fmt"

	"github.com/fr3fou/go-queue/queue"
)

func main() {
	q := queue.New()
	q.Enqueue(5)
	q.Enqueue(4)
	fmt.Println(q.Dequeue())
}
