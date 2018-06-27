package calc_test

import (
	"testing"

	"github.com/mbesseux/calc"
)

func TestAddOK(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{3, 1, 4},
		{3, 5, 8},
		{11, 5, 16},
	}
	for _, c := range cases {
		got := calc.Add(c.a, c.b)

		if got != c.expected {
			t.Errorf("We expected %d but got %d", c.expected, got)
		}
	}

}

func TestSubstractOK(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, -1},
		{2, 2, 0},
		{3, 1, 2},
		{3, 5, -2},
		{11, 5, 6},
	}
	for _, c := range cases {
		got := calc.Substract(c.a, c.b)

		if got != c.expected {
			t.Errorf("We expected %d but got %d", c.expected, got)
		}
	}

}
