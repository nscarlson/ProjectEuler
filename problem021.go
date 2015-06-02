package main

import (
	"fmt"
	"math"
)

//
//	If D(a) = b and D(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
//

func IsAmicable(a int) bool {
	b := D(a)

	if D(b) == a && a != b {
		return true
	}

	return false
}

//
//	Let D(n) be defined as the sum of proper divisors of n
//

func D(n int) int {

	//
	//	1 is a proper divisor of every integer
	//

	sum := 1

	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			sum += (n/i + i)
		}
	}

	return sum
}

//
//	Driver function.  Evaluate the sum of all the amicable numbers under 10000
//

func main() {
	var sum int64 = 0

	for i := 1; i < 10000; i++ {
		if IsAmicable(i) {
			sum += int64(i)
		}
	}

	fmt.Println("The sum of amicable numbers under 10000 is", sum)
}
