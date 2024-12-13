package main

import (
    "fmt"
    "strconv"
    "sync"

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

    stones := libs.StrToIntSlice(libs.AllFileContentAsString("2024/day11/input.txt"), " ")

    ans = getNumStones(stones, 75)

    fmt.Println("The answer to part 2 for day 11 is:", ans)
}

func getNumStones(stones []int, times int) int {
    wg := sync.WaitGroup{}
    checked := &sync.Map{}
    ch := make(chan int, len(stones))
    for _, stone := range stones {
        wg.Add(1)
        go func(stone int) {
            defer wg.Done()
            ch <- getNumStonesAfter(stone, times, checked)
        }(stone)
    }
    wg.Wait()
    close(ch)
    result := 0
    for value := range ch {
        result += value
    }
    return result
}

func getNumStonesAfter(stone, times int, checked *sync.Map) int {
    result := 0
    if times > 0 {
        if v, ok := checked.Load(fmt.Sprintf("%d:%d", stone, times)); ok {
            result += v.(int)
        } else {

            if stone == 0 {
                result += getNumStonesAfter(1, times-1, checked)
            } else if digits := strconv.Itoa(stone); len(digits)%2 == 0 {
                a, _ := strconv.Atoi(digits[:len(digits)/2])
                b, _ := strconv.Atoi(digits[len(digits)/2:])
                result += getNumStonesAfter(a, times-1, checked)
                result += getNumStonesAfter(b, times-1, checked)
            } else {
                result += getNumStonesAfter(stone*2024, times-1, checked)
            }

            checked.Store(fmt.Sprintf("%d:%d", stone, times), result)
        }
    } else {
        return 1
    }
    return result
}