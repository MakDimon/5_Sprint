package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	if len(datastring) < 6 {
		return fmt.Errorf("количество символов меньше ожидаемого") //минимально возможно 6 символов 1,0h1m
	}
	parts := strings.Split(datastring, ",") //делим строку на шаги, активность и время
	if len(parts) != 2 {
		return fmt.Errorf("количество разделителей не соответствует ожидаемому")
	}
	steps, err := strconv.Atoi(parts[0]) // число количество шагов
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании количества шагов")
	}
	if steps <= 0 { //шагов должно быть больше 0
		return fmt.Errorf("ошибка подсчета количества шагов")
	}
	dur, err := time.ParseDuration(parts[1]) //преобразуем время
	if err != nil {
		return fmt.Errorf("ошибка преобразования времени")
	}
	if dur <= 0 {
		return fmt.Errorf("неверно задано время")
	}
	ds.Steps = steps
	ds.Duration = dur
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cal), nil
}
