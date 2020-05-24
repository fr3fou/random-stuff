package main

import "fmt"

func main() {
	arr := []int{4, 5, 12, 6, 3, 8, 11, 15, 1, 2}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, low, high int) {
	if low < high {
		pivot := (low + high) / 2
		mergeSort(arr, low, pivot-1)
		mergeSort(arr, pivot+1, high)
		merge(arr, low, pivot, high)
	}
}

func merge(arr []int, low, pivot, high int) {
	for _, num := range arr {

	}
}
