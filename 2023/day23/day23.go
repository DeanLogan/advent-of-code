package main

import (
	"fmt"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main() {
    part1()
    part2()
}

func part1() {
    data := libs.FileToSlice("2023/day23/input.txt", "\n")

    numRows := len(data)
    numColumns := len(data[0])
    src := Point{1, 1}
    dst := Point{numRows - 2, numColumns - 2}

    g := graphFromGrid(data, src, dst, false)
    ans := longestPath(g, src, dst, 0, make(map[Point]struct{})) + 2

    fmt.Println("The answer to part 1 for day 23 is:", ans)
}

type Point struct {
    r, c int
}

type QueueItem struct {
    point    Point
    distance int
}

func neighbors(grid []string, r, c int, ignoreSlopes bool) []Point {
    var points []Point
    if ignoreSlopes || grid[r][c] == '.' {
        for _, rc := range []Point{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}} {
            if rc.r >= 0 && rc.r < len(grid) && rc.c >= 0 && rc.c < len(grid[0]) && grid[rc.r][rc.c] != '#' {
                points = append(points, rc)
            }
        }
    } else if grid[r][c] == 'v' && r+1 < len(grid) {
        points = append(points, Point{r + 1, c})
    } else if grid[r][c] == '^' && r-1 >= 0 {
        points = append(points, Point{r - 1, c})
    } else if grid[r][c] == '>' && c+1 < len(grid[0]) {
        points = append(points, Point{r, c + 1})
    } else if grid[r][c] == '<' && c-1 >= 0 {
        points = append(points, Point{r, c - 1})
    }
    return points
}

func numNeighbors(grid []string, r, c int, ignoreSlopes bool) int {
    if ignoreSlopes || grid[r][c] == '.' {
        count := 0
        for _, rc := range []Point{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}} {
            if rc.r >= 0 && rc.r < len(grid) && rc.c >= 0 && rc.c < len(grid[0]) && grid[rc.r][rc.c] != '#' {
                count++
            }
        }
        return count
    }
    return 1
}

func isNode(grid []string, rc Point, src, dst Point, ignoreSlopes bool) bool {
    return rc == src || rc == dst || numNeighbors(grid, rc.r, rc.c, ignoreSlopes) > 2
}

func adjacentNodes(grid []string, rc, src, dst Point, ignoreSlopes bool) []QueueItem {
    var queue []QueueItem
    seen := make(map[Point]struct{})
    queue = append(queue, QueueItem{rc, 0})

    var items []QueueItem
    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]
        rc, dist := item.point, item.distance
        seen[rc] = struct{}{}

        for _, n := range neighbors(grid, rc.r, rc.c, ignoreSlopes) {
            if _, ok := seen[n]; ok {
                continue
            }

            if isNode(grid, n, src, dst, ignoreSlopes) {
                items = append(items, QueueItem{n, dist + 1})
                continue
            }

            queue = append(queue, QueueItem{n, dist + 1})
        }
    }
    return items
}

func graphFromGrid(grid []string, src, dst Point, ignoreSlopes bool) map[Point][]QueueItem {
    g := make(map[Point][]QueueItem)
    var queue []Point
    seen := make(map[Point]struct{})
    queue = append(queue, src)

    for len(queue) > 0 {
        rc := queue[0]
        queue = queue[1:]
        if _, ok := seen[rc]; ok {
            continue
        }

        seen[rc] = struct{}{}

        for _, item := range adjacentNodes(grid, rc, src, dst, ignoreSlopes) {
            g[rc] = append(g[rc], item)
            queue = append(queue, item.point)
        }
    }
    return g
}

func longestPath(g map[Point][]QueueItem, cur, dst Point, distance int, seen map[Point]struct{}) int {
    if cur == dst {
        return distance
    }

    best := 0
    seen[cur] = struct{}{}

    for _, item := range g[cur] {
        if _, ok := seen[item.point]; ok {
            continue
        }

        best = max(best, longestPath(g, item.point, dst, distance+item.distance, seen))
    }

    delete(seen, cur)
    return best
}

func part2() {
    data := libs.FileToSlice("2023/day23/input.txt", "\n")

    numRows := len(data)
    numColumns := len(data[0])
    src := Point{1, 1}
    dst := Point{numRows - 2, numColumns - 2}

    g := graphFromGrid(data, src, dst, true)
    ans := longestPath(g, src, dst, 0, make(map[Point]struct{})) + 2

    fmt.Println("The answer to part 2 for day 23 is:", ans)
}