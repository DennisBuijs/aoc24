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

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := 0; i < inputSize; i++ {
		totalDistance += max(left[i], right[i]) - min(left[i], right[i])
	}

	log.Println(totalDistance)
}
