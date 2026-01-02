package main

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/DeanLogan/advent-of-code/libs"
)

func main() {
    part1()
}

type Tree struct {
    width, height int
    grid          [][]int
    piecesNeeded  []int
}

type Piece struct {
    matrix [][]int
    cells  [][2]int
    rots   [][][2]int
}

func part1() {
    ans := 0
    lines := libs.FileToSlice("2025/day12/input.txt", "\n\n")

    piecesBlocks := lines[:len(lines)-1]
    trees := strings.Split(strings.TrimSpace(lines[len(lines)-1]), "\n")

    pieces := strToPieces(piecesBlocks)

    pieceAreas := make([]int, len(pieces))
    for i, p := range pieces {
        pieceAreas[i] = computePieceArea(p.matrix)
    }

    for _, tline := range trees {
        tline = strings.TrimSpace(tline)
        if tline == "" {
            continue
        }

        width, height, counts := parseTreeLine(tline)

        maxArea := width * height
        neededArea := 0
        for i, cnt := range counts {
            if i >= len(pieceAreas) {
                continue
            }
            neededArea += cnt * pieceAreas[i]
        }

        if neededArea <= maxArea {
            ans++
        }
    }

    fmt.Println("🎄 The answer to part 1 for day 12 is:", ans, "🎄")
}

func strToPieces(blocks []string) []Piece {
    pieces := make([]Piece, 0, len(blocks))
    for _, block := range blocks {
        lines := strings.Split(strings.TrimSpace(block), "\n")
        // first line is "0:", "1:", etc., skip it
        lines = lines[1:]
        matrix := createMatrixFromPieceStr(lines)
        pieces = append(pieces, Piece{matrix: matrix})
    }
    return pieces
}

func createMatrixFromPieceStr(lines []string) [][]int {
    matrix := make([][]int, len(lines))
    for i, line := range lines {
        line = strings.TrimSpace(line)
        row := make([]int, len(line))
        for j, ch := range line {
            if ch == '#' {
                row[j] = 1
            } else {
                row[j] = 0
            }
        }
        matrix[i] = row
    }
    return matrix
}

func computePieceArea(m [][]int) int {
    area := 0
    for r := range m {
        for c := range m[r] {
            if m[r][c] == 1 {
                area++
            }
        }
    }
    return area
}

func parseTreeLine(line string) (int, int, []int) {
    parts := strings.Split(line, ":")
    dims := strings.TrimSpace(parts[0])
    countsStr := strings.TrimSpace(parts[1])

    wh := strings.Split(dims, "x")
    w, _ := strconv.Atoi(strings.TrimSpace(wh[0]))
    h, _ := strconv.Atoi(strings.TrimSpace(wh[1]))

    fields := strings.Fields(countsStr)
    counts := make([]int, len(fields))
    for i, f := range fields {
        v, _ := strconv.Atoi(f)
        counts[i] = v
    }

    return w, h, counts
}
