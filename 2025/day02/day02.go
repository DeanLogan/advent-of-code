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
        fmt.Println(idRangeStr)
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
            fmt.Println(idStr)
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
    fmt.Println("🎄 The answer to part 2 for day 02 is:", ans, "🎄")
}