package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var numbers []int
	for i := 0; i < 5; i++ {
		numbers = append(numbers, rand.Intn(41)+10)
	}

	fmt.Println("Numbers:", numbers)

	var sum int
	for _, n := range numbers {
		sum += n
	}
	average := float64(sum) / float64(len(numbers))

	fmt.Println("Average:", average)
}
