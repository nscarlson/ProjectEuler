package main

import (
	"fmt"
)

const (
	JANUARY 	= 0
	FEBRUARY	= 1
	MARCH		= 2
	APRIL		= 3
	MAY			= 4
	JUNE		= 5
	JULY		= 6
	AUGUST		= 7
	SEPTEMBER	= 8
	OCTOBER		= 9
	NOVEMBER	= 10
	DECEMBER	= 11

	SUNDAY		= 0
	MONDAY		= 1
	TUESDAY		= 2
	WEDNESDAY	= 3
	THURSDAY	= 4
	FRIDAY		= 5
	SATURDAY	= 6

	EPOCH_YEAR	= 1900
	EPOCH_MONTH	= 0
	EPOCH_DAY	= 1


)

var MONTHS_DAYS_ARR = [...]int {
	31,	// January
	28, // February
	31, // March
	30,	// April
	31,	// May
	30, // June
	31, // July
	31, // August
	30, // September
	31, // October
	30, // November
	31, // December
}

var WEEKDAY_NAMES_ARR = [...]string {
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func isLeapYear(year int) bool {
	if ((year % 4 == 0) && (year % 100 != 0) || (year % 400 == 0)) {
		return true
	} else {
		return false
	}
}


//
//
//
func daysSinceEpoch(month int, day int, year int) int {
	nDays := 0

	if year > EPOCH_YEAR {
		for i := EPOCH_YEAR; i < year; i++ {
			if isLeapYear(i) {
				nDays += 366

				fmt.Println("Adding leap year", i)
			} else {
				nDays += 365

				fmt.Println("Adding year", i)
			}
		}
	}

	if month > EPOCH_MONTH {
		for i := EPOCH_MONTH+1; i < month; i++ {
			nDays += MONTHS_DAYS_ARR[i]
		}
	}

	if day > EPOCH_DAY {
		for i := EPOCH_DAY+1; i < day; i++ {
			nDays++
		}
	}

	return nDays
}

//
//
//

func dayOfWeek(month, day, year int) int {
	return daysSinceEpoch(month, day, year) % 7
}

func dayOfWeekName(weekDay int) string {
	return WEEKDAY_NAMES_ARR[weekDay]
}

func main() {
	fmt.Println("January 2, 1900 falls on a", dayOfWeekName(dayOfWeek(JANUARY, 2, 1900)))
	fmt.Println("January 2, 1901 falls on a", dayOfWeekName(dayOfWeek(JANUARY, 2, 1901)))
}