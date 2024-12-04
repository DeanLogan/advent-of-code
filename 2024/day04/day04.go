package main

import (
    "fmt"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main() {
    part1()
    part2()
}

func part1() {
    ans := 0

    lines := libs.FileToSlice("2024/day04/input.txt", "\n")
    numRows := len(lines)
    numCols := len(lines[0])

    for row := 0; row < numRows; row++ {
        for col := 0; col < numCols; col++ {
            if col+3 < numCols && lines[row][col] == 'X' && lines[row][col+1] == 'M' && lines[row][col+2] == 'A' && lines[row][col+3] == 'S' {
                ans++
            }
            if row+3 < numRows && lines[row][col] == 'X' && lines[row+1][col] == 'M' && lines[row+2][col] == 'A' && lines[row+3][col] == 'S' {
                ans++
            }
            if row+3 < numRows && col+3 < numCols && lines[row][col] == 'X' && lines[row+1][col+1] == 'M' && lines[row+2][col+2] == 'A' && lines[row+3][col+3] == 'S' {
                ans++
            }
            if col+3 < numCols && lines[row][col] == 'S' && lines[row][col+1] == 'A' && lines[row][col+2] == 'M' && lines[row][col+3] == 'X' {
                ans++
            }
            if row+3 < numRows && lines[row][col] == 'S' && lines[row+1][col] == 'A' && lines[row+2][col] == 'M' && lines[row+3][col] == 'X' {
                ans++
            }
            if row+3 < numRows && col+3 < numCols && lines[row][col] == 'S' && lines[row+1][col+1] == 'A' && lines[row+2][col+2] == 'M' && lines[row+3][col+3] == 'X' {
                ans++
            }
            if row-3 >= 0 && col+3 < numCols && lines[row][col] == 'S' && lines[row-1][col+1] == 'A' && lines[row-2][col+2] == 'M' && lines[row-3][col+3] == 'X' {
                ans++
            }
            if row-3 >= 0 && col+3 < numCols && lines[row][col] == 'X' && lines[row-1][col+1] == 'M' && lines[row-2][col+2] == 'A' && lines[row-3][col+3] == 'S' {
                ans++
            }
        }
    }

    fmt.Println("The answer to part 1 for day 04 is:", ans)
}

func part2() {
    ans := 0

    lines := libs.FileToSlice("2024/day04/input.txt", "\n")
    numRows := len(lines)
    numCols := len(lines[0])

    for row := 0; row < numRows; row++ {
        for col := 0; col < numCols; col++ {
            if row+2 < numRows && col+2 < numCols && lines[row][col] == 'M' && lines[row+1][col+1] == 'A' && lines[row+2][col+2] == 'S' && lines[row+2][col] == 'M' && lines[row][col+2] == 'S' {
                ans++
            }
            if row+2 < numRows && col+2 < numCols && lines[row][col] == 'M' && lines[row+1][col+1] == 'A' && lines[row+2][col+2] == 'S' && lines[row+2][col] == 'S' && lines[row][col+2] == 'M' {
                ans++
            }
            if row+2 < numRows && col+2 < numCols && lines[row][col] == 'S' && lines[row+1][col+1] == 'A' && lines[row+2][col+2] == 'M' && lines[row+2][col] == 'M' && lines[row][col+2] == 'S' {
                ans++
            }
            if row+2 < numRows && col+2 < numCols && lines[row][col] == 'S' && lines[row+1][col+1] == 'A' && lines[row+2][col+2] == 'M' && lines[row+2][col] == 'S' && lines[row][col+2] == 'M' {
                ans++
            }
        }
    }

    fmt.Println("The answer to part 2 for day 04 is:", ans)
}