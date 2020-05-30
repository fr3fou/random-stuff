package main

import "fmt"

func main() {
	arr := []int{4, 5, 12, 6, 3, 8, 11, 15, 1, 2}
	arr = mergeSort(arr)
	fmt.Println(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := len(arr) / 2
	a := mergeSort(arr[:pivot])
	b := mergeSort(arr[pivot:])
	return merge(a, b)
}

func merge(a []int, b []int) []int {
	length := len(a) + len(b)
	sorted := make([]int, 0, length-2)

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}

	// drain the last elements
	for i < len(a) {
		sorted = append(sorted, a[i])
		i++
	}

	// drain the last elements
	for j < len(b) {
		sorted = append(sorted, b[j])
		j++
	}

	return sorted
}
