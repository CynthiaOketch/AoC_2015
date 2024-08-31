package main

import (
	"testing"
)

func TestCountHouses(t *testing.T) {
	tests := []struct {
		directions string
		expected   int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
		{"", 1},         // No moves, only the starting point
		{"^^>>vv<<", 8}, // Moves in a square and returns to the starting point
	}

	for _, test := range tests {
		result := CountHouses(test.directions)
		if result != test.expected {
			t.Errorf("For directions %q, expected %d, but got %d", test.directions, test.expected, result)
		}
	}
}

func TestCountHousesWithRoboSanta(t *testing.T) {
	tests := []struct {
		directions string
		expected   int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
		{"", 1},         // No moves, only the starting point
		{"^^>>vv<<", 4}, // Moves in a square and returns to the starting point
		{"^^>v<v<", 5},  // Complex path with both Santa and Robo-Santa moving
	}

	for _, test := range tests {
		result := CountHousesWithRoboSanta(test.directions)
		if result != test.expected {
			t.Errorf("For directions %q, expected %d, but got %d", test.directions, test.expected, result)
		}
	}
}
