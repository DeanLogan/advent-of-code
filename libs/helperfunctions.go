package libs

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
	This file contains a list of functions that are used frequently across all of the days so instead of copying and pasting them over
	I created a mini library of functions that I can use.
*/

// AllFileContent reads a file and returns its content as a byte slice.
func AllFileContent(filePath string) []byte {
    content, err := os.ReadFile(filePath)
	if err != nil {
        log.Fatal(err)
    }
    return content
}

// FileToSlice reads a file, replaces all "\r\n" with "\n", and splits the content by a given delimiter.
// It returns a slice of strings.
func FileToSlice(filePath string, delimiter string) []string {
    fileContent := AllFileContent(filePath)
    data := strings.ReplaceAll(string(fileContent), "\r\n", "\n")
    return strings.Split(data, delimiter)
}

// GetScannerForFile opens a file and returns a bufio.Scanner for it.
func GetScannerForFile(filePath string) *bufio.Scanner {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err, " Failed to open file")
        return nil
    }

    return bufio.NewScanner(file)
}

// SplitAtChar splits a string at the last occurrence of a given character.
// It returns two strings: the part before the character and the part after the character.
func SplitAtChar(str string, char rune) (string, string) {
    index := strings.LastIndex(str, string(char))
    if index != -1 {
        return str[:index], str[index+1:]
    }
    return str, str
}

// SplitAtStr splits a string at the last occurrence of a given string.
// It returns two strings: the part before the character and the part after the character.
func SplitAtStr(str string, strToSplitAt string) (string, string) {
    index := strings.LastIndex(str, strToSplitAt)
    if index != -1 {
        return str[:index], str[index+1:]
    }
    return str, str
}

// ReverseString reverses a string and returns the result.
func ReverseString(s string) (result string) {
    for _,v := range s {
        result = string(v) + result
    }
    return result 
}

// Min returns the smaller of two integers.
func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// Max returns the larger of two integers
func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// SearchForStrInSlice performs a linear search for a string in a slice of strings.
// It returns true if the string is found and is not an empty string, false otherwise.
func SearchForStrInSlice(str string, slice []string) bool {
    for _, sliceStr := range slice {
        if sliceStr == str && sliceStr != "" {
            return true
        }
    }
    return false
}

// converts any string into a slice of ints, if a rune/character in the string cannot be converted then it is ignored for the final slice
func StrToIntSlice(str string, delimiter string) []int {
    strSlice := strings.Split(str, delimiter)
    intSlice := []int{}
    for _, strNum := range strSlice{
        strNum = strings.ReplaceAll(strNum, " ", "")
        num, err := strconv.Atoi(strNum)
        // ignore any strings or characters that cannot be converted into an int
        if err == nil{
            intSlice = append(intSlice, num)
        }
    }
    return intSlice
}

// returns the index value of the maximum value in a slice
func IndexOfMax(steps []int) int {
    maxIndex := 0
    for i, value := range steps {
        if value > steps[maxIndex] {
            maxIndex = i
        }
    }
    return maxIndex
}

// finds the max value of a slice and returns it
func MaxOfSlice(steps []int) int {
    if len(steps) == 0 {
        return 0 // or return an error
    }
    maxIndex := IndexOfMax(steps)
    return steps[maxIndex]
}

// returns the index value of the minimum value in a slice
func IndexOfMin(steps []int) int {
    minIndex := 0
    for i, value := range steps {
        if value < steps[minIndex] {
            minIndex = i
        }
    }
    return minIndex
}

// finds the min value of a slice and returns it
func MinOfSlice(steps []int) int {
    if len(steps) == 0 {
        return 0 // or return an error
    }
    minIndex := IndexOfMin(steps)
    return steps[minIndex]
}

// Gcd calculates the Greatest Common Divisor (GCD) of two integers.
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Lcm calculates the Least Common Multiple (LCM) of two integers.
func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}

// Inserts a string into a specific index of a []string
func InsertIntoSlice(slice []string, index int, value string) []string {
    slice = append(slice, "")
    copy(slice[index+1:], slice[index:])
    slice[index] = value
    return slice
}

// Absolute of any int value
func Abs(val int) int {
    if val < 0 {
        return -val
    }
    return val
}