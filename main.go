package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var cal string
	var result, result2 float64

	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&cal)

	if cal == "+" {
		result = num1 + num2
		result2 = math.Floor(result*1000) / 1000
		fmt.Println(num1, "+", num2, "=", result2)
	} else if cal == "-" {
		result = num1 - num2
		result2 = math.Floor(result*1000) / 1000
		fmt.Println(num1, "-", num2, "=", result)
	} else if cal == "*" {
		result = num1 * num2
		result2 = math.Floor(result*1000) / 1000
		fmt.Println(num1, "*", num2, "=", result)
	} else if cal == "/" {
		result = num1 / num2
		result2 = math.Floor(result*1000) / 1000
		fmt.Println(num1, "/", num2, "=", result)
	} else if cal == "%" {
		result = math.Mod(num1, num2)
		result2 = math.Floor(result*1000) / 1000
		fmt.Println(num1, "%", num2, "=", result2)
	} else {
		fmt.Println("Error")
	}
}
