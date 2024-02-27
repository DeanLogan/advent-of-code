package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

// TODO: come back to this problem at the end and try to make part 2 faster

func main() {
	partOne()
	partTwo()
}

func partOne() {	
	maps := libs.FileToSlice("day5/input.txt", "\n\n")

	firstLine := maps[0]
	seeds := strings.Split(firstLine, " ")
	seeds = seeds[1:]
	ans := int(math.MaxInt64)
	for _, seedStr := range seeds{
		checkFor, err := strconv.Atoi(seedStr)
		if err != nil {
			log.Fatal(err, "Failed to convert string to int")
			return 
		}
		checkFor = gettingSeedLocation(checkFor, maps)
		if checkFor < ans {
			ans = checkFor
		}
	} 
	fmt.Println("The answer to part 1 for day 5 is: ", ans)
}

func calcMapping(mapping string, checkFor int) int {
	lines := strings.Split(mapping, "\n")
	for i:=1; i < len(lines); i++ {
		// convert each number to an int
		intSlice := []int{}
		for _, num := range strings.Split(lines[i], " ") {
			intNum, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err, "Failed to convert string to int")
				return -1
			}
			intSlice = append(intSlice, intNum)
		}
		if checkFor >= intSlice[1] && checkFor <= intSlice[1] + (intSlice[2]-1) {
			//figures out where checkFor falls in the above range and adds 0 index intSlice to match mapping
			return (checkFor - intSlice[1]) + intSlice[0] 
		}
	}
	return checkFor
}

func gettingSeedLocation(seed int, maps []string) int {
    results := make(chan int, len(maps)-1)
    for i := 1; i < len(maps); i++ {
        go func(i int, seed int) {
            mapped := calcMapping(maps[i], seed)
            results <- mapped
        }(i, seed)
        seed = <-results
    }

    return seed
}

func partTwo() {
	seeds, maps := parse()

	for loc := 0; ; loc++ {
		x := loc
		for m := len(maps) - 1; m >= 0; m-- {
			x = maps[m].InverseConvert(x)
		}

		for s := 0; s < len(seeds); s += 2 {
			if x >= seeds[s] && x <= seeds[s]+seeds[s+1] {
				fmt.Println("The answer to part 2 for day 5 is: ",loc)
				return
			}
		}
	}
}

func (m Map) InverseConvert(to int) int {
	for _, t := range m {
		if to >= t.Destination && to < t.Destination+t.Length {
			return t.Source + (to - t.Destination)
		}
	}
	return to
}

type Map []Transform

type Transform struct {
	Destination, Source, Length int
}

func parse() ([]int, [7]Map) {
	scanner := libs.GetScannerForFile("day5/input.txt")

	seeds := []int{}
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	for i := 1; i < len(fields); i++ {
		num, _ := strconv.Atoi(fields[i])
		seeds = append(seeds, num)
	}

	return seeds, parseMaps(scanner)
}

func parseMaps(scanner *bufio.Scanner) [7]Map {
	maps := [7]Map{}
	t := -1
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			t++
			continue
		}

		fields := strings.Fields(line)

		dest, _ := strconv.Atoi(fields[0])
		source, _ := strconv.Atoi(fields[1])
		length, _ := strconv.Atoi(fields[2])

		maps[t] = append(maps[t], Transform{
			Destination: dest,
			Source:      source,
			Length:      length,
		})
	}
	return maps
}

