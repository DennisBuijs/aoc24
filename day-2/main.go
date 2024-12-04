package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reports := make([][]int, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		report := make([]int, 0, len(levels))
		for _, l := range levels {
			level, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal(err)
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	output := calculateAmountOfSafeReports(reports)

	log.Println(output)
}

func calculateAmountOfSafeReports(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		levels := make([][]int, len(report)+1)
		for i := range report {
			level := make([]int, 0, len(report)-1)
			level = append(level, report[:i]...)
			level = append(level, report[i+1:]...)
			levels[i] = level
		}

		levels[len(levels)-1] = report

		for _, level := range levels {
			safe := false

			increasing := level[0] < level[1]

			for i := 0; i < len(level); i++ {
				if i > 0 {
					if difference(level[i], level[i-1]) > 3 {
						safe = false
						break
					}

					if increasing && level[i] < level[i-1] {
						safe = false
						break
					}

					if !increasing && level[i] > level[i-1] {
						safe = false
						break
					}

					if level[i] == level[i-1] {
						safe = false
						break
					}

					safe = true
				}
			}

			if safe {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func difference(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
