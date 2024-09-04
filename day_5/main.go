package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Helper function to check if a string is nice
func isNice1(s string) bool {
	return hasThreeVowels(s) && hasDoubleLetter(s) && !containsDisallowed(s)
}

// Helper function to check if a string contains at least three vowels
func hasThreeVowels(s string) bool {
	vowels := "aeiou"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
		if count >= 3 {
			return true
		}
	}
	return false
}

// Helper function to check if a string contains at least one letter that appears twice in a row
func hasDoubleLetter(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}

// Helper function to check if a string contains any of the disallowed substrings
func containsDisallowed(s string) bool {
	disallowed := []string{"ab", "cd", "pq", "xy"}
	for _, substr := range disallowed {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	niceCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			niceCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Number of nice strings: %d\n", niceCount)
}

// Helper function to check if a string is nice
func isNice(s string) bool {
	return hasPair(s) && hasRepeatWithOneBetween(s)
}

// Helper function to check if a string contains a pair of any two letters that appears at least twice without overlapping
func hasPair(s string) bool {
	pairs := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if _, exists := pairs[pair]; exists {
			if i-pairs[pair] > 1 {
				return true
			}
		}
		pairs[pair] = i
	}
	return false
}

// Helper function to check if a string contains at least one letter which repeats with exactly one letter between them
func hasRepeatWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
