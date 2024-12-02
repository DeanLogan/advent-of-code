package main

import (
    "fmt"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0

    reports := libs.FileToSlice("2024/day02/input.txt", "\n")
    for _, report := range reports {
        if checkSafety(report, false) {
            ans++
        }
    }

    fmt.Println("The answer to part 1 for day 02 is:", ans)
}

func checkSafety(report string, problemDampener bool) bool {
    reportInts := libs.StrToIntSlice(report, " ")
    increasingFlag := true
    if reportInts[0] > reportInts[1] {
        increasingFlag = false
    }

    for i:=0; i<len(reportInts)-1; i++ {
        levelDiff := reportInts[i] - reportInts[i+1]
        if (libs.Abs(levelDiff) < 1 || libs.Abs(levelDiff) > 3) || (increasingFlag && levelDiff > 0) || (!increasingFlag && levelDiff < 0) {
            if problemDampener {
                for j := 0; j < len(reportInts); j++ {
                    newReport := libs.RemoveElementFromSlice(reportInts, j)
                    if checkSafety(libs.IntSliceToStr(newReport, " "), false) {
                        return true
                    }
                }
            }
            return false
        }
    }
    return true
}

func part2(){
    ans := 0
    reports := libs.FileToSlice("2024/day02/input.txt", "\n")
    for _, report := range reports {
        if checkSafety(report, true) {
            ans++
        }
    }

    fmt.Println("The answer to part 2 for day 02 is:", ans)
}