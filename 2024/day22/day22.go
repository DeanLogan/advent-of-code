package main

import (
    "fmt"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0

    secretNums := libs.StrToIntSlice(libs.AllFileContentAsString("2024/day22/input.txt"), "\n")

    for _, secretNum := range secretNums {
        for i:=0; i<2000; i++ {
            secretNum = generateSecretNumer(secretNum)
        }
        ans += secretNum
    }

    fmt.Println("🎄 The answer to part 1 for day 22 is:", ans, "🎄")
}

func generateSecretNumer(secretNum int) int {
    secretNum = prune(mix(secretNum, secretNum * 64))
    secretNum = prune(mix(secretNum, secretNum / 32))
    secretNum = prune(mix(secretNum, secretNum * 2048))

    return secretNum
}

func mix(secretNum int, givenNum int) int {
    return secretNum ^ givenNum
}

func prune(secretNum int) int {
    return secretNum % 16777216
}

func part2(){
    ans := 0
    fmt.Println("🎄 The answer to part 2 for day 22 is:", ans, "🎄")
}