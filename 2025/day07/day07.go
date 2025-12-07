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
    grid := libs.FileToByteGrid("2025/day07/input.txt")
    splits, _ := simulateBeams(grid)
    fmt.Println("🎄 The answer to part 1 for day 07 is:", splits, "🎄")
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

func simulateBeams(grid [][]byte) (int, map[int]int) {
    splits := 0
    beamPositions := getBeamPositions(grid)

    for _, row := range grid {
        nextBeamPositions := make(map[int]int)
        
        for i, val := range row {
            if beamPositions[i] > 0 {
                if val == '^' {
                    nextBeamPositions[i-1] += beamPositions[i]
                    nextBeamPositions[i+1] += beamPositions[i]
                    splits++
                } else {
                    nextBeamPositions[i] += beamPositions[i]
                }
            }
        }
        beamPositions = nextBeamPositions
    }
    return splits, beamPositions
}

func part2(){
    grid := libs.FileToByteGrid("2025/day07/input.txt")
    _, beamPositions := simulateBeams(grid)

    ans := 0
    for _, val := range beamPositions {
        ans += val
    }
    fmt.Println("🎄 The answer to part 2 for day 07 is:", ans, "🎄")
}