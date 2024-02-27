package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main() {
    partOne()
    partTwo()
}

func partOne(){
	scanner := libs.GetScannerForFile("day2/input.txt")
	if err := scanner.Err(); err != nil {
		log.Fatal(err, "Failed to scan file")
		return 
	}

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		game, line := libs.SplitAtChar(line, ':')

		validGame := isGameValid(line)

		// if the game is valid, add the id to the answer
		if validGame {
			_, idString := libs.SplitAtChar(game, ' ')
			id, err := strconv.Atoi(strings.ReplaceAll(idString," ", ""))
			if err != nil {
				log.Fatal(err, " Failed to convert string to int")
			}
			ans = ans + id
		}
	}

	fmt.Println("The answer to part 1 for day 2 is: ", ans)
}

func isGameValid(line string) bool {
	first, second := libs.SplitAtChar(line, ';')

	bag := map[string]int{
		"red":		12,
		"green":	13,
		"blue": 	14,
	}

	// logic for checking if the set of cubes is valid
	cubes := strings.Split(second, ",")
	for _, cube := range cubes {
		numOfCubes, cubeColour := libs.SplitAtChar(cube, ' ')
		num, err := strconv.Atoi(strings.ReplaceAll(numOfCubes," ", ""))
		if err != nil {
			log.Fatal(err, " Failed to convert string to int")
		}
		if bag[cubeColour] < num {
			return false
		}
	}

	if first == second {
		return true
	} 
	return isGameValid(first)
}

func partTwo(){
	scanner := libs.GetScannerForFile("day2/input.txt")
	if err := scanner.Err(); err != nil {
		log.Fatal(err, "Failed to scan file")
		return 
	}

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		_, line = libs.SplitAtChar(line, ':')
		bag := map[string]int{
			"red":		0,
			"green":	0,
			"blue": 	0,
		}

		bag = getMinNumOfCubesForValidGame(line, bag)

		// if the game is valid, add the id to the answer
		ans = ans + (bag["red"] * bag["green"] * bag["blue"])
	}

	fmt.Println("The answer to part 2 for day 2 is: ", ans)
}

func getMinNumOfCubesForValidGame(line string, bag map[string]int) map[string]int {
	first, second := libs.SplitAtChar(line, ';')

	// logic for checking the set of cubes
	cubes := strings.Split(second, ",")
	for _, cube := range cubes {
		numOfCubes, cubeColour := libs.SplitAtChar(cube, ' ')
		num, err := strconv.Atoi(strings.ReplaceAll(numOfCubes," ", ""))
		if err != nil {
			log.Fatal(err, " Failed to convert string to int")
		}
		if bag[cubeColour] < num {
			bag[cubeColour] = num
		}
	}

	if first == second {
		return bag
	} 
	return getMinNumOfCubesForValidGame(first, bag)

}
