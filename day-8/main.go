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
var antinodes = make(map[string]Position)

func main() {
	grid = util.OpenFileAsStringGrid("test.txt")

	totalAntennas := 0

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

	for key, group := range antennaGroups {
		totalAntennas += len(group.antennas)
		for _, antenna := range group.antennas {
			for _, otherAntenna := range group.antennas {
				vector := antenna.position.getVector(otherAntenna.position)

				if vector.x == 0 && vector.y == 0 {
					continue
				}

				group.vectors[vector] = vector
			}
		}

		antennaGroups[key] = group

		for _, antenna := range group.antennas {
			for _, vector := range group.vectors {
				antinodePosition := antenna.position.applyVector(vector)
				if antinodePosition.inBounds() && !group.containsAntennaAtPosition(antinodePosition) {
					antinodes[antinodePosition.cacheKey()] = antinodePosition
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
			cellKey := fmt.Sprintf("%v-%v", x, y)
			if _, exists := antinodes[cellKey]; exists {
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
