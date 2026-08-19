package prime

import (
	"math/rand"
	"strconv"
)

const(
	RULE = "Answer `да` if given number is prime. Otherwise answer `нет`."
	MAX_GEN_NUMBER = 100
	MIN_GEN_NUMBER = 1
)

func isPrime(number int) bool {
	if number % 2 == 0{
		return false 
	} else if number == 1{
		return false
	}
	for i := 3; i == number ; i += 2 {
		if number % i == 0{
			return false
		}
	}
	return true
}

func GenerateQuestionAndAnswer() (string,string) {
	answer := "да"
	number := rand.Intn(MAX_GEN_NUMBER-MIN_GEN_NUMBER+1) + MIN_GEN_NUMBER
	if number == 2{
		return strconv.Itoa(number),answer
	}
	if isPrime(number) != true{
		answer = "нет"
	}
	return strconv.Itoa(number), answer
}