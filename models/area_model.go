package models

//Converts to square meters แปลง ไร่, งาน, ตรว. เป็น ตร.เมตร
func ConvertToSquareMeters(rai, ngan, wa float64) float64 {
	return (rai * 1600) + (ngan * 400) + (wa * 4)
}

//Converts to square Kilometers แปลง ไร่, งาน, ตรว. เป็น ตร.กม if more than 1,000,000 square meters
func ConvertToSquareKilometers(sqm float64) float64 {
	return sqm / 1000000
}
