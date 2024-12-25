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

type gate struct {
    gateName string
    wire1 string
    wire2 string
    operation string
}

func sortZGatesAndBinary(zGates, zGatesBinary []string) ([]string, []string) {
    type pair struct {
        gate  string
        value string
    }

    pairs := make([]pair, len(zGates))
    for i := range zGates {
        pairs[i] = pair{zGates[i], zGatesBinary[i]}
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].gate > pairs[j].gate
    })

    for i := range pairs {
        zGates[i] = pairs[i].gate
        zGatesBinary[i] = pairs[i].value
    }

    return zGates, zGatesBinary
}

func part1() {
    ans := 0

    sections := libs.FileToSlice("2024/day24/input.txt", "\n\n")
    wires := strings.Split(sections[0], "\n")
    gatesStr := strings.Split(sections[1], "\n")

    wireMap := createWireMap(wires)

    gates := createGateList(gatesStr)

    getAllGateResults(wireMap, gates)

    zGates := []string{}
    zGatesResult := []string{}
    for name, val := range wireMap {
        if name[0] == 'z' {
            zGates = append(zGates, name)
            zGatesResult = append(zGatesResult, val)
        }
    }

    zGates, zGatesResult = sortZGatesAndBinary(zGates, zGatesResult)
    

    zGatesBinary := strings.Join(zGatesResult, "")
    
    ans = libs.BinaryToDecimal(zGatesBinary)

    fmt.Println("🎄 The answer to part 1 for day 24 is:", ans, "🎄")
}

func createWireMap(wires []string) map[string]string {
    wireMap := make(map[string]string)
    for _, wire := range wires {
        wire, val := libs.SplitAtStr(wire, ": ")
        wireMap[wire] = val[1:]
    }
    return wireMap
}

func createGateList(gatesStr []string) []gate {
    gates := []gate{}
    for _, gateStr := range gatesStr {
        strParts := strings.Split(gateStr, " ")
        gates = append(gates, gate{wire1: strParts[0], operation: strParts[1], wire2: strParts[2], gateName: strParts[4]})
    }
    return gates
}

func gateResult(wireMap map[string]string, gate gate) {
    val1 := wireMap[gate.wire1]
    val2 := wireMap[gate.wire2]

    if val1 != "" && val2 != "" {
        result := "0"
        switch gate.operation {
        case "AND":
            if val1 == "1" && val2 == "1" {
                result = "1"
            }
        case "OR":
            if val1 == "1" || val2 == "1" {
                result = "1"
            }
        case "XOR":
            if (val1 == "1" && val2 == "0") || (val1 == "0" && val2 == "1") {
                result = "1"
            }
        }
        wireMap[gate.gateName] = result
    }
}

func getAllGateResults(wireMap map[string]string, gates []gate) {
    remainingGates := []gate{}
    for _, gate := range gates {
        gateResult(wireMap, gate)
        _, gotResult := wireMap[gate.gateName]
        if !gotResult {
            remainingGates = append(remainingGates, gate)
        }
    }
    if len(remainingGates) != 0 {
        getAllGateResults(wireMap, remainingGates)
    }
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 24 is:", ans, "🎄")
}