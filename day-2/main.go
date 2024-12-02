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
		safe := false
		increasing := report[0] < report[1]

		for i := 0; i < len(report); i++ {
			if i > 0 {
				if difference(report[i], report[i-1]) > 3 {
					safe = false
					break
				}

				if increasing && report[i] < report[i-1] {
					safe = false
					break
				}

				if !increasing && report[i] > report[i-1] {
					safe = false
					break
				}

				if report[i] == report[i-1] {
					safe = false
					break
				}

				safe = true
			}
		}

		if safe {
			safeReports++
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
