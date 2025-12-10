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

type Pattern struct {
    pattern []int
    cost    int
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

// can't take credit for the solution to part 2 found this on reddit in the reddit thread it links to a python solution which I just reimplemented in go
//  https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/

func part2(){
    score := 0
    lines := libs.FileToSlice("2025/day10/input.txt", "\n")
    
    for i, line := range lines {
        coeffs, goal := parseLine(line)
        subscore := solveSingle(coeffs, goal)
        fmt.Printf("Line %d/%d: answer %d\n", i+1, len(lines), subscore)
        score += subscore
    }
    
    fmt.Println("🎄 The answer to part 2 for day 10 is:", score, "🎄")
}

func parseLine(line string) ([][]int, []int) {
    joltage := extractJoltage(line)
    buttons := extractButtons(line)
    
    var coeffs [][]int
    for _, button := range buttons {
        coeff := make([]int, len(joltage))
        for _, idx := range button {
            coeff[idx] = 1
        }
        coeffs = append(coeffs, coeff)
    }
    
    return coeffs, joltage
}

func patterns(coeffs [][]int) map[string][]Pattern {
    numButtons := len(coeffs)
    numVariables := len(coeffs[0])
    
    out := make(map[string][]Pattern)
    
    for numPressed := 0; numPressed <= numButtons; numPressed++ {
        combinations := generateCombinations(numButtons, numPressed)
        
        for _, buttons := range combinations {
            pattern := make([]int, numVariables)
            for _, btnIdx := range buttons {
                for i := 0; i < numVariables; i++ {
                    pattern[i] += coeffs[btnIdx][i]
                }
            }
            
            parityPattern := make([]int, numVariables)
            for i := 0; i < numVariables; i++ {
                parityPattern[i] = pattern[i] % 2
            }
            parityKey := libs.IntSliceToStr(parityPattern, ",")
            
            exists := false
            for _, p := range out[parityKey] {
                if intSliceEquals(p.pattern, pattern) {
                    exists = true
                    break
                }
            }
            
            if !exists {
                out[parityKey] = append(out[parityKey], Pattern{
                    pattern: pattern,
                    cost:    numPressed,
                })
            }
        }
    }
    
    return out
}

func generateCombinations(n, k int) [][]int {
    if k == 0 {
        return [][]int{{}}
    }
    if k > n {
        return [][]int{}
    }
    
    var result [][]int
    var helper func(start int, current []int)
    
    helper = func(start int, current []int) {
        if len(current) == k {
            combo := make([]int, k)
            copy(combo, current)
            result = append(result, combo)
            return
        }
        
        for i := start; i < n; i++ {
            helper(i+1, append(current, i))
        }
    }
    
    helper(0, []int{})
    return result
}

func solveSingle(coeffs [][]int, goal []int) int {
    patternCosts := patterns(coeffs)
    cache := make(map[string]int)
    
    var solve func([]int) int
    solve = func(currentGoal []int) int {
        allZero := true
        for _, v := range currentGoal {
            if v != 0 {
                allZero = false
                break
            }
        }
        if allZero {
            return 0
        }
        
        key := libs.IntSliceToStr(currentGoal, ",")
        if val, ok := cache[key]; ok {
            return val
        }
        
        answer := 1000000
        
        parity := make([]int, len(currentGoal))
        for i := 0; i < len(currentGoal); i++ {
            parity[i] = currentGoal[i] % 2
        }
        parityKey := libs.IntSliceToStr(parity, ",")
        
        for _, p := range patternCosts[parityKey] {
            fits := true
            for i := 0; i < len(p.pattern); i++ {
                if p.pattern[i] > currentGoal[i] {
                    fits = false
                    break
                }
            }
            
            if fits {
                newGoal := make([]int, len(currentGoal))
                for i := 0; i < len(currentGoal); i++ {
                    newGoal[i] = (currentGoal[i] - p.pattern[i]) / 2
                }
                
                subResult := solve(newGoal)
                if p.cost + 2*subResult < answer {
                    answer = p.cost + 2*subResult
                }
            }
        }
        
        cache[key] = answer
        return answer
    }
    
    return solve(goal)
}

func intSliceEquals(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}