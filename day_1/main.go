package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	position, err := NotQuiteLisp2("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(position)
}

// NotQuiteLisp function calculates the position based on the content of the specified file
func NotQuiteLisp(filename string) (int, error) {
	countup := 0
	countdown := 0
	position := 0

	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		for _, char := range input {
			if char == '(' {
				countup++
			} else if char == ')' {
				countdown++
			}
			position = countup - countdown
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return position, nil
}



func NotQuiteLisp2(filename string) (int, error) {
	countup := 0
	countdown := 0
	position := 0

	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		for i, char := range input {
			if char == '(' {
				countup++
			} else if char == ')' {
				countdown++
			}
			if countdown > countup {
				position = i + 1
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return position, nil
}