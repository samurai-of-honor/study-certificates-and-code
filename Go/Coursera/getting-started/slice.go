package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var str string
	slice := make([]int, 0, 3)
	fmt.Println("Enter integer or X to quit...")
	for fmt.Scan(&str); str != "X"; fmt.Scan(&str) {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic("Enter integer!")
		}
		slice = append(slice, n)
		sort.Slice(slice, func(i, j int) bool {
			return slice[j] > slice[i]
		})
		for _, val := range slice {
			fmt.Print(val, " ")
		}
		fmt.Print("\n")
	}
}
