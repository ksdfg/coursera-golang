package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Name : struct to store names
type Name struct {
	fname string
	lname string
}

func main() {
	// take input of filepath
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter full path to file : ")
	scanner.Scan()
	path := scanner.Text()

	// read contents of the file, then split them by newline to get list of all names
	content, _ := ioutil.ReadFile(path)
	namesList := strings.Split(string(content), "\n")

	// make an slice of Names, then add all names read from file into it
	var names []Name
	for _, name := range namesList {
		if name != "" {
			temp := strings.Split(name, " ")
			names = append(names, Name{fname: temp[0], lname: temp[1]})
		}
	}

	// print names
	fmt.Println("\nNames in the given file are :")
	for _, name := range names {
		fmt.Println(name.fname + " " + name.lname)
	}
}
