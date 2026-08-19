package main

import (
	"brain_games/games/calc"
	"brain_games/games/even"
	"brain_games/games/gcd"
	"brain_games/games/prime"
	"brain_games/core"
)

const NUMBER_OF_QUESTIONS = 3
type GameName uint8
const(
	Even GameName = iota
	Calc
	Prime
	GCD
)

func main() {
	name := core.GreetUser()
	game_number := core.Choose()
	switch game_number {
	case int(Even):
		core.Start(even.GenerateQuestionAndAnswer, even.RULE, NUMBER_OF_QUESTIONS, name)
	case int(Calc):
		core.Start(calc.GenerateQuestionAndAnswer, calc.RULE, NUMBER_OF_QUESTIONS, name)
	case int(Prime):
		core.Start(prime.GenerateQuestionAndAnswer, prime.RULE, NUMBER_OF_QUESTIONS, name)
	case int(GCD):
		core.Start(gcd.GenerateQuestionAndAnswer, gcd.RULE, NUMBER_OF_QUESTIONS, name)
	}

}
