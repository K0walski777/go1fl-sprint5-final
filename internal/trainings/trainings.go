package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	sliceForDatastr := strings.Split(datastring, ",")

	if len(sliceForDatastr) != 3 {
		return fmt.Errorf("Строка слишком короткая для парсинга")
	}

	steps, err := strconv.Atoi(sliceForDatastr[0])
	if err != nil {
		return fmt.Errorf("Возникла ошибка преобразования шагов: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("Кол-во шагов должно быть больше нуля")
	}

	t.Steps = steps

	trainingType := sliceForDatastr[1]
	t.TrainingType = trainingType

	duration, err := time.ParseDuration(sliceForDatastr[2])
	if err != nil {
		return fmt.Errorf("Возникла ошибка преобразования длительности тренировки: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("Продолжительность тренировки должна быть больше нуля")
	}

	t.Duration = duration

	return err
}

func (t Training) ActionInfo() (string, error) {
	distanсe := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	durationHours := t.Duration.Hours()

	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		durationHours,
		distanсe,
		meanSpeed,
		calories,
	)

	return result, err
}
