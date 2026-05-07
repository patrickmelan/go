package main

import (
	"fmt"
)

func main() {

	num, err := testFunc()

	if err == nil {
		fmt.Println(num)
	}

}

func testFunc() (int8, error) {
	return 3, nil
}
