package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// initialize slice and string to  take user input
	slice := make([]int, 0, 3)
	var in string

	for true {
		// print the slice
		fmt.Println(slice)

		// take input
		fmt.Print("Enter a number, or X to exit : ")
		fmt.Scanln(&in)

		// check input to decide what to do
		if in == "X" || in == "x" {
			break
		}
		elt, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Please do not enter anything other than a number or 'X'")
			continue
		}

		// append the integer to the slice and sort
		slice = append(slice, elt)
		sort.Ints(slice)
	}
}
