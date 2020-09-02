package main

import (
	"fmt"
)

// Animal : base struct to describe an animal
type Animal struct {
	food, locomotion, sound string
}

// AnimalInterface : interface to describe the methods of animal
type AnimalInterface interface {
	Eat()
	Move()
	Speak()
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
	var command, animal, action string
	var generalAnimal AnimalInterface

	// infinite loop
	for {
		fmt.Print(">")
		fmt.Scan(&command, &animal, &action)

		// call function according to given inputs
		if command == "query" {
			generalAnimal = m[animal]
			switch action {
			case "eat":
				generalAnimal.Eat()
			case "move":
				generalAnimal.Move()
			case "speak":
				generalAnimal.Speak()
			default:
				fmt.Println("Please give a valid action")
			}
		} else if command == "newanimal" {
			m[animal] = m[action]
			fmt.Println("Created it!")
		}
	}
}
