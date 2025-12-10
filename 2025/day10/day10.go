package main

import (
    "fmt"
    "regexp"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type Machine struct {
    lightDiagram []int
    buttons [][]int
    joltage []int
}

type State struct {
    lights []int
    presses int
}

func part1(){
    ans := 0
    lines := libs.FileToSlice("2025/day10/input.txt", "\n")
    machines := convertLinesToMachine(lines)

    for _, machine := range machines {
        minPresses := solveBFS(machine)
        if minPresses != -1 {
            ans += minPresses
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 10 is:", ans, "🎄")
}

func extractlightDiagram(line string) []int {
    lightRegex := regexp.MustCompile(`\[([.\#]+)\]`)
    if match := lightRegex.FindStringSubmatch(line); len(match) > 1 {
        str := match[1]
        result := make([]int, len(str))
        for i, ch := range str {
            if ch == '#' {
                result[i] = 1
            }
        }
        return result
    }
    return []int{}
}

func extractButtons(line string) [][]int {
    buttons := [][]int{}
    buttonRegex := regexp.MustCompile(`\(([0-9,]+)\)`)
    buttonMatches := buttonRegex.FindAllStringSubmatch(line, -1)
    
    for _, match := range buttonMatches {
        buttons = append(buttons, libs.StrToIntSlice(match[1], ","))
    }
    
    return buttons
}

func extractJoltage(line string) []int {
    joltageRegex := regexp.MustCompile(`\{([0-9,]+)\}`)
    
    if match := joltageRegex.FindStringSubmatch(line); len(match) > 1 {
        return libs.StrToIntSlice(match[1], ",")
    }
    
    return []int{}
}

func convertLinesToMachine(lines []string) []Machine {
    machines := []Machine{}
    for _, line := range lines {
        lightDiagram := extractlightDiagram(line)
        
        machine := Machine{
            lightDiagram:   lightDiagram,
            buttons:        extractButtons(line),
            joltage:        extractJoltage(line),
        }
        machines = append(machines, machine)
    }
    return machines
}

func lightsToKey(lights []int) string {
    return libs.IntSliceToStr(lights, ",")
}

func applyButtonPress(currentLights []int, button []int) []int {
    newLights := make([]int, len(currentLights))
    copy(newLights, currentLights)
    
    for _, lightIdx := range button {
        newLights[lightIdx] = (newLights[lightIdx] + 1) % 2
    }
    
    return newLights
}

func initializeBFS(machine Machine) (State, map[string]bool, string) {
    startState := State{
        lights: make([]int, len(machine.lightDiagram)), 
        presses: 0,
    }
    
    visited := make(map[string]bool)
    visited[lightsToKey(startState.lights)] = true
    
    targetKey := lightsToKey(machine.lightDiagram)
    
    return startState, visited, targetKey
}

func isTargetReached(lights []int, target []int) bool {
    return libs.CompareSlices(lights, target)
}

func processButtonPress(button []int, current State, visited map[string]bool, targetKey string) (State, bool, bool) {
    newLights := applyButtonPress(current.lights, button)
    key := lightsToKey(newLights)
    
    if visited[key] {
        return State{}, false, false
    }
    visited[key] = true
    
    if key == targetKey {
        return State{lights: newLights, presses: current.presses + 1}, true, true
    }
    return State{lights: newLights, presses: current.presses + 1}, true, false
}

func solveBFS(machine Machine) int {
    startState, visited, targetKey := initializeBFS(machine)
    
    queue := []State{startState}
    queueIndex := 0
    
    for queueIndex < len(queue) {
        current := queue[queueIndex]
        queueIndex++
        
        if isTargetReached(current.lights, machine.lightDiagram) {
            return current.presses
        }
        
        for _, button := range machine.buttons {
            newState, shouldProcess, targetReached := processButtonPress(button, current, visited, targetKey)
            
            if targetReached {
                return newState.presses
            }
            
            if shouldProcess {
                queue = append(queue, newState)
            }
        }
    }
    
    return -1
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 10 is:", ans, "🎄")
}