//
//	Problem 23
//	Non-abundant sums
//

package main

import (
	"fmt"
)

const (
	LIMIT	=	28123
)

func sliceSum (slice []int) int {
	var sum int

	for _, i := range slice {
		sum += i
	}

	return sum
}

func bAbundant(n int) bool {

	divisors := []int{}
	sum := 0

	//
	//	Populate n's list of proper divisors
	//

	for i := 1; i < n / 2 + 1; i++ {
		if n % i == 0 {
			divisors = append(divisors, i)
		}
	}

	sum = sliceSum(divisors)

	if sum > n {
		return true
	} else {
		return false
	}
}

func main() {
	sum := 0

	abundantList 			:= []int{}
	canBeWrittenAsAbundant 	:= make([]bool, LIMIT + 1)


	//
	//	Initialize all abundants < 28123
	//

	for i := 1; i < LIMIT; i++ {

		if bAbundant(i) {
			
			abundantList = append(abundantList, i)
		}
	}

	//
	//	Initialize list of all sums of abundants < 28123
	//

	for i := 0; i < len(abundantList); i++ {
		for j := i; j < len(abundantList); j++ {
			if (abundantList[i] + abundantList[j] <= LIMIT) {
				canBeWrittenAsAbundant[abundantList[i] + abundantList[j]] = true
			} else {
				break
			}
		}
	}


	//
	//	Sum all that are NOT sums of abundants
	//

	for i := 1; i <= LIMIT; i++ {
		if !canBeWrittenAsAbundant[i] {
			sum += i
		}
	}

	fmt.Println("The sum of all integers not expressible as a sum of two abundants is:", sum)
}