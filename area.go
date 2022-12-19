package main

import "fmt"

func main() {
	var area, perimeter, radius, diameter float32
	fmt.Print("diameter : ")
	fmt.Scan(&diameter)
	radius = diameter / 2
	fmt.Println("Radius : ", radius)
	perimeter = 2 * 22 / 7 * radius
	fmt.Println("perimeter : ", perimeter)
	area = 22 / 7 * radius * radius
	fmt.Println("area : ", area)
}
