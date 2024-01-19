package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main(){
	partOne()
	partTwo()
}

func partOne(){
	ans := 1
	lines := libs.FileToSlice("day6/input.txt", "\n")
	times := libs.StrToIntSlice(lines[0], " ")
	distances := libs.StrToIntSlice(lines[1], " ")
	for i := 0; i < len(times); i++ {
		winningWays := 0
		for j := 0; j < times[i]; j++ {
			if calcDist(j, times[i]) > distances[i] {
				winningWays ++
			}
		}
		ans = ans * winningWays
	}
	fmt.Println("The answer to part 1 for day 6 is:", ans)
}

func calcDist(buttonTime int, totalTime int) int{
	return (totalTime - buttonTime) * buttonTime // time the button has to travel * the time the button was held down for 
}

func partTwo(){
	ans := 0
	lines := libs.FileToSlice("day6/input.txt", "\n")

	_, timeStr := libs.SplitAtChar(lines[0], ':')
	timeStr = strings.ReplaceAll(timeStr, " ", "")
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Fatal("cannot convert time to int")
	}

	_, distStr := libs.SplitAtChar(lines[1], ':')
	distStr = strings.ReplaceAll(distStr, " ", "")
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		log.Fatal("cannot convert dist to int")
	}

	for j := 0; j < time; j++ {
		if calcDist(j, time) > dist {
			ans ++
		}
	}
	fmt.Println("The answer to part 2 for day 6 is:", ans)
}