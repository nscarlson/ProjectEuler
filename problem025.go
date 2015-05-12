package main

import(
	"math"
	"fmt"
)

func Ceil(f float64) int {
	return int(math.Ceil(f))
}

func main() {
	//
	// F(n) = PHI^n / sqrt(n) for large n
	//
	// Solving for n with F(n) = 10^999, the first 1000-digit integer
	// n = (log(sqrt5) + 10log(999)) / log(PHI)
	//
	
	n := Ceil((math.Log(math.Sqrt(5)) + 999 * math.Log(10)) / math.Log(math.Phi))

	fmt.Println(n, "is the first 1000-digit fibonacci number")
}