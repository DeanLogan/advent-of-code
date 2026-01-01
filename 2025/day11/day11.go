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
    lines := libs.FileToSlice("2025/day11/input.txt", "\n")
    serverMap := getServerMap(lines)
    ans := countPaths("you", make(map[string]bool), serverMap)
    fmt.Println("🎄 The answer to part 1 for day 11 is:", ans, "🎄")
}

func getServerMap(lines []string) map[string][]string {
    serverMap := make(map[string][]string)
    for _, line := range lines {
        key, valueStr := libs.SplitAtStr(line, ": ")
        valueStr = strings.Trim(valueStr, " ")
        value := strings.Split(valueStr, " ")
        serverMap[key] = value
    }
    return serverMap
}

func countPaths(current string, visited map[string]bool, graph map[string][]string) int {
    if current == "out" {
        return 1
    }
    
    visited[current] = true
    count := 0
    for _, neighbor := range graph[current] {
        if !visited[neighbor] {
            count += countPaths(neighbor, visited, graph)
        }
    }
    visited[current] = false
    return count
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 11 is:", ans, "🎄")
}