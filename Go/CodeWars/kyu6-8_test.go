package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatDuration(t *testing.T) {
	var testCases = []struct {
		got    int64
		expect string
	}{
		{0, "now"},
		{1, "1 second"},
		{62, "1 minute and 2 seconds"},
		{120, "2 minutes"},
		{3600, "1 hour"},
		{3662, "1 hour, 1 minute and 2 seconds"},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, FormatDuration(tc.got))
	}
}

func TestListSquared(t *testing.T) {
	var testCases = []struct {
		gotM   int
		gotN   int
		expect [][]int
	}{
		{1, 250, [][]int{{1, 1}, {42, 2500}, {246, 84100}}},
		{250, 500, [][]int{{287, 84100}}},
		{300, 600, [][]int{}},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, ListSquared(tc.gotM, tc.gotN))
	}
}

func TestFindUniq(t *testing.T) {
	var testCases = []struct {
		got    []float32
		expect float32
	}{
		{[]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}, float32(2)},
		{[]float32{0, 0, 0.55, 0, 0}, float32(0.55)},
		{[]float32{3.0, 2.0, 2.0}, float32(3)},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, FindUniq(tc.got))
	}
}

func TestProductFib(t *testing.T) {
	var testCases = []struct {
		got    uint64
		expect [3]uint64
	}{
		{0, [3]uint64{0, 1, 1}},
		{800, [3]uint64{34, 55, 0}},
		{4895, [3]uint64{55, 89, 1}},
		{84049690, [3]uint64{10946, 17711, 0}},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, ProductFib(tc.got))
	}
}

func TestHumanReadableTime(t *testing.T) {
	var testCases = []struct {
		got    int
		expect string
	}{
		{59, "00:00:59"},
		{3599, "00:59:59"},
		{45296, "12:34:56"},
		{359999, "99:59:59"},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, HumanReadableTime(tc.got))
	}
}

func TestFindMissingLetter(t *testing.T) {
	var testCases = []struct {
		got    []rune
		expect rune
	}{
		{[]rune{'a', 'b', 'c', 'd', 'f'}, 'e'},
		{[]rune{'O', 'Q', 'R', 'S'}, 'P'},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, FindMissingLetter(tc.got))
	}
}

func TestDuplicateEncode(t *testing.T) {
	var testCases = []struct {
		got    string
		expect string
	}{
		{"din", "((("},
		{"recede", "()()()"},
		{"Success", ")())())"},
		{"(( @", "))(("},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, DuplicateEncode(tc.got))
	}
}

func TestWave(t *testing.T) {
	var testCases = []struct {
		got    string
		expect []string
	}{
		{" x yz", []string{" X yz", " x Yz", " x yZ"}},
		{"abc", []string{"Abc", "aBc", "abC"}},
		{"", []string{}},
		{"z", []string{"Z"}},
		{"a a a a a", []string{"A a a a a", "a A a a a", "a a A a a", "a a a A a", "a a a a A"}},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, Wave(tc.got))
	}
}

func TestAccum(t *testing.T) {
	var testCases = []struct {
		got    string
		expect string
	}{
		{"abcd", "A-Bb-Ccc-Dddd"},
		{"cwAt", "C-Ww-Aaa-Tttt"},
		{"NyffsGeyylB", "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb"},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, Accum(tc.got))
	}
}

func TestGetSum(t *testing.T) {
	var testCases = []struct {
		gotA   int
		gotB   int
		expect int
	}{
		{1, 5, 15},
		{2, 2, 2},
		{-2, 1, -2},
		{2, -1, 2},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, GetSum(tc.gotA, tc.gotB))
	}
}

func TestMinMax(t *testing.T) {
	var testCases = []struct {
		got    []int
		expect [2]int
	}{
		{[]int{8, 323, 6}, [2]int{6, 323}},
		{[]int{1}, [2]int{1, 1}},
		{[]int{-5, -4, -2, -6}, [2]int{-6, -2}},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expect, MinMax(tc.got))
	}
}
