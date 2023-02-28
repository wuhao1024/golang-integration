package slice

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
