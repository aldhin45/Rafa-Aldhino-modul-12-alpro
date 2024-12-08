package main

import (
	"fmt"
)

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("jumlah daerah  Hercules (n): ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf(" %d: ", i+1)

		var m int
		fmt.Scan(&m)

		angka := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&angka[j])
		}

		selectionSortDesc(angka)
		fmt.Printf(" %d: %v\n", i+1, angka)
	}
}
