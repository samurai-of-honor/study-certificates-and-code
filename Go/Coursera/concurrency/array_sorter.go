package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SubSort(n []int, c chan []int) {
	sort.Ints(n)
	fmt.Println(n)
	c <- n
}

func main() {
	/**/
	fmt.Println("Enter some integers, separated by space, and press enter to sort:")
	r := bufio.NewReader(os.Stdin)
	line, _, _ := r.ReadLine()

	strNums := strings.Split(string(line), " ")
	if len(strNums) < 4 {
		fmt.Println("Enter more numbers!")
		return
	}
	nums := make([]int, len(strNums))

	for i, val := range strNums {
		n, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("All numbers must be integer! Instead you entered: %s", val)
			return
		}
		nums[i] = n
	}
	/*
	 */
	// var nums = []int{2, 1, 4, 3, 6, 5, 8, 7}

	size := len(nums) / 4
	c := make(chan []int)
	fmt.Print("Sub sort:\n")
	for i := 0; i < 4; i++ {
		if i != 3 {
			go SubSort(nums[i*size:(i+1)*size], c)
		} else {
			go SubSort(nums[i*size:], c)
		}
	}

	sl1 := <-c
	sl2 := <-c
	sl3 := <-c
	sl4 := <-c

	var sorted []int
	sorted = append(sorted, sl1...)
	sorted = append(sorted, sl2...)
	sorted = append(sorted, sl3...)
	sorted = append(sorted, sl4...)

	sort.Ints(sorted)
	fmt.Print("Result:\n", sorted)
}
