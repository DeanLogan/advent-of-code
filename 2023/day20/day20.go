package main

import (
	"fmt"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

func part1(){
    ans := 0
    lines := libs.FileToSlice("2023/day20/input.txt", "\n")
    moduleMap := convertLinesToMap(lines)
    moduleMap = populateMemory(moduleMap)
    queue := Queue{}
    for i := 0; i<1000; i++ {
        lowPulse++
        queue.Enqueue(moduleMap["roadcaster"])
        queue, moduleMap = pulses(queue, moduleMap)
    }
    ans = lowPulse * highPulse
    fmt.Println("The answer to part 1 for day 20 is:", ans)
}
var lowPulse = 0
var highPulse = 0

type Queue struct {
    items []Module
}

type Module struct {
    Type string
    Name string
    Destinations []string
    SendPulse bool
    Memory map[string]bool
}

func (q *Queue) Enqueue(i Module) {
    q.items = append(q.items, i)
}

func (q *Queue) Dequeue() Module {
    toRemove := q.items[0]
    q.items = q.items[1:]
    return toRemove
}

func convertLinesToMap(lines []string) map[string]Module {
    moduleMap := make(map[string]Module)
    for _, line := range lines {
        key, valueStr := libs.SplitAtStr(line, " ->")
        valueStr = valueStr[3:]
        value := strings.Split(valueStr, ", ")
        moduleMap[key[1:]] = Module{string(key[0]), key[1:], value, false, make(map[string]bool)}
    }
    return moduleMap
}

func populateMemory(moduleMap map[string]Module) map[string]Module {
    for key, value := range moduleMap {
        for _, module := range value.Destinations {
            mod := moduleMap[module]
            if mod.Type == "&" {
                mod.Memory[key] = false
                moduleMap[module] = mod
            }
        }
    }
    return moduleMap
}

func pulse(moduleMap map[string]Module, queue Queue) Queue {
    mod := queue.Dequeue()
    for _, module := range mod.Destinations {
        destMod := moduleMap[module]
        addPulses(mod.SendPulse)
        if !mod.SendPulse && destMod.Type == "%" {
            destMod.SendPulse = !destMod.SendPulse
            moduleMap[module] = destMod
            queue.Enqueue(moduleMap[module])
        } 
        if destMod.Type == "&" {
            destMod.Memory[mod.Name] = mod.SendPulse
            destMod.SendPulse = !allTrue(destMod.Memory)
            moduleMap[module] = destMod
            queue.Enqueue(moduleMap[module])
        }
        if destMod.Name == "rx" && !mod.SendPulse{
            fmt.Println(mod.SendPulse, destMod.Name)
        }
    }
    return queue
}

func pulses(queue Queue, moduleMap map[string]Module) (Queue, map[string]Module) {
    if len(queue.items) != 0 {
        queue = pulse(moduleMap, queue)
        return pulses(queue, moduleMap)
    }
    return queue, moduleMap
}


func addPulses(pulseValue bool) {
    if pulseValue {
        highPulse++
    } else {
        lowPulse++
    }
}

func allTrue(memoryMap map[string]bool) bool {
    for _, val := range memoryMap {
        if !val {
            return false
        }
    }
    return true
}

func part2() {
    ans := 0
    lines := libs.FileToSlice("2023/day20/input.txt", "\n")
    moduleMap := convertLinesToMap(lines)
    moduleMap = populateMemory(moduleMap)
    queue := Queue{}
    seens := make(map[string][]int)
    for i := 0; i < 1000000000000000000; i++ {
        if i % 100000 == 0 && i > 0 {
            cycles := []int{}
            for _, v := range seens {
                v2 := []int{}
                for i := 0; i < len(v) - 1; i++ {
                    v2 = append(v2, v[i+1]-v[i])
                }
                cycles = append(cycles, v2[0])
            }
            ans = lcm(cycles)
            fmt.Println("The answer to part 2 for day 20 is:", ans)
            return
        }
        if i > 0 && allFalse(moduleMap) {
            panic("All modules are off")
        }
        queue.Enqueue(moduleMap["roadcaster"])
        queue, moduleMap, _ = pulses2(queue, moduleMap, i, seens)
    }
    fmt.Println("The answer to part 2 for day 20 is:", ans)
}

func pulses2(queue Queue, moduleMap map[string]Module, i int, seens map[string][]int) (Queue, map[string]Module, bool) {
    if len(queue.items) != 0 {
        stop := false
        queue, stop = pulse2(moduleMap, queue, i, seens)
        if stop {
            return queue, moduleMap, true
        }
        return pulses2(queue, moduleMap, i, seens)
    }
    return queue, moduleMap, false
}

func pulse2(moduleMap map[string]Module, queue Queue, i int, seens map[string][]int) (Queue, bool) {
    mod := queue.Dequeue()
    for _, module := range mod.Destinations {
        destMod := moduleMap[module]
        addPulses(mod.SendPulse)
        if mod.SendPulse && isSrc(mod.Name) {
            seens[mod.Name] = append(seens[mod.Name], i+1)
        }
        if !mod.SendPulse && destMod.Type == "%" {
            destMod.SendPulse = !destMod.SendPulse
            moduleMap[module] = destMod
            queue.Enqueue(moduleMap[module])
        } 
        if destMod.Type == "&" {
            destMod.Memory[mod.Name] = mod.SendPulse
            destMod.SendPulse = !allTrue(destMod.Memory)
            moduleMap[module] = destMod
            queue.Enqueue(moduleMap[module])
        }
        if destMod.Name == "rx" && !mod.SendPulse{
            fmt.Println(i + 1)
            return queue, true
        }
    }
    return queue, false
}

func isSrc(src string) bool {
    return src == "bm" || src == "cl" || src == "tn" || src == "dr"
}

func allFalse(moduleMap map[string]Module) bool {
    for _, mod := range moduleMap {
        if mod.SendPulse || allTrue(mod.Memory) {
            return false
        }
    }
    return true
}

func lcm(nums []int) int {
    result := nums[0]
    for i := 1; i < len(nums); i++ {
        result = result * nums[i] / libs.Gcd(result, nums[i])
    }
    return result
}