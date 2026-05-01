package main

import "fmt"

func main2() {
	// fmt.Println("hello")

	// declare celsius variable
	//var c float64

	//fmt.Println("Enter degrees Celsius (C):")

	// take user input using Scan()
	//fmt.Scan(&c)

	// program output using toFahrenheit(c) to convert to degrees F

	/*
		num := 0
		var x int

		fmt.Scan(&x)

		for num < x {
			if num%2 == 0 {
				fmt.Println("even")
			} else {
				fmt.Println("odd")
			}
			num++
		}*/

	arr := []string{"lebron", "stephon", "anthony", "labaron"}

	for _, value := range arr {
		fmt.Printf("%v is in the array\n", value)
	}

	fmt.Println(arr)

}

// declare func, VARNAME TYPE PARAMS, have the return type after params
func toFahrenheit2(tempCel float64) float64 {
	return (tempCel * 9 / 5) + 32
}
