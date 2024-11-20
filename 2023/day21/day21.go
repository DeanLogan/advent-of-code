package main

import (
	"fmt"
	"strings"
    "math"
    "gonum.org/v1/gonum/mat"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    lines := libs.FileToSlice("2023/day21/input.txt", "\n")
    fmt.Println("The answer to part 1 for day 21 is:", floodFill(lines, 64))
}

type Point struct {
    x, y, steps int
}

func floodFill(data []string, steps int) int {
    width := len(data[0])
    height := len(data)

    q := make([]Point, 0)
    sx, sy := width/2, height/2

    q = append(q, Point{sx, sy, 0})
    s64 := make(map[Point]bool)
    visited := make(map[Point]bool)

    for len(q) > 0 {
        point := q[0]
        q = q[1:]

        if visited[point] {
            continue
        }
        visited[point] = true

        if point.steps == steps {
            s64[Point{point.x, point.y, 0}] = true
        } else {
            if point.x > 0 && data[point.y][point.x-1] != '#' {
                q = append(q, Point{point.x - 1, point.y, point.steps + 1})
            }
            if point.x < width-1 && data[point.y][point.x+1] != '#' {
                q = append(q, Point{point.x + 1, point.y, point.steps + 1})
            }
            if point.y > 0 && data[point.y-1][point.x] != '#' {
                q = append(q, Point{point.x, point.y - 1, point.steps + 1})
            }
            if point.y < height-1 && data[point.y+1][point.x] != '#' {
                q = append(q, Point{point.x, point.y + 1, point.steps + 1})
            }
        }
    }

    if steps == 64 {
        for y, line := range data {
            ll := strings.Split(line, "")
            for point := range s64 {
                if point.y == y {
                    ll[point.x] = "O"
                }
            }
        }
    }

    return len(s64)
}

func part2(){
    lines := libs.FileToSlice("2023/day21/input.txt", "\n")
    newLines := []string{}
    for i := 0; i < 5; i++ {
        for _, line := range lines {
            newLines = append(newLines, strings.Repeat(strings.ReplaceAll(line, "S", "."), 5))
        }
    }

    a0 := float64(floodFill(newLines, 65))
    a1 := float64(floodFill(newLines, 65+131))
    a2 := float64(floodFill(newLines, 65+2*131))

    b := mat.NewVecDense(3, []float64{a0, a1, a2})

    vandermondeData := []float64{0, 0, 1, 1, 1, 1, 4, 2, 1}
    vandermonde := mat.NewDense(3, 3, vandermondeData)

    var x mat.VecDense
    x.SolveVec(vandermonde, b)

    n := 202300

    ans := int64(math.Round(x.At(0, 0)*float64(n*n)+x.At(1, 0)*float64(n)+x.At(2, 0)))
    fmt.Println("The answer to part 2 for day 21 is:", ans)
}
