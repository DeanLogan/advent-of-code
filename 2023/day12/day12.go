package main

import (
	"fmt"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    springs := libs.FileToSlice("2023/day12/input.txt", "\n")
    for _, spring := range springs {
        springConditionRecord, contiguousDamagedSprings := libs.SplitAtStr(spring, " ")
        ans += calcArrangementsPartOne(springConditionRecord, libs.StrToIntSlice(contiguousDamagedSprings, ","))
    }
    fmt.Println("🎄 The answer to part 1 for day 12 is:", ans, "🎄")
}

func calcArrangementsPartOne(springConditionRecord string, contiguousDamagedSprings []int) int {
    springConditionRecord = strings.Trim(springConditionRecord, ".")

    if len(contiguousDamagedSprings) == 0 {
        if springConditionRecord == "" || !strings.Contains(springConditionRecord, "#") {
            return 1
        }
        return 0
    }
    if springConditionRecord == "" {
        return 0
    }

    sum := 0
    if springConditionRecord[0] == '?' {
        sum += calcArrangementsPartOne(springConditionRecord[1:], contiguousDamagedSprings)
        springConditionRecord = "#" + springConditionRecord[1:]
    }

    if len(springConditionRecord) < contiguousDamagedSprings[0] || strings.ContainsRune(springConditionRecord[:contiguousDamagedSprings[0]], '.') {
        return sum
    }
    if len(springConditionRecord) > contiguousDamagedSprings[0] && springConditionRecord[contiguousDamagedSprings[0]] == '?' {
        springConditionRecord = springConditionRecord[:contiguousDamagedSprings[0]] + "." + springConditionRecord[contiguousDamagedSprings[0]+1:]
    } else if len(springConditionRecord) > contiguousDamagedSprings[0] && springConditionRecord[contiguousDamagedSprings[0]] == '#' {
        return sum
    }

    return sum + calcArrangementsPartOne(springConditionRecord[contiguousDamagedSprings[0]:], contiguousDamagedSprings[1:])
}

func part2(){
    ans := 0
    springs := libs.FileToSlice("2023/day12/input.txt", "\n")
    for _, spring := range springs {
        springConditionRecord, contiguousDamagedSprings := libs.SplitAtStr(spring, " ")
        springConditionRecord, contiguousDamagedSprings = unfoldSprings(springConditionRecord, contiguousDamagedSprings)

        ans += calcArrangementsPartTwo(springConditionRecord, libs.StrToIntSlice(contiguousDamagedSprings, ","))
    }
    fmt.Println("🎄 The answer to part 2 for day 12 is:", ans, "🎄")
}

func unfoldSprings(springs string, arrangements string) (string, string) {
    newSpring  := springs
    newArrangements := arrangements
    for i := 0; i< 4; i++ {
        newSpring += "?"+springs
        newArrangements += ","+arrangements
    }
    return newSpring, newArrangements
}

func clearMap(m map[[3]int]int) {
    for key := range m {
        delete(m, key)
    }
}

func calcArrangementsPartTwo(springConditionRecord string, contiguousDamagedSprings []int) int {
    position := 0
    currentStates := map[[3]int]int{{0, 0, 0}: 1}
    nextStates := map[[3]int]int{}

    for springIndex := 0; springIndex < len(springConditionRecord); springIndex++ {
        springCondition := springConditionRecord[springIndex]
        for state, count := range currentStates {
            damagedSpringIndex, damagedSpringCount, expectWorking := state[0], state[1], state[2]
            switch {
            case (springCondition == '#' || springCondition == '?') && damagedSpringIndex < len(contiguousDamagedSprings) && expectWorking == 0:
                if springCondition == '?' && damagedSpringCount == 0 {
                    nextStates[[3]int{damagedSpringIndex, damagedSpringCount, expectWorking}] += count
                }
                damagedSpringCount++
                if damagedSpringCount == contiguousDamagedSprings[damagedSpringIndex] {
                    damagedSpringIndex, damagedSpringCount, expectWorking = damagedSpringIndex+1, 0, 1
                }
                nextStates[[3]int{damagedSpringIndex, damagedSpringCount, expectWorking}] += count
            case (springCondition == '.' || springCondition == '?') && damagedSpringCount == 0:
                expectWorking = 0
                nextStates[[3]int{damagedSpringIndex, damagedSpringCount, expectWorking}] += count
            }
        }
        currentStates, nextStates = nextStates, currentStates
        for key := range nextStates {
            delete(nextStates, key)
        }
    }

    for state, value := range currentStates {
        if state[0] == len(contiguousDamagedSprings) {
            position += value
        }
    }
    return position
}