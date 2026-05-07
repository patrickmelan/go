package main

import (
	"fmt"
	"runtime"
)

// no values assigned
// CALLED: ZERO VALUES
// will default to 0, false, ""
// var a int, patrick bool, myStr string

func main() {

	// multi-initialize variables
	var a, b, message = 1, 6, "The answer is: %v"

	// test multi initialize with helper() call
	fmt.Println(helper(a, b, message))

	fmt.Println(convertToFloat(4))

	// user's system
	fmt.Println(runtime.GOOS)

	testSwitchStatement()

}

func helper(x, y int, msg string) (message string) {

	// assign x + y to num
	num := x + y

	// assign the string to message
	// don't need := because it was already declared as string in func declaration line
	// can use := when implicit type such as:
	//		num := 0 (initialized as int)
	message = fmt.Sprintf(msg, num)

	// naked return
	// ONLY USE IN SHORT FUNCTIONS LIKE THIS
	return
}

func convertToFloat(x int) float32 {
	// can convert from type to type by using T(v)
	// var i int = 67
	// var floatI float32 = float32(i)
	new := float32(x)
	const Constant = 4.0

	return (new * Constant)
}

func testSwitchStatement() {
	x := 42

	// wait til rest of function executes
	// LIFO order; last-in, first-out
	
	defer fmt.Println("first defer") // prints last
	defer fmt.Println("last defer")  // prints first

	switch x {
	case 10:
		fmt.Println("equal to 10.")
	case 41:
		fmt.Println("equal to 41.")
	default:
		fmt.Println("not in choices.")
	}
}
