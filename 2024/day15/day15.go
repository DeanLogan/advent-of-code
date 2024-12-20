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

    lines := libs.AllFileContentAsString("2024/day15/input.txt")
    mapStr, dirs := libs.SplitAtStr(lines, "\n\n")
    dirs = strings.ReplaceAll(dirs[1:], "\n", "")
    dirToPos := makeDirToPosMap()
    warehouseMap := strings.Split(mapStr, "\n")
    warehouseMap = enlargeWarehouseMap(warehouseMap)

    robotLoc := locateRobot(warehouseMap)

    for _, dir := range dirs {
        ogMap := make([]string, len(warehouseMap))
        copy(ogMap, warehouseMap)
        robotLoc, warehouseMap = moveEnlarged(robotLoc, byte(dir), warehouseMap, dirToPos, ogMap)
    }

    for y, line := range warehouseMap {
        for x, char := range line {
            if char == '[' {
                ans += (100 * y) + x
            }
        }
    }

    fmt.Println("🎄 The answer to part 2 for day 15 is:", ans, "🎄")
}

func enlargeWarehouseMap(warehouseMap []string) []string {
    for i := range warehouseMap {
        warehouseMap[i] = strings.ReplaceAll(warehouseMap[i], "#", "##")
        warehouseMap[i] = strings.ReplaceAll(warehouseMap[i], "O", "[]")
        warehouseMap[i] = strings.ReplaceAll(warehouseMap[i], ".", "..")
        warehouseMap[i] = strings.ReplaceAll(warehouseMap[i], "@", "@.")
    }
    return warehouseMap
}

func moveEnlarged(itemToMove Pos, dir byte, warehouseMap []string, dirToPos map[byte]Pos, ogMap []string) (Pos, []string) {
    isRobot := false
    if warehouseMap[itemToMove.y][itemToMove.x] == '@' {
        isRobot = true
    } 

    newPos := itemToMove.Add(dirToPos[dir])
    if warehouseMap[newPos.y][newPos.x] == '#' {
        return itemToMove, warehouseMap
    }

    if warehouseMap[newPos.y][newPos.x] == '[' || warehouseMap[newPos.y][newPos.x] == ']' {
        boxMoved := true
        if dir == 'v'|| dir == '^' {
            boxMoved = moveBoxUpDown(newPos, dir, warehouseMap, dirToPos, ogMap)
        }
        if boxMoved {
            newBoxLoc, warehouseMap := moveEnlarged(newPos, dir, warehouseMap, dirToPos, ogMap)
            if newBoxLoc == newPos {
                if isRobot && errorsInMap(warehouseMap) {
                    return itemToMove, ogMap
                }
                return itemToMove, warehouseMap
            } 
        } else {
            if isRobot && errorsInMap(warehouseMap) {
                return itemToMove, ogMap
            }
            return itemToMove, warehouseMap
        }
    }

    if isRobot && errorsInMap(warehouseMap) {
        return itemToMove, ogMap
    }
    warehouseMap = updateWarehouseMapEnlarged(warehouseMap, itemToMove, newPos)
    return newPos, warehouseMap
}

func moveBoxUpDown(currentPos Pos, dir byte, warehouseMap []string, dirToPos map[byte]Pos, ogMap []string) (bool) {
    if warehouseMap[currentPos.y][currentPos.x] == ']' {
        currentPos.x = currentPos.x - 1
    } else {
        currentPos.x = currentPos.x + 1
    }

    newPos, _ := moveEnlarged(currentPos, dir, warehouseMap, dirToPos, ogMap)
    return newPos != currentPos
}

func errorsInMap(warehouseMap []string) bool {
    for _, line := range warehouseMap {
        for x:=1; x<len(line); x++ {
            if (line[x] == ']' && line[x-1] != '[') {
                return true
            }
        }
    }
    return false
}

func updateWarehouseMapEnlarged(warehouseMap []string, oldPos Pos, newPos Pos) []string{
    char := warehouseMap[oldPos.y][oldPos.x]
    warehouseMap[newPos.y] = libs.ReplaceCharAtIndex(warehouseMap[newPos.y], newPos.x, rune(char))
    warehouseMap[oldPos.y] = libs.ReplaceCharAtIndex(warehouseMap[oldPos.y], oldPos.x, '.')
    return warehouseMap
}
