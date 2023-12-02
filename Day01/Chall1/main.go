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
		log.Println("Line : ", line)
		res := findFdigit(line)
		log.Println("Digit 1: ", res)
		res1 := findLdigit(line)
		log.Println("Digit 2:", res1)
		tmpresult, err := strconv.Atoi(res)
		tmpresult1, err := strconv.Atoi(res1)

		result += tmpresult + tmpresult1
		if err != nil {
			log.Fatal(err)
		}
	}
	println(result)
}

func findLdigit(line string) string {
	number := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return string(line[i])
		}
		for name, digit := range number {
			if i-len(name)-1 <= len(line)-1 && i-len(name) >= 0 {
				if line[i-len(name):len(name)] == name {
					return digit
				}
			}
		}
	}
	return ""
}

func findFdigit(line string) string {
	number := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for i, char := range line {
		if unicode.IsDigit(char) {
			return string(char)
		}
		for name, digit := range number {
			if line[i:i+len(name)] == name {
				return digit
			}
		}
	}
	return ""
}
