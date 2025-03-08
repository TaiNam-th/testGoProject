package views

import "fmt"

//Get User input
func GetUserInput() (float64, float64, float64) {
	var rai, ngan, wa float64
	fmt.Print("Enter rai: ")
	fmt.Scanln(&rai)
	fmt.Print("Enter ngan: ")
	fmt.Scanln(&ngan)
	fmt.Print("Enter wa: ")
	fmt.Scanln(&wa)
	return rai, ngan, wa
}

//Display Area Result
func DisplayAreaResult(sqm float64, km2 bool) {
	if km2 {
		fmt.Println("Area in Square Kilometers: ", sqm, " km²")
	} else {
		fmt.Println("Area in Square Meters: ", sqm, " m²")
	}
}
