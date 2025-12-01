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
    lines := libs.FileToSlice("2025/day01/input.txt", "\n")

    dialPos := 50
    for _, line := range lines {
        amountToTurn, _ := strconv.Atoi(line[1:])
        dialPos = turnDial(dialPos, amountToTurn, line[0])
        if dialPos == 0 {
            ans += 1
        }
    }
    fmt.Println("🎄 The answer to part 1 for day 01 is:", ans, "🎄")
}

func turnDial(dialPos int, amountToTurn int, direction byte) int {
    if direction == 'R' {
        dialPos += amountToTurn
    } else {
        dialPos -= amountToTurn
    }
    return libs.WrapToRange(dialPos, 0, 100)
}

func part2() {
    ans := 0
    lines := libs.FileToSlice("2025/day01/input.txt", "\n")

    dialPos := 50
    for _, line := range lines {
        amountToTurn, _ := strconv.Atoi(line[1:])
        newPos, crosses := turnDialPart2(dialPos, amountToTurn, line[0])
        ans += crosses
        dialPos = newPos
    }
    fmt.Println("🎄 The answer to part 2 for day 01 is:", ans, "🎄")
}

func turnDialPart2(dialPos int, amountToTurn int, direction byte) (int, int) {
    if direction == 'R' {
        return applyRightTurn(dialPos, amountToTurn)
    }
    return applyLeftTurn(dialPos, amountToTurn)
}

func applyRightTurn(dialPos int, amountToTurn int) (int, int) {
    dialPos += amountToTurn
    crossings := dialPos / 100
    dialPos %= 100
    return dialPos, crossings
}

func applyLeftTurn(dialPos int, amountToTurn int) (int, int) {
    wasZero := dialPos == 0
    dialPos -= amountToTurn
    crossings := countLandingOnZero(dialPos)

    dialPos, wrapCrossings := normaliseLeftTurn(dialPos)
    
    crossings += wrapCrossings
    if wasZero {
        crossings--
    }
    return dialPos, crossings
}

func countLandingOnZero(rawPos int) int {
    if rawPos%100 == 0 {
        return 1
    }
    return 0
}

func normaliseLeftTurn(rawPos int) (int, int) {
    crossings := 0
    for rawPos < 0 {
        rawPos += 100
        crossings++
    }
    return rawPos, crossings
}