package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(raw)

	// replace with count of "don't()"s in the input
	amountOfDisableStatements := 36
	for i := 0; i < amountOfDisableStatements; i++ {
		input = strings.Replace(input, getDisabled(input), "", -1)
	}

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

func getDisabled(str string) string {
	start := "don't()"
	end := "do()"

	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}

	e := strings.Index(str[s+len(start):], end)
	if e == -1 {
		return str[s:]
	}

	return str[s : s+len(start)+e]
}
