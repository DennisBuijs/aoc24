package util

import (
	"bufio"
	"log"
	"os"
)

func OpenFileAsStringGrid(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	grid := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for col, r := range line {
			row[col] = string(r)
		}
		grid = append(grid, row)
	}

	return grid
}

func OpenFileAsStringSlice(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	result := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
