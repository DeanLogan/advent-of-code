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
    width := len(grid[0])
    for y := range width {
        for x := range width {
            if grid[y][x] == '@' {
                pos := libs.Pos{X: x, Y: y}
                if libs.CountAdjacentInByteGrid(grid, pos, '@') < 4 {
                    ans++
                }
            }
        }
    }
    fmt.Println("🎄 The answer to part 1 for day 04 is:", ans, "🎄")
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 04 is:", ans, "🎄")
}