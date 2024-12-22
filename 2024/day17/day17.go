package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/DeanLogan/advent-of-code/libs"
)

func main(){
    part1()
    part2()
}

type register struct {
    A int
    B int
    C int
}

func part1(){
    file := libs.FileToSlice("2024/day17/input.txt", "\n\n")
    reg := getRegister(file[0])
    program := getProgram(file[1])
    opcodes, operands := programToOpcodesAndOperands(program)

    _, _, out := runProgram(opcodes, operands, reg, 0, []int{})
    ans := libs.IntSliceToStr(out, ",")

    fmt.Println("🎄 The answer to part 1 for day 17 is:", ans, "🎄")
}

func runProgram(opcodes []int, operands []int, reg register, ip int, out []int) (register, int, []int) {
    if ip >= len(opcodes) {
        return reg, ip, out
    }

    switch opcodes[ip] {
    case 0: // adv
        reg.A = reg.A / int(math.Pow(2, float64(comboOperand(operands[ip], reg))))
    case 1: // bxl
        reg.B = reg.B ^ operands[ip]
    case 2: // bst
        reg.B = comboOperand(operands[ip], reg) % 8
    case 3: // jnz
        if reg.A != 0 {
            ip = operands[ip]
            return runProgram(opcodes, operands, reg, ip, out)
        }
    case 4: // bxc
        reg.B = reg.B ^ reg.C
    case 5: // out
        out = append(out, (comboOperand(operands[ip], reg) % 8))
    case 6: // bdv
        reg.B = reg.A / int(math.Pow(2, float64(comboOperand(operands[ip], reg))))
    case 7: // cdv
        reg.C = reg.A / int(math.Pow(2, float64(comboOperand(operands[ip], reg))))
    }
    
    ip++

    return runProgram(opcodes, operands, reg, ip, out)
}

func comboOperand(operand int, reg register) int {
    if operand >= 0 && operand <= 3 {
        return operand
    } else if operand == 4 {
        return reg.A
    } else if operand == 5 {
        return reg.B
    } else if operand == 6 {
        return reg.C
    }
    return -1
}

func programToOpcodesAndOperands(program []int) ([]int, []int) {
    opcodes := []int{}
    operands := []int{}

    for i, val := range program {
        if i%2 == 0 {
            opcodes = append(opcodes, val)
        } else {
            operands = append(operands, val)
        }
    }

    return opcodes, operands
}

func getProgram(programStr string) []int {
    programStr = programStr[9:]
    return libs.StrToIntSlice(programStr,",")
}

func getRegister(regStr string) register {
    regVals := make([]int, 3)
    regParts := strings.Split(regStr, "\n")

    for i, part := range regParts {
        _, numStr := libs.SplitAtStr(part, ": ")
        regVals[i], _ = strconv.Atoi(numStr[1:])
    }

    return register{A:regVals[0], B:regVals[1], C:regVals[2]}
}

func part2(){
    ans := 0

    file := libs.FileToSlice("2024/day17/input.txt", "\n\n")
    program := getProgram(file[1])
    opcodes, operands := programToOpcodesAndOperands(program)

    ans = expect(opcodes, operands, program, 0)

    fmt.Println("🎄 The answer to part 2 for day 17 is:", ans, "🎄")
}

func expect(opcodes []int, operands, out []int, prevA int) int {
    if len(out) == 0 {
        return prevA
    }
    for a := 0; a < (1 << 10); a++ {
		_, _, result := runProgram(opcodes, operands, register{a,0,0}, 0, []int{})
        if a>>3 == prevA&127 && result[0] == out[len(out)-1] {
            ret := expect(opcodes, operands, out[:len(out)-1], (prevA<<3)|(a%8))
            if ret != -1 {
                return ret
            }
        }
    }
    return -1
}