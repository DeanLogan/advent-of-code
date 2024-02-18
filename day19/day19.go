package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    file := string(libs.AllFileContent("day19/input.txt"))
    workflowsStr, partsStr := libs.SplitAtStr(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n\n")
    partsStr = partsStr[1:]
    parts := strings.Split(partsStr, "\n")
    workflows := strings.Split(workflowsStr, "\n")
    groupings := generateGroupings(workflows)

    for _, part := range parts {
        partMap := createPartMap(part)
        if sortPart(partMap, "in", groupings) == "A" {
            for _, val := range partMap {
                ans += val
            }
        }
    }

    fmt.Println("The answer to part 1 for day 19 is:", ans)
}

func createPartMap(part string) map[string]int {
    part = part[1:len(part)-1] // remove '{' and '}' from str
    ratings := strings.Split(part, ",")
    partMap := make(map[string]int)
    for _, rating := range ratings {
        partPiece, valueStr := libs.SplitAtChar(rating, '=')
        value, _ := strconv.Atoi(valueStr)
        partMap[partPiece] = value
    }
    return partMap
}

func generateGroupings(workflows []string) map[string][]string{
    groupings := make(map[string][]string)
    for _, workflow := range workflows{
        group, conditions := libs.SplitAtChar(workflow, '{')
        groupings[group] = strings.Split(conditions[:len(conditions)-1], ",") 
    } 
    return groupings
}

func sortPart(part map[string]int, group string, groupings map[string][]string) string {
    conditions := groupings[group]
    finalVal := conditions[len(conditions)-1]
    for _, condition := range conditions[:len(conditions)-1] {
        variable, conditional := libs.SplitAtChar(condition, '<')
        // if the 2 sides of the condition are equal then < is not in the string so > must be in it 
        if variable == conditional {
            variable, conditional = libs.SplitAtChar(condition, '>')
            numStr, goToNext := libs.SplitAtChar(conditional, ':')
            num, _ := strconv.Atoi(numStr)
            if part[variable] > num {
                if goToNext == "A" || goToNext == "R"{
                    return goToNext
                }
                return sortPart(part, goToNext, groupings)
            }
        } else {
            numStr, goToNext := libs.SplitAtChar(conditional, ':')
            num, _ := strconv.Atoi(numStr)
            if part[variable] < num {
                if goToNext == "A" || goToNext == "R"{
                    return goToNext
                }
                return sortPart(part, goToNext, groupings)
            }
        }
    }
    if finalVal == "A" || finalVal == "R"{
        return finalVal
    }
    return sortPart(part, finalVal, groupings)
}

func part2(){
	ans := 0
	file := string(libs.AllFileContent("day19/input.txt"))
    file = strings.ReplaceAll(file, "\r\n", "\n")
    workflowsStr, _ := libs.SplitAtStr(file, "\n\n")
    workflowsLines := strings.Split(workflowsStr, "\n")
    workflows := readWorkflows(workflowsLines)

    ans = processParts(workflows)
    fmt.Println("The answer to part 2 for day 19 is:",ans)
}

var partMapping = map[byte]int{
    'x': 0,
    'm': 1,
    'a': 2,
    's': 3,
}

type workflowArm struct {
	s, n int
	o    byte
	t    string
}

func readWorkflows(workflowsLines []string) map[string]([]workflowArm) {

    var workflows = make(map[string]([]workflowArm))

    for _, line := range workflowsLines {
        if len(line) == 0 {
            continue
        }
        n, line, _ := strings.Cut(line, "{")
        operations := strings.Split(strings.Trim(line, "{}"), ",")
        workflow := mapSlices(operations[:len(operations)-1], func(line string) workflowArm {
            s, o, r := line[0], line[1], line[2:]
            n, t, _ := strings.Cut(r, ":")
            return workflowArm{partMapping[s], stringToInt(n), o, t}
        })
        workflow = append(workflow, workflowArm{0, 0, '=', operations[len(operations)-1]})
        workflows[n] = workflow
    }
    return workflows
}


func processParts(workflows map[string]([]workflowArm)) int {
    type partRange struct {
        rangeArray  [4][2]int
        currentWorkflow  string
        currentWorkflowIndex int
    }
    var parts, nextParts []partRange
    parts = []partRange{{[4][2]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}, "in", 0}}
    sum := 0
    for len(parts) > 0 {
        for _, part := range parts {
            switch part.currentWorkflow {
            case "A":
                sum += part.rangeArray[0][1] * part.rangeArray[1][1] * part.rangeArray[2][1] * part.rangeArray[3][1]
                continue
            case "R":
                continue
            }
            workflow := workflows[part.currentWorkflow][part.currentWorkflowIndex]
            var successfulRange, failedRange [][2]int
            switch workflow.o {
            case '=':
                successfulRange = [][2]int{part.rangeArray[workflow.s]}
            case '<':
                failedRange, successfulRange, _ = intersectIntervals(part.rangeArray[workflow.s], [2]int{1, workflow.n - 1})
            case '>':
                failedRange, successfulRange, _ = intersectIntervals(part.rangeArray[workflow.s], [2]int{workflow.n + 1, 4000 - workflow.n})
            }
            part.currentWorkflowIndex++
            for _, r := range failedRange {
                part.rangeArray[workflow.s] = r
                nextParts = append(nextParts, part)
            }
            part.currentWorkflow, part.currentWorkflowIndex = workflow.t, 0
            for _, r := range successfulRange {
                part.rangeArray[workflow.s] = r
                nextParts = append(nextParts, part)
            }
        }
        parts, nextParts = nextParts, parts[:0]
    }
    return sum
}

func mapSlices[T, U any](inputSlice []T, mapFunc func(T) U) []U {
    outputSlice := make([]U, len(inputSlice))
    for index := range inputSlice {
        outputSlice[index] = mapFunc(inputSlice[index])
    }
    return outputSlice
}

func stringToInt(s string) int {
    intValue, _ := strconv.Atoi(s)
    return intValue
}

func intersectIntervals(intervalA, intervalB [2]int) (leftoverA, intersection, leftoverB [][2]int) {
    startA, endA, startB, endB := intervalA[0], intervalA[0]+intervalA[1], intervalB[0], intervalB[0]+intervalB[1]
    if (endA < startB) || (endB < startA) {
        leftoverA = [][2]int{intervalA}
        leftoverB = [][2]int{intervalB}
        return
    }
    intersectionStart, intersectionEnd := max(startA, startB), min(endA, endB)
    intersection = [][2]int{{intersectionStart, intersectionEnd - intersectionStart}}
    if startA < intersectionStart {
        leftoverA = append(leftoverA, [2]int{startA, intersectionStart - startA})
    } else if startB < intersectionStart {
        leftoverB = append(leftoverB, [2]int{startB, intersectionStart - startB})
    }
    if intersectionEnd < endA {
        leftoverA = append(leftoverA, [2]int{intersectionEnd, endA - intersectionEnd})
    } else if intersectionEnd < endB {
        leftoverB = append(leftoverB, [2]int{intersectionEnd, endB - intersectionEnd})
    }
    return
}