package main

import (
	"flag"
	"fmt"
	"strconv"
)

func NumberOfLetters(n int) int {
	//fmt.Println(n)
	if n < 20 {
		switch n {
		case 0:
			return 0
		case 1, 2, 6, 10:
			return 3
		case 4, 5, 9:
			return 4
		case 3, 7, 8:
			return 5
		case 11, 12:
			return 6
		case 13, 14, 18, 19:
			return 8
		case 15, 16:
			return 7
		case 17:
			return 9
		}

	} else if n >= 20 && n <= 99 {
		switch n / 10 {
		case 2, 3, 8, 9:
			return 6 + NumberOfLetters(n%10)
		case 4, 5, 6:
			return 5 + NumberOfLetters(n%10)
		case 7:
			return 7 + NumberOfLetters(n%10)
		}

	} else if n >= 100 && n <= 999 {
		nLetters := 0
		nLetters += NumberOfLetters(n / 100)
		nLetters += len("hundred")

		//fmt.Println("Mod is", n%100)
		if n%100 > 0 {
			nLetters += len("and")

			//fmt.Println("and")

			nLetters += NumberOfLetters(n % 100)
		}

		return nLetters

	} else if n >= 1000 && n < 10000 {
		nLetters := 0
		nLetters += NumberOfLetters(n / 1000)
		nLetters += len("thousand")

		//fmt.Println("thousand")

		remainder := n % 1000
		if remainder >= 1 {
			nLetters += NumberOfLetters(n % 1000)
		}

		return nLetters
	}

	return 0
}

func TotalNumberOfLetters(begin, end int) int {

	sum := 0

	for i := begin; i <= end; i++ {
		sum += NumberOfLetters(i)
	}

	return sum
}

func main() {
	flag.Parse()

	if len(flag.Arg(1)) > 0 {
		var arg, _ = strconv.Atoi(flag.Arg(0))
		fmt.Println(arg, "has:", NumberOfLetters(arg))
	}

	var begin = 1
	var end = 1000

	for i := 1; i <= 1000; i++ {
		fmt.Println(i, "-", NumberOfLetters(i))
	}

	fmt.Println("Total number of letters from 1 to 1000 is:", TotalNumberOfLetters(begin, end))
}
