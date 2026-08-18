package core

import "fmt"

func GreetUser() string {
	var name string
	fmt.Println("Ну здарова бродяга!")
	fmt.Printf("Как звать тебя то? ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return "Error has occured"
	}
	fmt.Println("Ну здарова " + name + "!")
	return name
}

func print_defeat_message(user_answer string, answer string, user_name string) {
	fmt.Printf("Ты проебал лошара! Твой ответ: %s. А правильный: %s\n",user_answer,answer)
}

func congrats()  {
	fmt.Println("Поздравляю, ты выйграл в игре, хз возьми с полки пирожок")
}

func Choose() int {
	var game int
	fmt.Println("А теперь, выбери игру, пиши только число:")
	fmt.Printf("1.Игра на чётное/нечёное\n2.Простая математика\n3.Простое ли число\n4.Найти общий множитель\n")
	fmt.Printf("Твой выбор? ")
	_, err := fmt.Scanln(&game)
	if err != nil {
		fmt.Println("НАПИШИ ЧИСЛО ЁПТА ОТ 1 ДО 4! БЕСИИИШЬ")
		fmt.Println()
		return Choose()
	}
	if game > 4 {
		fmt.Println("НАПИШИ ЧИСЛО ЁПТА ОТ 1 ДО 4! БЕСИИИШЬ")
		fmt.Println()
		return Choose()
	}
	return game
}

func Start(game func() (string, string), rule string, questions int, user_name string) bool {
	var user_answer string
	fmt.Println(rule)
	for i := 0; i < questions; i++ {
		number, answer := game()
		fmt.Println("Вопрос: ", number)
		fmt.Printf("Твой ответ? ")
		fmt.Scanln(&user_answer)
		if answer != user_answer {
			print_defeat_message(user_answer, answer, user_name)
			return false
		}
		fmt.Println("Малаца, правильно\n")
	}
	congrats()
	return true
}
