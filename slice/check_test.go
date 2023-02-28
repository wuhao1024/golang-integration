package slice

import "testing"

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
