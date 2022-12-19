package main

import "fmt"

func main() {
	var fahrenheit, centigrade float32
	fmt.Print("enter degree in c : ")
	fmt.Scan(&centigrade)
	fahrenheit = centigrade*(9/5) + 32
	fmt.Println("fahrenheit : ", fahrenheit)
}
