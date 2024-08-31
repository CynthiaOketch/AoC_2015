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
		for _,char := range input {
			if char
		}
	}
}
