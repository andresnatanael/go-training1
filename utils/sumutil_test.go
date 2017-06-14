/*
sumutils

@Author Andres Natanael Soria <andresnatanael@gmail.com>
*/

package utils

import "testing"

func TestSumValues(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{3, 3},
		{10, 33},
	}
	for _, c := range cases {
		got := SumValues(c.in)
		if got != c.want {
			t.Errorf("SumValues(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
