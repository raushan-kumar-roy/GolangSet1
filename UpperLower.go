package main

import (
	"fmt"
	"strings"
)

func main() {
	var place string
	fmt.Print("enter the place : ")
	fmt.Scan(&place)
	var str_lower = strings.ToLower(place)
	var str_upper = strings.ToUpper(place)

	fmt.Println("place i like to visit the most is : ", str_lower)
	fmt.Println("place i like to visit the most is : ", str_upper)

}
