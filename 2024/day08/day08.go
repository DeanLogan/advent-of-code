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
    X int
    Y int
}

func part1(){
    ans := 0

    mapSlice := libs.FileToSlice("2024/day08/input.txt", "\n")
    antennas := createAntennaMap(mapSlice)
    width := len(mapSlice[0])
    height := len(mapSlice)
    antinodeLocations := make(map[Pos]bool)
    
    for _, antennaPosList := range antennas {
        antinodesList := findAntinodesForAntennas(antennaPosList, width, height)
        ans += len(antinodesList)
        for _, antinode := range antinodesList {
            antinodeLocations[antinode] = true
        }
    }

    ans = len(antinodeLocations)

    fmt.Println("The answer to part 1 for day 08 is:", ans)
}

func createAntennaMap(mapSlice []string) map[rune][]Pos {
    antennas := make(map[rune][]Pos)
    for y, row := range mapSlice {
        for x, char := range row {
            if char != '.' {
                antennas[char] = append(antennas[char], Pos{x, y})
            }
        }
    } 
    return antennas
}

func findAntinodesForAntennas(antennaPos []Pos, width int, height int) []Pos {
    antinodes := []Pos{}

    for i, centerPos := range antennaPos {
        for j, otherPos := range antennaPos {
            if i != j {
                inverseX := 2*centerPos.X - otherPos.X
                inverseY := 2*centerPos.Y - otherPos.Y
                if inverseX >= 0 && inverseY >= 0 && inverseX < width && inverseY < height  {
                    antinodes = append(antinodes, Pos{inverseX, inverseY})
                } 
            }
        }
    }

    return antinodes
}

func part2() {
    ans := 0

    mapSlice := libs.FileToSlice("2024/day08/input.txt", "\n")
    antennas := createAntennaMap(mapSlice)
    width := len(mapSlice[0])
    height := len(mapSlice)

    antinodeLocations := findAntinodeLocations(antennas, width, height)
    mapSlice = placePointsOnMap(mapSlice, antinodeLocations)
    
    for _, row := range mapSlice {
        for _, char := range row {
            if char != '.' {
                ans++
            }
        }
    }

    fmt.Println("The answer to part 2 for day 08 is:", ans)
}

func findAntinodeLocations(antennas map[rune][]Pos, width int, height int) []Pos {
    antinodes := []Pos{}

    for _, antennaPosList := range antennas {
        for i, centerPos := range antennaPosList {
            for j, otherPos := range antennaPosList {
                if i != j {
                    diffX := otherPos.X - centerPos.X
                    diffY := otherPos.Y - centerPos.Y
                    antiX := otherPos.X + diffX
                    antiY := otherPos.Y + diffY

                    for 0 <= antiX && antiX < width && 0 <= antiY && antiY < height {
                        antinodes = append(antinodes, Pos{antiX, antiY})
                        antiX += diffX
                        antiY += diffY
                    }
                }
            }
        }
    }

    return antinodes
}

func placePointsOnMap(mapSlice []string, points []Pos) []string {
    for _, point := range points {
        mapSlice[point.Y] = libs.ReplaceCharAtIndex(mapSlice[point.Y], point.X, '#')
    }
    return mapSlice
}