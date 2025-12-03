package main

import (
    "fmt"
    "strconv"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    banks := libs.FileToSlice("2025/day03/input.txt", "\n")
    for _, bank := range banks {
        largestCharTenCol, pos := findLargestCharacter(bank[:len(bank)-1])
        largestCharUnitCol, _ := findLargestCharacter(bank[pos+1:])
        numStr := string(largestCharTenCol)+string(largestCharUnitCol)
        num, _ := strconv.Atoi(numStr)
        ans += num
    }
    fmt.Println("🎄 The answer to part 1 for day 03 is:", ans, "🎄")
}

func findLargestCharacter(bank string) (rune, int) {
    largestChar := '0'
    largestCharIndex := 0
    for i:=0; i<len(bank); i++ {
        char := rune(bank[i])
        if char > largestChar {
            largestChar = char
            largestCharIndex = i
        }
    }
    return largestChar, largestCharIndex
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 03 is:", ans, "🎄")
}