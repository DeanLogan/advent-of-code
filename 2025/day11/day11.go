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
    lines := libs.FileToSlice("2025/day11/input.txt", "\n")
    serverMap := getServerMap(lines)
    
    svrToFft := pathsDP(serverMap, "svr", "fft")
    fftToDac := pathsDP(serverMap, "fft", "dac")
    dacToOut := pathsDP(serverMap, "dac", "out")
    
    ans := svrToFft * fftToDac * dacToOut
    fmt.Println("🎄 The answer to part 2 for day 11 is:", ans, "🎄")
}

func pathsDP(graph map[string][]string, start, end string) int {
    dp := make(map[string]int)
    for node := range graph {
        dp[node] = 0
    }
    dp[end] = 1
    
    visited := make(map[string]bool)
    visited[end] = true
    
    return dfsDP(graph, dp, visited, start)
}

func dfsDP(graph map[string][]string, dp map[string]int, visited map[string]bool, curr string) int {
    if visited[curr] {
        return dp[curr]
    }
    
    visited[curr] = true
    
    for _, node := range graph[curr] {
        if !visited[node] {
            dfsDP(graph, dp, visited, node)
        }
        dp[curr] += dp[node]
    }
    
    return dp[curr]
}