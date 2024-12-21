package main

import (
    "fmt"
    "container/heap"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type Pos struct {
	x int
	y int
	dir int
}

type Edge struct {
	p Pos
	score int
}

type Item struct { 
	pos Pos 
	priority int 
	index int 
} 

type PriorityQueue []*Item 
func (pq PriorityQueue) Len() int { 
	return len(pq) 
} 

func (pq PriorityQueue) Less(i, j int) bool { 
	return pq[i].priority < pq[j].priority 
} 

func (pq PriorityQueue) Swap(i, j int) { 
	pq[i], pq[j] = pq[j], pq[i] 
	pq[i].index = i 
	pq[j].index = j 
} 

func (pq *PriorityQueue) Push(x interface{}) { 
	n := len(*pq) 
	item := x.(*Item) 
	item.index = n 
	*pq = append(*pq, item) 
} 

func (pq *PriorityQueue) Pop() interface{} { 
	old := *pq 
	n := len(old) 
	item := old[n-1] 
	old[n-1] = nil 
	item.index = -1 
	*pq = old[0 : n-1] 
	return item 
} 

func part1(){
	ans := 0

    maze := libs.FileToSlice("2024/day16/input.txt", "\n")
    startPos := getPosForChar(maze, 'S')
    endPos := getPosForChar(maze, 'E')

    weightedMaze := weightMaze(maze)
    ans = dijkstra(weightedMaze, startPos, endPos)

    fmt.Println("🎄 The answer to part 1 for day 16 is:", ans, "🎄")
}

func weightMaze(maze []string) map[Pos][]Edge {
	weightedMaze := make(map[Pos][]Edge)
	
	for i := 0; i < len(maze); i++{
		for j := 0; j < len(maze[i]); j++{
			if maze[i][j] == '.' || maze[i][j] == 'S' || maze[i][j] == 'E'{
				for k := 0; k < 4; k++{
					locEdges := determineEdges(maze, Pos{x:j, y: i, dir: k})
					weightedMaze[Pos{x:j, y: i, dir: k}] = locEdges
				}
			} 
		}
	}

	return weightedMaze
}

func determineEdges(maze []string, pos Pos) []Edge {
    edges := []Edge{}

    valid, nextPos := nextPosValid(maze, pos)
    if valid {
        edges = append(edges, Edge{p: nextPos, score: 1})
    }

    newDir := pos.dir - 1
    if newDir < 0 {
        newDir = 3
    }
    valid, _ = nextPosValid(maze, Pos{x: pos.x, y: pos.y, dir: newDir})
    if valid {
        edges = append(edges, Edge{p: Pos{x: pos.x, y: pos.y, dir: newDir}, score: 1000})
    }

    newDir = pos.dir + 1
    if newDir > 3 {
        newDir = 0
    }
    valid, _ = nextPosValid(maze, Pos{x: pos.x, y: pos.y, dir: newDir})
    if valid {
        edges = append(edges, Edge{p: Pos{x: pos.x, y: pos.y, dir: newDir}, score: 1000})
    }

    return edges
}

func nextPosValid(maze []string, pos Pos) (bool, Pos) {
	nextPos := pos

	switch pos.dir {
	case 0:
		nextPos.y -= 1
	case 1:
		nextPos.x += 1
	case 2:
		nextPos.y += 1
	case 3:
		nextPos.x -= 1
	default:
		break
	}

	if (nextPos.y > 0 && nextPos.y < len(maze) && nextPos.x > 0 && nextPos.x < len(maze[0])) && (maze[nextPos.y][nextPos.x] == '.' || maze[nextPos.y][nextPos.x] == 'S' || maze[nextPos.y][nextPos.x] == 'E') { 
		return true, nextPos
	}
	return false, pos
}

func dijkstra(weightedMaze map[Pos][]Edge, start Pos, end Pos) int { 
    distances := make(map[Pos]int) 
    previousNodes := make(map[Pos][]Pos) 

    for pos := range weightedMaze { 
        distances[pos] = int(^uint(0) >> 1) 
    } 

    distances[start] = 0 
    priorityQueue := make(PriorityQueue, 0) 
    heap.Init(&priorityQueue) 
    heap.Push(&priorityQueue, &Item{ pos: start, priority: 0, }) 
    visited := make(map[Pos]bool) 

    var endPos *Pos 

    for priorityQueue.Len() > 0 { 
        currentPos := heap.Pop(&priorityQueue).(*Item).pos 
        if currentPos.x == end.x && currentPos.y == end.y { 
            endPos = &currentPos 
            break 
        } 
        if visited[currentPos] { 
            continue 
        } 
        visited[currentPos] = true 

        for _, edge := range weightedMaze[currentPos] { 
            if visited[edge.p] { 
                continue 
            } 

            alternativeDist := distances[currentPos] + edge.score 
			
            if alternativeDist < distances[edge.p] { 
                distances[edge.p] = alternativeDist 
                previousNodes[edge.p] = []Pos{currentPos} 
                heap.Push(&priorityQueue, &Item{ pos: edge.p, priority: alternativeDist, }) 
            } else if alternativeDist == distances[edge.p] { 
                    previousNodes[edge.p] = append(previousNodes[edge.p], currentPos) 
            } 
        } 
    } 

    if endPos == nil { 
        return -1
    } 

    return distances[*endPos]
}

func getPosForChar(maze []string, char rune) Pos {
    for y, line := range maze {
        for x, checkChar := range line {
            if checkChar == char {
                return Pos{x, y, 1}
            }
        }
    }
    return Pos{}
}

func part2(){
    ans := 0

	maze := libs.FileToSlice("2024/day16/input.txt", "\n")
    startPos := getPosForChar(maze, 'S')
    endPos := getPosForChar(maze, 'E')

    weightedMaze := weightMaze(maze)

	bestPaths := dijkstra2(weightedMaze, startPos, endPos)
	uniqueTiles := make(map[Pos]bool)

	for _, path := range bestPaths {
		for _, tile := range path {
			uniqueTiles[Pos{x: tile.x, y: tile.y, dir:0}] = true
		}
	}

	ans = len(uniqueTiles)

    fmt.Println("🎄 The answer to part 2 for day 16 is:", ans,"🎄")
}

func dijkstra2(weightedMaze map[Pos][]Edge, start Pos, end Pos) [][]Pos { 
    distances := make(map[Pos]int) 
    previousNodes := make(map[Pos][]Pos) 

    for pos := range weightedMaze { 
        distances[pos] = int(^uint(0) >> 1) 
    } 

    distances[start] = 0 
    priorityQueue := make(PriorityQueue, 0) 
    heap.Init(&priorityQueue) 
    heap.Push(&priorityQueue, &Item{ pos: start, priority: 0, }) 
    visited := make(map[Pos]bool) 

    var endPos *Pos 

    for priorityQueue.Len() > 0 { 
        currentPos := heap.Pop(&priorityQueue).(*Item).pos 
        if currentPos.x == end.x && currentPos.y == end.y { 
            endPos = &currentPos 
            break 
        } 
        if visited[currentPos] { 
            continue 
        } 
        visited[currentPos] = true 

        for _, edge := range weightedMaze[currentPos] { 
            if visited[edge.p] { 
                continue 
            } 

            alternativeDist := distances[currentPos] + edge.score 

            if alternativeDist < distances[edge.p] { 
                distances[edge.p] = alternativeDist 
                previousNodes[edge.p] = []Pos{currentPos} 
                heap.Push(&priorityQueue, &Item{ pos: edge.p, priority: alternativeDist, }) 
            } else if alternativeDist == distances[edge.p] { 
                    previousNodes[edge.p] = append(previousNodes[edge.p], currentPos) 
            } 
        } 
    } 

    if endPos == nil { 
        return nil 
    } 
    
    paths := reconstructPaths(previousNodes, *endPos) 
    return paths 
}

func reconstructPaths(previous map[Pos][]Pos, end Pos) [][]Pos { 
	var paths [][]Pos 
	var stack [][]Pos 

	stack = append(stack, []Pos{end}) 
	for len(stack) > 0 { 
		path := stack[len(stack)-1] 
		stack = stack[:len(stack)-1] 
		cur := path[0] 

		if len(previous[cur]) == 0 { 
			paths = append(paths, path) 
		} else { 
			for _, prev := range previous[cur] { 
				newPath := make([]Pos, len(path)+1) 
				copy(newPath[1:], path) 
				newPath[0] = prev 
				stack = append(stack, newPath) 
			} 
		} 
	} 
	
	for i, path := range paths { 
		for j := 0; j < len(path)/2; j++ { 
			path[j], path[len(path)-1-j] = path[len(path)-1-j], path[j] 
		} 

		paths[i] = path 
	} 
	return paths
}