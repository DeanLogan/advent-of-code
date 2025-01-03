package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main() {
	part1()
	part2()
}

type Vector3 struct {
	x, y, z float64
}

type Vector2 struct {
	x, y float64
}

type Hailstone struct {
	pos, vel Vector3
}

func part1() {
	lines := libs.FileToSlice("2023/day24/input.txt", "\n")
	hailStones := parseHailstones(lines)

	areaMin, areaMax := float64(200000000000000), float64(400000000000000)
	intersectCount := 0
	for i := 0; i < len(hailStones)-1; i++ {
		for j := i + 1; j < len(hailStones); j++ {
			a, b := hailStones[i], hailStones[j]
			if point, does := hailstonesIntersect(a, b); does {
				if point.x >= areaMin && point.x <= areaMax &&
					point.y >= areaMin && point.y <= areaMax {
					dx := point.x - a.pos.x
					dy := point.y - a.pos.y
					if (dx > 0) == (a.vel.x > 0) && (dy > 0) == (a.vel.y > 0) {
						dx = point.x - b.pos.x
						dy = point.y - b.pos.y
						if (dx > 0) == (b.vel.x > 0) && (dy > 0) == (b.vel.y > 0) {
							intersectCount++
						}
					}
				}
			}
		}
	}

	ans := intersectCount
	fmt.Println("🎄 The answer to part 1 for day 24 is:", ans, "🎄")
}

func hailstonesIntersect(a, b Hailstone) (Vector2, bool) {
	a2 := Vector2{a.vel.x, a.vel.y}
	b2 := Vector2{b.vel.x, b.vel.y}
	d2 := Vector2{b.pos.x - a.pos.x, b.pos.y - a.pos.y}

	det := vectorCross(a2, b2)
	
	if det == 0 {
		return Vector2{-1, -1}, false
	}

	u := vectorCross(d2, b2) / det
	return Vector2{a.pos.x + a.vel.x*u, a.pos.y + a.vel.y*u}, true
}

func vectorCross(a, b Vector2) float64 {
	return (a.x * b.y) - (a.y * b.x)
}

func parseHailstones(lines []string) []Hailstone {
	hailStones := make([]Hailstone, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, " @ ")
		coords := libs.StrToIntSlice(split[0], ",")
		vels := libs.StrToIntSlice(split[1],",")
		hailStone := Hailstone{Vector3{float64(coords[0]), float64(coords[1]), float64(coords[2])}, Vector3{float64(vels[0]), float64(vels[1]), float64(vels[2])}}
		hailStones = append(hailStones, hailStone)
	}
	return hailStones
}

func part2() {
	lines := libs.FileToSlice("2023/day24/input.txt", "\n")
	hailStones := parseHailstones(lines)
    maybeX, maybeY, maybeZ := []int{}, []int{}, []int{}
    for i := 0; i < len(hailStones)-1; i++ {
        for j := i + 1; j < len(hailStones); j++ {
            a, b := hailStones[i], hailStones[j]
            if a.vel.x == b.vel.x {
                nextMaybe := findMatchingVel(int(b.pos.x-a.pos.x), int(a.vel.x))
                if len(maybeX) == 0 {
                    maybeX = nextMaybe
                } else {
                    maybeX = getIntersect(maybeX, nextMaybe)
                }
            }
            if a.vel.y == b.vel.y {
                nextMaybe := findMatchingVel(int(b.pos.y-a.pos.y), int(a.vel.y))
                if len(maybeY) == 0 {
                    maybeY = nextMaybe
                } else {
                    maybeY = getIntersect(maybeY, nextMaybe)
                }
            }
            if a.vel.z == b.vel.z {
                nextMaybe := findMatchingVel(int(b.pos.z-a.pos.z), int(a.vel.z))
                if len(maybeZ) == 0 {
                    maybeZ = nextMaybe
                } else {
                    maybeZ = getIntersect(maybeZ, nextMaybe)
                }
            }
        }
    }
    
    ans := 0
    if len(maybeX) == len(maybeY) && len(maybeY) == len(maybeZ) && len(maybeZ) == 1 {
        // only one possible velocity in all dimensions
        rockVel := Vector3{float64(maybeX[0]), float64(maybeY[0]), float64(maybeZ[0])}
        hailStoneA, hailStoneB := hailStones[0], hailStones[1]
        mA, mB := 0.0, 0.0
        if (hailStoneA.vel.x - rockVel.x) != 0 {
            mA = (hailStoneA.vel.y - rockVel.y) / (hailStoneA.vel.x - rockVel.x)
        }
        if (hailStoneB.vel.x - rockVel.x) != 0 {
            mB = (hailStoneB.vel.y - rockVel.y) / (hailStoneB.vel.x - rockVel.x)
        }
        cA := hailStoneA.pos.y - (mA * hailStoneA.pos.x)
        cB := hailStoneB.pos.y - (mB * hailStoneB.pos.x)
        if (mA - mB) != 0 {
            xPos := (cB - cA) / (mA - mB)
            yPos := mA*xPos + cA
            if (hailStoneA.vel.x - rockVel.x) != 0 {
                time := (xPos - hailStoneA.pos.x) / (hailStoneA.vel.x - rockVel.x)
                zPos := hailStoneA.pos.z + (hailStoneA.vel.z-rockVel.z)*time
                ans = int(xPos + yPos + zPos)
            }
        }
    }
    
    fmt.Println("🎄 The answer to part 2 for day 24 is:", ans, "🎄")
}

func findMatchingVel(dvel, pv int) []int {
	match := []int{}
	for v := -1000; v < 1000; v++ {
		if v != pv && dvel%(v-pv) == 0 {
			match = append(match, v)
		}
	}
	return match
}

func getIntersect(a, b []int) []int {
	result := []int{}
	for _, val := range a {
		if slices.Contains(b, val) {
			result = append(result, val)
		}
	}
	return result
}
