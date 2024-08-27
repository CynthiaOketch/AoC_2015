package main

import (
	"testing"
)

func TestNotQuiteLisp(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    int
		wantErr bool
	}{
		{"Balanced parentheses", "()", 0, false},
		{"More opening parentheses", "(((", 3, false},
		{"More closing parentheses", "(()))", -1, false},
		{"Invalid character", "(()a)", 0, true},
		{"Empty string", "", 0, false},
		{"Single open parenthesis", "(", 1, false},
		{"Single close parenthesis", ")", -1, false},
		{"Nested parentheses", "(((())))", 0, false},
		{"Mixed valid characters", "(()())", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NotQuiteLisp(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotQuiteLisp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NotQuiteLisp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotQuiteLisp2(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    int
		wantErr bool
	}{
		{"Valid case - excess closing", "(()))", 5, false},
		{"Valid case - excess closing at end", "((())", 0, false},
		{"Valid case - closing exceeds in the middle", "(()))(", 5, false},
		{"Invalid character", "(()a)", 0, true},
		{"No closing excess", "((())", 0, false},
		{"Empty string", "", 0, false},
		{"Single character open", "(", 0, false},
		{"Single character close", ")", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NotQuiteLisp2(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotQuiteLisp2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NotQuiteLisp2() = %v, want %v", got, tt.want)
			}
		})
	}
}
