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
    inputFile := libs.AllFileContentAsString("2025/day02/input.txt")
    idRanges := strings.Split(inputFile, ",")
    for _, idRangeStr := range idRanges {
        idRange := libs.StrToIntSlice(idRangeStr, "-")
        ans += getInvalidIdTotalForRange(idRange[0],idRange[1])
    }
    fmt.Println("🎄 The answer to part 1 for day 02 is:", ans, "🎄")
}

func getInvalidIdTotalForRange(lowerRange int, higherRange int) int {
    invalidIds := 0
    for i := lowerRange; i<higherRange+1; i++ {
        idStr := strconv.Itoa(i)
        if checkRepeatSequenceTwice(idStr) {
            invalidIds += i
        }
    }
    return invalidIds
}

func checkRepeatSequenceTwice(str string) bool {
    if len(str)%2 != 0 {
        return false
    }
    subString := str[:len(str)/2]
    count := strings.Count(str, subString)
    return count == 2
}

func part2(){
    ans := 0
    inputFile := libs.AllFileContentAsString("2025/day02/input.txt")
    idRanges := strings.Split(inputFile, ",")
    for _, idRangeStr := range idRanges {
        idRange := libs.StrToIntSlice(idRangeStr, "-")
        ans += getInvalidIdTotalForRange2(idRange[0],idRange[1])
    }
    fmt.Println("🎄 The answer to part 2 for day 02 is:", ans, "🎄")
}

func getInvalidIdTotalForRange2(lowerRange int, higherRange int) int {
    invalidIds := 0
    for i := lowerRange; i<higherRange+1; i++ {
        idStr := strconv.Itoa(i)
        if repeatingPatternFactors(idStr) {
            invalidIds += i
        }
    }
    return invalidIds
}

func repeatingPatternFactors(str string) bool {
    // checking if all of the characters are the same ensuring it is only counted once
    if checkPatternMatch(str, 1, len(str)) {
        return true
    }
    factorPairs := libs.PrimeFactorPairs(len(str))
    factorPairs = factorPairs[1:]
    for _, factorPair := range factorPairs {
        if checkPatternMatch(str, factorPair[0], factorPair[1]) {
            return true
        }
        if checkPatternMatch(str, factorPair[1], factorPair[0]) {
            return true
        }
    }
    return false
}

func checkPatternMatch(str string, patternLen int, expectedCount int) bool {
    subString := str[:patternLen]
    count := strings.Count(str, subString)
    return count == expectedCount && count >= 2
}