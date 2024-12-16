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

func part1(){
    ans := 0
    universe := libs.FileToSlice("2023/day11/input.txt", "\n")
    universe = expandUniversePartOne(universe)
    galaxies := getGalaxies(universe)
    for i := 0; i<len(galaxies); i++ {
        ans += calcDistToAllOtherGalaxiesPartOne(galaxies[i], galaxies)
    }
    fmt.Println("🎄 The answer to part 1 for day 11 is:", ans / 2, "🎄")
}

// prints what the slice would look like if it was in a .txt
func printUniverse(universe []string){
    for _, line := range universe {
        fmt.Println(line)
    }
}

type Point struct {
    X, Y int
}

func expandUniversePartOne(universe []string) []string {
    // expands the universe by row
    for i := 0; i<len(universe); i++ {
        if !strings.Contains(universe[i], "#") {
            universe = libs.InsertIntoSlice(universe, i, universe[i])
            i++ // skip the line that was just inserted
        }
    }

    // expands the universe by column
    for i := 0; i<len(universe[0]); i++ { // columns
        foundGalaxy := false
        for j := 0; j<len(universe); j++ { // rows
            if universe[j][i] == '#' {
                foundGalaxy = true
                break
            }
        }
        if !foundGalaxy {
            for j := 0; j<len(universe); j++ {
                universe[j] = universe[j][:i] + "." + universe[j][i:] 
            }
            i++ // skips the next col as it has just been inserted
        }
    }
    return universe
}

func getGalaxies(universe []string) []Point {
    pointsOfGalaxies := []Point{}
    for i := 0; i<len(universe); i++ {
        for j := 0; j<len(universe[i]); j++ { 
            if universe[i][j] == '#' {
                pointsOfGalaxies = append(pointsOfGalaxies, Point{j,len(universe)-i-1}) // the len(universe)-i is so that Point reads the file bottom-up rather than top-down
            }
        }
    }
    return pointsOfGalaxies
}

func shortestPathLength(point1 Point, point2 Point) int {
    return libs.Abs(point1.X - point2.X) + libs.Abs(point1.Y - point2.Y)
}

func calcDistToAllOtherGalaxiesPartOne(galaxy Point, galaxies []Point) int {
    dist := 0
    for _, gal := range galaxies {
        dist += shortestPathLength(galaxy, gal)
    }
    return dist
}

func part2(){
    ans := 0
    universe := libs.FileToSlice("2023/day11/input.txt", "\n")
    rowsExpanding, colsExpanding := expandUniversePartTwo(universe)
    galaxies := getGalaxies(universe)
    for i := 0; i<len(galaxies); i++ {
        ans += calcDistToAllOtherGalaxiesPartTwo(galaxies[i], galaxies, rowsExpanding, colsExpanding)
    }
    fmt.Println("🎄 The answer to part 2 for day 1 is:", ans / 2, "🎄")
}

func expandUniversePartTwo(universe []string) (map[int]bool, map[int]bool) {
    // expands the universe by col
    colsExpanding := make(map[int]bool)
    for i := 0; i<len(universe); i++ {
        if !strings.Contains(universe[i], "#") {
            colsExpanding[len(universe)-i-1] = true 
        }
    }

    // expands the universe by row
    rowsExpanding := make(map[int]bool)
    for i := 0; i<len(universe[0]); i++ { 
        foundGalaxy := false
        for j := 0; j<len(universe); j++ { 
            if universe[j][i] == '#' {
                foundGalaxy = true
                break
            }
        }
        if !foundGalaxy {
            rowsExpanding[i] = true 
        }
    }
    return rowsExpanding, colsExpanding
}

func calcDistToAllOtherGalaxiesPartTwo(galaxy Point, galaxies []Point, rowsExpanding map[int]bool, colsExpanding map[int]bool) int{
    dist := calcDistToAllOtherGalaxiesPartOne(galaxy, galaxies)
    for _, gal := range galaxies {
        dist += addGap(galaxy.X, gal.X, rowsExpanding)
        dist += addGap(galaxy.Y, gal.Y, colsExpanding)
    }
    return dist
}

func addGap(val1 int, val2 int, expandMap map[int]bool) int {
    if val1 > val2 {
        val1, val2 = val2, val1
    }

    dist := 0
    for i := val1+1; i < val2; i++ {
        if expandMap[i] {
            dist += 999999
        }
    }

    return dist
}