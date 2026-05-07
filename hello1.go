package main

// go run hello.go

import (
	"fmt"
	//"time"
	//"math/rand"
)

func main() {

	conversions := 0

	var input float32
	var numToConvert float32

	for {

		fmt.Print("Convert: (1) C → F  (2) F → C  (0) Quit: ")
		fmt.Scan(&input)

		if input == 0 {
			fmt.Printf("Done! %v conversions made.", conversions)
			break
		} else if input == 1 {

			fmt.Print("Enter Temperature in Celsius (°C): ")
			fmt.Scan(&numToConvert)

			numF := toFahrenheit(numToConvert)
			conversions++

			formatOutput(numToConvert, numF, true)
		} else if input == 2 {

			fmt.Print("Enter temperature in Fahrenheit (°F): ")
			fmt.Scan(&numToConvert)
			numC := toCelsius(numToConvert)
			conversions++

			formatOutput(numC, numToConvert, false)
		} else {
			fmt.Println("Incorrect input. Please try again!")
		}

	}

}

func toFahrenheit(tempC float32) float32 {
	return (tempC * 9 / 5) + 32
}

func toCelsius(tempF float32) float32 {
	return (tempF - 32) * 5 / 9
}

func formatOutput(c float32, f float32, direction bool) {
	if direction {
		fmt.Printf("%.2f°C >>> %.2f°F\n\n", c, f)
	} else {
		fmt.Printf("%.2f°F >>> %.2f°C\n\n", f, c)
	}
}
