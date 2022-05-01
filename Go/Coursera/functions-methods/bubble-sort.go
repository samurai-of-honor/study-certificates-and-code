package main

import "fmt"

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
	fmt.Println(sli)
}

func Swap(sli []int, j int) {
	sli[j], sli[j+1] = sli[j+1], sli[j]
}

func main() {
	var n int
	slice := make([]int, 0)
	fmt.Print("Enter int element or '0' to sort: ")
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		slice = append(slice, n)
		fmt.Print("Enter int element or '0' to sort: ")
	}
	BubbleSort(slice)
}
