package main

import (
	"fmt"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    lines := libs.FileToSlice("day10/input.txt", "\n")
    lineOn, posAt := findS(lines)
    startLine, startPos, direction := checkCharSurrondings(lineOn, posAt, lines)
    ans := goRoundMaze(startLine, startPos, lines, direction, 1) / 2 // divide by 2 to get the number of steps furthest away in the loop
    fmt.Println("The answer to part 1 for day 10 is:", ans)
}

func findS(data []string) (int, int) {
    for i := 0; i < len(data); i++ {
        for j := 0; j < len(data[i]); j++ {
            if data[i][j] == 'S' {
                return i, j // i = index of line the character was found on, j = index within the line where the character was found
            }
        }
    }
    return -1, -1
}

// finds the first valid pipe operation from each direction and returns one as a starting point to the maze
func checkCharSurrondings(lineOn int, posAt int, lines []string) (int, int, rune){
    // check west
    if posAt != 0 && (rune(lines[lineOn][posAt-1]) == '-' || rune(lines[lineOn][posAt-1]) == 'F' || rune(lines[lineOn][posAt-1]) == 'L') {
        return lineOn, posAt-1, 'e'
    } else if rune(lines[lineOn][posAt+1]) == '-' || rune(lines[lineOn][posAt+1]) == '7' || rune(lines[lineOn][posAt+1]) == 'J'  { // check east
        return lineOn, posAt+1, 'w'
    } else if lineOn != 0 && (rune(lines[lineOn-1][posAt])  == '|' || rune(lines[lineOn-1][posAt])  == '7' || rune(lines[lineOn-1][posAt])  == 'F') { // check north
        return lineOn-1, posAt, 's'
    } else if rune(lines[lineOn+1][posAt])  == '|' || rune(lines[lineOn-1][posAt])  == 'L' || rune(lines[lineOn-1][posAt])  == 'J' { // check south
        return lineOn+1, posAt, 'n'
    }
    return -1, -1, 'W'
}

func goRoundMaze(lineOn int, posAt int, lines []string, directionFrom rune, charsInMaze int) int {
    lineOn, posAt, directionFrom = handlePipeCharacter(lineOn, posAt, rune(lines[lineOn][posAt]), directionFrom)
    charsInMaze++;
    if  rune(lines[lineOn][posAt]) == 'S'{
        return charsInMaze
    }
    return goRoundMaze(lineOn, posAt, lines, directionFrom, charsInMaze)
}

func handlePipeCharacter(lineOn int, posAt int, char rune, directionFrom rune) (int, int, rune) {
    switch char {
    case '|':
        if directionFrom == 'n'{
            return lineOn+1, posAt, 'n' // moves to the next line, with the same character position from the north
        } else {
            return lineOn-1, posAt, 's'
        }
    case '-':
        if directionFrom == 'w'{
            return lineOn, posAt+1, 'w'
        } else {
            return lineOn, posAt-1, 'e'
        }
    case 'L':
        if directionFrom == 'n'{
            return lineOn, posAt+1, 'w'
        } else {
            return lineOn-1, posAt, 's'
        }
    case 'J':
        if directionFrom == 'n'{
            return lineOn, posAt-1, 'e'
        } else {
            return lineOn-1, posAt, 's'
        }
    case '7':
        if directionFrom == 's'{
            return lineOn, posAt-1, 'e'
        } else {
            return lineOn+1, posAt, 'n'
        }
    case 'F':
        if directionFrom == 's'{
            return lineOn, posAt+1, 'w'
        } else {
            return lineOn+1, posAt, 'n'
        }
    default:
        return 0,0, 'W'
    }
}

type Point struct {
    X, Y int
}

func part2(){
    lines := libs.FileToSlice("day10/input.txt", "\n")
    lineOn, posAt := findS(lines)
    startLine, startPos, direction := checkCharSurrondings(lineOn, posAt, lines)
    mazePoints := []Point{{posAt, lineOn}, {startPos, startLine}}
    mazePoints = mazeToSlice(startLine, startPos, lines, direction, mazePoints)
    ans := countSurroundedPoints(mazePoints, len(lines[0])-1, len(lines)-1)
    fmt.Println("The answer to part 2 for day 10 is:", ans)
}

// adds all of the pipes into a 2d slice, the lineOn and posAt act like points on a graph (the graph being the maze/pipes)
func mazeToSlice(lineOn int, posAt int, lines []string, directionFrom rune, points []Point) []Point {
    lineOn, posAt, directionFrom = handlePipeCharacter(lineOn, posAt, rune(lines[lineOn][posAt]), directionFrom)
    points = append(points, Point{posAt, lineOn})
    if  rune(lines[lineOn][posAt]) == 'S'{
        return points
    }
    return mazeToSlice(lineOn, posAt, lines, directionFrom, points)
}

func countSurroundedPoints(line1 []Point, maxX int, maxY int) int {
    pointMap := make(map[Point]bool)

    for _, point := range line1 {
        pointMap[point] = true
    }

    count := 0
    for x := 0; x <= maxX; x++ {
        for y := 0; y <= maxY; y++ {
            point := Point{x, y}
            if !pointMap[point] {
                if isInside(point, line1) {
                    count++
                }
            }
        }
    }
    return count
}

func isInside(point Point, polygon []Point) bool {
    n := len(polygon)
    inside := false

    p1 := polygon[0]
    for i := 1; i < n+1; i++ {
        p2 := polygon[i%n]
        if point.Y > libs.Min(p1.Y, p2.Y) {
            if point.Y <= libs.Max(p1.Y, p2.Y) {
                if point.X <= libs.Max(p1.X, p2.X) {
                    if p1.Y != p2.Y {
                        xinters := (point.Y-p1.Y)*(p2.X-p1.X)/(p2.Y-p1.Y) + p1.X
                        if p1.X == p2.X || point.X <= xinters {
                            inside = !inside
                        }
                    }
                }
            }
        }
        p1 = p2
    }
    return inside
}