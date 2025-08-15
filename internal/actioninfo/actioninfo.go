package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, value := range dataset { //перебираем значения слайса dataset
		err := dp.Parse(value) //парсим
		if err != nil {        //обработаем ошибку, если найдем, то в лог ее
			log.Println(err)
		} else {
			a, err := dp.ActionInfo() //запускаем метод ActionInfo, который сформирует строку для вывода
			if err != nil {           //обработаем ошибку, если нету- выведем на экран
				log.Println(err)
			} else {
				fmt.Println(a)
			}
		}
	}
}
