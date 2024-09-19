package main

import "fmt"

// // bubbleSort ordena um slice de inteiros
// func bubbleSort(arr []int) {
// 	n := len(arr)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// }

// metodo que faz o merge sort
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

	// Adiciona os elementos restantes
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
	// arr := []int{64, 34, 25, 12, 22, 11, 90}
	// fmt.Println("Array antes de ordenar:", arr)
	// bubbleSort(arr)
	// fmt.Println("Array depois de ordenar:", arr)

	left := []int{1, 3, 5}
	right := []int{2, 4, 6}
	merged := merge(left, right)
	fmt.Println(merged) // Output: [1 2 3 4 5 6]
}
