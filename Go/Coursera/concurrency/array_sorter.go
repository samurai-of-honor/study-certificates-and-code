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
	fmt.Println(n)
	sort.Ints(n)
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

	for k := 0; k < len(nums); k++ {
		i := 0
		j := 0
		if len(sl1) != 0 {
			i = sl1[0]
			j = 1
		}
		if len(sl2) != 0 {
			if j == 0 {
				i = sl2[0]
				j = 2
			} else {
				if sl2[0] < i {
					i = sl2[0]
					j = 2
				}
			}
		}
		if len(sl3) != 0 {
			if j == 0 {
				i = sl3[0]
				j = 3
			} else {
				if sl3[0] < i {
					i = sl3[0]
					j = 3
				}
			}
		}
		if len(sl4) != 0 {
			if j == 0 {
				i = sl4[0]
				j = 4
			} else {
				if sl4[0] < i {
					i = sl4[0]
					j = 4
				}
			}
		}

		sorted = append(sorted, i)
		switch j {
		case 1:
			sl1 = sl1[1:]
		case 2:
			sl2 = sl2[1:]
		case 3:
			sl3 = sl3[1:]
		case 4:
			sl4 = sl4[1:]
		}
	}

	fmt.Print("Result:\n", sorted)
}
