package main

import (
    "fmt"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type Pos struct {
    x int
    y int
}

func part1(){
    ans := 0

    lines := libs.FileToSlice("2024/day12/input.txt", "\n")
    gardenPlotMap := createGardenPlotMap(lines)

    regions := [][]Pos{}

    for _, plotSlice := range gardenPlotMap {
        regions = append(regions, findRegions(plotSlice)...)
    }

    for _, region := range regions {
        ans += (calculateArea(region) * calculatePerimeter(region))
    }

    fmt.Println("The answer to part 1 for day 12 is:", ans)
}

func calculateArea(points []Pos) int {
    return len(points)
}

func calculatePerimeter(points []Pos) int {
    pointSet := make(map[Pos]bool)
    for _, p := range points {
        pointSet[p] = true
    }

    perimeter := 0
    directions := []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    for _, p := range points {
        for _, d := range directions {
            neighbor := Pos{p.x + d.x, p.y + d.y}
            if !pointSet[neighbor] {
                perimeter++
            }
        }
    }

    return perimeter
}


func createGardenPlotMap(lines []string) map[rune][]Pos {
    gardenPlotMap := make(map[rune]([]Pos))
    for y, line := range lines {
        for x, char := range line {
            gardenPlotMap[char] = append(gardenPlotMap[char], Pos{x, y})
        }
    }
    return gardenPlotMap
}

func findRegions(points []Pos) [][]Pos {
    visited := make(map[Pos]bool)
    regions := [][]Pos{}

    directions := []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    for _, pos := range points {
        if !visited[pos] {
            region := []Pos{}
            dfs(pos, &region, directions, points, visited)
            regions = append(regions, region)
        }
    }

    return regions
}

func dfs(pos Pos, group *[]Pos, directions []Pos, points []Pos, visited map[Pos]bool) {
    visited[pos] = true
    *group = append(*group, pos)

    for _, d := range directions {
        neighbor := Pos{pos.x + d.x, pos.y + d.y}
        if !visited[neighbor] && contains(points, neighbor) {
            dfs(neighbor, group, directions, points, visited)
        }
    }
}

func contains(points []Pos, p Pos) bool {
    for _, point := range points {
        if point == p {
            return true
        }
    }
    return false
}

func part2(){
    ans := 0

    lines := libs.FileToSlice("2024/day12/input.txt", "\n")
    gardenPlotMap := createGardenPlotMap(lines)

    regions := [][]Pos{}

    for _, plotSlice := range gardenPlotMap {
        regions = append(regions, findRegions(plotSlice)...)
    }

    for _, region := range regions {
        ans += (calculateArea(region) * calculateFilteredPerimeter(region))
    }

    fmt.Println("The answer to part 2 for day 12 is:", ans)
}

func calculateFilteredPerimeter(points []Pos) int {
    pointSet := make(map[Pos]bool)
    for _, p := range points {
        pointSet[p] = true
    }

    perimeter := 0
    directions := []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    edges := make(map[[2]Pos]bool)
    for _, p := range points {
        for _, d := range directions {
            neighbor := Pos{p.x + d.x, p.y + d.y}
            if !pointSet[neighbor] {
                edges[[2]Pos{p, neighbor}] = true
            }
        }
    }

    filteredEdges := make(map[[2]Pos]bool)
    for edge := range edges {
        keep := true
        for _, d := range []Pos{{1, 0}, {0, 1}} {
            p1 := Pos{edge[0].x + d.x, edge[0].y + d.y}
            p2 := Pos{edge[1].x + d.x, edge[1].y + d.y}
            if edges[[2]Pos{p1, p2}] {
                keep = false
            }
        }
        if keep {
            filteredEdges[edge] = true
        }
    }

    perimeter = len(filteredEdges)
    return perimeter
}