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
    return -1 
}

func part2() {
    lines := libs.FileToSlice("2024/day18/input.txt", "\n")
    start := Pos{0,0}
    end := Pos{70,70}

    // modified binary search
    left, right := 0, len(lines)-1
    lastPossible := -1

    for left <= right {
        mid := (left + right) / 2
        byteMap := linesToByteMap(lines, mid)
        if isPathPossible(start, end, byteMap) {
            lastPossible = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    ans := lines[lastPossible]

    fmt.Println("🎄 The answer to part 2 for day 18 is:", ans, "🎄")
}

func isPathPossible(start, end Pos, byteMap map[Pos]bool) bool {
    directions := []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    visited := make(map[Pos]bool)
    queue := list.New()
    queue.PushBack(start)
    visited[start] = true

    for queue.Len() > 0 {
        current := queue.Remove(queue.Front()).(Pos)
        if current == end {
            return true
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
    return false
}