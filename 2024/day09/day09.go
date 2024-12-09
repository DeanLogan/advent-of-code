package main

import (
	"fmt"
	"strconv"
    "strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1() {
    ans := 0

    diskMap := parseLine(strings.Split(libs.AllFileContentAsString("2024/day09/input.txt"), ""))
    diskMap = adjustDiskMap(diskMap)
    
    for i := range diskMap {
        if diskMap[i] != "." {
            checkSum, _ := strconv.Atoi(diskMap[i])
            ans += i * checkSum
        }
    }

    fmt.Println("The answer to part 1 for day 09 is:", ans)
}

func adjustDiskMap(diskMap []string) []string {
    for i, j := 0, len(diskMap)-1; i <= j; {
        if diskMap[i] == "." && diskMap[j] != "." {
            diskMap[i], diskMap[j] = diskMap[j], diskMap[i]
            i++
            j--
        } else {
            if diskMap[i] != "." {
                i++
            } else if diskMap[j] != "." {
                j--
            } else if diskMap[i] == "." && diskMap[j] == "." {
                j--
            }
        }
    }
    return diskMap
}


func parseLine(line []string) []string {
    diskMap := make([]string, 0)
    for i := range line {
        if i%2 == 0 {
            id := i / 2
            fileSpace, _ := strconv.ParseInt(line[i], 10, 32)
            diskMap = addSpace(diskMap, int(fileSpace), fmt.Sprintf("%d", id))
        } else {
            freeSpace, _ := strconv.ParseInt(line[i], 10, 32)
            diskMap = addSpace(diskMap, int(freeSpace), ".")
        }
    }
    return diskMap
}

func part2() {
    ans := 0

    fileSpaces := libs.StrToIntSlice(libs.AllFileContentAsString("2024/day09/input.txt"), "")
    diskMap, files, spaces := parseFileSpaces(fileSpaces)
    diskMap = adjustFilesInSpaces(diskMap, files, spaces)

    for i := range diskMap {
        if diskMap[i] != "." {
            checkSum, _ := strconv.Atoi(diskMap[i])
            ans += i * checkSum
        }
    }

    fmt.Println("The answer to part 2 for day 09 is:", ans)
}

func addSpace(diskMap []string, numOfChars int, chars string) []string {
    for j := 0; j < numOfChars; j++ {
        diskMap = append(diskMap, chars)
    }
    return diskMap
}

func parseFileSpaces(fileSpaces []int) ([]string, [][]int, [][]int) {
    diskMap := make([]string, 0)
    files := make([][]int, 0)
    spaces := make([][]int, 0)
    for i, fileSpace := range fileSpaces {
        if i%2 == 0 {
            id := i / 2
            files = append(files, []int{len(diskMap), len(diskMap) - 1 + fileSpace})
            diskMap = addSpace(diskMap, fileSpace, fmt.Sprintf("%d", id))
        } else {
            spaces = append(spaces, []int{len(diskMap), len(diskMap) - 1 + fileSpace})
            diskMap = addSpace(diskMap, fileSpace, ".")
        }
    }
    return diskMap, files, spaces
}

func adjustFilesInSpaces(diskMap []string, files [][]int, spaces [][]int) []string {
    for j := len(files) - 1; j >= 0; j-- {
        for i := 0; i < len(spaces); i++ {
            space := spaces[i]
            if space[1] < files[j][1] && space[1]-space[0] >= files[j][1]-files[j][0] {
                for k := space[0]; k <= space[0]+files[j][1]-files[j][0]; k++ {
                    diskMap[k] = fmt.Sprintf("%d", j)
                }
                for k := files[j][0]; k <= files[j][1]; k++ {
                    diskMap[k] = "."
                }
                space[0] = space[0] + files[j][1] - files[j][0] + 1
                break
            }
        }
    }
    return diskMap
}