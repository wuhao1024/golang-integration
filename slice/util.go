package slice

import "errors"

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

// StringsRemove remove +s+ form +ss+
func StringsRemove(ss []string, s string) []string {
	ns := make([]string, 0, len(ss))
	for _, v := range ss {
		if v != s {
			ns = append(ns, v)
		}
	}
	return ns
}
