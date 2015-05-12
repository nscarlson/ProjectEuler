//
//	Non-abundant sums
//

//
//	A number n is called deficient if the sum of its proper divisors is less than n,
//	and it is called abundant if this sum exceeds n.
//

package main

import (
	"fmt"
)

func sliceSum (divisors []int) int {
	var sum int

	for _, i := range divisors {
		sum += i
	}

	return sum
}

func bSumOfTwoAbundants(n) bool {

}

func bAbundant(n int) bool {
	divisors := []int{}

	//
	//	Populate n's list of proper divisors
	//

	for i := 1; i < n / 2; i++ {
		if n % i == 0 {
			divisors = append(divisors, i)

			fmt.Println(i, divisors, sliceSum(divisors))
		}
	}

	//
	// If sum > n, it's abundant
	// Else it is either perfect or deficient
	//

	fmt.Println("Sum is:", sliceSum(divisors))
	return sliceSum(divisors) > n

}

func sumOfNonAbundantSums {
	var lastAbundant int
	
}

func main() {
	fmt.Println("945 Abundant?", abundant(945))

}