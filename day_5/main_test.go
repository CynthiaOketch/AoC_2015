package main

import "testing"

func TestIsNice1(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},  // Nice string example
		{"aaa", true},               // Nice string example with overlapping rules
		{"jchzalrnumimnmhp", false}, // No double letter
		{"haegwjzuvuyypxyu", false}, // Contains disallowed substring "xy"
		{"dvszwmarrgswjxmb", false}, // Not enough vowels
	}

	for _, test := range tests {
		result := isNice1(test.input)
		if result != test.expected {
			t.Errorf("isNice1(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsNice2(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"qjhvhtzxzqqjkmpb", true},  // Nice string with both conditions met
		{"xxyxx", true},             // Overlapping letters with both conditions met
		{"uurcxstgmygtbstg", false}, // No repeating letter with one between
		{"ieodomkazucvgmuy", false}, // No pair appearing twice without overlapping
		{"aabcdefgaa", false},       // Pair "aa" and "efe" with one between
		{"aaa", false},              // Overlapping pair "aa"
	}

	for _, test := range tests {
		result := isNice2(test.input)
		if result != test.expected {
			t.Errorf("isNice2(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestHasThreeVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"aeiou", true},             // Exactly three vowels
		{"ugknbfddgicrmopn", true},  // Contains more than three vowels
		{"dvszwmarrgswjxmb", false}, // Contains only one vowel
		{"aei", true},               // Contains exactly three vowels
	}

	for _, test := range tests {
		result := hasThreeVowels(test.input)
		if result != test.expected {
			t.Errorf("hasThreeVowels(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestHasDoubleLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ugknbfddgicrmopn", true}, // Contains "dd"
		{"abcdefg", false},         // No double letters
		{"aabbccdd", true},         // Multiple double letters
		{"abcddcba", true},         // Double letters in the middle
	}

	for _, test := range tests {
		result := hasDoubleLetter(test.input)
		if result != test.expected {
			t.Errorf("hasDoubleLetter(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestContainsDisallowed(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"haegwjzuvuyypxyu", true}, // Contains "xy"
		{"abcdef", true},           // No disallowed substrings
		{"pqrsab", true},           // Contains "ab"
		{"xyzabc", true},           // Contains "cd"
	}

	for _, test := range tests {
		result := containsDisallowed(test.input)
		if result != test.expected {
			t.Errorf("containsDisallowed(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestHasPair(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"xyxy", true},             // Pair "xy" appears twice
		{"aabcdefgaa", true},       // Pair "aa" appears twice
		{"aaa", false},             // Overlapping pair "aa"
		{"uurcxstgmygtbstg", true}, // Pair "tg" appears twice without overlapping
		{"abcdefg", false},         // No repeating pairs
	}

	for _, test := range tests {
		result := hasPair(test.input)
		if result != test.expected {
			t.Errorf("hasPair(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestHasRepeatWithOneBetween(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"xyx", true},        // "x" repeats with "y" between
		{"abcdefeghi", true}, // "e" repeats with "f" between
		{"aaa", true},        // "a" repeats with "a" between
		{"abcdefg", false},   // No repeating letter with one between
		{"aaba", true},       // "a" repeats but not with one between
	}

	for _, test := range tests {
		result := hasRepeatWithOneBetween(test.input)
		if result != test.expected {
			t.Errorf("hasRepeatWithOneBetween(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
