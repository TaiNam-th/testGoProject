package controllers

import (
	"testGoProject/models"
	"testGoProject/views"
)

// CalculateArea calculates the area of land and displays
func CalculateandDisplayArea() {
	rai, ngan, wa := views.GetUserInput()
	sqm := models.ConvertToSquareMeters(rai, ngan, wa)
	km2 := false

	//Check if area is more than 1,000,000 square meters
	if sqm > 1000000 {
		km2 = true
		sqm = models.ConvertToSquareKilometers(sqm)
	}
	views.DisplayAreaResult(sqm, km2)
}
