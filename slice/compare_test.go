package slice

import (
	"testing"
)

func TestDifferenceFromLeftStringSlice(t *testing.T) {
	v1 := []string{"b", "c", "d", "e", "f"}
	s1 := []string{"a", "b", "c", "d"}

	ideal := []string{"e", "f"}
	res := DifferenceFromLeftStringSlice(v1, s1)

	if !(len(res) == len(ideal)) || !SameValueStringSlice(ideal, res) {
		t.Fatal("Not Meet Expectations")
	}
}

func TestSameValueStringSlice(t *testing.T) {
	v1 := []string{"a", "b", "c", "d"}
	s1 := []string{"a", "b", "c", "d"}

	res := SameValueStringSlice(v1, s1)

	if !res == true {
		t.Fatal("Not Meet Expectations")
	}

	v2 := []string{"a", "b", "c", "d"}
	s2 := []string{"a", "b", "c", "d", "f"}

	res2 := SameValueStringSlice(v2, s2)

	if !res2 == false {
		t.Fatal("Not Meet Expectations")
	}
}
