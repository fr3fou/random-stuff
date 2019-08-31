package queue

// Queue is a FIFO data structure
type Queue struct {
	array []int
}

// New returns a queue instance
func New() *Queue {
	return &Queue{}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(n int) {
	q.array = append(q.array, n)
}

// Dequeue removes the first element from the array and returns it
func (q *Queue) Dequeue() int {
	el := q.array[0]
	q.array = q.array[1:]
	return el
}

// Peek returns the first element without removing it
func (q *Queue) Peek() int {
	return q.array[0]
}

// Contains checks if the given element is inside the array
func (q *Queue) Contains(n int) bool {
	for _, el := range q.array {
		if el == n {
			return true
		}
	}

	return false
}

// Clear clears all of the elements in the array
func (q *Queue) Clear() {
	q.array = []int{}
}

// Length returns the length of the queue
func (q *Queue) Length() int {
	return len(q.array)
}
