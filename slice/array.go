package slice

import (
	"errors"
	"strings"
)

// IsStringArrayInclude Returns +true+ if the given +obj+ is present in +arr+
func IsStringArrayInclude(arr []string, obj string) bool {
	for _, e := range arr {
		if e == obj {
			return true
		}
	}
	return false
}

// UintArrayContains same to IsStringArrayInclude, but it is uint
func UintArrayContains(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// StringContainsAny Returns +true+ if any substr in +substrArr+ are within +s+
func StringContainsAny(s string, substrArr []string) bool {
	for _, substr := range substrArr {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// DifferenceFromLeftStringSlice returns the elements in `a` that aren't in `b`.
func DifferenceFromLeftStringSlice(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, e := range b {
		mb[e] = struct{}{}
	}

	var diff []string
	for _, e := range a {
		if _, found := mb[e]; !found {
			diff = append(diff, e)
		}
	}
	return diff
}

// DifferenceFromLeftUintSlice same as DifferenceFromLeftStringSlice but uint
func DifferenceFromLeftUintSlice(a, b []uint) []uint {
	mb := make(map[uint]struct{}, len(b))
	for _, e := range b {
		mb[e] = struct{}{}
	}

	var diff []uint
	for _, e := range a {
		if _, found := mb[e]; !found {
			diff = append(diff, e)
		}
	}
	return diff
}

// SameValueStringSlice Returns +true+ if a == b
func SameValueStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	var flag bool

	for _, v := range a {
		if !IsStringArrayInclude(b, v) {
			return false
		}
		flag = true
	}

	return flag
}

// MaxIntSlice return the maximum value in +slice+
func MaxIntSlice(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("get an empty slice")
	}

	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max, nil
}
