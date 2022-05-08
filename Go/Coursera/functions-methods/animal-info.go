package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func main() {
	aMap := map[string]Animal{
		"cow": {
			"grass",
			"walk",
			"moo",
		},
		"bird": {
			"worms",
			"fly",
			"peep",
		},
		"snake": {
			"mice",
			"slither",
			"hsss",
		}}
	var who Animal
	var name, method, res string

	for {
		fmt.Print("Enter animal name(cow, bird or snake):\n> ")
		fmt.Scan(&name)
		if val, ok := aMap[name]; ok {
			who = val
		} else {
			fmt.Print("You entered a non-existent animal. Try again\n\n")
			continue
		}

		fmt.Print("Enter the required info(eat, move or speak):\n> ")
		fmt.Scan(&method)
		switch {
		case method == "eat":
			res = who.Eat()
		case method == "move":
			res = who.Move()
		case method == "speak":
			res = who.Speak()
		default:
			fmt.Print("You entered a non-existent method!\n\n")
			continue
		}

		fmt.Printf("-----------------------\n %s %s: \"%s\"\n-----------------------\n", name, method, res)
	}
}
