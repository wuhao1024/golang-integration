package math

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	x, y := 8, 10

	if MinInt(x, y) != 8 {
		t.Fatal("Not Meet Expectations")
	}
}

func TestMaxInt(t *testing.T) {
	x, y := 8, 10

	if MaxInt(x, y) != 10 {
		t.Fatal("Not Meet Expectations")
	}
}

func TestRoundN(t *testing.T) {
	var f float64 = 1.234567

	if RoundN(f, 2) != 1.23 {
		t.Fatal("Not Meet Expectations")
	}

	if RoundN(f, 5) != 1.23457 {
		t.Fatal("Not Meet Expectations")
	}
}
