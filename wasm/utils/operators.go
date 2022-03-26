package utils

func Ternary(test bool, a, b float64) float64 {
	if test {
		return a
	} else {
		return b
	}
}
