package main

import (
	"fmt"
)

const (
	JANUARY = iota
	FEBRUARY
	MARCH
	APRIL
	MAY
	JUNE
	JULY
	AUGUST
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER
)

//var monthsDaysArr = [...]int{31,28, 30, 31, 30}

func isNewYear(year int) bool {

	if ((year % 4 == 0) && (year % 100 != 0) || (year % 400 == 0)) {
		return true
	} else {
		return false
	}
}

func main() {
	
}