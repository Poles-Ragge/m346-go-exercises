package main

import (
	"fmt"
)

func main() {

	var firstName, lastName = "Ano", "nym"
	var dayOfBirth, monthOfBirth, yearOfBirth = 01, 01, 1970
	var numberOfSiblings int = 0
	var heightInMeters float32 = 1.970
	var zodiacSign = '\u2650'

	fmt.Printf("Vor- und Nachname: %s %s \n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
