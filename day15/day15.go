package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code-2023/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    initializationSequence := libs.FileToSlice("day15/input.txt", ",")
    for _, value := range initializationSequence {
        ans += hashAlgorithm(value)
    }
    fmt.Println("The answer to part 1 for day 15 is:", ans)
}

func hashAlgorithm(val string) int {
    hashValue := 0
    for _, char := range val {
        hashValue = ((hashValue + int(char)) * 17) % 256
    }
    return hashValue
}

func part2(){
    ans := 0
    initializationSequence := libs.FileToSlice("day15/input.txt", ",")
    
    // initalise boxes
    boxes := []LinkedList{}
    for i:=0; i<=255; i++ {
        boxes = append(boxes, LinkedList{})
    }

    for _, value := range initializationSequence {
        boxes = placeInBox(value, boxes)
    }

    for boxNum, box := range boxes {
        if box.Length != 0 {
            current := box.Head
            currentSlot := 1
            for current != nil {
                ans += (boxNum + 1) * currentSlot * current.FocalLength
                currentSlot += 1
                current = current.Next
            }
        }
    }
    fmt.Println("The answer to part 2 for day 15 is:", ans)
}

type Node struct {
    Lens string
    FocalLength int
    Next *Node
}

type LinkedList struct {
    Head   *Node
    Length int
}

func (l *LinkedList) Add(lens string, focalLength int) {
    newNode := &Node{Lens: lens, FocalLength: focalLength}
    if l.Head == nil {
        l.Head = newNode
    } else {
        current := l.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    l.Length++
}

func (l *LinkedList) Delete(lens string) {
    if l.Head == nil {
        return
    }

    if l.Head.Lens == lens {
        l.Head = l.Head.Next
        l.Length--
        return
    }

    current := l.Head
    for current.Next != nil && current.Next.Lens != lens {
        current = current.Next
    }

    if current.Next != nil {
        current.Next = current.Next.Next
        l.Length--
    }
}

func (l *LinkedList) Replace(lens string, newFocalLength int) bool {
    current := l.Head
    for current != nil {
        if current.Lens == lens {
            current.FocalLength = newFocalLength
            return true
        }
        current = current.Next
    }
    return false
}

func placeInBox(val string, boxes []LinkedList) []LinkedList {
    if val[len(val)-1:] == "-"{
        // remove from box
        boxNum := hashAlgorithm(val[:len(val)-1])
        boxes[boxNum].Delete(val[:len(val)-1])
        return boxes
    }
    // add to box
    splitVal := strings.Split(val, "=")
    boxNum := hashAlgorithm(splitVal[0])
    num, _ := strconv.Atoi(splitVal[1])

    alreadyExists := boxes[boxNum].Replace(splitVal[0], num)
    if !alreadyExists {
        boxes[boxNum].Add(splitVal[0], num)
    }
    return boxes
}