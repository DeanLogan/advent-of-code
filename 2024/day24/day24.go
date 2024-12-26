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

func part1() {
    ans := 0

    sections := libs.FileToSlice("2024/day24/input.txt", "\n\n")
    wires := strings.Split(sections[0], "\n")
    gatesStr := strings.Split(sections[1], "\n")

    wireMap := createWireMap(wires)
    gates := createGateList(gatesStr)
    getAllGateResults(wireMap, gates)

    ans = libs.BinaryToDecimal(getBinaryForGatesWithChar(wireMap, 'z'))

    fmt.Println("🎄 The answer to part 1 for day 24 is:", ans, "🎄")
}

func getBinaryForGatesWithChar(wireMap map[string]string, char byte) string {
    gates := []string{}
    gatesResults := []string{}
    for name, val := range wireMap {
        if name[0] == char {
            gates = append(gates, name)
            gatesResults = append(gatesResults, val)
        }
    }

    _, gatesResults = sortGatesAndBinary(gates, gatesResults)
    return strings.Join(gatesResults, "")
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

func sortGatesAndBinary(zGates, zGatesBinary []string) ([]string, []string) {
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

func part2() {
    sections := libs.FileToSlice("2024/day24/input.txt", "\n\n")
    gatesStr := strings.Split(sections[1], "\n")
    gates := createGateList(gatesStr)

    ans := processGatesAndSwaps(gates)

    fmt.Println("🎄 The answer to part 2 for day 24 is:", ans, "🎄")
}

func find(wire1, wire2, operator string, gates []gate) string {
    for _, gate := range gates {
        if (gate.wire1 == wire1 && gate.operation == operator && gate.wire2 == wire2) ||
            (gate.wire1 == wire2 && gate.operation == operator && gate.wire2 == wire1) {
            return gate.gateName
        }
    }
    return ""
}

func processGatesAndSwaps(gates []gate) string {
    swapped := []string{}
    carryIn := ""

    for i := 0; i < 45; i++ {
        index := fmt.Sprintf("%02d", i)
        var sum, carryOut, intermediateCarry, finalSum, nextCarry string

        sum = find("x"+index, "y"+index, "XOR", gates)
        carryOut = find("x"+index, "y"+index, "AND", gates)

        if carryIn != "" {
            intermediateCarry = find(carryIn, sum, "AND", gates)
            if intermediateCarry == "" {
                sum, carryOut = carryOut, sum
                swapped = append(swapped, sum, carryOut)
                intermediateCarry = find(carryIn, sum, "AND", gates)
            }

            finalSum = find(carryIn, sum, "XOR", gates)

            if strings.HasPrefix(sum, "z") {
                sum, finalSum = finalSum, sum
                swapped = append(swapped, sum, finalSum)
            }

            if strings.HasPrefix(carryOut, "z") {
                carryOut, finalSum = finalSum, carryOut
                swapped = append(swapped, carryOut, finalSum)
            }

            if strings.HasPrefix(intermediateCarry, "z") {
                intermediateCarry, finalSum = finalSum, intermediateCarry
                swapped = append(swapped, intermediateCarry, finalSum)
            }

            nextCarry = find(intermediateCarry, carryOut, "OR", gates)
        }

        if strings.HasPrefix(nextCarry, "z") && nextCarry != "z45" {
            nextCarry, finalSum = finalSum, nextCarry
            swapped = append(swapped, nextCarry, finalSum)
        }

        if carryIn == "" {
            carryIn = carryOut
        } else {
            carryIn = nextCarry
        }
    }

    sort.Strings(swapped)
    return strings.Join(swapped, ",")
}