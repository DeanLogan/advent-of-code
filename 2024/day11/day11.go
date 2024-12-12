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

    stones := libs.StrToIntSlice(libs.AllFileContentAsString("2024/day11/input.txt"), " ")

    for i:=0; i<25; i++ {
        stones = blink(stones)
    }

    ans = len(stones)

    fmt.Println("The answer to part 1 for day 11 is:", ans)
}

func blink(stones []int) []int {
    newStones := []int{}
    for _, stone := range stones {
        if stone == 0 {
            newStones = append(newStones, 1)
        } else {
            numStr := strconv.Itoa(stone)
            digitLen := len(numStr)
            if digitLen % 2 == 0 {
                halfDigitLen := digitLen / 2
                leftStone, _ := strconv.Atoi(numStr[:halfDigitLen])
                rightStone, _ := strconv.Atoi(numStr[halfDigitLen:])
                newStones = append(newStones, leftStone)
                newStones = append(newStones, rightStone)
            } else {
                newStones = append(newStones, stone*2024)
            }
        }
    }
    return newStones
}

func part2(){
    ans := 0
    fmt.Println("The answer to part 2 for day 11 is:", ans)
}