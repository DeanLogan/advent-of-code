package main

import (
	"fmt"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    lines := libs.FileToSlice("day09/input.txt", "\n")

    for _, line := range lines{
        historyOfValue := [][]int{libs.StrToIntSlice(line, " ")}
        ans += calcForwardExtrapolated(historyOfValue)
    }
    fmt.Println("The answer to part 1 for day 9 is:", ans)
}

func calcForwardExtrapolated(slices [][]int) int {
    diffsSlice := calcDiffsForIntSlice(slices[len(slices)-1])
    slices = append(slices, diffsSlice)
    allZero := true
    for _, val := range diffsSlice {
        if val != 0 {
            allZero = false
        }
    }
    if allZero {
        extrapolatedValue := 0 
        for _, slice := range slices {
            extrapolatedValue += slice[len(slice)-1]
        }
        return extrapolatedValue
    } 
    return calcForwardExtrapolated(slices)
}

func calcDiffsForIntSlice(slice []int) []int {
    diffs := []int{}

    for i := 1; i < len(slice); i++ {
        diffs = append(diffs, slice[i] - slice[i-1])
    }
    return diffs
}

func part2(){
    ans := 0
    lines := libs.FileToSlice("day09/input.txt", "\n")

    for _, line := range lines{
        historyOfValue := [][]int{libs.StrToIntSlice(line, " ")}
        ans += calcBackwardExtrapolated(historyOfValue)
    }
    fmt.Println("The answer to part 2 for day 9 is:", ans)
}

func calcBackwardExtrapolated(slices [][]int) int {
    diffsSlice := calcDiffsForIntSlice(slices[len(slices)-1])
    slices = append(slices, diffsSlice)
    allZero := true
    for _, val := range diffsSlice {
        if val != 0 {
            allZero = false
        }
    }
    if allZero {
        extrapolatedValue := 0 
        for i := len(slices)-1; i > -1; i-- {
            extrapolatedValue = slices[i][0] - extrapolatedValue
        }
        return extrapolatedValue
    } 
    return calcBackwardExtrapolated(slices)
}