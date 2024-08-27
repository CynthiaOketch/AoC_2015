package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	s := ScanFile("input.txt")
	position, err := NotQuiteLisp2(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(position)
}

// NotQuiteLisp function calculates the position based on the content of the specified file
func NotQuiteLisp(content string) (int, error) {
	countup := 0
	countdown := 0
	position := 0
	for _, char := range content {
		if char == '(' {
			countup++
		} else if char == ')' {
			countdown++
		} else {
			return 0, errors.New("invalid character")
		}
		position = countup - countdown
	}

	return position, nil
}

func NotQuiteLisp2(content string) (int, error) {
	countup := 0
	countdown := 0
	position := 0
	for i, char := range content {
		if char == '(' {
			countup++
		} else if char == ')' {
			countdown++
		} else {
			return 0, errors.New("invalid character")
		}
		if countdown > countup {
			position = i + 1
			break
		}
	}

	return position, nil
}

func ScanFile(filename string) string {
	file, err := os.Open(filename)
	result := ""
	if err != nil {
		return "error opening file"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "error scanning file"
	}
	return result
}
