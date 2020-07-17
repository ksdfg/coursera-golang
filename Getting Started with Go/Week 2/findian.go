package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// take input of the string
	fmt.Print("Enter a string : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	// check if string satisfies conditions using regex
	matched, _ := regexp.MatchString("^[iI].*[aA].*[nN]$", str)

	// print output accordingly
	if matched {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
