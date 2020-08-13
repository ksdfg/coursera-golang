package main

import "fmt"

// GenDisplaceFn : generate a function that will calculate displacement after given period of time
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		fmt.Println(a * t * t / 2)
		return s0 + (v0 * t) + (a * t * t / 2)
	}
}

func main() {
	// take input of acceleration, initial velocity and initial displacement
	var a, v0, s0 float64
	fmt.Print("Enter Acceleration : ")
	fmt.Scanln(&a)
	fmt.Print("Enter Initial Velocity : ")
	fmt.Scanln(&v0)
	fmt.Print("Enter Initial Displacement : ")
	fmt.Scanln(&s0)

	// generate function to calculate displacement after given period of time
	DisplaceFn := GenDisplaceFn(a, v0, s0)

	// take input of time
	var t float64
	fmt.Print("Enter time passed : ")
	fmt.Scanln(&t)

	// print the displacement
	fmt.Printf("The amount of displacement in given time is %f units\n", DisplaceFn(t))
}
