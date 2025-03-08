package main

import "fmt"

func main() {
	fmt.Println("Hello, Go World!")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello, ", name)
}
