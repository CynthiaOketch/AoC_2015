package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Process the input file
	totalRibbon, err := FindRibbonSize("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the total wrapping paper/ribbon required
	// fmt.Printf("Total wrapping paper required: %d square feet\n", totalPaper)
	fmt.Printf("Total wrapping paper required: %d square feet\n", totalRibbon)
}

func calculateWrappingPaper(l, w, h int) int {
	// Calculate surface area
	surfaceArea := 2*l*w + 2*w*h + 2*h*l

	// Calculate slack (smallest side area)
	slack := min(l*w, w*h, h*l)

	// Total paper required
	return surfaceArea + slack
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func totalWrappingPaper(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	totalPaper := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split dimensions by 'x' and convert to integers
		dimensions := strings.Split(line, "x")
		if len(dimensions) != 3 {
			return 0, fmt.Errorf("invalid dimension format: %s", line)
		}
		l, err := strconv.Atoi(dimensions[0])
		if err != nil {
			return 0, fmt.Errorf("error converting dimension to integer: %v", err)
		}
		w, err := strconv.Atoi(dimensions[1])
		if err != nil {
			return 0, fmt.Errorf("error converting dimension to integer: %v", err)
		}
		h, err := strconv.Atoi(dimensions[2])
		if err != nil {
			return 0, fmt.Errorf("error converting dimension to integer: %v", err)
		}

		// Calculate paper needed for the present
		totalPaper += calculateWrappingPaper(l, w, h)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return totalPaper, nil
}

func FindRibbonSize(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	totalRibbon := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split dimensions by 'x' and convert to integers
		dimensions := strings.Split(line, "x")
		if len(dimensions) != 3 {
			return 0, fmt.Errorf("invalid dimension format: %s", line)
		}
		l, err := strconv.Atoi(dimensions[0])
		if err != nil {
			return 0, fmt.Errorf("error converting dimension to integer: %v", err)
		}
		w, err := strconv.Atoi(dimensions[1])
		if err != nil {
			return 0, fmt.Errorf("error converting dimension to integer: %v", err)
		}
		h, err := strconv.Atoi(dimensions[2])
		if err != nil {
			return 0, fmt.Errorf("error converting dimension to integer: %v", err)
		}

		// Calculate ribbon needed for the present
		totalRibbon += min(2*l+2*w, 2*w+2*h, 2*h+2*l) + l*w*h

	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return totalRibbon, nil
}
