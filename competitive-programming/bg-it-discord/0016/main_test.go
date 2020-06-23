package main

import "testing"

var input = [][]int{
	{1, 2, 3, 4},
	{4, 4, 6},
	{6, 7, 8, 9},
	{1, 5, 9},
	{10, 11, 12},
	{12, 20, 80},
	{30, 50, 60, 80},
	{1, 43},
	{1, 1, 1, 1, 1, 1, 1, 1},
}

func BenchmarkGood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeK(input)
	}
}

func BenchmarkBad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeKBad(input)
	}
}
