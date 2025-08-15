package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training содержит количество шагов, тип тренировки, длительность тренировки, имя, рост и вес
type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	if len(datastring) < 10 {
		return fmt.Errorf("количество символов меньше ожидаемого") //минимально возможно 10 символов: 1,Бег,0h1m
	}
	parts := strings.Split(datastring, ",") //делим строку на шаги, активность и время
	if len(parts) != 3 {
		return fmt.Errorf("количество разделителей не соответствует ожидаемому")
	}
	steps, err := strconv.Atoi(parts[0]) // число количество шагов
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании количества шагов")
	}
	if steps <= 0 { //шагов должно быть больше 0
		return fmt.Errorf("ошибка подсчета количества шагов")
	}
	dur, err := time.ParseDuration(parts[2]) //преобразуем время
	if err != nil {
		return fmt.Errorf("ошибка преобразования времени")
	}
	if dur <= 0 {
		return fmt.Errorf("неверно задано время")
	}
	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = dur
	return nil
}

func (t Training) ActionInfo() (string, error) {
	//здесь должен быть вызов метода расчета дистанции, но я его впихнул в return, чтобы не плодить переменные
	//здесь должен быть вызов метода расчета средней скорости, но я его впихнул в return, чтобы не плодить переменные
	var cal float64
	var err error
	switch t.TrainingType { //в соответствии с видом тренировки, считаем калории
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height), spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), cal), err
}
