package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	rules := [][]int{}
	manuals := [][]int{}

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

	validManuals := [][]int{}
	invalidManuals := [][]int{}

	for _, manual := range manuals {
		if isManualValid(manual, rules) {
			validManuals = append(validManuals, manual)
		} else {
			invalidManuals = append(invalidManuals, manual)
		}
	}

	outputForValidManuals := calculateSumOfMiddlePageNumbers(validManuals)

	fixedInvalidManuals := fixInvalidManuals(invalidManuals, rules)
	outputForFixedInvalidManuals := calculateSumOfMiddlePageNumbers(fixedInvalidManuals)

	log.Println("Valid", outputForValidManuals)
	log.Println("Invalid", outputForFixedInvalidManuals)

	elapsed := time.Since(start)
	log.Printf("Execution time: %s\n", elapsed)
}

func mapLineToRule(line string) []int {
	parts := strings.Split(line, "|")

	var err error

	rule := make([]int, 2)

	rule[0], _ = strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	rule[1], err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

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
	firstIndex, secondIndex := findPageNumberIndices(manual, first, second)

	return firstIndex < secondIndex
}

func filterRulesForManual(manual []int, rules [][]int) [][]int {
	pageSet := make(map[int]bool)
	for _, pageNumber := range manual {
		pageSet[pageNumber] = true
	}

	result := [][]int{}

	for _, rule := range rules {
		if pageSet[rule[0]] && pageSet[rule[1]] {
			result = append(result, rule)
		}
	}

	return result
}

func fixInvalidManuals(invalidManuals [][]int, rules [][]int) [][]int {
	manuals := [][]int{}

	for manualIndex := 0; manualIndex < len(invalidManuals); manualIndex++ {
		invalidManual := invalidManuals[manualIndex]
		relevantRules := filterRulesForManual(invalidManual, rules)

		for i := 0; i < len(relevantRules); i++ {
			rule := relevantRules[i]
			if !isPageNumberBefore(invalidManual, rule[0], rule[1]) {
				invalidManual = swap(invalidManual, rule[0], rule[1])
				i = 0
			}
		}

		if !isManualValid(invalidManual, rules) {
			manualIndex = manualIndex - 1
			continue
		}

		manuals = append(manuals, invalidManual)
	}

	return manuals
}

func swap(manual []int, first int, second int) []int {
	firstIndex, secondIndex := findPageNumberIndices(manual, first, second)

	manual[firstIndex] = second
	manual[secondIndex] = first

	return manual
}

func findPageNumberIndices(manual []int, first int, second int) (int, int) {
	var firstIndex, secondIndex int
	for i, pageNumber := range manual {
		if pageNumber == first {
			firstIndex = i
		}

		if pageNumber == second {
			secondIndex = i
		}
	}

	return firstIndex, secondIndex
}
