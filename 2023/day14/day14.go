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
    lines := libs.FileToSlice("2023/day14/input.txt", "\n")
    lines = libs.TransposeStringSlice(lines)
    for _, line := range lines{
        rolledLine := libs.ReverseString((rollLine(line)))
        for i, char := range rolledLine {
            if char == 'O' {
                ans += i + 1
            }
        }
    }
    fmt.Println("The answer to part 1 for day 14 is:", ans)
}

func rollLine(line string) string {
    splitLine := strings.Split(line, "#")
    line = ""
    for i, partOfLine := range splitLine {
        if i != 0 {
            line += "#"
        }
        line +=  strings.ReplaceAll(partOfLine, ".", "") + strings.ReplaceAll(partOfLine, "O", "")
    }
    return line
}

func part2(){
    s := libs.FileToSlice("2023/day14/input.txt", "\n")

	seen := map[string]int{}

	var cycleStart, cycleEnd int
	for n := 1; ; n++ {
		cycle(s)

		if lastSeen, ok := seen[strings.Join(s, "\n")]; ok {
			cycleStart = lastSeen
			cycleEnd = n
			break
		}
		seen[strings.Join(s, "\n")] = n
	}

	remaining := (1e9 - cycleStart) % (cycleEnd - cycleStart)
	for n := 0; n < remaining; n++ {
		cycle(s)
	}

	ans := 0
	for j := range s {
		for i := range s[j] {
			if s[j][i] == uint8(Round) {
				ans += len(s) - j
			}
		}
	}

    fmt.Println("The answer to part 2 for day 14 is:", ans)
}

const (
	None  rock = '.'
	Round rock = 'O'
	Cube  rock = '#'
)

type rock uint8

func cycle(s []string) {
	x2 := len(s[0])
	y2 := len(s)

	for i := 0; i < x2; i++ {
		topFree := 0
		for j := 0; j < y2; j++ {
			switch s[j][i] {
			case uint8(None):
				continue
			case uint8(Cube):
				topFree = j + 1
			case uint8(Round):
				s[j] = s[j][:i] + string(None) + s[j][i+1:]
				s[topFree] = s[topFree][:i] + string(Round) + s[topFree][i+1:]
				topFree++
			}
		}
	}

	for j := 0; j < y2; j++ {
		leftFree := 0
		for i := 0; i < x2; i++ {
			switch s[j][i] {
			case uint8(None):
				continue
			case uint8(Cube):
				leftFree = i + 1
			case uint8(Round):
				s[j] = s[j][:i] + string(None) + s[j][i+1:]
				s[j] = s[j][:leftFree] + string(Round) + s[j][leftFree+1:]
				leftFree++
			}
		}
	}

	for i := x2 - 1; i >= 0; i-- {
		bottomFree := y2 - 1
		for j := y2 - 1; j >= 0; j-- {
			switch s[j][i] {
			case uint8(None):
				continue
			case uint8(Cube):
				bottomFree = j - 1
			case uint8(Round):
				s[j] = s[j][:i] + string(None) + s[j][i+1:]
				s[bottomFree] = s[bottomFree][:i] + string(Round) + s[bottomFree][i+1:]
				bottomFree--
			}
		}
	}

	for j := y2 - 1; j >= 0; j-- {
		rightFree := x2 - 1
		for i := x2 - 1; i >= 0; i-- {
			switch s[j][i] {
			case uint8(None):
				continue
			case uint8(Cube):
				rightFree = i - 1
			case uint8(Round):
				s[j] = s[j][:i] + string(None) + s[j][i+1:]
				s[j] = s[j][:rightFree] + string(Round) + s[j][rightFree+1:]
				rightFree--
			}
		}
	}
}