package main

import (
	"fmt"
)

func main() {
	number1, number2 := 10, 3

	fmt.Println(number1, " = ", number2, " = ", number1 == number2)
	fmt.Println(number1, " != ", number2, " = ", number1 != number2)
	fmt.Println(number1, " > ", number2, " = ", number1 >= number2)
	fmt.Println(number1, " < ", number2, " = ", number1 <= number2)

}
