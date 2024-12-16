package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/DeanLogan/advent-of-code/libs"
)

type Brick struct {
    x1, y1, z1, x2, y2, z2, id int
}

func main() {
    part1()
	part2()
}

func readBricks(filename string) []Brick {
    lines := libs.FileToSlice(filename, "\n")
    var bricks []Brick
    regex := regexp.MustCompile(`\d+`)
    for i, line := range lines {
        matches := regex.FindAllString(line, -1)
        x1, _ := strconv.Atoi(matches[0])
        y1, _ := strconv.Atoi(matches[1])
        z1, _ := strconv.Atoi(matches[2])
        x2, _ := strconv.Atoi(matches[3])
        y2, _ := strconv.Atoi(matches[4])
        z2, _ := strconv.Atoi(matches[5])
        bricks = append(bricks, Brick{x1, y1, z1, x2, y2, z2, i})
    }
    return bricks
}

func processBricks(bricks []Brick) ([]Brick, map[Brick][]Brick, map[Brick][]Brick) {
    sort.Slice(bricks, func(i, j int) bool {
        return bricks[i].z1 < bricks[j].z1
    })

    occupied := make(map[string]Brick)
    var fallen []Brick
    for _, brick := range bricks {
        for brick.z1 > 0 && allPositionsFree(occupied, down(brick)) {
            brick = down(brick)
        }
        for _, pos := range positions(brick) {
            occupied[pos] = brick
        }
        fallen = append(fallen, brick)
    }

    above := make(map[Brick][]Brick)
    below := make(map[Brick][]Brick)
    for _, brick := range fallen {
        inThisBrick := positions(brick)
        for _, pos := range positions(down(brick)) {
            if occupiedBrick, ok := occupied[pos]; ok && !contains(inThisBrick, pos) {
                above[occupiedBrick] = append(above[occupiedBrick], brick)
                below[brick] = append(below[brick], occupiedBrick)
            }
        }
    }

    return fallen, above, below
}

func down(brick Brick) Brick {
    return Brick{brick.x1, brick.y1, brick.z1 - 1, brick.x2, brick.y2, brick.z2 - 1, brick.id}
}

func positions(brick Brick) []string {
    var pos []string
    for x := brick.x1; x <= brick.x2; x++ {
        for y := brick.y1; y <= brick.y2; y++ {
            for z := brick.z1; z <= brick.z2; z++ {
                pos = append(pos, fmt.Sprintf("%d,%d,%d", x, y, z))
            }
        }
    }
    return pos
}

func allPositionsFree(occupied map[string]Brick, brick Brick) bool {
    for _, pos := range positions(brick) {
        if _, ok := occupied[pos]; ok {
            return false
        }
    }
    return true
}

func whatIf(above, below map[Brick][]Brick, disintegrated Brick) int {
    falling := make(map[Brick]bool)
    var falls func(brick Brick)
    falls = func(brick Brick) {
        if falling[brick] {
            return
        }
        falling[brick] = true
        for _, parent := range above[brick] {
            allFalling := true
            for _, child := range below[parent] {
                if !falling[child] {
                    allFalling = false
                    break
                }
            }
            if allFalling {
                falls(parent)
            }
        }
    }
    falls(disintegrated)
    return len(falling)
}

func contains(slice []string, item string) bool {
    for _, a := range slice {
        if a == item {
            return true
        }
    }
    return false
}

func part1() {
	ans := 0
    bricks := readBricks("2023/day22/input.txt")
    fallen, above, below := processBricks(bricks)

    for _, brick := range fallen {
        wouldFall := whatIf(above, below, brick)
        if wouldFall == 1 {
            ans++
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 22 is:", ans, "🎄")
}

func part2() {
	ans := 0
    bricks := readBricks("2023/day22/input.txt")
    fallen, above, below := processBricks(bricks)

    for _, brick := range fallen {
        wouldFall := whatIf(above, below, brick)
        ans += wouldFall - 1
    }

    fmt.Println("🎄 The answer to part 2 for day 22 is:", ans, "🎄")
}
