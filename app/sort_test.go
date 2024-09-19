package main

import (
	"testing"
)

func BenchmarkSort(b *testing.B) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	for i := 0; i < b.N; i++ {
		bubbleSort(arr)
	}
}
