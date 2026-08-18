package calc

import (
	"math/rand"
	"strconv"
	"strings"
)

const (
	Rule           = "What is the result of the expression?"
	MAX_GEN_NUMBER = 50
	MIN_GEN_NUMBER = 5
)

func calculate(number1 int, number2 int, operation string) string {
	answer := 0
	switch operation {
	case "*":
		answer = number1 * number2
	case "+":
		answer = number1 + number2
	case "-":
		answer = number1 - number2
	}
	result := strconv.Itoa(answer)
	return result
}

func Generate_question_and_answer() (string, string) {
	var sb strings.Builder
	operations := []string{"*", "+", "-"}
	number1 := rand.Intn(MAX_GEN_NUMBER-MIN_GEN_NUMBER+1) + MIN_GEN_NUMBER
	number2 := rand.Intn(MAX_GEN_NUMBER-MIN_GEN_NUMBER+1) + MIN_GEN_NUMBER
	randomIndex := rand.Intn(len(operations))
	operation := operations[randomIndex]
	answer := calculate(number1, number2, operation)
	sb.WriteString(strconv.Itoa(number1))
	sb.WriteString(operation)
	sb.WriteString(strconv.Itoa(number2))
	return sb.String(), answer

}
