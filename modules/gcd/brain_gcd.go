package gcd

import (
	"math/rand"
	"strconv"
	"strings"
)

const (
	Rule = "Find the greatest common divisor of given numbers."
	MAX_GEN_NUMBER = 50
	MIN_GEN_NUMBER = 1
)

func find_gcd(number1 int, number2 int) string {
	max_number := max(number1, number2)
	min_number := min(number1, number2)
	for min_number != 0 {
		rest := max_number % min_number
		max_number, min_number = min_number, rest
	}
	return strconv.Itoa(max_number)
}

func Generate_question_and_answer() (string, string) {
	var sb strings.Builder
	number1 := rand.Intn(MAX_GEN_NUMBER-MIN_GEN_NUMBER+1) + MIN_GEN_NUMBER
	number2 := rand.Intn(MAX_GEN_NUMBER-MIN_GEN_NUMBER+1) + MIN_GEN_NUMBER
	answer := find_gcd(number1, number2)
	sb.WriteString(strconv.Itoa(number1))
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(number2))
	return sb.String(), answer
}
