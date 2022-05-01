package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, adr string
	fmt.Print("Enter your name: \n")
	fmt.Scanln(&name)
	fmt.Print("Enter your address: \n")
	fmt.Scanln(&adr)
	m := map[string]string{
		"name":    name,
		"address": adr,
	}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", data)
}
