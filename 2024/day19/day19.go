package main

import (
	"fmt"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0

    file := libs.FileToSlice("2024/day19/input.txt", "\n\n")
    towels := createTowelMap(strings.Split(file[0], ", "))
    patterns := strings.Split(file[1], "\n")
    
    for _, pattern := range patterns {
        if isPatternPossible(pattern, towels) {
            ans++
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 19 is:", ans, "🎄")
}

func isPatternPossible(pattern string, towels map[string]bool) bool {
    if pattern == "" {
        return true
    }
    for i := 1; i <= len(pattern); i++ {
        if towels[pattern[:i]] && isPatternPossible(pattern[i:], towels) {
            return true
        }
    }
    return false
}

func createTowelMap(towelSlice []string) map[string]bool {
    towelMap := make(map[string]bool)
    for _, towel := range towelSlice {
        towelMap[towel] = true
    }
    return towelMap
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 19 is:", ans, "🎄")
}