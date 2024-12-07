package main

import (
	"fmt"
	"strconv"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0

    lines := libs.FileToSlice("2024/day07/input.txt", "\n")

    for _, line := range lines {
        testValueStr, equationStr := libs.SplitAtStr(line, ": ")
        testValue, _ := strconv.Atoi(testValueStr)
        equationList := libs.StrToIntSlice(equationStr, " ")

        if validateLine(testValue, equationList) {
            ans += testValue
        }
    }

    fmt.Println("The answer to part 1 for day 07 is:", ans)
}

func validateLine(testValue int, equationList []int) bool {
    if len(equationList) == 1 {
        return equationList[0] == testValue
    }

    last := equationList[len(equationList)-1]
    equationList = equationList[:len(equationList)-1]

    possibleMul, possibleAdd := false, false

    if testValue%last == 0 {
        possibleMul = validateLine(testValue/last, equationList)
    } else {
        possibleMul = false
    }

    possibleAdd = validateLine(testValue-last, equationList)

    return possibleMul || possibleAdd
}

func part2(){
    ans := 0

    lines := libs.FileToSlice("2024/day07/input.txt", "\n")

    for _, line := range lines {
        testValueStr, equationStr := libs.SplitAtStr(line, ": ")
        testValue, _ := strconv.Atoi(testValueStr)
        equationList := libs.StrToIntSlice(equationStr, " ")

        if validateLineWithConcat(testValue, equationList) {
            ans += testValue
        }
    }

    fmt.Println("The answer to part 2 for day 07 is:", ans)
}

func validateLineWithConcat(testValue int, equationList []int) bool {
    if len(equationList) == 1 {
        return equationList[0] == testValue
    }

    last := equationList[len(equationList)-1]
    equationList = equationList[:len(equationList)-1]

    possibleMul, possibleConcat, possibleAdd := false, false, false

    if testValue%last == 0 {
        possibleMul = validateLineWithConcat(testValue/last, equationList)
    } else {
        possibleMul = false
    }

    nextPowerOf10 := 1
    for nextPowerOf10 <= last {
        nextPowerOf10 *= 10
    }
    if (testValue-last)%nextPowerOf10 == 0 {
        possibleConcat = validateLineWithConcat((testValue-last)/nextPowerOf10, equationList)
    } else {
        possibleConcat = false
    }

    possibleAdd = validateLineWithConcat(testValue-last, equationList)

    return possibleMul || possibleAdd || possibleConcat
}