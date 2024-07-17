package main

import (
	"fmt"
	"math/rand"
)

func main() {

	num1 := rand.Intn(100)

	if num1%3 == 0 && num1%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if num1%3 == 0 {
		fmt.Println("Fizz")
	} else if num1%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Printf("Else number: %v", num1)
	}

}
