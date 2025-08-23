package test2
package main

import "github.com/01-edu/z01"

func main() {
    EightQueens()
}

func EightQueens() {
    solve([]int{})
}

// Recursive function to try placing queens row by row
func solve(positions []int) {
    row := len(positions)
    if row == 8 {
        printSolution(positions)
        return
    }

    for col := 1; col <= 8; col++ {
        if isSafe(positions, row, col) {
            solve(append(positions, col))
        }
    }
}

// Check if a queen can be safely placed at (row, col)
func isSafe(positions []int, row, col int) bool {
    for r := 0; r < row; r++ {
        c := positions[r]
        if c == col || abs(c-col) == abs(r-row) {
            return false
        }
    }
    return true
}

// Print a valid solution using z01.PrintRune
func printSolution(positions []int) {
    for _, col := range positions {
        z01.PrintRune(rune(col + '0'))
    }
    z01.PrintRune('\n')
}

// Return absolute value
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
