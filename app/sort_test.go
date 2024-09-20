package main

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(100)
		Sort(arr)
	}
}
