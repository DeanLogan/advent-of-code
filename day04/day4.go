package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {	
	scanner := libs.GetScannerForFile("day04/input.txt")
	if err := scanner.Err(); err != nil {
		log.Fatal(err, "Failed to scan file")
		return 
	}

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		_, line = libs.SplitAtChar(line, ':')

		winnningNums, card := libs.SplitAtChar(line, '|')
		cardSlice := strings.Split(card, " ")
		winnningNumsSlice := strings.Split(winnningNums, " ")
		
		points := 0
		firstMatch := false
		matches := 0
		for _, numStr := range cardSlice {
			if(libs.SearchForStrInSlice(numStr, winnningNumsSlice)){
				matches ++
				if(!firstMatch){
					points ++
					firstMatch = true
				} else {
					points *= 2
				}
			}
		}
		ans += points
	}

	fmt.Println("The answer to part 1 for day 4 is: ", ans)
}

func partTwo() {
	lines := libs.FileToSlice("day04/input.txt", "\n")

	cards := make([]int, len(lines))
	for i := range cards {
		cards[i] = 1
	}
	for i, line := range lines {
		parts := strings.Split(line, "|")
		x := strings.Fields(parts[0])
		y := strings.Fields(parts[1])

		n := intersectionSize(x, y)
		for j := i + 1; j < libs.Min(i+1+n, len(lines)); j++ {
			cards[j] += cards[i]
		}
	}
	ans := 0
	for _, v := range cards {
		ans += v
	}
	fmt.Println("The answer to part 2 for day 4 is: ", ans)
}
	
func intersectionSize(x, y []string) int {
	set := make(map[string]bool)
	for _, v := range x {
		set[v] = true
	}
	count := 0
	for _, v := range y {
		if set[v] {
			count++
		}
	}
	return count
}
