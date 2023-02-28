package math

import "math"

// MinInt returns the min of x or y.
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MaxInt returns the max of x or y.
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// RoundN 保留n位小数
func RoundN(value float64, n int) float64 {
	return math.Round((value)*math.Pow10(n)) / math.Pow10(n)
}
