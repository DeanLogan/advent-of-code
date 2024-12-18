package main

import (
	"fmt"
	"strconv"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type Pos struct {
    x int
    y int
}

type Robot struct {
    currentPos Pos
    velocity Pos
}

const roomWidth = 101
const roomHeight = 103

func part1(){
    ans := 0

    robotsStr := libs.FileToSlice("2024/day14/input.txt", "\n")

    robotsInQuads := (make(map[int][]Robot))
    quadLines := Pos{roomWidth / 2, roomHeight / 2}

    for _, robotStr := range robotsStr {
        robot := strToRobot(robotStr)
        robot = moveNumOfSecs(robot, 100, roomWidth, roomHeight)
        robotInQuad := getQuadrant(robot.currentPos, quadLines)
        robotsInQuads[robotInQuad] = append(robotsInQuads[robotInQuad], robot)
    }

    ans = len(robotsInQuads[1]) * len(robotsInQuads[2]) * len(robotsInQuads[3]) * len(robotsInQuads[4])
    
    fmt.Println("🎄 The answer to part 1 for day 14 is:", ans, "🎄")
}

func strToRobot(robotStr string) Robot {
    currentPosStr, velocityStr := libs.SplitAtStr(robotStr, " v=")

    return Robot{strToPos(currentPosStr), strToPos(velocityStr)}
}

func strToPos(str string) Pos {
    xStr, yStr := libs.SplitAtChar(str[2:], ',')
    x, _ := strconv.Atoi(xStr)
    y, _ := strconv.Atoi(yStr)

    return Pos{x, y}
}

func moveNumOfSecs(robot Robot, secs int, roomWidth int, roomHeight int) Robot {
    robot.currentPos.x = (((robot.currentPos.x+(robot.velocity.x*secs)) % roomWidth) + roomWidth) % roomWidth
	robot.currentPos.y = (((robot.currentPos.y+(robot.velocity.y*secs)) % roomHeight) + roomHeight) % roomHeight

    return robot
}

func getQuadrant(robotPos Pos, quadLines Pos) int {
    if robotPos.x == quadLines.x || robotPos.y == quadLines.y {
        return 0 // represents robots on the quadrant lines
    } else if robotPos.x < quadLines.x && robotPos.y < quadLines.y {
        return 1
    } else if robotPos.x > quadLines.x && robotPos.y < quadLines.y {
        return 2
    } else if robotPos.x < quadLines.x && robotPos.y > quadLines.y {
        return 3
    } else {
        return 4
    }
}

func part2(){
    ans := 0

    robotsStr := libs.FileToSlice("2024/day14/input.txt", "\n")
    quadLines := Pos{roomWidth / 2, roomHeight / 2}
    
    minSafety := 215987200
    for i := 1; i < 10000; i++{
        var updatedRobots []Robot
        robotsInQuads := make(map[int]int)
        for j, robotStr := range robotsStr {
            robot := strToRobot(robotStr)
            updatedRobot := moveNumOfSecs(robot, i, roomWidth, roomHeight)
            updatedRobot.velocity = Pos{0, 0} // Set velocity to 0
            updatedRobots = append(updatedRobots, updatedRobot)
            quadrant := getQuadrant(updatedRobots[j].currentPos, quadLines)
            robotsInQuads[quadrant] += 1
        }
        count := robotsInQuads[1] * robotsInQuads[2] * robotsInQuads[3] * robotsInQuads[4]
        if count < minSafety {
            minSafety = count
            ans = i
        }
    }

    fmt.Println("🎄 The answer to part 2 for day 14 is:", ans, "🎄")
}