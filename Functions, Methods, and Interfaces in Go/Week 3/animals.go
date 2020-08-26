package main

import (
	"fmt"
)

// Animal : base struct to describe an animal
type Animal struct {
	food, locomotion, sound string
}

// Eat : function to print what an animal eats
func (v Animal) Eat() {
	fmt.Println(v.food)
}

// Move : function to print how an animal moves
func (v Animal) Move() {
	fmt.Println(v.locomotion)
}

// Speak : function to print what sound an animal makes
func (v Animal) Speak() {
	fmt.Println(v.sound)
}

func main() {
	// make a map of all the possible animals
	m := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	// vars to store input in
	var animal, action string

	// infinite loop
	for {
		fmt.Print(">")
		fmt.Scan(&animal, &action)

		// call function according to given inputs
		if animal != "cow" && animal != "bird" && animal != "snake" {
			fmt.Println("Please give a valid animal")
		} else if action == "eat" {
			m[animal].Eat()
		} else if action == "move" {
			m[animal].Move()
		} else if action == "speak" {
			m[animal].Speak()
		} else {
			fmt.Println("Please give a valid action")
		}
	}
}
