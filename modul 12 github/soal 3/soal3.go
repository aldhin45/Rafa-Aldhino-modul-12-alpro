package main

import "fmt"

func selectionSort(arr []int) {
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

func findMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	}
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	var num int
	var angka []int

	for {
		fmt.Scan(&num)
		if num == -5313 {
			break
		}
		if num != 0 {
			angka = append(angka, num)
		} else {
			if len(angka) > 0 {
				selectionSort(angka)
				median := findMedian(angka)
				fmt.Println(median)
			}
		}
	}
}
