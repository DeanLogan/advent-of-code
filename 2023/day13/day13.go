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
    lines := strings.ReplaceAll(string(libs.AllFileContent("2023/day13/input.txt")), "\r\n", "\n")
    patterns := strings.Split(string(lines), "\n\n")
    for _, pattern := range patterns {
        if v, found := findHorizontalSymmetry(libs.TransposeString(pattern)); found {
            ans += v + 1
        }
        if v, found := findHorizontalSymmetry(pattern); found {
            ans += (v + 1) * 100
        }
    }
    fmt.Println("The answer to part 1 for day 13 is:", ans)
}

func findHorizontalSymmetry(pattern string) (int, bool) {
    lines := strings.Split(pattern, "\n")
    NextLine:
    for line := 0; line < len(lines)-1; line++ {
        for delta := 0; ; delta++ {
            up, down := line-delta, line+delta+1
            if up < 0 || down >= len(lines) {
                return line, true
            }

            if lines[up] != lines[down] {
                continue NextLine
            }
        }
    }

    return 0, false
}

func part2(){
    ans := 0
    lines := strings.ReplaceAll(string(libs.AllFileContent("2023/day13/input.txt")), "\r\n", "\n")
    patterns := strings.Split(string(lines), "\n\n")
    for _, pattern := range patterns {
        if v, found := findHorizontalSymmetryWithSmudges(libs.TransposeString(pattern), 1); found {
            ans += v + 1
        }
        if v, found := findHorizontalSymmetryWithSmudges(pattern, 1); found {
            ans += (v + 1) * 100
        }
    }
    fmt.Println("The answer to part 2 for day 13 is:", ans)
}

func findHorizontalSymmetryWithSmudges(pattern string, smudges int) (int, bool) {
    lines := strings.Split(pattern, "\n")
    NextLine:
    for line := 0; line < len(lines)-1; line++ {
        changes := smudges
        for delta := 0; ; delta++ {
            up, down := line-delta, line+delta+1
            if up < 0 || down >= len(lines) {
                if changes == 0 {
                    return line, true
                }
                continue NextLine
            }

            diff := levenshtein(lines[up], lines[down])
            if diff > changes {
                continue NextLine
            }

            changes -= diff
        }
    }

    return 0, false
}

func levenshtein(str1, str2 string) int {
	diff := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			diff++
		}
	}
	return diff
}