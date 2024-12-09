package main

import (
	"aoc24/util"
	"fmt"
)

type Vector struct {
	x, y int
}

type Position struct {
	x, y int
}

type Antenna struct {
	position Position
}

type AntennaGroup struct {
	antennas []Antenna
	vectors  map[Vector]Vector
}

var grid = [][]string{}

var antennaGroups = make(map[string]AntennaGroup)
var antinodes = make(map[Position]Position)

var emptyVector = Vector{0, 0}

func main() {
	grid = util.OpenFileAsStringGrid("input.txt")

	for y, row := range grid {
		for x, cell := range row {
			if cell != "." {
				antenna := Antenna{
					position: Position{
						x: x,
						y: y,
					},
				}

				if _, exist := antennaGroups[cell]; !exist {
					antennaGroups[cell] = AntennaGroup{
						antennas: []Antenna{},
						vectors:  make(map[Vector]Vector),
					}
				}

				group := antennaGroups[cell]
				group.antennas = append(group.antennas, antenna)
				antennaGroups[cell] = group
			}
		}
	}

	for _, group := range antennaGroups {
		for _, antenna := range group.antennas {
			for _, otherAntenna := range group.antennas {
				vector := antenna.position.getVector(otherAntenna.position)
				otherVector := otherAntenna.position.getVector(antenna.position)

				antinodePosition := antenna.position.applyVector(vector)
				otherAntinodePosition := otherAntenna.position.applyVector(otherVector)

				if antinodePosition.inBounds() && vector != emptyVector {
					antinodes[antinodePosition] = antinodePosition
				}

				if otherAntinodePosition.inBounds() && otherVector != emptyVector {
					antinodes[otherAntinodePosition] = otherAntinodePosition
				}
			}
		}
	}

	drawGrid(grid)
	fmt.Println(len(antinodes))
}

func (this Position) getVector(other Position) Vector {
	return Vector{
		x: this.x - other.x,
		y: this.y - other.y,
	}
}

func (p Position) applyVector(v Vector) Position {
	return Position{
		x: p.x + v.x,
		y: p.y + v.y,
	}
}

func (p Position) inBounds() bool {
	return p.x >= 0 && p.x < len(grid[0]) && p.y >= 0 && p.y < len(grid)
}

func (p Position) cacheKey() string {
	return fmt.Sprintf("%v-%v", p.x, p.y)
}

func drawGrid(grid [][]string) {
	for y, row := range grid {
		for x, cell := range row {
			if _, exists := antinodes[Position{x, y}]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(cell)
			}
		}

		fmt.Print("\n")
	}
}

func (g AntennaGroup) containsAntennaAtPosition(p Position) bool {
	for _, a := range g.antennas {
		if a.position == p {
			return true
		}
	}

	return false
}