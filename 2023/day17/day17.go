package main

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1() {
	ans := 0

    lines := libs.FileToSlice("2023/day17/input.txt", "\n")
    graph := createMatrix(lines) 

	ans = dijktras(graph, 0, 3)
	fmt.Println("🎄 The answer to part 1 for day 17 is:", ans, "🎄")
}

func createMatrix(input []string) [][]int {
	output := make([][]int, len(input))
	for i, line := range input {
		output[i] = make([]int, len(line))
		for j, char := range line {
			num, _ := strconv.ParseInt(string(char), 10, 64)
			output[i][j] = int(num)
		}
	}
	return output
}

type step struct {
	x       int
	y       int
	lastDir direction
	count   int
}

type queueNode struct {
	step
	heatLoss int
}

type direction string

const (
	up    direction = "up"
	down  direction = "down"
	left  direction = "left"
	right direction = "right"
)

type PriorityQueue []*queueNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(queueNode) 
    *pq = append(*pq, &item) 
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return *item 
}

func dijktras(matrix [][]int, minX, maxX int) int {
    priorityQueue := &PriorityQueue{}
    heap.Init(priorityQueue)

    heap.Push(priorityQueue, queueNode{step{0, 1, right, 1}, matrix[0][1]})
    heap.Push(priorityQueue, queueNode{step{1, 0, down, 1}, matrix[1][0]})

    visited := make(map[step]int)

    for priorityQueue.Len() > 0 {
        element := heap.Pop(priorityQueue)

        currentNode := element.(queueNode)

        if _, ok := visited[currentNode.step]; ok {
            continue
        }

        if currentNode.x == len(matrix)-1 && currentNode.y == len(matrix[0])-1 {
            if currentNode.count < minX {
                continue
            }
            return currentNode.heatLoss
        }

        dirs := fetchPossibleDirections(currentNode, minX, maxX) 
        possibleNextNodes := getNodesFromDirection(dirs, currentNode, matrix)
        for _, p := range possibleNextNodes {
            if _, ok := visited[p.step]; ok {
                continue
            }
            heap.Push(priorityQueue, p)
        }

        visited[currentNode.step] = currentNode.heatLoss
    }

    return -1
}

type point struct {
	x int
	y int
}

var positionMap = map[direction]point{
	up:    {-1, 0},
	down:  {1, 0},
	left:  {0, -1},
	right: {0, 1},
}

func isValid(p point, input [][]int) bool {
	if p.x < 0 || p.x > len(input)-1 || p.y < 0 || p.y > len(input[0])-1 {
		return false
	}
	return true
}

func fetchPossibleDirections(node queueNode, min, max int) []direction {
	output := []direction{}
	switch node.lastDir {
	case up, down:
		if node.count < max {
			output = append(output, node.lastDir)
		}
		if node.count >= min {
			output = append(output, left)
			output = append(output, right)
		}
	case left, right:
		if node.count < max {
			output = append(output, node.lastDir)
		}
		if node.count >= min {
			output = append(output, up)
			output = append(output, down)
		}
	}
	return output
}

func getNodesFromDirection(dirs []direction, startNode queueNode, matrix [][]int) []queueNode {
	nodes := []queueNode{}
	for _, dir := range dirs {
		addPoint := positionMap[dir]
		newPoint := point{startNode.x + addPoint.x, startNode.y + addPoint.y}

		if !isValid(newPoint, matrix) {
			continue
		}

		count := 1
		if startNode.lastDir == dir {
			count = startNode.count + 1
		}

		nodes = append(nodes, queueNode{
			step{
				x:       startNode.x + addPoint.x,
				y:       startNode.y + addPoint.y,
				lastDir: dir,
				count:   count,
			},
			startNode.heatLoss + matrix[newPoint.x][newPoint.y],
		})
	}
	return nodes
}

func part2() {
	ans := 0

    lines := libs.FileToSlice("2023/day17/input.txt", "\n")
    graph := createMatrix(lines) 

    ans = dijktras(graph, 4, 10)
	fmt.Println("🎄 The answer to part 2 for day 17 is:", ans, "🎄")
}