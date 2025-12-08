package main

import (
    "fmt"
    "sort"
    "strconv"
    "strings"

    "github.com/DeanLogan/advent-of-code/libs"
)

type edge struct {
    node1, node2 int
    distance float64
}

func main(){
    part1()
    part2()
}

func part1(){
    lines := libs.FileToSlice("2025/day08/input.txt", "\n")
    coords := parseCoords(lines)
    edges := computeEdges(coords)
    
    // stated in problem 1000 pairs of connections
    graph := buildGraph(edges, 1000, len(coords))
    
    visited := make([]bool, len(coords))
    sizes := []int{}
    
    for i := 0; i < len(coords); i++ {
        if !visited[i] {
            size := bfsComponentSize(i, graph, visited)
            sizes = append(sizes, size)
        }
    }
    
    sort.Slice(sizes, func(i, j int) bool {
        return sizes[i] > sizes[j]
    })

    ans := sizes[0] * sizes[1] * sizes[2]
    fmt.Println("🎄 The answer to part 1 for day 08 is:", ans, "🎄")
}

func parseCoords(lines []string) []libs.Pos3D {
    coords := make([]libs.Pos3D, 0, len(lines))
    for _, line := range lines {
        parts := strings.Split(line, ",")
        x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
        y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
        z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
        coords = append(coords, libs.Pos3D{X: x, Y: y, Z: z})
    }
    return coords
}

func computeEdges(coords []libs.Pos3D) []edge {
    edges := []edge{}
    for i := range coords {
        for j := i + 1; j < len(coords); j++ {
            distance := libs.EuclideanDistance3D(coords[i], coords[j])
            edges = append(edges, edge{
                node1:    i,
                node2:    j,
                distance: distance,
            })
        }
    }
    
    sort.Slice(edges, func(i, j int) bool {
        return edges[i].distance < edges[j].distance
    })
    
    return edges
}

func buildGraph(edges []edge, connections int, nodeCount int) map[int][]int {
    graph := make(map[int][]int)
    
    for i := range nodeCount {
        graph[i] = []int{}
    }
    
    for i:=0; i<connections && i<len(edges); i++ {
        graph[edges[i].node1] = append(graph[edges[i].node1], edges[i].node2)
        graph[edges[i].node2] = append(graph[edges[i].node2], edges[i].node1)
    }
    
    return graph
}

func bfsComponentSize(start int, graph map[int][]int, visited []bool) int {
    queue := []int{start}
    visited[start] = true
    size := 0
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        size++
        
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
    
    return size
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 08 is:", ans, "🎄")
}

