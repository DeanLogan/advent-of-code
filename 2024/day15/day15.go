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

type Pos struct {
    x int
    y int
}

func (p Pos) Add(other Pos) Pos {
    return Pos{
        x: p.x + other.x,
        y: p.y + other.y,
    }
}

func part1(){
    ans := 0

    lines := libs.AllFileContentAsString("2024/day15/input.txt")
    mapStr, dirs := libs.SplitAtStr(lines, "\n\n")
    dirs = strings.ReplaceAll(dirs[1:], "\n", "")
    dirToPos := makeDirToPosMap()
    warehouseMap := strings.Split(mapStr, "\n")

    robotLoc := locateRobot(warehouseMap)

    for _, dir := range dirs {
        robotLoc, warehouseMap = move(robotLoc, byte(dir), warehouseMap, dirToPos)
    }

    for y, line := range warehouseMap {
        for x, char := range line {
            if char == 'O' {
                ans += (100 * y) + x
            }
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 15 is:", ans, "🎄")
}

func move(itemToMove Pos, dir byte, warehouseMap []string, dirToPos map[byte]Pos) (Pos, []string) {
    newPos := itemToMove.Add(dirToPos[dir])
    if warehouseMap[newPos.y][newPos.x] == '#' {
        return itemToMove, warehouseMap
    }
    if warehouseMap[newPos.y][newPos.x] == 'O' {
        newBoxLoc, warehouseMap := move(newPos, dir, warehouseMap, dirToPos)
        if newBoxLoc == newPos {
            return itemToMove, warehouseMap
        } 

    }
    updateWarehouseMap(warehouseMap, itemToMove, newPos)
    return newPos, warehouseMap
}

func updateWarehouseMap(warehouseMap []string, oldPos Pos, newPos Pos) []string{
    char := warehouseMap[oldPos.y][oldPos.x]
    warehouseMap[newPos.y] = libs.ReplaceCharAtIndex(warehouseMap[newPos.y], newPos.x, rune(char))
    warehouseMap[oldPos.y] = libs.ReplaceCharAtIndex(warehouseMap[oldPos.y], oldPos.x, '.')
    return warehouseMap
}

func locateRobot(warehouseMap []string) Pos {
    for y, line := range warehouseMap {
        for x, char := range line {
            if char == '@' {
                return Pos{x, y}
            }
        }
    }
    return Pos{}
}

func makeDirToPosMap() map[byte]Pos {
    dirToPos := make(map[byte]Pos)

    dirToPos['^'] = Pos{0, -1}
    dirToPos['v'] = Pos{0, 1}
    dirToPos['>'] = Pos{1, 0}
    dirToPos['<'] = Pos{-1, 0}

    return dirToPos
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 15 is:", ans, "🎄")
}