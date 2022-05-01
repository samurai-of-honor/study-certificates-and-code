package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	fmt.Scan(&in)
	s := strings.ToLower(in)
	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
