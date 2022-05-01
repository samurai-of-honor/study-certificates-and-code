package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var title string
	fmt.Println("Enter the name of the file with first and last names,\nwritten on a new line and separated by a space: ")
	// title = "read_file.txt"
	fmt.Scanln(&title)
	file, err := os.Open(title)
	if err != nil {
		panic("file does not exist")
	}
	defer file.Close()

	type Person struct {
		fname string
		lname string
	}
	slice := make([]Person, 0)

	str := bufio.NewScanner(file)
	for str.Scan() {
		name := strings.Split(str.Text(), " ")
		p := Person{
			fname: name[0],
			lname: name[1],
		}
		slice = append(slice, p)
	}
	for i, val := range slice {
		fmt.Printf("%d. %s %s \n", i+1, val.fname, val.lname)
	}
}
