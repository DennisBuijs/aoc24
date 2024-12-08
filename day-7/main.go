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
	Value    int
	Children []*Node
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
	tree := e.buildTree(1, e.values[0])
	return tree.hasResult(e.result)
}

func (e Equation) buildTree(depth int, currentValue int) *Node {
	if depth >= len(e.values) {
		return &Node{Value: currentValue}
	}

	node := &Node{Value: currentValue}

	multiplyValue := currentValue * e.values[depth]
	addValue := currentValue + e.values[depth]
	concatValue := concatenateNumbers(currentValue, e.values[depth])

	node.Children = append(node.Children, e.buildTree(depth+1, multiplyValue))
	node.Children = append(node.Children, e.buildTree(depth+1, addValue))
	node.Children = append(node.Children, e.buildTree(depth+1, concatValue))

	return node
}

func (n *Node) hasResult(result int) bool {
	if n == nil {
		return false
	}

	if len(n.Children) == 0 {
		return n.Value == result
	}

	for _, child := range n.Children {
		if child.hasResult(result) {
			return true
		}
	}

	return false
}

func concatenateNumbers(a, b int) int {
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)

	concatenated := sa + sb

	result, err := strconv.Atoi(concatenated)
	if err != nil {
		log.Panic(err)
	}

	return result
}

func (n *Node) print(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("--")
	}

	fmt.Printf("%v\n", n.Value)

	for _, child := range n.Children {
		child.print(depth + 1)
	}
}
