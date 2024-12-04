package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(raw)

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	regex := regexp.MustCompile(pattern)
	muls := regex.FindAllString(input, -1)

	total := 0

	for _, mul := range muls {
		regex = regexp.MustCompile("[0-9]+")
		values := regex.FindAllString(mul, -1)

		a, _ := strconv.Atoi(values[0])
		b, _ := strconv.Atoi(values[1])

		total += a * b
	}

	log.Println(total)
}
