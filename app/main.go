package main

import "fmt"

// Função para ordenar um array de inteiros
func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := Sort(arr[:middle])
	right := Sort(arr[middle:])

	return merge(left, right)
}

// Função para mesclar dois arrays
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Array antes de ordenar:", arr)
	arr = Sort(arr)
	fmt.Println("Array depois de ordenar:", arr)
}
