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
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	sliceForDatastr := strings.Split(datastring, ",")

	if len(sliceForDatastr) != 2 {
		return fmt.Errorf("Строка слишком короткая для парсинга")
	}

	steps, err := strconv.Atoi(sliceForDatastr[0])
	if err != nil {
		return fmt.Errorf("Возникла ошибка преобразования шагов: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("Кол-во шагов должно быть больше нуля")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(sliceForDatastr[1])
	if err != nil {
		return fmt.Errorf("Возникла ошибка преобразования длительности тренировки: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("Продолжительность тренировки должна быть больше нуля")
	}

	ds.Duration = duration

	return err
}

func (ds DaySteps) ActionInfo() (string, error) {
	var err error

	distanсe := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distanсe,
		spentCalories,
	)

	return result, err
}
