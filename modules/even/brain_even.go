package even

import "math/rand"



const(
	Rule = "Отвечай `да` если число чётное, `нет` если нет XD"
	MAX_GEN_NUMBER = 50
	MIN_GEN_NUMBER = 5
)

func Generate_question_and_answer() (int,string) {
	number := rand.Intn(MAX_GEN_NUMBER-MIN_GEN_NUMBER+1)+MIN_GEN_NUMBER
	answer := "нет"
	if number % 2 == 0{
		answer = "да"
	}
	return number,answer
}