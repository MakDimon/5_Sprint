package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	rsc, err := RunningSpentCalories(steps, weight, height, duration)
	return walkingCaloriesCoefficient * rsc, err
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0.0, fmt.Errorf("ошибка в количестве шагов")
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("ошибка в весе пользователя")
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("ошибка в росте пользователя")
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("ошибка в длительности занятия")
	}
	averSpeed := MeanSpeed(steps, height, duration)
	return weight * averSpeed * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 { //проверка на больше 0
		return 0.0
	}
	return Distance(steps, height) / duration.Hours() // Вычисляем среднюю скорость, переводим в часы
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return stepLengthCoefficient * height * float64(steps) / mInKm
}
