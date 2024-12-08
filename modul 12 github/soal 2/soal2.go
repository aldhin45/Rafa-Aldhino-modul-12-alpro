package main

import "fmt"

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

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat Hercules (n): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf(" %d: ", i+1)

		var m int
		fmt.Scan(&m)

		angka := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&angka[j])
		}

		var odd, even []int
		for _, num := range angka {
			if num%2 == 0 {
				even = append(even, num)
			} else {
				odd = append(odd, num)
			}
		}

		selectionSortDesc(odd)
		selectionSortAsc(even)

		fmt.Printf(" %d: ", i+1)
		for _, num := range odd {
			fmt.Printf("%d ", num)
		}
		for _, num := range even {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
