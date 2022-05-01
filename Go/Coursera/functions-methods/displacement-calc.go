package main

import "fmt"

// Formula:
// s = 1/2 at^2 + v0t + s0

func main() {
	var a, t, v0, s0 float64
	fmt.Println("Enter value in this format: a v0 s0")
	_, err := fmt.Scanf("%f %f %f\n", &a, &v0, &s0)
	if err != nil {
		fmt.Println("Try again:", err)
		return
	}

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Print("Enter t: ")
	for {
		fmt.Scan(&t)
		fmt.Printf("Result: s = %.2f\nEnter another t: ", fn(t))
	}
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		s := (a*t*t)/2 + v0*t + s0
		return s
	}
	return fn
}
