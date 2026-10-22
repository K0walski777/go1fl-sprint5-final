package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	ErrNegativeSteps   = errors.New("Количество шагов не может быть отрицательным")
	ErrNonPositiveTime = errors.New("Продолжительность должна быть больше нуля")
)

// Расчет калорий, потраченых при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, ErrNegativeSteps
	}

	if weight <= 0 {
		return 0, errors.New("Вес должен быть больше нуля")
	}

	if height <= 0 {
		return 0, errors.New("Рост должен быть больше нуля")
	}

	if duration <= 0 {
		return 0, ErrNonPositiveTime
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()

	calories := (weight * meanSpeed * durationInMin) / minInH

	return calories * walkingCaloriesCoefficient, nil
}

// Расчет калорий, потраченых при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, ErrNegativeSteps
	}

	if weight <= 0 {
		return 0, errors.New("Вес должен быть больше нуля")
	}

	if height <= 0 {
		return 0, errors.New("Рост должен быть больше нуля")
	}

	if duration <= 0 {
		return 0, ErrNonPositiveTime
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()

	calories := (weight * meanSpeed * durationInMin) / minInH

	return calories, nil
}

// Расчет средней скорости в км/ч
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 || height <= 0 {
		return 0
	}

	dist := Distance(steps, height)

	return dist / duration.Hours()

}

// Расчет дистанции в км. на основе шагов и роста
func Distance(steps int, height float64) float64 {
	if steps < 0 || height <= 0 {
		return 0
	}

	return ((height * stepLengthCoefficient) * float64(steps)) / mInKm
}
