package main

import (
	"fmt"
	"math/rand"
)

func main() {

	num := rand.Intn(10)
	guesses := 0

	fmt.Println("Welcome to my guessing game! Pick a number between 1 and 10.")
	fmt.Println("BTW. You get 3 guesses.\n")
	fmt.Println(num)

	for guesses < 3 {
		correct, msg := askInput(num)

		guesses++

		if !correct {
			fmt.Println(msg)
		} else {
			fmt.Println(msg)
			fmt.Printf("You got it in %v guesses!", guesses)
			break
		}
	}

}

func askInput(target int) (bool, string) {

	var guess int

	fmt.Print("Guess: ")
	fmt.Scan(&guess)

	if guess == target {
		return true, "Correct answer!"
	} else if guess < target {
		return false, "Target number is higher!"
	} else {
		return false, "Target number is lower!"
	}

}
