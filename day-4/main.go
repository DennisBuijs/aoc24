package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	gridHeight := 140

	grid := make([][]string, 0, gridHeight)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := make([]string, len(scanner.Text()))
		for i, char := range scanner.Text() {
			line[i] = string(char)
		}
		grid = append(grid, line)
	}

	output := findAmountOfXmasInGrid(grid)
	fmt.Println(output)
}

func findAmountOfXmasInGrid(grid [][]string) int {
	total := 0

	gridHeight := len(grid)
	gridWidth := len(grid[0])

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			if grid[y][x] != "X" {
				continue
			}

			// Left to right
			if x+3 < gridWidth && grid[y][x+1] == "M" && grid[y][x+2] == "A" && grid[y][x+3] == "S" {
				total++
			}

			// Right to left
			if x-3 >= 0 && grid[y][x-1] == "M" && grid[y][x-2] == "A" && grid[y][x-3] == "S" {
				total++
			}

			// Top to bottom
			if y+3 < gridHeight && grid[y+1][x] == "M" && grid[y+2][x] == "A" && grid[y+3][x] == "S" {
				total++
			}

			// Bottom to top
			if y-3 >= 0 && grid[y-1][x] == "M" && grid[y-2][x] == "A" && grid[y-3][x] == "S" {
				total++
			}

			// Diagonal NE
			if x+3 < gridWidth && y-3 >= 0 && grid[y-1][x+1] == "M" && grid[y-2][x+2] == "A" && grid[y-3][x+3] == "S" {
				total++
			}

			// Diagonal SE
			if x+3 < gridWidth && y+3 < gridHeight && grid[y+1][x+1] == "M" && grid[y+2][x+2] == "A" && grid[y+3][x+3] == "S" {
				total++
			}

			// Diagonal NW
			if x-3 >= 0 && y-3 >= 0 && grid[y-1][x-1] == "M" && grid[y-2][x-2] == "A" && grid[y-3][x-3] == "S" {
				total++
			}

			// Diagonal SW
			if x-3 >= 0 && y+3 < gridHeight && grid[y+1][x-1] == "M" && grid[y+2][x-2] == "A" && grid[y+3][x-3] == "S" {
				total++
			}
		}
	}

	return total
}
