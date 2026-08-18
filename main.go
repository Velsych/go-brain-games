package main

import (
	"brain_games/modules/calc"
	"brain_games/modules/even"
	"brain_games/modules/gcd"
	"brain_games/modules/prime"
	"brain_games/core"
)

const (
	NUMBER_OF_QUESTIONS = 3
)

func main() {
	name := core.GreetUser()
	game_number := core.Choose()
	if game_number == 0 {
		return
	}
	switch game_number {
	case 1:
		core.Start(even.Generate_question_and_answer, even.Rule, NUMBER_OF_QUESTIONS, name)
	case 2:
		core.Start(calc.Generate_question_and_answer, calc.Rule, NUMBER_OF_QUESTIONS, name)
	case 3:
		core.Start(prime.Generate_question_and_answer, prime.Rule, NUMBER_OF_QUESTIONS, name)
	case 4:
		core.Start(gcd.Generate_question_and_answer, gcd.Rule, NUMBER_OF_QUESTIONS, name)
	}

}
