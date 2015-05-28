package main

import (
	"fmt"
	"math"
)

func multiply(number []int, arg int) []int {

	carry := 0

	for i := len(number) - 1; i >= 0; i-- {
		digit := number[i]*arg + carry

		number[i] = digit % 10

		carry = digit / 10
	}
	return number
}

func power_2(power int) []int {
	tmplength := float64(power) * math.Log10(float64(2))
	length := int(math.Ceil(tmplength))
	result := make([]int, length, length)
	result[len(result)-1] = 2

	for i := 0; i < power-1; i++ {
		result = multiply(result, 2)
	}

	return result
}

func sumOfDigits(number []int) int {
	sum := 0

	for i := 0; i < len(number); i++ {
		sum += number[i]
	}

	return sum
}

func main() {

	fmt.Println("The sum of the digits of 2^100 is:", sumOfDigits(power_2(1000)))
}
