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
    lines := libs.FileToSlice("2025/day09/input.txt", "\n")
    positions := getPosSlice(lines)

    for i, pos1 := range positions {
        for j:=i; j<len(positions); j++ {
            pos2 := positions[j]
            currentArea := area(pos1, pos2)
            if currentArea > ans {
                ans = currentArea
            }
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 09 is:", ans, "🎄")
}

func getPosSlice(lines []string) []libs.Pos {
    posSlice := []libs.Pos{}
    for _, line := range lines {
        xAndYSlice := libs.StrToIntSlice(line, ",")
        posSlice = append(posSlice, libs.Pos{X: xAndYSlice[0], Y: xAndYSlice[1]})
    }
    return posSlice
}

func area(pos1 libs.Pos, pos2 libs.Pos) int {
    return libs.Abs(pos1.X - pos2.X + 1) * libs.Abs(pos1.Y - pos2.Y + 1)
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 09 is:", ans, "🎄")
}