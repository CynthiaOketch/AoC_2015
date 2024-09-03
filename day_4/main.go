package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	secretKey := "bgvyzdsv"  // Replace this with your actual secret key
	targetPrefix := "000000" // The hash should start with this many zeroes
	result := findLowestNumber(secretKey, targetPrefix)
	fmt.Printf("The lowest number for secret key %s that produces an MD5 hash starting with %s is: %d\n", secretKey, targetPrefix, result)
}

func findLowestNumber(secretKey, targetPrefix string) int {
	for i := 1; ; i++ {
		// Concatenate the secret key with the current number
		data := secretKey + strconv.Itoa(i)

		// Generate the MD5 hash
		hash := md5.Sum([]byte(data))
		hashString := hex.EncodeToString(hash[:])

		// Check if the hash starts with the target prefix
		if hashString[:len(targetPrefix)] == targetPrefix {
			return i
		}
	}
}
