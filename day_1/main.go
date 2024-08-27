package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countup := 0
	countdown := 0
	position := 0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
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
	fmt.Println(position)
}

// part2 solution

// for scanner.Scan() {
// 	input := scanner.Text()
// 	for i, char := range input {
// 		if char == '(' {
// 			countup++
// 		} else if char == ')' {
// 			countdown++
// 		}
// 		if countdown > countup {
// 			position = i + 1
// 			break
// 		}
// 	}
// }
