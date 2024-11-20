package main

import (
	"fmt"
	"strconv"
	"strings"
    "math"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type Point struct {
    X int 
    Y int
}

func part1(){
    ans := 0
    digPlan := libs.FileToSlice("2023/day18/input.txt", "\n")
    polygon:= createPolygon(digPlan)
    ans = perimeter(polygon) / 2 + shoelace(polygon) + 1
    fmt.Println("The answer to part 1 for day 18 is:", ans)
}

func createPolygon(digPlan []string) []Point {
    newPoint := Point{0,0}
    polygon := []Point{newPoint}
    for _, line := range digPlan {
        parts := strings.Split(line, " ")
        distance, _ := strconv.Atoi(parts[1])
        switch parts[0]{
        case "R":
            newPoint.X += distance
        case "L":
            newPoint.X -= distance
        case "U":
            newPoint.Y -= distance
        case "D":
            newPoint.Y += distance
        default:
            fmt.Println("Error")
        }
        
        polygon = append(polygon, Point{newPoint.X, newPoint.Y})
    }
    return polygon
}

func shoelace(points []Point) int {
    area := 0
    n := len(points)

    for i := 0; i < n; i++ {
        j := (i + 1) % n
        area += points[i].X * points[j].Y
        area -= points[j].X * points[i].Y
    }

    return int(math.Abs(float64(area)) / 2)
}

func distance(p1, p2 Point) int {
    dx := p2.X - p1.X
    dy := p2.Y - p1.Y
    return int(math.Sqrt(float64(dx*dx + dy*dy)))
}

func perimeter(points []Point) int {
    perim := 0
    for i := 0; i < len(points); i++ {
        nextIndex := (i + 1) % len(points)
        perim += distance(points[i], points[nextIndex])
    }
    return perim
}

func part2(){
    ans := 0
    digPlan := libs.FileToSlice("2023/day18/input.txt", "\n")
    polygon := createPolygonFromHexCodes(digPlan)
    ans = perimeter(polygon) / 2 + shoelace(polygon) + 1
    fmt.Println("The answer to part 2 for day 18 is:", ans)
}

func createPolygonFromHexCodes(digPlan []string) []Point {
    newPoint := Point{0,0}
    polygon := []Point{newPoint}
    for _, line := range digPlan {
        parts := strings.Split(line, " ")
        hexNumStr := parts[2][2:len(parts[2])-2]
        direction := parts[2][len(parts[2])-2:len(parts[2])-1]
        distance := hexToDecimal(hexNumStr)

        switch direction {
        case "0":
            newPoint.X += distance
        case "2":
            newPoint.X -= distance
        case "3":
            newPoint.Y -= distance
        case "1":
            newPoint.Y += distance
        default:
            fmt.Println("Error")
        }
        
        polygon = append(polygon, Point{newPoint.X, newPoint.Y})
    }
    return polygon
}

var hexCharMap = map[string]int{
    "0": 0,
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "a": 10,
    "b": 11,
    "c": 12,
    "d": 13,
    "e": 14,
    "f": 15,
}

func hexToDecimal(hexNum string) int{
    hexNum = libs.ReverseString(hexNum)
    num := 0
    for i, char := range hexNum {
        num += hexCharMap[string(char)] * int(math.Pow(16, float64(i)))
    }
    return num
}