package day0302

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

func SolvePart2() int {
	fmt.Println("Solving...")
	enableOperation := true
	mulOperations := getInstructions(input)
	acc := 0
	for _, o := range mulOperations {

		if o == "don't()" {
			enableOperation = false
			continue
		}

		if o == "do()" {
			enableOperation = true
			continue
		}

		if !enableOperation || !strings.HasPrefix(o, "mul") {
			continue
		}

		multiples := getMultiples(o)
		acc += multiples[0] * multiples[1]
	}

	fmt.Println("========")
	fmt.Println(acc)
	fmt.Println("========")

	return acc
}

func getInstructions(input string) []string {
	r, _ := regexp.Compile("mul\\([\\d]{1,3},[\\d]{1,3}\\)|don't\\(\\)|do\\(\\)")
	return r.FindAllString(input, -1)
}

func getMultiples(mulOperation string) [2]int {
	r, _ := regexp.Compile("[\\d]{1,3}")
	multiples := r.FindAllString(mulOperation, -1)
	m1, _ := strconv.Atoi(multiples[0])
	m2, _ := strconv.Atoi(multiples[1])
	return [2]int{m1, m2}
}
