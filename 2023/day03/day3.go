// TODO: Currently getFullNum returns a string, and the result is added to numStr, however if multiple numbers are found for the same character then only one of these numbers is saved.


package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"github.com/DeanLogan/advent-of-code/libs"
)

func main() {
    partOne()
    partTwo()
}

func partOne(){
	reg, err := regexp.Compile(`[^a-zA-Z0-9_.]`)
	if err != nil {
		log.Fatal(err, "Failed to compile regex")
		return
	}

	ans := 0

	dataSlice := libs.FileToSlice("2023/day03/input.txt", "\n")
	numsFound := []string{}

	// checks first line characters then finds the corresponding numbers for these characters
	numsFound = getNumsForAllCharsOnLine(reg.FindAllStringIndex(dataSlice[0], -1), numsFound, dataSlice[0], "", dataSlice[1])
	// checks last line characters then finds the corresponding numbers for these characters
	numsFound = getNumsForAllCharsOnLine(reg.FindAllStringIndex(dataSlice[len(dataSlice)-1], -1), numsFound, dataSlice[len(dataSlice)-1], dataSlice[len(dataSlice)-2], "")
	
	for i := 1; i < len(dataSlice)-1; i++ {
		currentLine := dataSlice[i]
		charactersFound := reg.FindAllStringIndex(currentLine, -1)
		numsFound = getNumsForAllCharsOnLine(charactersFound, numsFound, currentLine, dataSlice[i-1], dataSlice[i+1])
	}
	for _, numStr := range numsFound {
		num, _ := strconv.Atoi(string(numStr))
		ans += num
	}

	fmt.Println("🎄 The answer to part 1 for day 03 is:", ans, "🎄")
}

func partTwo(){
	reg, err := regexp.Compile(`[*]`)
	if err != nil {
		log.Fatal(err, "Failed to compile regex")
		return
	}

	dataSlice := libs.FileToSlice("2023/day03/input.txt", "\n")
	//numsFound := [][]string{}

	// checks first line characters then finds the corresponding numbers for these characters
	ans := findMultiplesAndAdd(reg.FindAllStringIndex(dataSlice[0], -1), 0, dataSlice[0], "", dataSlice[1])
	/// checks last line characters then finds the corresponding numbers for these characters
	ans = findMultiplesAndAdd(reg.FindAllStringIndex(dataSlice[len(dataSlice)-1], -1), ans, dataSlice[len(dataSlice)-1], dataSlice[len(dataSlice)-2], "")
	
	for i := 1; i < len(dataSlice)-1; i++ {
		currentLine := dataSlice[i]
		charactersFound := reg.FindAllStringIndex(currentLine, -1)
		ans = findMultiplesAndAdd(charactersFound, ans, currentLine, dataSlice[i-1], dataSlice[i+1])
	}

	fmt.Println("🎄 The answer to part 2 for day 03 is:", ans, "🎄")
}

// if a character only has two numbers, multiply them and add to ans
func findMultiplesAndAdd(charactersFound [][]int, ans int, currentLine string, prevLine string, nextLine string) int {
	for _, character := range charactersFound {
		numsForChar := getAllNumsForChar(character, []string{}, currentLine, prevLine, nextLine)
		if(len(numsForChar) == 2){
			firstNum, _ := strconv.Atoi(string(numsForChar[0]))
			secondNum, _ := strconv.Atoi(string(numsForChar[1]))
			ans += firstNum * secondNum
		}
	}
	return ans
}

// for each character found, check for any adjacent numbers and add them to numsFound
func getNumsForAllCharsOnLine(charactersFound [][]int, numsFound []string, currentLine string, prevLine string, nextLine string) []string {
	for _, character := range charactersFound {
		numsFound = getAllNumsForChar(character, numsFound, currentLine, prevLine, nextLine)
	}
	return numsFound
}

func getAllNumsForChar(character []int, numsFound []string, currentLine string, prevLine string, nextLine string) []string {
	// checks left
	if(currentLine[character[0]-1] != '.'){
		numsFound = append(numsFound, getFullNum(currentLine[:character[0]], false))
	}
	// checks right
	if(currentLine[character[1]] != '.'){
		numsFound = append(numsFound, getFullNum(currentLine[character[1]:], true))
	}
	// checks up
	if(prevLine[character[0]] != '.'){
		numsFound = append(numsFound, getFullNumFromMiddle(prevLine, character[0]))
	} else {
		// checks up right
		if(prevLine[character[1]] != '.'){ 
			numsFound = append(numsFound, getFullNum(prevLine[character[1]:], true))
		} 
		// checks up left
		if(prevLine[character[0]-1] != '.'){ 
			numsFound = append(numsFound, getFullNum(prevLine[:character[0]], false))
		}
	}
	// checks below
	if(nextLine[character[0]] != '.'){
		numsFound = append(numsFound, getFullNumFromMiddle(nextLine, character[0]))
	} else {
		// checks below right
		if(nextLine[character[1]] != '.'){ 
			numsFound = append(numsFound, getFullNum(nextLine[character[1]:], true))
		}
		// checks below left
		if(nextLine[character[0]-1] != '.'){ 
			numsFound = append(numsFound, getFullNum(nextLine[:character[0]], false))
		}
	}
	return numsFound
}

// startingSide is true if the number is on the left side of the line, false if on the right
func getFullNum(line string, startingSide bool) string {
	numStr := ""
	if(!startingSide){
		line = libs.ReverseString(line)
	}

	for _, character := range line {
		if(character != '.'){
			numStr += string(character)
		} else {
			break
		}
	}

	if(!startingSide){
		numStr = libs.ReverseString(numStr)
	}

	return numStr
}

func getFullNumFromMiddle(line string, middleIndex int) string {
	// gets the middle number
	numStr := string(line[middleIndex])

	// walk left to get left side of number
	numStr = getFullNum(line[:middleIndex], false) + numStr

	// walk right to get right side of number
	numStr = numStr + getFullNum(line[middleIndex+1:], true)

	return numStr
}

