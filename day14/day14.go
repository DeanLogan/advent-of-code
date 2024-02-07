package main

import (
	"fmt"
	"strings"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    lines := libs.FileToSlice("day14/input.txt", "\n")
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
    ans := 0
    fmt.Println("The answer to part 2 for day 14 is:", ans)
}