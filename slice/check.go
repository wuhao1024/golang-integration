package slice

import "strings"

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
