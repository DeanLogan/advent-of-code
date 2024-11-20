package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"github.com/DeanLogan/advent-of-code/libs"
)

func main() {
    partOne()
	partTwo()
}

func partOne(){
    scanner := libs.GetScannerForFile("2023/day01/input1.txt")
	if err := scanner.Err(); err != nil {
		log.Fatal(err, "Failed to scan file")
		return 
	}

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err, "Failed to compile regex")
		return
	}
	total := 0
    for scanner.Scan() {
		processedString := reg.ReplaceAllString(scanner.Text(), "")
		num, err := strconv.Atoi(string(processedString[0]) + string(processedString[len(processedString)-1]))
		if err != nil {
			log.Fatal(err, "Failed to convert string to int")
			return
		}
		total += num
    }
	fmt.Println("The answer to part 1 for day 1 is: ", total)
}

func partTwo(){
	scanner := libs.GetScannerForFile("2023/day01/input1.txt")

	if err := scanner.Err(); err != nil {
		log.Fatal(err, "Failed to scan file")
		return 
	}

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err, "Failed to compile regex")
		return
	}

	total := 0
    for scanner.Scan() {
		text := replaceWordWithDigit(scanner.Text())
		processedString := reg.ReplaceAllString(text, "")
		num, err := strconv.Atoi(string(processedString[0]) + string(processedString[len(processedString)-1]))
		if err != nil {
			log.Fatal(err, "Failed to convert string to int")
			return
		}
		total += num
    }
	fmt.Println("The answer to part 2 for day 1 is: ", total)
}

func replaceWordWithDigit(text string) string {
	// converts words to digits but keeps the first and last letters of the word in place for edge cases like eighthree where it needs to 
	// be converted to 83, but if we just replace the word with the digit alone the t in eighthree will be lost when replacing three with 3
	numberMap := map[string]string{
        "one":   "o1e",
        "two":   "t2o",
        "three": "t3e",
		"four":  "f4r",
        "five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
        "eight": "e8t",
        "nine":  "n9e",
    }

	for word, digit := range numberMap {
		text = strings.ReplaceAll(text, word, digit)
	}
	return text
}
