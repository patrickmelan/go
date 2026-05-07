package main

import (
	"fmt"
)

func main() {

	var choice int

	for {
		fmt.Println("What do you want to convert?")
		fmt.Print("(1) Temperature  (2) Weight  (3) Distance  (0) Quit:  ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handleTemp()
		case 2:
			handleWeight()
		case 3:
			handleDistance()
		case 0:
			break
		default:
			fmt.Println("Incorrect input. Try again.\n")
		}

		if choice == 0 {
			break
		}
	}

}

func handleTemp() {

	var choice int
	var temp float32

	fmt.Print("(1) F >> C | (2) C >> F: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter the temperature you want to convert to Celsius: ")
		fmt.Scan(&temp)
		fmt.Println(toCelsius(temp))
	case 2:
		fmt.Print("Enter the temperature you want to convert to Fahrenheit: ")
		fmt.Scan(&temp)
		fmt.Println(toFahrenheit(temp))
	}

}

func handleWeight() {

	var choice int
	var weight float32

	fmt.Print("(1) Pounds (lb) >> Grams (g) | (2) Grams (g) >> Pounds (lb): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter the weight you want to convert to grams: ")
		fmt.Scan(&weight)
		fmt.Println(toGrams(weight))
	case 2:
		fmt.Print("Enter the weight you want to convert to pounds: ")
		fmt.Scan(&weight)
		fmt.Println(toPounds(weight))
	}

}

func handleDistance() {

	var choice int
	var distance float32

	fmt.Print("(1) Miles (mi) >> Meters (m) | (2) Meters(m) >> Miles(mi): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter the distance you want to convert to meters: ")
		fmt.Scan(&distance)
		fmt.Println(toMeters(distance))
	case 2:
		fmt.Print("Enter the distance you want to convert to miles: ")
		fmt.Scan(&distance)
		fmt.Println(toMiles(distance))
	}

}

func toFahrenheit(tempC float32) float32 {
	return (tempC * 9 / 5) + 32
}

func toCelsius(tempF float32) float32 {
	return (tempF - 32) * 5 / 9
}

func toGrams(weightLbs float32) float32 {
	return weightLbs * 453.6
}

func toPounds(weightGrams float32) float32 {
	return weightGrams / 453.6
}

func toMeters(distMiles float32) float32 {
	return distMiles * 1609.34
}

func toMiles(distMeters float32) float32 {
	return distMeters / 1609.34
}
