package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now.Day(), "/", now.Month(), "/", now.Year())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())

}
