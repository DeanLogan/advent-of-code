package main

import (
	"fmt"
	// "regexp"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0

    sections := libs.FileToSlice("2024/day05/input.txt", "\n\n")
    pageOrderingRules := strings.Split(sections[0], "\n")
    pages := strings.Split(sections[1], "\n")

    for _, page := range pages {
        pageNums := strings.Split(page, ",")
        if checkOrdering(pageNums, pageOrderingRules) {
            middleIndex := pageNums[(len(pageNums) / 2)]
            num, _ := strconv.Atoi(middleIndex)
            ans += num
        }
    }

    fmt.Println("The answer to part 1 for day 05 is:", ans)
}

func checkOrdering(page []string, rules []string) bool {
    for _, rule := range rules {
        parts := strings.Split(rule, "|")
        first := parts[0]
        second := parts[1]

        firstIndex := libs.SearchForStrInSlice(first, page)
        secondIndex := libs.SearchForStrInSlice(second, page)

        if firstIndex != -1 && secondIndex != -1 {
            if firstIndex > secondIndex {
                return false
            }
        }
    }
    return true
}

func part2(){
    ans := 0

    sections := libs.FileToSlice("2024/day05/input.txt", "\n\n")
    pageOrderingRules := strings.Split(sections[0], "\n")
    pages := strings.Split(sections[1], "\n")

    for _, page := range pages {
        pageNums := strings.Split(page, ",")
        if !checkOrdering(pageNums, pageOrderingRules) {
            reorderPage(pageNums, pageOrderingRules)
            middleIndex := pageNums[(len(pageNums) / 2)]
            num, _ := strconv.Atoi(middleIndex)
            ans += num
        }
    }

    fmt.Println("The answer to part 2 for day 05 is:", ans)
}

func reorderPage(page []string, rules []string) []string {
    problem := false
    for _, rule := range rules {
        parts := strings.Split(rule, "|")
        first := parts[0]
        second := parts[1]

        firstIndex := libs.SearchForStrInSlice(first, page)
        secondIndex := libs.SearchForStrInSlice(second, page)

        if firstIndex != -1 && secondIndex != -1 {
            if firstIndex > secondIndex {
                temp := page[firstIndex]
                page[firstIndex] = page[secondIndex]
                page[secondIndex] = temp
                problem = true
            }
        }
    }

    if problem {
        return reorderPage(page, rules)
    }
    return page
}