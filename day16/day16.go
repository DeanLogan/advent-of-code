package main

import (
	"fmt"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main(){
    part1()
    part2()
}

type BeamPath struct {
    row, col, deltaRow, deltaCol int
}

type GridPoint struct {
    row, col int
}

func part1() {
    graph := libs.FileToSlice("day16/input.txt", "\n")
    beamPaths := []BeamPath{{0, -1, 0, 1}}
    visitedPaths := make(map[BeamPath]bool)
    energizedTiles := make(map[GridPoint]bool)

    for len(beamPaths) > 0 {
        currentPath := beamPaths[0]
        beamPaths = beamPaths[1:]

        newPaths := processBeamPath(currentPath, graph, visitedPaths, energizedTiles)
        beamPaths = append(beamPaths, newPaths...)
    }

    fmt.Println("The answer to part 1 for day 16 is:", len(energizedTiles))
}

func processBeamPath(path BeamPath, contraptionLayout []string, visitedPaths map[BeamPath]bool, energizedTiles map[GridPoint]bool) []BeamPath {
    path.row += path.deltaRow
    path.col += path.deltaCol

    if path.row < 0 || path.row >= len(contraptionLayout) || path.col < 0 || path.col >= len(contraptionLayout[0]) {
        return nil
    }

    tile := rune(contraptionLayout[path.row][path.col])

    if tile == '.' || (tile == '-' && path.deltaCol != 0) || (tile == '|' && path.deltaRow != 0) {
        if !visitedPaths[path] {
            visitedPaths[path] = true
            energizedTiles[GridPoint{path.row, path.col}] = true
            return []BeamPath{path}
        }
    } else if tile == '/' {
        path.deltaRow, path.deltaCol = -path.deltaCol, -path.deltaRow
        if !visitedPaths[path] {
            visitedPaths[path] = true
            energizedTiles[GridPoint{path.row, path.col}] = true
            return []BeamPath{path}
        }
    } else if tile == '\\' {
        path.deltaRow, path.deltaCol = path.deltaCol, path.deltaRow
        if !visitedPaths[path] {
            visitedPaths[path] = true
            energizedTiles[GridPoint{path.row, path.col}] = true
            return []BeamPath{path}
        }
    } else {
        splitterDirections := [][2]int{{1, 0}, {-1, 0}}
        if tile == '-' {
            splitterDirections = [][2]int{{0, 1}, {0, -1}}
        }
        newPaths := []BeamPath{}
        for _, direction := range splitterDirections {
            newPath := BeamPath{path.row, path.col, direction[0], direction[1]}
            if !visitedPaths[newPath] {
                visitedPaths[newPath] = true
                energizedTiles[GridPoint{newPath.row, newPath.col}] = true
                newPaths = append(newPaths, newPath)
            }
        }
        return newPaths
    }

    return nil
}

func part2(){
    ans := 0
    graph := libs.FileToSlice("day16/input.txt", "\n")
    for r := 0; r < len(graph); r++ {
        ans = libs.Max(ans, findOptimalBeamConfiguration(BeamPath{r, -1, 0, 1}, graph))
        ans = libs.Max(ans, findOptimalBeamConfiguration(BeamPath{r, len(graph[0]), 0, -1}, graph))
    }
    
    for c := 0; c < len(graph[0]); c++ {
        ans = libs.Max(ans, findOptimalBeamConfiguration(BeamPath{-1, c, 1, 0}, graph))
        ans = libs.Max(ans, findOptimalBeamConfiguration(BeamPath{len(graph), c, -1, 0}, graph))
    }
    fmt.Println("The answer to part 2 for day 16 is:", ans)
}

func findOptimalBeamConfiguration(path BeamPath, contraptionLayout []string) int {
    beamPaths := []BeamPath{path}
    visitedPaths := make(map[BeamPath]bool)
    energizedTiles := make(map[GridPoint]bool)

    for len(beamPaths) > 0 {
        currentPath := beamPaths[0]
        beamPaths = beamPaths[1:]

        currentPath.row += currentPath.deltaRow
        currentPath.col += currentPath.deltaCol

        if currentPath.row < 0 || currentPath.row >= len(contraptionLayout) || currentPath.col < 0 || currentPath.col >= len(contraptionLayout[0]) {
            continue
        }

        tile := rune(contraptionLayout[currentPath.row][currentPath.col])

        if tile == '.' || (tile == '-' && currentPath.deltaCol != 0) || (tile == '|' && currentPath.deltaRow != 0) {
            if !visitedPaths[currentPath] {
                visitedPaths[currentPath] = true
                energizedTiles[GridPoint{currentPath.row, currentPath.col}] = true
                beamPaths = append(beamPaths, currentPath)
            }
        } else if tile == '/' {
            currentPath.deltaRow, currentPath.deltaCol = -currentPath.deltaCol, -currentPath.deltaRow
            if !visitedPaths[currentPath] {
                visitedPaths[currentPath] = true
                energizedTiles[GridPoint{currentPath.row, currentPath.col}] = true
                beamPaths = append(beamPaths, currentPath)
            }
        } else if tile == '\\' {
            currentPath.deltaRow, currentPath.deltaCol = currentPath.deltaCol, currentPath.deltaRow
            if !visitedPaths[currentPath] {
                visitedPaths[currentPath] = true
                energizedTiles[GridPoint{currentPath.row, currentPath.col}] = true
                beamPaths = append(beamPaths, currentPath)
            }
        } else {
            splitterDirections := [][2]int{{1, 0}, {-1, 0}}
            if tile == '-' {
                splitterDirections = [][2]int{{0, 1}, {0, -1}}
            }
            for _, direction := range splitterDirections {
                newPath := BeamPath{currentPath.row, currentPath.col, direction[0], direction[1]}
                if !visitedPaths[newPath] {
                    visitedPaths[newPath] = true
                    energizedTiles[GridPoint{newPath.row, newPath.col}] = true
                    beamPaths = append(beamPaths, newPath)
                }
            }
        }
    }

    return len(energizedTiles)
}