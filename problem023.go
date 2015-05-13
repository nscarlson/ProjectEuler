//
//	Problem 23
//	Non-abundant sums
//

package main

import (
	"fmt"
	//"time"
	//"log"
)

const (
	LIMIT	:= 28123
)

type List struct {
	data []int 
}

var abundantList = List{make([]int, 0, 2600)}

func(list *List) append(n int) {
	list.data = append(list.data, n)
}

//
// Returns the sum of all
//

func sliceSum (divisors []int) int {
	var sum int

	for _, i := range divisors {
		sum += i
	}

	return sum
}





func bAbundant(n int) bool {

	//start := time.Now()

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

	//
	// If sum > n, it's abundant
	// Else it is either perfect or deficient
	//

	sum = sliceSum(divisors)

	//elapsed := time.Since(start)

	if sum > n {
		//fmt.Println("Binomial took %s", elapsed)
		return true
	} else {
		//fmt.Println("Binomial took %s", elapsed)
		return false
	}
}



func sumOfNonAbundantSums() int {
	//
	//	The first odd abundant number is 945
	//

	//first abundant number

	sum := 0
	//sums := []int{}

	for i := 24; i < LIMIT; i++ {
		if i % 2 == 0 && i > 46 {

		}
	}

	return sum
}

func main() {
	//
	//	Initialize all abundants < 28123
	//
	for i := 12; i < LIMIT; i++ {

		if bAbundant(i) {
			
			abundantList.append(i)
		}
	}


	for n := 2; n < LIMIT {

	}
}

