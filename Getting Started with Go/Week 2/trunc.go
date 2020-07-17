package main

import (
	"fmt"
)

func main() {
	// take input of the floating point integer
	var x float64
	fmt.Print("Enter floating point number : ")
	_, _ = fmt.Scan(&x)

	// truncate floating point to int
	var xInt = int64(x)

	// print the truncated int
	fmt.Printf("Truncated floating point number is : %d\n", xInt)
}
