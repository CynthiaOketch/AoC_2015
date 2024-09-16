package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const gridSize = 1000

// Create a 1000x1000 grid of lights, all initially OFF (0)
var grid [gridSize][gridSize]int

// Function to apply an instruction to the grid
func applyInstruction1(instruction string, x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			switch instruction {
			case "turn on":
				grid[x][y] = 1
			case "turn off":
				grid[x][y] = 0
			case "toggle":
				grid[x][y] = 1 - grid[x][y] // Toggle between 0 and 1
			}
		}
	}
}

// Function to apply an instruction to the grid based on brightness
func applyInstruction2(instruction string, x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			switch instruction {
			case "turn on":
				grid[x][y] += 1
			case "turn off":
				if grid[x][y] > 0 {
					grid[x][y] -= 1
				}
			case "toggle":
				grid[x][y] += 2
			}
		}
	}
}

// Parse instruction line and apply the correct action to the grid
func processInstruction1(line string) {
	var x1, y1, x2, y2 int
	var instruction string

	// Parse the input string to extract instruction and coordinates
	if strings.HasPrefix(line, "turn on") {
		fmt.Sscanf(line, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		instruction = "turn on"
	} else if strings.HasPrefix(line, "turn off") {
		fmt.Sscanf(line, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		instruction = "turn off"
	} else if strings.HasPrefix(line, "toggle") {
		fmt.Sscanf(line, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		instruction = "toggle"
	}

	// Apply the parsed instruction to the grid
	applyInstruction1(instruction, x1, y1, x2, y2)
}

// Parse instruction line and apply the correct action to the grid
func processInstruction2(line string) {
	var x1, y1, x2, y2 int
	var instruction string

	// Parse the input string to extract instruction and coordinates
	if strings.HasPrefix(line, "turn on") {
		fmt.Sscanf(line, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		instruction = "turn on"
	} else if strings.HasPrefix(line, "turn off") {
		fmt.Sscanf(line, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		instruction = "turn off"
	} else if strings.HasPrefix(line, "toggle") {
		fmt.Sscanf(line, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		instruction = "toggle"
	}

	// Apply the parsed instruction to the grid
	applyInstruction2(instruction, x1, y1, x2, y2)
}

// Count the number of lights that are ON (1) in the grid
func countLitLights() int {
	count := 0
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			count += grid[x][y]
		}
	}
	return count
}

// Read instructions from the input file and process them
func readInstructionsFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processInstruction1(scanner.Text())
		processInstruction2(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// Calculate the total brightness of all the lights in the grid
func totalBrightness() int {
	brightness := 0
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			brightness += grid[x][y]
		}
	}
	return brightness
}

func main() {
	// File that contains the instructions
	const filename = "input.txt"

	// Read and process the instructions from the file
	err := readInstructionsFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}

	// Count and display the number of lights that are ON
	litLights := countLitLights()
	fmt.Printf("Number of lights that are ON: %d\n", litLights)

	// Calculate and display the total brightness of all lights
	brightness := totalBrightness()
	fmt.Printf("Total brightness of all lights: %d\n", brightness)
}
