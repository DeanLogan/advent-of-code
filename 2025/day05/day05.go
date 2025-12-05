package main

import (
    "fmt"
    "sort"
    "strconv"
    "strings"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type Range struct {
    start int
    end int
}

func part1(){
    ans := 0

    fileSections := libs.FileToSlice("2025/day05/input.txt", "\n\n")
    rangesStr := strings.Split(fileSections[0], "\n")
    ranges := createRangeArray(rangesStr)

    sort.Slice(ranges, func(i, j int) bool {
        return ranges[i].start < ranges[j].start
    })

    ranges = mergeRanges(ranges)
    
    ids := libs.StrToIntSlice(fileSections[1], "\n")
    for _, id := range ids {
        if checkIdInRanges(id, ranges) {
            ans++
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 05 is:", ans, "🎄")
}

func createRangeArray(rangesStr []string) []Range {
    ranges := []Range{}
    for _, rangeStr := range rangesStr {
        startstr, endStr := libs.SplitAtChar(rangeStr, '-')
        start, _ := strconv.Atoi(startstr)
        end, _ := strconv.Atoi(endStr)
        ranges = append(ranges, Range{start, end})
    }
    return ranges
}

func mergeRanges(ranges []Range) []Range {
    newRanges := []Range{}
    for i := range ranges {
        if i<len(ranges)-1 && ranges[i].end >= ranges[i+1].start {
            ranges[i+1] = Range{ranges[i].start, libs.Max(ranges[i].end, ranges[i+1].end)}
        } else {
            newRanges = append(newRanges, ranges[i])
        }
    }
    return newRanges
}

func checkIdInRanges(id int, ranges []Range) bool {
    for _, idRange := range ranges {
        if id >= idRange.start && id <= idRange.end {
            return true
        }
    }
    return false
}

func part2(){
    ans := 0

    fileSections := libs.FileToSlice("2025/day05/input.txt", "\n\n")
    rangesStr := strings.Split(fileSections[0], "\n")
    ranges := createRangeArray(rangesStr)

    sort.Slice(ranges, func(i, j int) bool {
        return ranges[i].start < ranges[j].start
    })

    ranges = mergeRanges(ranges)

    for _, idRange := range ranges {
        ans += idRange.end - idRange.start + 1 // + 1 for the number left over in the range calculation
    }

    fmt.Println("🎄 The answer to part 2 for day 05 is:", ans, "🎄")
}