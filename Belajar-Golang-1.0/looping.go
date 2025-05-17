package main

// SumRange returns the sum of all integers from start to end inclusive
func SumRange(start, end int) int {
	// Initialize sum variable to accumulate the total
	sum := 0

	// Loop variable explicitly declared for clarity
	var i int

	// Iterate from start to end inclusive
	for i = start; i <= end; i++ {
		// Add current value of i to sum
		sum += i
	}

	// Return the final accumulated sum
	return sum
}

// Factorial calculates the factorial of a non-negative integer n
func Factorial(n int) int {
	// Check for invalid input: factorial is not defined for negative numbers
	if n < 0 {
		// Return -1 to indicate error case
		return -1
	}

	// Initialize result variable to 1 (factorial of 0 is 1)
	result := 1

	// Loop variable explicitly declared for clarity
	var i int

	// Calculate factorial by multiplying result with each integer from 2 to n
	for i = 2; i <= n; i++ {
		result *= i
	}

	// Return the computed factorial value
	return result
}

// CountVowels counts the number of vowels in a given string
func CountVowels(s string) int {
	// Initialize count variable to keep track of vowels found
	count := 0

	// Define a map for quick lookup of vowels (both lowercase and uppercase)
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	// Loop variable explicitly declared for clarity
	var ch rune

	// Iterate over each character in the string
	for _, ch = range s {
		// Check if the character is a vowel using the map
		if vowels[ch] {
			// Increment count if vowel found
			count++
		}
	}

	// Return the total count of vowels in the string
	return count
}
