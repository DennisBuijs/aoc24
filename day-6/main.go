// INCOMPLETE -- PART 2 is incorrect

package main

import (
	"aoc24/util"
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type Grid struct {
	cells [][]string
}

type Guard struct {
	position  Position
	direction string
}

type Position struct {
	x, y int
}

type Vector struct {
	x, y int
}

var grid Grid
var blocks []Position
var guard Guard

var debug = false
var autoStep = true

var directionVectors = map[string]Vector{
	"UP":    {x: 0, y: -1},
	"DOWN":  {x: 0, y: 1},
	"LEFT":  {x: -1, y: 0},
	"RIGHT": {x: 1, y: 0},
}

var directionCharacters = map[string]string{
	"UP":    "^",
	"DOWN":  "v",
	"LEFT":  "<",
	"RIGHT": ">",
}

var nextDirection = map[string]string{
	"UP":    "RIGHT",
	"DOWN":  "LEFT",
	"LEFT":  "UP",
	"RIGHT": "DOWN",
}

func (g Guard) cacheKey() string {
	return fmt.Sprintf("%v-%v", g.position.x, g.position.y)
}

func (g Guard) fullCacheKey() string {
	return fmt.Sprintf("%v-%v-%s", g.position.x, g.position.y, g.direction)
}

var guardHistory = map[string]bool{}
var fullGuardHistory = map[string]bool{}

var loopBlocks = map[string]bool{}

func main() {
	grid = Grid{util.OpenFileAsStringGrid("test.txt")}
	guard := Guard{direction: "UP"}

	for y := 0; y < len(grid.cells); y++ {
		for x := 0; x < len(grid.cells[0]); x++ {
			cell := grid.cells[y][x]

			if cell == "#" {
				blocks = append(blocks, Position{x: x, y: y})
			}

			if cell == "^" {
				guard.position = Position{x: x, y: y}
			}
		}
	}

	grid.cells[guard.position.y][guard.position.x] = "."

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		if debug {
			clearTerminal()
			grid.draw(guard)
			fmt.Printf("[GUARD X:%v Y:%v]", guard.position.x, guard.position.y)
			fmt.Printf("[GUARD DIR:%s]", guard.direction)
			fmt.Printf("[GUARD STEPS:%v]", len(guardHistory))
		}

		if guard.outOfBounds() {
			clearTerminal()
			fmt.Printf("Guard out of bounds after %v steps\n", len(guardHistory))
			fmt.Printf("Found %v loop blocks\n", len(loopBlocks))
			break
		} else {
			guardHistory[guard.cacheKey()] = true
			fullGuardHistory[guard.fullCacheKey()] = true
			guard.step()
		}

		if !autoStep {
			scanner.Scan()
			scanner.Text()
		}
	}
}

func (g *Guard) step() {
	if g.nextStepIsBlocked() {
		g.turnRight()
		return
	}

	g.position = g.position.applyVector(directionVectors[g.direction])
	g.nextStepBlockedWouldCreateLoop()
}

func (g *Guard) turnRight() {
	g.direction = nextDirection[g.direction]
}

func (g Guard) nextStepIsBlocked() bool {
	nextPosition := g.position.applyVector(directionVectors[g.direction])

	for _, b := range blocks {
		if b == nextPosition {
			return true
		}
	}

	return false
}

func (p Position) applyVector(v Vector) Position {
	return Position{x: p.x + v.x, y: p.y + v.y}
}

func (g Grid) draw(guard Guard) {
	for y, row := range g.cells {
		for x, cell := range row {
			if guard.position.x == x && guard.position.y == y {
				fmt.Print(directionCharacters[guard.direction])
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Print("\n")
	}
}

func (g Guard) outOfBounds() bool {
	return g.position.x < 0 || g.position.x == len(grid.cells[0]) || g.position.y < 0 || g.position.y == len(grid.cells)
}

func (g Guard) nextStepBlockedWouldCreateLoop() {
	g.turnRight()
	if _, exists := fullGuardHistory[g.fullCacheKey()]; exists {
		g.position = g.position.applyVector(directionVectors[g.direction])
		loopBlocks[g.fullCacheKey()] = true
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
