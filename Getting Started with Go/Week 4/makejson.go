package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// make scanner to take input
	scanner := bufio.NewScanner(os.Stdin)

	// Take input of name
	fmt.Print("Enter Name : ")
	scanner.Scan()
	name := scanner.Text()

	// Take input of address
	fmt.Print("Enter Address : ")
	scanner.Scan()
	address := scanner.Text()

	// Make map and store details
	details := make(map[string]string)
	details["name"] = name
	details["address"] = address

	// Convert to JSON object and print
	jsonObject, _ := json.Marshal(details)
	fmt.Println("\nThe created JSON object is\n" + string(jsonObject))
}
