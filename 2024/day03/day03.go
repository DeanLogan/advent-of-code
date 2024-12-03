package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1() {
    ans := 0
    memoryChunks := libs.FileToSlice("2024/day03/input.txt", "\n")

    for _, chunk := range memoryChunks {
        ans += checkChunk(chunk)
    }

    fmt.Println("The answer to part 1 for day 03 is:", ans)
}

func checkChunk(chunk string) int {
    result := 0
    for i:=0; i<len(chunk)-7; i++ {
        // looks for the char m (comparing bytes)
        if chunk[i] == 109 {
            result, i = checkForMul(chunk, i, result)
        }
    }
    return result
}

func checkForMul(chunk string, startPoint int, result int) (int, int) {
    for i:=8; i<13; i++ {
        checkUntil := startPoint+i
        if checkUntil > len(chunk) {
            return result, startPoint
        }

        operation := chunk[startPoint:checkUntil]
        if isValidMul(operation) {
            num1, num2 := extractNumbers(operation)
            result += num1 * num2
            startPoint = checkUntil-1 // move i to end of the operation to skip checking unneccassary chars
        }
    }
    return result, startPoint
}

func extractNumbers(s string) (int, int) {
    re := regexp.MustCompile(`^mul\((\d{1,3}),(\d{1,3})\)$`)
    matches := re.FindStringSubmatch(s)
    if len(matches) != 3 {
        return 0, 0
    }
    num1, _ := strconv.Atoi(matches[1])
    num2, _ := strconv.Atoi(matches[2])
    return num1, num2
}

func isValidMul(s string) bool {
    re := regexp.MustCompile(`^mul\(\d{1,3},\d{1,3}\)$`)
    return re.MatchString(s)
}

func part2(){
    ans := 0

    memoryChunks := libs.FileToSlice("2024/day03/input.txt", "\n")
    doMul := true
    chunkAns := 0
    for _, chunk := range memoryChunks {
        chunkAns, doMul = checkChunkWithInstructions(chunk, doMul)
        ans += chunkAns
    }

    fmt.Println("The answer to part 2 for day 03 is:", ans)
}

func checkChunkWithInstructions(chunk string, doMul bool) (int, bool) {
    result := 0
    for i:=0; i<len(chunk)-7; i++ {
        // looks for the char d (comparing bytes)
        if chunk[i] == 100 {
            if chunk[i:i+7] == "don't()" {
                doMul = false
                i += 7
            } else if chunk[i:i+4] == "do()" {
                doMul = true
                i += 4
            }
        }

        if doMul && chunk[i] == 109 {
            result, i = checkForMul(chunk, i, result)
        }
    }
    return result, doMul
}