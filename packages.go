package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("hi patrick. my random number is", rand.Intn(10))

	var greater5 float32 = 0
	var less5 float32 = 0
	x := 0

	for x < 100000 {
		if rand.Intn(100) > 50 {
			greater5++
		} else {
			less5++
		}

		x++
	}

	fmt.Printf("> 50: %.2f%% of the time\n", (greater5/100000)*100)
	fmt.Printf("< 50: %.2f%% of the time", (less5/100000)*100)
}
