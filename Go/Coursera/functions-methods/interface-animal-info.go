package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (a Cow) Eat() {
	fmt.Print("grass\n")
}

func (a Cow) Move() {
	fmt.Print("walk\n")
}

func (a Cow) Speak() {
	fmt.Print("moo\n")
}

type Bird struct{ name string }

func (a Bird) Eat() {
	fmt.Print("worms\n")
}

func (a Bird) Move() {
	fmt.Print("fly\n")
}

func (a Bird) Speak() {
	fmt.Print("peep\n")
}

type Snake struct{ name string }

func (a Snake) Eat() {
	fmt.Print("mice\n")
}

func (a Snake) Move() {
	fmt.Print("slither\n")
}

func (a Snake) Speak() {
	fmt.Print("hsss\n")
}

func main() {
	var newA Animal
	var comm, name, method string
	slice := make([]Animal, 0)
	for {
		fmt.Print("> ")
		_, err := fmt.Scanf("%s %s %s\n", &comm, &name, &method)
		if err != nil {
			fmt.Print("Argument failure. Try again\n")
			continue
		}

		if comm == "newanimal" {
			switch {
			case method == "cow":
				newA = Cow{name}
			case method == "bird":
				newA = Bird{name}
			case method == "snake":
				newA = Snake{name}
			default:
				fmt.Print("You entered a non-existent animal!\n")
				continue
			}
			slice = append(slice, newA)
			fmt.Print("Created it!\n")
		} else if comm == "query" {
			var need string
			for _, val := range slice {
				switch an := val.(type) {
				case Cow:
					need = an.name
				case Bird:
					need = an.name
				case Snake:
					need = an.name
				}
				if need == name {
					switch {
					case method == "eat":
						val.Eat()
					case method == "move":
						val.Move()
					case method == "speak":
						val.Speak()
					}
				}
			}
		} else {
			fmt.Print("You entered a non-existent argument. Try again\n")
			continue
		}
	}
}
