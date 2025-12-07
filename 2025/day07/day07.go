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
    grid := libs.FileToByteGrid("2025/day07/input.txt")

    beamPositions := getBeamPositions(grid)

    for _, row := range grid {
        nextBeamPositions := make(map[int]int)
        
        for i, val := range row {
            if beamPositions[i] > 0 {
                if val == '^' {
                    nextBeamPositions[i-1] += beamPositions[i]
                    nextBeamPositions[i+1] += beamPositions[i]
                    ans++
                } else {
                    nextBeamPositions[i] += beamPositions[i]
                }
            }
        }
        beamPositions = nextBeamPositions
    }

    fmt.Println("🎄 The answer to part 1 for day 07 is:", ans, "🎄")
}

func getBeamPositions(grid [][]byte) map[int]int {
    beamPositions := make(map[int]int)
    for i, char := range grid[0] {
        if char == 'S' {
            beamPositions[i] = 1
            break
        }
    }
    return beamPositions
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 07 is:", ans, "🎄")
}