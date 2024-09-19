package main

import (
	"testing"
)

func BenchmarkSort(b *testing.B) {
	left := []int{1, 3, 5}
	right := []int{2, 4, 6}
	for i := 0; i < b.N; i++ {
		merge(left, right)
	}
}
