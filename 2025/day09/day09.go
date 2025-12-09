package main

import (
    "fmt"
    "sort"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    lines := libs.FileToSlice("2025/day09/input.txt", "\n")
    positions := getPosSlice(lines)

    for i, pos1 := range positions {
        for j:=i; j<len(positions); j++ {
            pos2 := positions[j]
            currentArea := area(pos1, pos2)
            if currentArea > ans {
                ans = currentArea
            }
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 09 is:", ans, "🎄")
}

func getPosSlice(lines []string) []libs.Pos {
    posSlice := []libs.Pos{}
    for _, line := range lines {
        xAndYSlice := libs.StrToIntSlice(line, ",")
        posSlice = append(posSlice, libs.Pos{X: xAndYSlice[0], Y: xAndYSlice[1]})
    }
    return posSlice
}

func area(pos1 libs.Pos, pos2 libs.Pos) int {
    return libs.Abs(pos1.X - pos2.X + 1) * libs.Abs(pos1.Y - pos2.Y + 1)
}

type Rectangle struct {
    pos1, pos2 libs.Pos
    area       int
}

type Bounds struct {
    minX, maxX int
    processed  bool
}

func part2(){
    lines := libs.FileToSlice("2025/day09/test.txt", "\n")
    positions := getPosSlice(lines)
    
    boundariesSlice, boundsMinY := createBoundariesSlice(positions)
    rectangles := getRectangles(positions)

    maxArea := 0

    sort.Slice(rectangles, func(i, j int) bool {
        return rectangles[i].area > rectangles[j].area
    })

    for _, rectangle := range rectangles {
        minX := libs.Min(rectangle.pos1.X, rectangle.pos2.X)
        maxX := libs.Max(rectangle.pos1.X, rectangle.pos2.X)
        minY := libs.Min(rectangle.pos1.Y, rectangle.pos2.Y)
        maxY := libs.Max(rectangle.pos1.Y, rectangle.pos2.Y)

        if isRectangleValid(minX, maxX, minY, maxY, boundariesSlice, boundsMinY) {
            maxArea = rectangle.area
            break
        }
    }

    fmt.Println("🎄 The answer to part 2 for day 09 is:", maxArea, "🎄")
}

func createBoundariesSlice(coords []libs.Pos) ([]Bounds, int) {
    if len(coords) == 0 {
        return []Bounds{}, 0
    }

    minY, maxY := coords[0].Y, coords[0].Y
    for _, coord := range coords {
        minY = libs.Min(minY, coord.Y)
        maxY = libs.Max(maxY, coord.Y)
    }

    height := maxY - minY + 1
    bounds := make([]Bounds, height)

    for i := range bounds {
        bounds[i] = Bounds{minX: 1<<31 - 1, maxX: -1<<31, processed: false}
    }

    for i := 0; i < len(coords); i++ {
        curr := coords[i]
        next := coords[(i+1)%len(coords)]

        startY := libs.Min(curr.Y, next.Y)
        endY := libs.Max(curr.Y, next.Y)

        for y := startY; y <= endY; y++ {
            index := y - minY
            if curr.X == next.X {
                bounds[index].minX = libs.Min(bounds[index].minX, curr.X)
                bounds[index].maxX = libs.Max(bounds[index].maxX, curr.X)
            } else {
                startX := libs.Min(curr.X, next.X)
                endX := libs.Max(curr.X, next.X)
                bounds[index].minX = libs.Min(bounds[index].minX, startX)
                bounds[index].maxX = libs.Max(bounds[index].maxX, endX)
            }
            bounds[index].processed = true
        }
    }

    return bounds, minY
}

func getRectangles(coords []libs.Pos) []Rectangle {
    rectangles := []Rectangle{}
    for i := range len(coords) {
        p1 := coords[i]
        for j := i + 1; j < len(coords); j++ {
            p2 := coords[j]
            dx := libs.Abs(p2.X-p1.X) + 1
            dy := libs.Abs(p2.Y-p1.Y) + 1
            area := dx * dy
            rectangles = append(rectangles, Rectangle{pos1: p1, pos2: p2, area: area})
        }
    }
    return rectangles
}

func isRectangleValid(minX, maxX, minY, maxY int, bounds []Bounds, boundsMinY int) bool {
    for y := minY; y <= maxY; y++ {
        index := y - boundsMinY
        if index < 0 || index >= len(bounds) {
            return false
        }
        b := bounds[index]
        if !b.processed || minX < b.minX || maxX > b.maxX {
            return false
        }
    }
    return true
}