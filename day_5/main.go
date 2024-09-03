package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt") // Replace with your input file name
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	niceCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if isNiceString(line) {
			niceCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Number of nice strings: %d\n", niceCount)
}

func isNiceString(s string) bool {
	return hasNonOverlappingPair(s) && hasRepeatWithOneBetween(s)
}

func hasNonOverlappingPair(s string) bool {
	pairs := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		pairs[pair]++
	}
	for _, count := range pairs {
		if count > 1 {
			return true
		}
	}
	return false
}

func hasRepeatWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
