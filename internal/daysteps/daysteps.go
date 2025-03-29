package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"yp-go1fl/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	sData := strings.Split(data, ",")
	if len(sData) != 2 {
		return 0, 0, errors.New("")
	}

	steps, err := strconv.Atoi(sData[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("count steps %d < 0", steps)
	}

	duration, err := time.ParseDuration(sData[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * StepLength
	distance = distance / 1000

	spentCalories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, spentCalories)
}
