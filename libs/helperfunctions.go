package libs

import (
	"bufio"
	"log"
	"os"
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