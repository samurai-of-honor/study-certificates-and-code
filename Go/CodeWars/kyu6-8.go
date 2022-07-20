package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// FormatDuration formats the number of seconds into an easy-to-read form
func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	var (
		res       []string
		formatRes string
		times     = [5]string{" second", " minute", " hour", " day", " year"}
	)

	dur := [5]int64{
		seconds % 60,
		seconds % 3600 / 60,
		seconds % 86400 / 3600,
		seconds % 31536000 / 86400,
		seconds / 31536000,
	}

	for i, v := range dur {
		if v == 1 {
			res = append(res, strconv.Itoa(int(v))+times[i])
		} else if v > 1 {
			res = append(res, strconv.Itoa(int(v))+times[i]+"s")
		}
	}

	switch len(res) {
	case 1:
		return res[0]
	case 2:
		return res[1] + " and " + res[0]
	default:
		for i := len(res) - 1; i > 1; i-- {
			formatRes += res[i] + ", "
		}
		formatRes += res[1] + " and " + res[0]
	}
	return formatRes
}

// ListSquared returns a number and the sum of the squares of its divisors
// if this sum is the square of any integer number
func ListSquared(m, n int) [][]int {
	var res = [][]int{}

	for m <= n {
		var sum int
		for i := 1; i <= n; i++ {
			if m%i == 0 {
				sum += i * i
			}
		}

		_, r := math.Modf(math.Sqrt(float64(sum)))
		if r == 0 {
			res = append(res, []int{m, sum})
		}
		m++
	}
	return res
}

// FindUniq searches in a slice of identical numbers unique
func FindUniq(arr []float32) float32 {
	if arr[1] != arr[0] && arr[2] != arr[0] {
		return arr[0]
	}
	var cache = [1]float32{arr[0]}
	for _, v := range arr[1:] {
		if v != cache[0] {
			return v
		}
	}
	return arr[0]
}

// ProductFib finds out whether the passed number is a product of numbers from the
// Fibonacci series, if not, displays the closest to the desired ones
func ProductFib(prod uint64) [3]uint64 {
	var i uint64
	cache := make(map[uint64]uint64)
	first := fibonacci(cache, i)
	next := fibonacci(cache, i+1)
	for first*next < prod {
		i++
		first, next = fibonacci(cache, i), fibonacci(cache, i+1)
	}
	if first*next == prod {
		return [3]uint64{first, next, 1}
	}
	return [3]uint64{first, next, 0}
}

// sequence calculation with caching
func fibonacci(cache map[uint64]uint64, n uint64) uint64 {
	if n < 2 {
		return n
	}
	if result, ok := cache[n]; ok {
		return result
	}

	result := fibonacci(cache, n-1) + fibonacci(cache, n-2)
	cache[n] = result
	return result
}

// HumanReadableTime formats the number of seconds into an easy-to-read form
func HumanReadableTime(seconds int) string {
	h := seconds / 3600
	m := seconds % 3600 / 60
	s := seconds % 60
	return fmt.Sprintf("%.2d:%.2d:%.2d", h, m, s)
}

// FindMissingLetter looking for a missing letter of the alphabet
func FindMissingLetter(chars []rune) rune {
	char := chars[0]
	for i := range chars {
		if chars[i] != char {
			break
		}
		char++
	}
	return char
}

// DuplicateEncode encodes strings depending on whether they contain repeated letters
func DuplicateEncode(word string) string {
	var res string
	s := strings.ToLower(word)
	for i := range s {
		if strings.Count(s, string(s[i])) > 1 {
			res += ")"
		} else {
			res += "("
		}
	}
	return res
}

// Wave creates a wave from the entered string
func Wave(words string) []string {
	var res = []string{}
	for i := range words {
		s := strings.Split(words, "")
		if s[i] != " " {
			s[i] = strings.ToUpper(s[i])
			res = append(res, strings.Join(s, ""))
		}
	}
	return res
}

// Accum repeats with a capital letter an increasing number of times each letter
func Accum(s string) string {
	var res string
	s = strings.ToLower(s)
	str := strings.Split(s, "")
	for n, v := range str {
		res += strings.ToUpper(v)
		for i := 2; i <= n+1; i++ {
			res += v
		}
		if n != len(str)-1 {
			res += "-"
		}
	}
	return res
}

// GetSum return sum of all numbers from a to b
func GetSum(a, b int) int {
	var sum, min, max int
	if a == b {
		return a
	}
	if a > b {
		min, max = b, a
	} else {
		min, max = a, b
	}
	for min <= max {
		sum += min
		min++
	}
	return sum
}

// MinMax returns the minimum and maximum number from the slice
func MinMax(arr []int) [2]int {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return [2]int{min, max}
}
