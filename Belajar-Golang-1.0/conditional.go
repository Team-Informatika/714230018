package main

// CheckEvenOdd returns "even" if the number is even, otherwise "odd"
func CheckEvenOdd(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}

// GradeClassifier returns a grade based on the score
func GradeClassifier(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// IsPositive returns true if the number is positive, false otherwise
func IsPositive(num int) bool {
	return num > 0
}

// IsLeapYear returns true if the year is a leap year, false otherwise
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

// DescribeNumber returns "zero" if num is 0, "positive" if > 0, "negative" if < 0
func DescribeNumber(num int) string {
	if num == 0 {
		return "zero"
	} else if num > 0 {
		return "positive"
	} else {
		return "negative"
	}
}

// TemperatureClassification classifies temperature into categories
func TemperatureClassification(temp float64) string {
	if temp <= 0 {
		return "Freezing"
	} else if temp > 0 && temp <= 15 {
		return "Cold"
	} else if temp > 15 && temp <= 25 {
		return "Warm"
	} else if temp > 25 && temp <= 35 {
		return "Hot"
	} else {
		return "Unknown"
	}
}
