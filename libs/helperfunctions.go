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
    fileContent := strings.ReplaceAll(string(AllFileContent(filePath)), "\r\n", "\n")
    return strings.Split(fileContent, delimiter)
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
// It returns the index it was found at if the string is found and is not an empty string, -1 otherwise.
func SearchForStrInSlice(str string, slice []string) int {
    for i, sliceStr := range slice {
        if sliceStr == str && sliceStr != "" {
            return i
        }
    }
    return -1
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

// converts a slice of ints into a string with a delimiter
func IntSliceToStr(slice []int, delimiter string) string {
    strSlice := make([]string, len(slice))
    for i, num := range slice {
        strSlice[i] = strconv.Itoa(num)
    }
    return strings.Join(strSlice, delimiter)
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

// Inserts a value into a specific index of a slice
func InsertIntoSlice[T any](slice []T, index int, value T) []T {
    if index < 0 || index > len(slice) {
        return slice 
    }
    slice = append(slice, value) 
    copy(slice[index+1:], slice[index:]) 
    slice[index] = value 
    return slice
}

// Removes the element at a given index from a slice
func RemoveElementFromSlice[T any](slice []T, index int) []T {
    if index < 0 || index >= len(slice) {
        return slice 
    }
    newSlice := make([]T, len(slice))
    copy(newSlice, slice)
    return append(newSlice[:index], newSlice[index+1:]...)
}

// Absolute of any int value
func Abs(val int) int {
    if val < 0 {
        return -val
    }
    return val
}

// Transposes a string, this has the effect of rotating the original string 90 degrees counterclockwise (flipping it from horizontal to vertical orientation).
func TransposeString(str string) string {
    lines := strings.Split(str, "\n")
    newStr := []string{}
    for c := 0; c < len(lines[0]); c++ {
        row := ""
        for r := 0; r < len(lines); r++ {
            row += string(lines[r][c])
        }
        newStr = append(newStr, row)
    }
    return strings.Join(newStr, "\n")
}

// Transposes a []string, this has the effect of rotating the original string 90 degrees counterclockwise (flipping it from horizontal to vertical orientation). 
// Assumes the []string represents each line with a different element
func TransposeStringSlice(slice []string) []string {
    if len(slice) == 0 {
        return []string{}
    }

    maxLength := 0
    for _, str := range slice {
        if len(str) > maxLength {
            maxLength = len(str)
        }
    }

    transposed := make([]string, maxLength)
    for i := range transposed {
        transposed[i] = ""
    }

    for _, str := range slice {
        for i, char := range str {
            transposed[i] += string(char)
        }
    }

    return transposed
}

func Rotate45(slice []string, clockwise bool) []string {
    if len(slice) == 0 {
        return []string{}
    }

    maxLen := 0
    for _, str := range slice {
        if len(str) > maxLen {
            maxLen = len(str)
        }
    }

    resultLen := len(slice) + maxLen - 1
    result := make([]string, resultLen)

    for i := range result {
        result[i] = ""
    }

    if clockwise {
        for i, str := range slice {
            for j, char := range str {
                result[i+j] += string(char)
            }
        }
    } else {
        for i, str := range slice {
            for j, char := range str {
                result[len(slice)-1-i+j] += string(char)
            }
        }
    }

    return result
}

// perform a linear search on a string for a given character
func SearchForCharInStr(str string, charToFind rune) bool {
    for _, char := range str {
        if char == charToFind {
            return true
        }
    }
    return false
}

// abs for a float value
func AbsFloat(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}