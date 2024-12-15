package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main() {
    part1()
    part2()
}

type button struct {
    x int
    y int
}

type machineInfo struct {
    buttonA button
    buttonB button
    prizeX  int
    prizeY  int
}

type state struct {
    x, y, aPresses, bPresses int
}

func part1() {
    ans := 0

    machinesStr := libs.FileToSlice("2024/day13/input.txt", "\n\n")
    for _, machineStr := range machinesStr {
        aPresses, bPresses := minTokenToWinPrizeBfs(extractMachineInfo(machineStr), 100)
        ans += (aPresses * 3) + bPresses
    }

    fmt.Println("🎄 The answer to part 1 for day 13 is:", ans, "🎄")
}

func minTokenToWinPrizeBfs(machine machineInfo, maxPresses int) (int, int) {
    queue := []state{{0, 0, 0, 0}}
    visited := make(map[state]bool)
    visited[queue[0]] = true

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current.x == machine.prizeX && current.y == machine.prizeY {
            return current.aPresses, current.bPresses
        }

        if current.aPresses < maxPresses {
            nextA := state{
                x:        current.x + machine.buttonA.x,
                y:        current.y + machine.buttonA.y,
                aPresses: current.aPresses + 1,
                bPresses: current.bPresses,
            }
            if !visited[nextA] {
                visited[nextA] = true
                queue = append(queue, nextA)
            }
        }

        if current.bPresses < maxPresses {
            nextB := state{
                x:        current.x + machine.buttonB.x,
                y:        current.y + machine.buttonB.y,
                aPresses: current.aPresses,
                bPresses: current.bPresses + 1,
            }
            if !visited[nextB] {
                visited[nextB] = true
                queue = append(queue, nextB)
            }
        }
    }

    return 0, 0
}

func extractMachineInfo(machineStr string) machineInfo {
    machine := machineInfo{}
    lines := strings.Split(machineStr, "\n")
    machine.buttonA = extractButton(lines[0])
    machine.buttonB = extractButton(lines[1])

    lines[2] = lines[2][9:]
    xStrPrize, yStrPrize := libs.SplitAtStr(lines[2], ", Y=")

    machine.prizeX, _ = strconv.Atoi(xStrPrize)
    machine.prizeY, _ = strconv.Atoi(yStrPrize[3:])
    return machine
}

func extractButton(buttonStr string) button {
    button := button{}

    buttonStr = buttonStr[12:]
    xStr, yStr := libs.SplitAtStr(buttonStr, ", Y+")

    button.x, _ = strconv.Atoi(xStr)
    button.y, _ = strconv.Atoi(yStr[3:])

    return button
}

func part2() {
    ans := 0

    machinesStr := libs.FileToSlice("2024/day13/input.txt", "\n\n")
    for _, machineStr := range machinesStr {
        machine := extractMachineInfo(machineStr)
        aPresses, bPresses := minTokenToWinPrizeLinearEquations(machine, 10000000000000)
        ans += (aPresses * 3) + bPresses
    }

    fmt.Println("🎄 The answer to part 2 for day 13 is:", ans, "🎄")
}

func minTokenToWinPrizeLinearEquations(machine machineInfo, delta int) (int, int) {
    pX := machine.prizeX + delta
    pY := machine.prizeY + delta
    aX := machine.buttonA.x
    aY := machine.buttonA.y
    bX := machine.buttonB.x
    bY := machine.buttonB.y

    a := float64(pX*bY - pY*bX) / float64(aX*bY - aY*bX)
    b := float64(pY*aX - pX*aY) / float64(aX*bY - aY*bX)

    if a == math.Trunc(a) && b == math.Trunc(b) {
        return int(a), int(b)
    }
    return 0, 0
}