package main

import (
	"container/list"
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

    lines := libs.FileToSlice("2024/day18/input.txt", "\n")
    byteMap := linesToByteMap(lines, 1024)

    ans = shortestPath(Pos{0,0}, Pos{70,70}, byteMap)

    fmt.Println("🎄 The answer to part 1 for day 18 is:", ans, "🎄")
}

func printMap(mapStr []string) {
    for _, line := range mapStr {
        fmt.Println(line)
    }
}

func createMap(end int) []string {
    mapStr := make([]string,end)
    newLine := ""
    for i:=0; i<end; i++ {
        newLine += "."
    }
    for i := range mapStr {
        mapStr[i] = newLine
    }
    return mapStr
}

func populateMap(mapStr []string, posList interface{}, replacementChar rune) []string {
    switch v := posList.(type) {
    case []Pos:
        for _, pos := range v {
            if pos.y >= 0 && pos.y < len(mapStr) && pos.x >= 0 && pos.x < len(mapStr[pos.y]) {
                mapStr[pos.y] = libs.ReplaceCharAtIndex(mapStr[pos.y], pos.x, replacementChar)
            }
        }
    case map[Pos]bool:
        for pos := range v {
            if pos.y >= 0 && pos.y < len(mapStr) && pos.x >= 0 && pos.x < len(mapStr[pos.y]) {
                mapStr[pos.y] = libs.ReplaceCharAtIndex(mapStr[pos.y], pos.x, replacementChar)
            }
        }
    default:
        fmt.Println("Unsupported type")
    }
    return mapStr
}

func linesToByteMap(lines []string, stopAt int) map[Pos]bool {
    byteMap := make(map[Pos]bool)
    for i:=0; i<stopAt; i++ {
        nums := libs.StrToIntSlice(lines[i], ",")
        byteMap[Pos{x:nums[0], y:nums[1]}] = true
    }
    return byteMap
}

func shortestPath(start, end Pos, byteMap map[Pos]bool) int {
    directions := []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    visited := make(map[Pos]bool)
    queue := list.New()
    queue.PushBack(start)
    visited[start] = true
    steps := 0

    for queue.Len() > 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            current := queue.Remove(queue.Front()).(Pos)
            if current == end {
                return steps
            }
            for _, dir := range directions {
                next := Pos{current.x + dir.x, current.y + dir.y}
                if next.x >= start.x && next.x <= end.x && next.y >= start.y && next.y <= end.y {
                    if !byteMap[next] && !visited[next] {
                        queue.PushBack(next)
                        visited[next] = true
                    }
                }
            }
        }
        steps++
    }
    return -1 // return -1 if there is no path
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 18 is:", ans, "🎄")
}