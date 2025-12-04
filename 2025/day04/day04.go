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
    grid := libs.FileToByteGrid("2025/day04/input.txt")
    replacedPaper := findPaperRolls(grid)

    ans = len(replacedPaper)
    fmt.Println("🎄 The answer to part 1 for day 04 is:", ans, "🎄")
}

func findPaperRolls(grid [][]byte) []libs.Pos {
    width := len(grid[0])

    replacedPaper := []libs.Pos{}
    for y := range width {
        for x := range width {
            if grid[y][x] == '@' {
                pos := libs.Pos{X: x, Y: y}
                if len(libs.FindAdjacentInGrid(grid, pos, '@')) < 4 {
                    replacedPaper = append(replacedPaper, pos)
                }
            }
        }
    }
    return replacedPaper
}

func part2(){
    ans := 0
    grid := libs.FileToByteGrid("2025/day04/input.txt")

    numRollsRemoved := 1 // defaulting to 1 just to make sure that the for loop executes

    for numRollsRemoved > 0 {
        replacedPaper := findPaperRolls(grid)
        numRollsRemoved = len(replacedPaper)
        replaceFoundPaper(replacedPaper, grid)
        ans += numRollsRemoved
    }

    fmt.Println("🎄 The answer to part 2 for day 04 is:", ans, "🎄")
}

func replaceFoundPaper(found []libs.Pos, grid [][]byte) {
    for _, pos := range found {
        libs.ReplaceCharAtPos(grid, pos, 'x')
    }
}