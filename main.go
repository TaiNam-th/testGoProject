package main

import (
	"fmt"
	"testGoProject/controllers"
)

func main() {
	fmt.Println("Hello, Go World!")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello, ", name)

	//Area Calculation
	fmt.Println("Area Calculation")
	controllers.CalculateandDisplayArea()
}
