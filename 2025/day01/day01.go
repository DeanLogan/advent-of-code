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

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 01 is:", ans, "🎄")
}