package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	inputSize := 1000

	left := make([]int, 0, inputSize)
	right := make([]int, 0, inputSize)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		leftVal, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}

		left = append(left, leftVal)

		rightVal, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		right = append(right, rightVal)
	}

	// output := calculateTotalDistance(left, right)
	output := calculateTotalSimilarityScore(left, right)

	log.Println(output)
}

func calculateTotalDistance(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		totalDistance += max(left[i], right[i]) - min(left[i], right[i])
	}

	return totalDistance
}

func calculateTotalSimilarityScore(left []int, right []int) int {
	totalSimilarityScore := 0

	for i := 0; i < len(left); i++ {
		frequencyOfLeftItemInRightSlice := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				frequencyOfLeftItemInRightSlice++
			}
		}

		totalSimilarityScore += left[i] * frequencyOfLeftItemInRightSlice
	}

	return totalSimilarityScore
}
