package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Guard struct {
	x, y      int
	direction string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	gridHeight := 100

	grid := make([][]string, 0, gridHeight)
	var startX, startY int
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := make([]string, len(scanner.Text()))
		for i, char := range scanner.Text() {
			line[i] = string(char)
			if line[i] == "^" {
				startX = i
				startY = y
			}
		}
		grid = append(grid, line)
		y++
	}

	guard := Guard{x: startX, y: startY, direction: "UP"}
	positionsVisited := [][]int{}

	for !outOfBounds(grid, guard) {
		if guard.direction == "UP" {
			if nextPositionIsAvailable(grid, guard.x, guard.y-1) {
				guard.y--
				positionsVisited = append(positionsVisited, []int{guard.x, guard.y})
				continue
			} else {
				guard.turnRight()
				continue
			}
		}

		if guard.direction == "RIGHT" {
			if nextPositionIsAvailable(grid, guard.x+1, guard.y) {
				guard.x++
				positionsVisited = append(positionsVisited, []int{guard.x, guard.y})
				continue
			} else {
				guard.turnRight()
				continue
			}
		}

		if guard.direction == "DOWN" {
			if nextPositionIsAvailable(grid, guard.x, guard.y+1) {
				guard.y++
				positionsVisited = append(positionsVisited, []int{guard.x, guard.y})
				continue
			} else {
				guard.turnRight()
				continue
			}
		}

		if guard.direction == "LEFT" {
			if nextPositionIsAvailable(grid, guard.x-1, guard.y) {
				guard.x--
				positionsVisited = append(positionsVisited, []int{guard.x, guard.y})
				continue
			} else {
				guard.turnRight()
				continue
			}
		}
	}

	uniquePositions := [][]int{}

	for _, position := range positionsVisited {
		notUnique := false
		for _, uniquePosition := range uniquePositions {
			if position[0] == uniquePosition[0] && position[1] == uniquePosition[1] {
				notUnique = true
				break
			}
		}

		if !notUnique {
			uniquePositions = append(uniquePositions, position)
		}
	}

	fmt.Println(len(uniquePositions))
}

func outOfBounds(grid [][]string, guard Guard) bool {
	if guard.direction == "UP" {
		return guard.y == 0
	}

	if guard.direction == "RIGHT" {
		return guard.x == len(grid[guard.y])
	}

	if guard.direction == "DOWN" {
		return guard.y == len(grid)
	}

	if guard.direction == "LEFT" {
		return guard.x == 0
	}

	return false
}

func nextPositionIsAvailable(grid [][]string, x int, y int) bool {
	return grid[y][x] != "#"
}

func (g *Guard) turnRight() {
	if g.direction == "UP" {
		g.direction = "RIGHT"
		return
	}

	if g.direction == "RIGHT" {
		g.direction = "DOWN"
		return
	}

	if g.direction == "DOWN" {
		g.direction = "LEFT"
		return
	}

	if g.direction == "LEFT" {
		g.direction = "UP"
		return
	}
}
