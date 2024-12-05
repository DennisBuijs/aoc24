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

	rules := make([][]int, 0)
	manuals := make([][]int, 0)

	readingMode := "RULES"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingMode = "MANUALS"
			continue
		}

		if readingMode == "RULES" {
			rules = append(rules, mapLineToRule(line))
		} else {
			manuals = append(manuals, mapLineToManual(line))
		}
	}

	validManuals := make([][]int, 0)

	for _, manual := range manuals {
		if isManualValid(manual, rules) {
			validManuals = append(validManuals, manual)
		}
	}

	output := calculateSumOfMiddlePageNumbers(validManuals)

	log.Println(output)
}

func mapLineToRule(line string) []int {
	parts := strings.Split(line, "|")

	rule := make([]int, 2)
	rule[0], _ = strconv.Atoi(parts[0])
	rule[1], _ = strconv.Atoi(parts[1])

	return rule
}

func mapLineToManual(line string) []int {
	parts := strings.Split(line, ",")

	manual := make([]int, len(parts))
	for i, part := range parts {
		pageNumber, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		}

		manual[i] = pageNumber
	}

	return manual
}

func calculateSumOfMiddlePageNumbers(manuals [][]int) int {
	total := 0

	for _, manual := range manuals {
		total += manual[(len(manual)-1)/2]
	}

	return total
}

func isManualValid(manual []int, rules [][]int) bool {
	relevantRules := filterRulesForManual(manual, rules)

	for _, rule := range relevantRules {
		if !isPageNumberBefore(manual, rule[0], rule[1]) {
			return false
		}
	}

	return true
}

func isPageNumberBefore(manual []int, first int, second int) bool {
	var firstIndex, secondIndex int
	for i, pageNumber := range manual {
		if pageNumber == first {
			firstIndex = i
		}

		if pageNumber == second {
			secondIndex = i
		}
	}

	return firstIndex < secondIndex
}

func filterRulesForManual(manual []int, rules [][]int) [][]int {
	result := make([][]int, 0)

	for _, rule := range rules {
		containsLeft := false
		containsRight := false

		for _, pageNumber := range manual {
			if rule[0] == pageNumber {
				containsLeft = true
			}

			if rule[1] == pageNumber {
				containsRight = true
			}
		}

		if containsLeft && containsRight {
			result = append(result, rule)
		}
	}

	return result
}
