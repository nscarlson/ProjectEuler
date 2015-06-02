package main

import "testing"

func TestIsAmicable(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{220, true},
		{284, true},
	}

	for _, c := range cases {
		result := IsAmicable(c.in)
		if result != c.want {
			t.Errorf("IsAmicable(%v) == %v, want %v", c.in, result, c.want)
		}
	}
}

func TestD(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{220, 284},
		{284, 220},
	}

	for _, c := range cases {
		result := D(c.in)
		if result != c.want {
			t.Errorf("D(%v) == %v, want %v", c.in, result, c.want)
		}
	}
}
