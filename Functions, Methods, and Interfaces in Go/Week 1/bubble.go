package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap : swap two elements in an array
func Swap(arr []int, index int) {
	temp := arr[index]
	arr[index] = arr[index+1]
	arr[index+1] = temp
}

// BubbleSort : sort an array in ascending order using bubble sort
func BubbleSort(arr []int) {
	// iterate till one less than length of array; array will always be sorted in last pass
	for i := 0; i < len(arr)-1; i++ {
		// iterate till ith element from the end; the last i elements will always be in the correct position
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {
	// take input of the numbers
	fmt.Println("Please enter the numbers in the array, separated by a space :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")

	// convert into int array
	var arr []int
	for _, value := range values {
		num, _ := strconv.Atoi(value)
		arr = append(arr, num)
	}

	// sort the array and convert it to a string array for printing
	BubbleSort(arr)
	for index, num := range arr {
		values[index] = strconv.Itoa(num)
	}

	// print the results
	fmt.Println("\nThe sorted array is :")
	fmt.Println(strings.Join(values, " "))
}
