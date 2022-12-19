package main

import "fmt"

func main() {
	var Fname, Lname string

	fmt.Print("enter first name : ")
	fmt.Scan(&Fname)

	fmt.Print("enter last name : ")
	fmt.Scan(&Lname)

	fmt.Println(Fname + " " + Lname)

}
