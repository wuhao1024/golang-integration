package slice

import "testing"

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

func TestStringsRemove(t *testing.T) {
	s1 := []string{"a", "ab", "a", "c"}
	res1 := []string{"ab", "c"}

	if !SameValueStringSlice(StringsRemove(s1, "a"), res1) {
		t.Fatal("Not Meet Expectations")
	}

	s2 := []string{"a", "aa", "a", "c"}
	res2 := []string{"aa", "c"}

	if !SameValueStringSlice(StringsRemove(s2, "a"), res2) {
		t.Fatal("Not Meet Expectations")
	}
}
