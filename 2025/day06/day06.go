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
        str = strings.ReplaceAll(str, " ", "")
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

    lines := libs.FileToSlice("2025/day06/input.txt", "\n")
    transposed := libs.TransposeStringSlice(lines)
    orderedLines := removeEmptyStrings(transposed)
    reforamtedLines := reformatTransposed(orderedLines)

    for i:=1; i<len(reforamtedLines); i++ {
        ans += solveCephalopodMathsEquation(reforamtedLines[i])
    }

    fmt.Println("🎄 The answer to part 2 for day 06 is:", ans, "🎄")
}

func reformatTransposed(transposed[]string) [][]string {
    newSlice := [][]string{}
    tmp := []string{}
    op := ""
    for _, str := range transposed {
        op = string(str[len(str)-1])
        if str[len(str)-1] == '+' || str[len(str)-1] == '*' {
            libs.SwapInSlice(tmp, 0, len(tmp)-1)
            newSlice = append(newSlice, tmp)
            tmp = []string{}
            tmp = append(tmp, op)
            str = str[:len(str)-1]
        }
        tmp = append(tmp, str)
    }
    libs.SwapInSlice(tmp, 0, len(tmp)-1)
    newSlice = append(newSlice, tmp)
    return newSlice
}