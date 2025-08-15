package personaldata

import "fmt"

// Personal структура содержит имя, вес и рост
type Personal struct {
	// TODO: добавить поля
	Name   string
	Weight float64
	Height float64
}

// метод имени, веса и роста на экран
func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
