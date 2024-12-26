package main

import (
	"fmt"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    fmt.Println("There is no part 2, all challanges have been completed :) 🎄🎅🐧❄️⛄")
}

func part1(){
    ans := 0

    locksAndKeys := libs.FileToSlice("2024/day25/input.txt", "\n\n")
    locks, keys := seperateLocksAndKeys(locksAndKeys)
    fmt.Println(locks)
    fmt.Println(keys)
    for _, lock := range locks {
        for _, key := range keys {
            if keyFitsLock(lock, key) {
                ans++
            }
        }
    }
    fmt.Println("🎄 The answer to part 1 for day 25 is:", ans, "🎄")
}

func seperateLocksAndKeys(locksAndKeys []string) ([][]int, [][]int) {
    var locks, keys [][]int 

    for _, lockOrKey := range locksAndKeys {
        if lockOrKey[:5] == "#####" {
            locks = append(locks, calcHeightMap(lockOrKey))
        } else {
            keys = append(keys, calcHeightMap(lockOrKey))
        }
    }

    return locks, keys
}

func calcHeightMap(lockOrKey string) []int {
    lockOrKey = libs.TransposeString(lockOrKey)
    lockOrKey = strings.ReplaceAll(lockOrKey, ".", "")
    lockOrKeySlice := strings.Split(lockOrKey, "\n")

    heights := make([]int, len(lockOrKeySlice))
    for i, line := range lockOrKeySlice {
        heights[i] = len(line)-1
    }
    return heights
}

func keyFitsLock(lock []int, key []int) bool {
    for i:=0; i<len(lock); i++ {
        if lock[i] + key[i] > 5 {
            return false
        }
    } 
    return true
}
