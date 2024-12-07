package main

import (
    "fmt"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

func (d Direction) rightTurn() Direction {
    return (d + 1) % 4
}

type Pos struct {
    X int
    Y int
}

type Guard struct {
    currentPos Pos
    facing Direction
    traveledPos map[Pos]bool
}

func part1(){
    ans := 0

    mapSlice :=  libs.FileToSlice("2024/day06/input.txt", "\n")
    guard := getGuard(mapSlice)

    movedBy := -1
    // if the currentPos of the guard is the same as the old one then there are no obstacles in their way
    for movedBy != 0 {
        guard.currentPos, movedBy = facingOpsticalPos(mapSlice, guard) // get new pos location
        guard.facing = guard.facing.rightTurn()
    }

    for _, traveledPos := range guard.traveledPos {
        if traveledPos {
            ans++
        }
    }

    fmt.Println("The answer to part 1 for day 06 is:", ans)
}

func getGuard(mapSlice []string) Guard {
    rowPos := 0
    colPos := -1
    for i, row := range mapSlice {
        colPos = libs.SearchForCharInStr(row, '^')
        if colPos != -1 {
            rowPos = i
            break
        }
    }
    currentPos := Pos{rowPos, colPos}
    traveledMap := make(map[Pos]bool)
    traveledMap[currentPos] = true
    return Guard{currentPos, Direction(0), traveledMap}
}

func facingOpsticalPos(mapSlice []string, guard Guard) (Pos, int) {
    posMoved := 0
    newPos := guard.currentPos
    switch guard.facing {
    case 0: // North
        for i:=guard.currentPos.X; i>=0; i-- {
            guard.traveledPos[Pos{i, guard.currentPos.Y}] = true
            if mapSlice[i][guard.currentPos.Y] == '#' {
                newPos.X = i+1
                posMoved = guard.currentPos.X - newPos.X
                guard.traveledPos[Pos{i, guard.currentPos.Y}] = false
                break
            }
        }
    case 1: // East
        for i:=guard.currentPos.Y; i<len(mapSlice[guard.currentPos.X]); i++ {
            guard.traveledPos[Pos{guard.currentPos.X, i}] = true
            if mapSlice[guard.currentPos.X][i] == '#' {
                newPos.Y = i-1
                posMoved = guard.currentPos.Y + newPos.Y
                guard.traveledPos[Pos{guard.currentPos.X, i}] = false
                break
            }
        }
        
    case 2: // South
        for i:=guard.currentPos.X; i<len(mapSlice); i++ {
            guard.traveledPos[Pos{i, guard.currentPos.Y}] = true
            if mapSlice[i][guard.currentPos.Y] == '#' {
                newPos.X = i-1
                posMoved = guard.currentPos.X + newPos.X
                guard.traveledPos[Pos{i, guard.currentPos.Y}] = false
                break
            } 
        }
    case 3: // West
        for i:=guard.currentPos.Y; i>=0; i-- {
            guard.traveledPos[Pos{guard.currentPos.X, i}] = true
            if mapSlice[guard.currentPos.X][i] == '#' {
                newPos.Y = i+1
                posMoved = guard.currentPos.Y - newPos.Y
                guard.traveledPos[Pos{guard.currentPos.X, i}] = false
                break
            }
        }
    }
    return newPos, posMoved
}

func part2() {
    ans := 0

    mapSlice := libs.FileToSlice("2024/day06/input.txt", "\n")
    guard := getGuard(mapSlice)

    dirIdx := 0
    visited := make(map[Pos]bool)
    visited[guard.currentPos] = true

    rowObstacles := scanRowObstacles(mapSlice)
    colObstacles := scanColObstacles(mapSlice)

    for {
        newPos := moveGuard(guard.currentPos, dirIdx)
        if !isInside(mapSlice, newPos) {
            break
        }
        if mapSlice[newPos.X][newPos.Y] == '#' {
            dirIdx = (dirIdx + 1) % 4 // Turn right.
        } else {
            if !visited[newPos] {
                mapSlice[newPos.X] = libs.ReplaceCharAtIndex(mapSlice[newPos.X], newPos.Y, '#')
                rowObstacles[newPos.X] = scanRowObstacles(mapSlice)[newPos.X]
                colObstacles[newPos.Y] = scanColObstacles(mapSlice)[newPos.Y]
                if isLooped(rowObstacles, colObstacles, guard.currentPos, dirIdx) {
                    ans++
                }
                mapSlice[newPos.X] = libs.ReplaceCharAtIndex(mapSlice[newPos.X], newPos.Y, '.')
                rowObstacles[newPos.X] = scanRowObstacles(mapSlice)[newPos.X]
                colObstacles[newPos.Y] = scanColObstacles(mapSlice)[newPos.Y]
                visited[newPos] = true
            }
            guard.currentPos = newPos
        }
    }
    
    fmt.Println("The answer to part 2 for day 06 is:", ans)
}

func scanRowObstacles(mapSlice []string) [][]int {
    rowObstacles := make([][]int, len(mapSlice))
    for row := range mapSlice {
        for col := range mapSlice[row] {
            if mapSlice[row][col] == '#' {
                rowObstacles[row] = append(rowObstacles[row], col)
            }
        }
    }
    return rowObstacles
}

func scanColObstacles(mapSlice []string) [][]int {
    colObstacles := make([][]int, len(mapSlice[0]))
    for col := range mapSlice[0] {
        for row := range mapSlice {
            if len(mapSlice[row]) > 0 {
                if mapSlice[row][col] == '#' {
                    colObstacles[col] = append(colObstacles[col], row)
                }
            }
        }
    }
    return colObstacles
}

func moveGuard(pos Pos, dirIdx int) Pos {
    switch dirIdx {
    case 0: // North
        return Pos{pos.X - 1, pos.Y}
    case 1: // East
        return Pos{pos.X, pos.Y + 1}
    case 2: // South
        return Pos{pos.X + 1, pos.Y}
    case 3: // West
        return Pos{pos.X, pos.Y - 1}
    default:
        return pos
    }
}

func isInside(mapSlice []string, pos Pos) bool {
    return pos.X >= 0 && pos.X < len(mapSlice) && pos.Y >= 0 && pos.Y < len(mapSlice[0])
}

func isLooped(rowObstacles, colObstacles [][]int, guardPos Pos, dirIdx int) bool {
    type State struct {
        pos    Pos
        dirIdx int
    }
    state := State{pos: guardPos, dirIdx: dirIdx}
    visitedUp := make(map[Pos]bool)
    for state.dirIdx != 0 || !visitedUp[state.pos] {
        switch state.dirIdx {
        case 0: // North
            visitedUp[state.pos] = true
            state.pos.X = lowerBound(colObstacles[state.pos.Y], state.pos.X, -100) + 1
        case 1: // East
            state.pos.Y = upperBound(rowObstacles[state.pos.X], state.pos.Y, -100) - 1
        case 2: // South
            state.pos.X = upperBound(colObstacles[state.pos.Y], state.pos.X, -100) - 1
        case 3: // West
            state.pos.Y = lowerBound(rowObstacles[state.pos.X], state.pos.Y, -100) + 1
        default:
            panic("invalid direction")
        }
        if state.pos.X < 0 || state.pos.Y < 0 {
            return false
        }
        state.dirIdx = (state.dirIdx + 1) % 4
    }
    return true
}

func lowerBound(arr []int, value int, defaultValue int) int {
    ans := defaultValue
    for _, x := range arr {
        if x >= value {
            break
        }
        ans = x
    }
    return ans
}

func upperBound(arr []int, value int, defaultValue int) int {
    for _, x := range arr {
        if x > value {
            return x
        }
    }
    return defaultValue
}