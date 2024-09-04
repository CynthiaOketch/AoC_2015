package main

import (
	"testing"
)

// TestFindLowestNumber tests the findLowestNumber function with different inputs
func TestFindLowestNumber(t *testing.T) {
	tests := []struct {
		secretKey    string
		targetPrefix string
		expected     int
	}{
		{"abcdef", "00000", 609043},     // From the problem example
		{"pqrstuv", "00000", 1048970},   // From the problem example
		{"bgvyzdsv", "000000", 1038736}, // Expected result for a secretKey and targetPrefix pair
	}

	for _, test := range tests {
		result := findLowestNumber(test.secretKey, test.targetPrefix)
		if result != test.expected {
			t.Errorf("For secretKey %s and targetPrefix %s, expected %d but got %d", test.secretKey, test.targetPrefix, test.expected, result)
		}
	}
}
