package main

import (
	"os"
	"testing"
)

func TestCalculateWrappingPaper(t *testing.T) {
	tests := []struct {
		l, w, h int
		want    int
	}{
		{2, 3, 4, 58},  // 2*6 + 2*12 + 2*8 + 6 (slack) = 58
		{1, 1, 10, 43}, // 2*1 + 2*10 + 2*10 + 1 (slack) = 43
		{5, 5, 5, 175}, // 2*25 + 2*25 + 2*25 + 25 (slack) = 175
		{1, 2, 3, 24},  // 2*2 + 2*3 + 2*6 + 2 (slack) = 24
	}

	for _, tt := range tests {
		got := calculateWrappingPaper(tt.l, tt.w, tt.h)
		if got != tt.want {
			t.Errorf("calculateWrappingPaper(%d, %d, %d) = %d; want %d", tt.l, tt.w, tt.h, got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b, c int
		want    int
	}{
		{10, 20, 30, 10},
		{50, 10, 20, 10},
		{15, 25, 5, 5},
		{7, 7, 7, 7},
	}

	for _, tt := range tests {
		got := min(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("min(%d, %d, %d) = %d; want %d", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestTotalWrappingPaper(t *testing.T) {
	content := `2x3x4
1x1x10
5x5x5`
	filename := "test_input.txt"
	err := os.WriteFile(filename, []byte(content), 0o644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	got, err := totalWrappingPaper(filename)
	if err != nil {
		t.Fatalf("totalWrappingPaper() error = %v", err)
	}

	want := 58 + 43 + 175
	if got != want {
		t.Errorf("totalWrappingPaper() = %d; want %d", got, want)
	}
}

func TestFindRibbonSize(t *testing.T) {
	content := `2x3x4
1x1x10
5x5x5`
	filename := "test_input.txt"
	err := os.WriteFile(filename, []byte(content), 0o644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	got, err := FindRibbonSize(filename)
	if err != nil {
		t.Fatalf("FindRibbonSize() error = %v", err)
	}

	want := (10 + 24) + (4 + 10) + (20 + 125)
	if got != want {
		t.Errorf("FindRibbonSize() = %d; want %d", got, want)
	}
}
