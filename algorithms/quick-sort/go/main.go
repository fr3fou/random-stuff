package main

import "fmt"

func main() {
	arr := []int{4, 5, 12, 6, 3, 8, 11, 15, 1, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low int, high int) (i int) {
	pivot := arr[high]
	minIndex := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			minIndex++
			// swap
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
	// swap
	arr[minIndex+1], arr[high] = arr[high], arr[minIndex+1]
	return minIndex + 1

}
