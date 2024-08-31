package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		Totalhouses := CountHouses(input)
		fmt.Println(Totalhouses)
		TotalRoboSantaHouses := CountHousesWithRoboSanta(input)
		fmt.Println(TotalRoboSantaHouses)
	}
}

func CountHouses(directions string) int {
	// Start at the origin (0, 0)
	x, y := 0, 0
	visitedHouses := map[string]bool{
		"0,0": true,
	}

	// Move according to each direction
	for _, direction := range directions {
		switch direction {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		}
		// Convert the position to a string and mark it as visited
		pos := fmt.Sprintf("%d,%d", x, y)
		visitedHouses[pos] = true
	}

	// Return the number of unique houses visited
	return len(visitedHouses)
}

func CountHousesWithRoboSanta(directions string) int {
	// Maps for Santa and Robo-Santa
	santaVisitedHouses := map[string]bool{}
	roboVisitedHouses := map[string]bool{}

	// Start both Santa and Robo-Santa at the origin (0, 0)
	santaX, santaY := 0, 0
	roboX, roboY := 0, 0

	// Mark the starting point in both maps
	santaVisitedHouses["0,0"] = true
	roboVisitedHouses["0,0"] = true

	// Move according to each direction, alternating between Santa and Robo-Santa
	for i, direction := range directions {
		if i%2 == 0 {
			// Santa's turn
			switch direction {
			case '>':
				santaX++
			case '<':
				santaX--
			case '^':
				santaY++
			case 'v':
				santaY--
			}
			pos := fmt.Sprintf("%d,%d", santaX, santaY)
			santaVisitedHouses[pos] = true
		} else {
			// Robo-Santa's turn
			switch direction {
			case '>':
				roboX++
			case '<':
				roboX--
			case '^':
				roboY++
			case 'v':
				roboY--
			}
			pos := fmt.Sprintf("%d,%d", roboX, roboY)
			roboVisitedHouses[pos] = true
		}
	}

	// Combine the maps to get unique houses visited by either Santa or Robo-Santa
	combinedVisitedHouses := map[string]bool{}
	for house := range santaVisitedHouses {
		combinedVisitedHouses[house] = true
	}
	for house := range roboVisitedHouses {
		combinedVisitedHouses[house] = true
	}

	// Return the number of unique houses visited
	return len(combinedVisitedHouses)
}
