package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1900, false},
		{2000, true},
		{2001, false},
		{1990, false},
		{4000, true},
	}

	for _, c := range cases {
		result := IsLeapYear(c.in)
		if result != c.want {
			t.Errorf("IsLeapYear(%v) == %v, want %v", c.in, result, c.want)
		}
	}
}

func TestDaysSinceEpoch(t *testing.T) {
	cases := []struct {
		month, day, year, want int
	}{
		{JANUARY, 1, 1900, 0},
		{JANUARY, 2, 1900, 1},
		{FEBRUARY, 8, 1900, 37},
		{MARCH, 1, 1900, 59},
		{JANUARY, 1, 1900, 0},
		{JANUARY, 1, 1900, 0},
	}

	for _, c := range cases {
		result := DaysSinceEpoch(c.month, c.day, c.year)
		if result != c.want {
			t.Errorf("DaysSinceEpoch(%v %v, %v) == %v, want %v", MONTH_NAMES_ARR[c.month], c.day, c.year, result, c.want)
		}
	}
}

func TestNumberFirstDayOfMonthSundaysInCentury(t *testing.T) {
	cases := []struct {
		centuryIn, want int
	}{
		{20, 171},
	}

	for _, c := range cases {
		result := NumberFirstDayOfMonthSundaysInCentury(c.centuryIn)
		if result != c.want {
			t.Errorf("NumberFirstDayOfMonthSundaysInCentury(%v) == %v, want %v", c.centuryIn, result, c.want)
		}
	}
}
