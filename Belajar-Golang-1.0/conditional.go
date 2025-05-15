package main

// CheckEvenOdd returns "even" if the number is even, otherwise "odd"
func CheckEvenOdd(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

// GradeClassifier returns a grade classification based on score
func GradeClassifier(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
