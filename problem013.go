package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func nFirstDigitsOfSum(nDigits int, str []string) string {
	nNumbers := len(str)
	sum := 0

	for i := 0; i < len(str); i++ {
		length := len(str[i])
		if nNumbers == 1 {
			if nDigits > length {
				return str[i][:len(str[0])]
			} else {
				return str[i][:nDigits]
			}
		}

		if nNumbers > 1 && nNumbers <= 100 {
			n, err := strconv.Atoi(str[i][:nDigits+2])
			check(err)

			sum += n
		}
	}

	return strconv.Itoa(sum)[:nDigits]
}

func main() {
	file, err := os.Open("problem013.txt")
	check(err)

	// close on exit
	defer func() {
		err := file.Close()
		check(err)
	}()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	nDigits := len(lines[0])
	nNumbers := len(lines)

	fmt.Println("Number of digits:", nDigits)
	fmt.Println("Number of numbers:", nNumbers)

	fmt.Println("The 10 first digits of the large sum are:", nFirstDigitsOfSum(10, lines))
}
