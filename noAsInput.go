package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}
	result := num * 10
	resultString := strconv.Itoa(result)
	fmt.Println("Result:", resultString)
}
