package main

import (
	"fmt"
	"strings"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

type Node struct {
    left  string
    right string
}

var puzzleMap = map[string]Node{}

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    lines := string(libs.AllFileContent("day08/input.txt"))
    directions, tree := libs.SplitAtStr(lines, "\n\n")
    tree = tree[1:] // remove first line as it is just an empty space
    treeSlice := strings.Split(tree, "\n")

    // adds the left and right value to the map, where the key is the "root" node
    for _, line := range treeSlice {
        puzzleMap[line[:3]] = Node{left: line[7:10], right: line[12:15]}
    }

    ans = findZZZ(directions, "AAA", 0)

    fmt.Println("The answer to part 1 for day 8 is:", ans)
}

func findZZZ(directions string, startingPoint string, stepsNum int) int {
    direction := rune(directions[0])
    nextValue := findNextValue(direction, startingPoint)
    if startingPoint == "ZZZ" {
        return stepsNum
    }
    stepsNum++
    directions = directions[1:] + string(direction)
    return findZZZ(directions, nextValue, stepsNum)
}

func findNextValue(direction rune, nodeParent string) string {
    if direction == 'L'{
        return puzzleMap[nodeParent].left
    } else { // assume the only other option is R
        return puzzleMap[nodeParent].right
    }
}

func part2(){
    lines := string(libs.AllFileContent("day08/input.txt"))
    directions, tree := libs.SplitAtStr(lines, "\n\n")
    tree = tree[1:] // remove first line as it is just an empty space
    treeSlice := strings.Split(tree, "\n")
    startingPoints := []string{}

    // adds the left and right value to the map, where the key is the "root" node
    for _, line := range treeSlice {
        puzzleMap[line[:3]] = Node{left: line[7:10], right: line[12:15]}
        if line[2] == 'A'{
            startingPoints = append(startingPoints, line[:3])
        }
    }

    results := []int{}
    for i := 0; i < len(startingPoints); i++{
        _, _, newStepsTaken := findEndInZ(directions, startingPoints[i], 0)
        results = append(results, newStepsTaken)
    }

    ans := results[0]
	for i := 1; i < len(results); i++ {
		ans = libs.Lcm(ans, results[i])
	}

    fmt.Println("The answer to part 2 for day 8 is:", ans)
}

// On reddit some people were frustracted (https://www.reddit.com/r/adventofcode/comments/18dh4p8/2023_day_8_part_2_im_a_bit_frustrated/) that part 2 could only 
// be solved by abusing a rule where the LCM of inital steps taken gives the correct answer, however this might not always be a case with a different input. 
// So as I thought it would be interesting to give it a go. I created a way that would work for other inputs (however it is extremely slow). It follows how the 
// problem is phrased closer, meaning that it will perform each of the searching steps until all the steps completed to find Z for each of the starting 
// points are the same.
func part2AltWay(){
    lines := string(libs.AllFileContent("day08/input.txt"))
    directions, tree := libs.SplitAtStr(lines, "\n\n")
    tree = tree[1:] // remove first line as it is just an empty space
    treeSlice := strings.Split(tree, "\n")
    startingPoints := []string{}

    // adds the left and right value to the map, where the key is the "root" node
    for _, line := range treeSlice {
        puzzleMap[line[:3]] = Node{left: line[7:10], right: line[12:15]}
        if line[2] == 'A'{
            startingPoints = append(startingPoints, line[:3])
        }
    }

    directionsSlice := make([]string, len(startingPoints))
    for i := range directionsSlice {
        directionsSlice[i] = directions
    }

    startingSteps := make([]int, len(startingPoints))

    ans := findEndForAllStartingPoints(directionsSlice, startingPoints, startingSteps)

    fmt.Println("The answer to part 2 for day 8 is:", ans)
}

func findEndForAllStartingPoints(directions []string, startingPoints []string, steps []int) int {
    results := []int{}

    highestStepNum := libs.MaxOfSlice(steps) 

    // get the results and update the necessary values for the next search iteration
    for i := 0; i < len(startingPoints); i++{
        var stepsTaken int
        if highestStepNum != steps[i] || steps[i] == 0 {
            updatedDirections, nextValue, newStepsTaken := findEndInZ(directions[i], startingPoints[i], steps[i])
            stepsTaken = newStepsTaken // update stepsTaken with the new value
            startingPoints[i] = nextValue // update the new starting point with the Z index found
            directions[i] = updatedDirections // update the new direction string with the current set of directions found
        } else {
            stepsTaken = steps[i]
        }
        results = append(results, stepsTaken)
    }

    allSame := true
    first := results[0]
    for _, result := range results {
        if result != first {
            allSame = false
            break
        }
    }

    if allSame {
        return first
    } else {
        return findEndForAllStartingPoints(directions, startingPoints, results)
    }
}

func findEndInZ(directions string, startingPoint string, stepsNum int) (string, string, int) { // returns each of the params in the same order
    direction := rune(directions[0])
    directions = directions[1:] + string(direction)
    nextValue := findNextValue(direction, startingPoint)
    stepsNum++
    // checks if last character is Z
    if nextValue[2] == 'Z' {
        return directions, nextValue, stepsNum
    } 
    return findEndInZ(directions, nextValue, stepsNum)
}