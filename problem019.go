package main

import (
	"fmt"
)

const (
	JANUARY   = 0
	FEBRUARY  = 1
	MARCH     = 2
	APRIL     = 3
	MAY       = 4
	JUNE      = 5
	JULY      = 6
	AUGUST    = 7
	SEPTEMBER = 8
	OCTOBER   = 9
	NOVEMBER  = 10
	DECEMBER  = 11

	SUNDAY    = 0
	MONDAY    = 1
	TUESDAY   = 2
	WEDNESDAY = 3
	THURSDAY  = 4
	FRIDAY    = 5
	SATURDAY  = 6

	EPOCH_YEAR  = 1900
	EPOCH_MONTH = JANUARY
	EPOCH_DAY   = 1
)

var MONTH_DAYS_ARR = [...]int{
	31, // January
	28, // February
	31, // March
	30, // April
	31, // May
	30, // June
	31, // July
	31, // August
	30, // September
	31, // October
	30, // November
	31, // December
}

var WEEKDAY_NAMES_ARR = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var MONTH_NAMES_ARR = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func IsLeapYear(year int) bool {
	if (year%4 == 0) && (year%100 != 0) || (year%400 == 0) {
		return true
	} else {
		return false
	}
}

func DaysSinceEpoch(month int, day int, year int) int {
	nDays := 0

	if year > EPOCH_YEAR {
		for i := EPOCH_YEAR; i < year; i++ {
			if IsLeapYear(i) {
				nDays += 366
			} else {
				nDays += 365
			}
		}
	}

	if month == EPOCH_MONTH+1 {
		nDays += MONTH_DAYS_ARR[EPOCH_MONTH] - 1

	} else if month > EPOCH_MONTH+1 {
		for i := EPOCH_MONTH; i < month; i++ {
			nDays += MONTH_DAYS_ARR[i]

			if IsLeapYear(year) && i == FEBRUARY {
				nDays++
			}
		}
	}

	if day > EPOCH_DAY {
		for i := EPOCH_DAY; i < day; i++ {
			nDays++
		}
	}

	return nDays
}

func DayOfWeek(month, day, year int) int {
	return DaysSinceEpoch(month, day, year) % 7
}

func DayOfWeekName(weekDay int) string {
	return WEEKDAY_NAMES_ARR[weekDay]
}

func NumberFirstDayOfMonthSundaysInCentury(century int) int {
	numberSundays := 0
	i := (century - 1) * 100

	for year := i; year < i+100; year++ {
		for month := JANUARY; month <= DECEMBER; month++ {
			if DayOfWeek(month, 1, year) == SUNDAY {
				numberSundays++
			}
		}
	}
	return numberSundays
}

func main() {

	//
	//	Return the number of first-month
	//

	fmt.Println("The number of Sundays falling on the first of a month is:", NumberFirstDayOfMonthSundaysInCentury(20))
}
