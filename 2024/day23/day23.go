package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0

    lines := libs.FileToSlice("2024/day23/input.txt", "\n")
    connections := make(map[string][]string)

    for _, line := range lines {
        addConnection(line, connections)
    }

    allConnectionSets := findConnectedTriplets(connections)

    for _, connSet := range allConnectionSets {
        for _, comp := range connSet {
            if comp[0] == 't' {
                ans++
                break
            }
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 23 is:", ans, "🎄")
}

func addConnection(connStr string, connections map[string][]string) {
    comp1, comp2 := libs.SplitAtChar(connStr, '-')
    connections[comp1] = append(connections[comp1], comp2)
    connections[comp2] = append(connections[comp2], comp1)
}

func findConnectedTriplets(connections map[string][]string) [][]string {
    triplets := [][]string{}
    for comp1, neighbors1 := range connections {
        for _, comp2 := range neighbors1 {
            for _, comp3 := range connections[comp2] {
                if comp3 != comp1 && contains(connections[comp3], comp1) {
                    triplet := []string{comp1, comp2, comp3}
                    if !containsTriplet(triplets, triplet) {
                        triplets = append(triplets, triplet)
                    }
                }
            }
        }
    }
    return triplets
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

func containsTriplet(triplets [][]string, triplet []string) bool {
    for _, t := range triplets {
        if (t[0] == triplet[0] && t[1] == triplet[1] && t[2] == triplet[2]) ||
            (t[0] == triplet[0] && t[1] == triplet[2] && t[2] == triplet[1]) ||
            (t[0] == triplet[1] && t[1] == triplet[0] && t[2] == triplet[2]) ||
            (t[0] == triplet[1] && t[1] == triplet[2] && t[2] == triplet[0]) ||
            (t[0] == triplet[2] && t[1] == triplet[0] && t[2] == triplet[1]) ||
            (t[0] == triplet[2] && t[1] == triplet[1] && t[2] == triplet[0]) {
            return true
        }
    }
    return false
}

func part2() {
    lines := libs.FileToSlice("2024/day23/input.txt", "\n")
    connections := make(map[string][]string)

    for _, line := range lines {
        addConnection(line, connections)
    }

    allConnectionSets := findAllCliques(connections)

    longestConnSet := []string{}
    for _, connSet := range allConnectionSets {
        if len(connSet) > len(longestConnSet) {
            longestConnSet = connSet
        }
    }
    sort.Strings(longestConnSet)
    ans := strings.Join(longestConnSet, ",")

    fmt.Println("🎄 The answer to part 2 for day 23 is:", ans, "🎄")
}

func findAllCliques(connections map[string][]string) [][]string {
    var cliques [][]string
    visited := make(map[string]bool)

    for node := range connections {
        if !visited[node] {
            clique := []string{}
            findClique(node, connections, visited, &clique)
            if len(clique) > 1 {
                cliques = append(cliques, clique)
            }
        }
    }

    return cliques
}

func findClique(node string, connections map[string][]string, visited map[string]bool, clique *[]string) {
    visited[node] = true
    *clique = append(*clique, node)

    for _, neighbor := range connections[node] {
        if !visited[neighbor] && isFullyConnected(neighbor, *clique, connections) {
            findClique(neighbor, connections, visited, clique)
        }
    }
}

func isFullyConnected(node string, clique []string, connections map[string][]string) bool {
    for _, member := range clique {
        if !contains(connections[member], node) {
            return false
        }
    }
    return true
}