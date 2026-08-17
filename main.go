package main

import (
	"brain_games/modules/even"

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
	}

}
