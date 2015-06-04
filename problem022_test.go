package main

import "testing"

func TestValue(t *testing.T) {
	cases := []struct {
		name  string
		index int
		want  int64
	}{
		{"COLIN", 938, 49714},
	}

	for _, c := range cases {
		result := Value(c.name, c.index)
		if result != c.want {
			t.Errorf("Value(%v, %v) == %v, want %v", c.name, c.index, result, c.want)
		}
	}
}
