package llist

type node struct {
	value int
	next  *node
}

// AddToEnd adds an element to the th end of the node "chain" - this is O(n)
func (n *node) AddToEnd(v int) {
	if n.next == nil {
		n.next = &node{
			value: v,
		}
		return
	}

	n.next.AddToEnd(v)
}

type LinkedList struct {
	head *node
}

// New returns a pointer to a new linked list
func New(v int) *LinkedList {
	return &LinkedList{
		head: &node{
			value: v,
		},
	}
}

// Add adds an element to the beginning of the LinkedList - this is O(1)
func (l *LinkedList) Add(v int) {
	next := l.head.next
	l.head.next = &node{value: v}
	l.head.next.next = next
}

// AddToEnd adds an element to the th end of the LinkedList - this is O(n)
func (l *LinkedList) AddToEnd(v int) {
	if l.head.value == 0 {
		l.head.value = v
		return
	}

	if l.head.next == nil {
		l.head.next = &node{
			value: v,
		}
		return
	}

	l.head.next.AddToEnd(v)
}

// ToArray turns the linked list to an array
func (l *LinkedList) ToArray() []int {
	a := []int{}

	for n := l.head; n != nil; n = n.next {
		a = append(a, n.value)
	}

	return a
}

// Reverse reverses the elements of the linked list
func (l *LinkedList) Reverse() {
	var (
		prev, current, next *node
	)

	prev, current, next = nil, l.head, l.head.next

	for next != nil {
		current.next = prev
		temp := next.next
		next.next = current

		prev, current, next = current, next, temp
	}

	// change the head
	l.head = current
}
