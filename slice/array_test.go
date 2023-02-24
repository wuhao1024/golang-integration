package slice

import (
	"testing"
)

func TestIsStringArrayInclude(t *testing.T) {
	testV := "abcd"
	testSlice := []string{"a", "ab", "abc", "abcd"}

	res := IsStringArrayInclude(testSlice, testV)

	if !res == true {
		t.Fatal("Not Meet Expectations")
	}
}

func TestUintArrayContains(t *testing.T) {
	var testV uint = 10086
	testSlice := []uint{10010, 10086}

	res := UintArrayContains(testSlice, testV)

	if !res == true {
		t.Fatal("Not Meet Expectations")
	}
}

func TestStringContainsAny(t *testing.T) {
	v1 := "abcd"
	s1 := []string{"a", "b", "c", "d"}

	res := StringContainsAny(v1, s1)

	if !res == true {
		t.Fatal("Not Meet Expectations")
	}

	v2 := "ttt"
	s2 := []string{"a", "b", "c", "d", "e"}

	res1 := StringContainsAny(v2, s2)

	if !res1 == false {
		t.Fatal("Not Meet Expectations")
	}
}

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

func TestMaxIntSlice(t *testing.T) {
	s1 := []int{-1, 2, 3, 4, 5}
	res, err := MaxIntSlice(s1)

	if err != nil || res != 5 {
		t.Fatal("Not Meet Expectations")
	}

	s2 := []int{}
	_, err = MaxIntSlice(s2)
	if err == nil {
		t.Fatal("Not Meet Expectations")
	}
}
