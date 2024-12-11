package main

import (
    "fmt"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main() {
    part1()
    part2()
}

type Pos struct {
    x int
    y int
}

var maxX int
var maxY int

func part1() {
    ans := 0

    topographicMapStr := libs.FileToSlice("2024/day10/input.txt", "\n")

    topographicMap := [][]int{}
    for _, line := range topographicMapStr {
        topographicMap = append(topographicMap, libs.StrToIntSlice(line, ""))
    }

    maxX = len(topographicMap[0]) - 1
    maxY = len(topographicMap) - 1

    trailHeads := findTrailHeadLocations(topographicMap)

    for _, trailHead := range trailHeads {
        ans += countPathsToNine(trailHead, topographicMap)
    }

    fmt.Println("The answer to part 1 for day 10 is:", ans)
}

func findTrailHeadLocations(topoMap [][]int) []Pos {
    locs := []Pos{}
    for y, line := range topoMap {
        for x, height := range line {
            if height == 0 {
                locs = append(locs, Pos{x, y})
            }
        }
    }
    return locs
}

func fetchPossibleDirections(currentHeight int, currentPos Pos, topoMap [][]int) []Pos {
    directions := []Pos{}
    targetHeight := currentHeight + 1
    if currentPos.x < maxX && targetHeight == topoMap[currentPos.y][currentPos.x+1] {
        directions = append(directions, Pos{currentPos.x + 1, currentPos.y})
    }
    if currentPos.y < maxY && targetHeight == topoMap[currentPos.y+1][currentPos.x] {
        directions = append(directions, Pos{currentPos.x, currentPos.y + 1})
    }
    if currentPos.x > 0 && targetHeight == topoMap[currentPos.y][currentPos.x-1] {
        directions = append(directions, Pos{currentPos.x - 1, currentPos.y})
    }
    if currentPos.y > 0 && targetHeight == topoMap[currentPos.y-1][currentPos.x] {
        directions = append(directions, Pos{currentPos.x, currentPos.y - 1})
    }
    return directions
}

func countPathsToNine(start Pos, topoMap [][]int) int {
    maxX = len(topoMap[0]) - 1
    maxY = len(topoMap) - 1
    return bfs(start, topoMap)
}

func bfs(start Pos, topoMap [][]int) int {
    queue := []Pos{start}
    visited := make(map[Pos]bool)
    visited[start] = true
    paths := 0

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if topoMap[current.y][current.x] == 9 {
            paths++
            continue
        }

        directions := fetchPossibleDirections(topoMap[current.y][current.x], current, topoMap)
        for _, nextPos := range directions {
            if !visited[nextPos] {
                visited[nextPos] = true
                queue = append(queue, nextPos)
            }
        }
    }

    return paths
}

func part2() {
    ans := 0
    fmt.Println("The answer to part 2 for day 10 is:", ans)
}