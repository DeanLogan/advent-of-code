package main

import (
	"fmt"
	"strings"
    "strconv"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    
    lines := libs.FileToSlice("2025/day06/input.txt", "\n")
    newLines := create2dSlice(lines)
    orderedLines := libs.Transpose2DSlice(newLines)

    for _, line := range orderedLines {
        ans += solveCephalopodMathsEquation(line)
    }

    fmt.Println("🎄 The answer to part 1 for day 06 is:", ans, "🎄")
}

func create2dSlice(lines []string) [][]string {
    newSlice := [][]string{}
    for _, line := range lines {
        newSlice = append(newSlice, removeEmptyStrings(strings.Split(line, " ")))
    } 
    return newSlice
}

func removeEmptyStrings(slice []string) []string {
    newSlice := []string{}
    for _, str := range slice {
        if str != "" {
            newSlice = append(newSlice, str)
        }
    }
    return newSlice
}

func solveCephalopodMathsEquation(equation []string) int {
    ans, _ := strconv.Atoi(equation[0])
    op := equation[len(equation)-1]
    
    for i:=1; i<len(equation)-1; i++ {
        num, _ := strconv.Atoi(equation[i])
        
        switch op {
        case "*":
            ans = ans * num
        case "+":
            ans = ans + num
        }
    }
    return ans
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 06 is:", ans, "🎄")
}