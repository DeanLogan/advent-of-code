package main

import (
    "container/heap"
    "fmt"
    "math"

    "github.com/DeanLogan/advent-of-code/libs"
)

type Pos struct {
    x int 
    y int
}

func (p Pos) Manhattan(other Pos) int {
    return int(math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y)))
}

func (p Pos) CardinalNeighbors() []Pos {
    return []Pos{
        {p.x + 1, p.y},
        {p.x - 1, p.y},
        {p.x, p.y + 1},
        {p.x, p.y - 1},
    }
}

type Item struct {
    priority int
    pos    Pos
}

type Cheat struct {
    start Pos
    end Pos
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func main() {
    part1()
    part2()
}

func part1() {
    ans := 0

    grid := libs.FileToSlice("2024/day20/input.txt", "\n")
    start := findChar(grid, 'S')
    end := findChar(grid, 'E')

    normalTime, normalCameFrom := aStar(grid, start, end)
    normalPath := constructPath(normalCameFrom, start, end)

    normalPathIndices := make(map[Pos]int)
    for i, point := range normalPath {
        normalPathIndices[point] = i
    }

    cheatSavings := calculateCheatSavings(grid, normalPath, normalPathIndices, normalTime)

    for savings := range cheatSavings {
        if savings >= 100 {
            ans += len(cheatSavings[savings])
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 20 is:", ans, "🎄")
}

func calculateCheatSavings(grid []string, normalPath []Pos, normalPathIndices map[Pos]int, normalTime int) map[int][]Cheat {
    cheatSavings := make(map[int][]Cheat)

    for costSoFar, point := range normalPath {
        for _, cheat := range viableCheats(grid, point) {
            _, cheatEnd := cheat.start, cheat.end
            if jumpTo, ok := normalPathIndices[cheatEnd]; ok {
                cheatTime := costSoFar + 2 + (normalTime - jumpTo)
                if cheatTime < normalTime {
                    savings := normalTime - cheatTime
                    cheatSavings[savings] = append(cheatSavings[savings], cheat)
                }
            }
        }
    }

    return cheatSavings
}

func viableCheats(grid []string, curr Pos) []Cheat {
    var cheats []Cheat
    for _, first := range curr.CardinalNeighbors() {
        for _, end := range first.CardinalNeighbors() {
            if gridGet(grid, end, "#") != "#" && curr != end {
                cheats = append(cheats, Cheat{curr, end})
            }
        }
    }
    return cheats
}

func findChar(grid []string, charToFind rune) Pos {
    for y, line := range grid {
        for x, char := range line {
            if char == charToFind {
                return Pos{x, y}
            }
        }
    }
    return Pos{}
}

func gridGet(grid []string, p Pos, defaultVal string) string {
    if p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[p.y]) {
        return string(grid[p.y][p.x])
    }
    return defaultVal
}

func aStar(grid []string, start, end Pos) (int, map[Pos]Pos) {
    heuristic := func(next Pos) int {
        return next.Manhattan(end)
    }

    frontier := &PriorityQueue{&Item{0, start}}
    heap.Init(frontier)
    cameFrom := make(map[Pos]Pos)
    costSoFar := make(map[Pos]int)
    cameFrom[start] = Pos{-1, -1}
    costSoFar[start] = 0

    for frontier.Len() > 0 {
        current := heap.Pop(frontier).(*Item).pos
        if current == end {
            break
        }
        for _, next := range current.CardinalNeighbors() {
            if gridGet(grid, next, "#") == "#" {
                continue
            }
            newCost := costSoFar[current] + 1
            if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
                costSoFar[next] = newCost
                priority := newCost + heuristic(next)
                heap.Push(frontier, &Item{priority, next})
                cameFrom[next] = current
            }
        }
    }
    return costSoFar[end], cameFrom
}

func constructPath(cameFrom map[Pos]Pos, start, end Pos) []Pos {
    path := []Pos{end}
    for path[len(path)-1] != start {
        prev := cameFrom[path[len(path)-1]]
        path = append(path, prev)
    }
    for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
        path[i], path[j] = path[j], path[i]
    }
    return path
}

func part2() {
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 20 is:", ans, "🎄")
}