package slice

import (
	"fmt"
	"testing"
)

func TestIsStringArrayInclude(t *testing.T) {
	testV := "abcd"
	testSlice := []string{"a", "ab", "abc", "abcd"}

	res := IsStringArrayInclude(testSlice, testV)

	fmt.Println(res == true)
}

func TestUintArrayContains(t *testing.T) {
	var testV uint = 10086
	testSlice := []uint{10010, 10086}

	res := UintArrayContains(testSlice, testV)

	fmt.Println(res == true)
}

func TestStringContainsAny(t *testing.T) {
	v1 := "abcd"
	s1 := []string{"a", "b", "c", "d"}

	res := StringContainsAny(v1, s1)

	fmt.Println(res == true)

	v2 := "ttt"
	s2 := []string{"a", "b", "c", "d", "e"}

	res1 := StringContainsAny(v2, s2)

	fmt.Println(res1 == false)
}

func TestDifferenceFromLeftStringSlice(t *testing.T) {
	v1 := []string{"b", "c", "d", "e", "f"}
	s1 := []string{"a", "b", "c", "d"}

	ideal := []string{"e", "f"}
	res := DifferenceFromLeftStringSlice(v1, s1)

	fmt.Println(len(res) == len(ideal))
	fmt.Println(SameValueStringSlice(ideal, res))
}

func TestSameValueStringSlice(t *testing.T) {
	v1 := []string{"a", "b", "c", "d"}
	s1 := []string{"a", "b", "c", "d"}

	res := SameValueStringSlice(v1, s1)

	fmt.Println(res == true)

	v2 := []string{"a", "b", "c", "d"}
	s2 := []string{"a", "b", "c", "d", "f"}

	res2 := SameValueStringSlice(v2, s2)

	fmt.Println(res2 == false)
}
