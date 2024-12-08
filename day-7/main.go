package main

import (
	"aoc24/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Equation struct {
	result int
	values []int
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

var equations = []Equation{}

func main() {
	input := util.OpenFileAsStringSlice("input.txt")
	for _, raw := range input {
		equation := Equation{values: []int{}}
		parts := strings.Split(raw, ": ")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Panic(err)
		}

		equation.result = result

		values := strings.Split(parts[1], " ")
		for _, s := range values {
			value, err := strconv.Atoi(s)
			if err != nil {
				log.Panic(err)
			}
			equation.values = append(equation.values, value)
		}

		equations = append(equations, equation)
	}

	sumOfPossibleEquations := 0
	for _, e := range equations {
		if e.possible() {
			sumOfPossibleEquations += e.result
		}
	}

	fmt.Printf("Sum of possible equations: %v", sumOfPossibleEquations)
}

func (e Equation) possible() bool {
	fmt.Printf("\nTrying for %v\n", e.result)

	tree := e.buildTree(1, e.values[0])
	possible := tree.hasResult(e.result)
	if possible {
		fmt.Printf("%v is possible\n", e.result)
	}
	return possible
}

func (e Equation) buildTree(depth int, currentValue int) *Node {
	if depth >= len(e.values) {
		return &Node{Value: currentValue}
	}

	node := &Node{Value: currentValue}

	multiplyValue := currentValue * e.values[depth]
	addValue := currentValue + e.values[depth]

	node.Left = e.buildTree(depth+1, multiplyValue)
	node.Right = e.buildTree(depth+1, addValue)

	return node
}

func (n *Node) hasResult(result int) bool {
	if n == nil {
		return false
	}

	if n.Value == result {
		return true
	}

	if n.Left.hasResult(result) {
		return true
	}

	if n.Right.hasResult(result) {
		return true
	}

	return false
}
