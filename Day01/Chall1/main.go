package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Get the file path from command-line argument
	filePath := os.Args[1]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Read each line
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		res := findFdigit(line)
		res += findLdigit(line)
		tmpresult, err := strconv.Atoi(res)
		result += tmpresult
		if err != nil {
			log.Fatal(err)
		}
	}
	println(result)
}

func findLdigit(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return string(line[i])
		}
	}
	return ""
}

func findFdigit(line string) string {
	for _, char := range line {
		if unicode.IsDigit(char) {
			return string(char)
		}
	}
	return ""
}
