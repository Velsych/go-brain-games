package even

import (
	"math/rand"
	"strconv"
)



const(
	RULE = "Отвечай `да` если число чётное, `нет` если нет XD"
	MAX_GEN_NUMBER = 50
	MIN_GEN_NUMBER = 5
)

func GenerateQuestionAndAnswer() (string,string) {
	number := rand.Intn(MAX_GEN_NUMBER-MIN_GEN_NUMBER+1)+MIN_GEN_NUMBER
	answer := "нет"
	if number % 2 == 0{
		answer = "да"
	}
	result := strconv.Itoa(number)
	return result,answer
}