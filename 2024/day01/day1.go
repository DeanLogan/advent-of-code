package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0

    lines := libs.FileToSlice("2024/day01/input.txt", "\n")
    leftList, rightList, err := getBothLists(lines)
    if err != nil {
        fmt.Println(err)
        return 
    }

    sort.Ints(leftList)
    sort.Ints(rightList)

    for i:=0; i<len(leftList); i++ {
        ans += libs.Abs(leftList[i] - rightList[i])
    }

    fmt.Println("The answer to part 1 for day 1 is:", ans)
}

func getBothLists(lines []string) ([]int, []int, error) {
    leftList, rightList := []int{}, []int{}

    for _, line := range lines {
        strNum1, strNum2 := libs.SplitAtStr(line, "   ")

        num1, err := strconv.Atoi(strNum1)
        if err != nil {
            return leftList, rightList, err
        }
        num2, err := strconv.Atoi(strNum2[2:])
        if err != nil {
            return leftList, rightList, err
        }

        leftList = append(leftList, num1)
        rightList = append(rightList, num2)
    }

    return leftList, rightList, nil
}

func part2(){
    ans := 0

    lines := libs.FileToSlice("2024/day01/input.txt", "\n")
    leftList, rightList, err := getBothLists(lines)
    if err != nil {
        fmt.Println(err)
        return 
    }

    occrrenceMap := createOccurrenceMap(rightList)

    for i:=0; i<len(leftList); i++ {
        ans += (leftList[i] * occrrenceMap[leftList[i]])
    }

    fmt.Println("The answer to part 2 for day 1 is:", ans)
}

func createOccurrenceMap(slice []int) map[int]int {
    occurrenceMap := make(map[int]int)
    for _, element := range slice {
        occurrenceMap[element]++
    }
    return occurrenceMap
}