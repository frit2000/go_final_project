package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// функция поиска числа для переноса задачи, если надо переносить на заданное количество дней
func addDays(now time.Time, dateInTimeFormat time.Time, repeat string) (string, error) {
	daySlice := strings.Split(repeat, " ")
	if len(daySlice) != 2 {
		return "", fmt.Errorf("ошибка формата repeat, количество дней некорректно задано")
	}
	//переводим в число параметр, на сколько дней надо перенести задачу
	dayCount, err := strconv.Atoi(daySlice[1])
	if err != nil {
		return "", fmt.Errorf("ошибка формата repeat, количество дней не число: %w", err)
	}
	if dayCount > 400 {
		return "", fmt.Errorf("ошибка формата repeat, количество дней больше 400")
	}

	//устанавливаем новую дату для задачи
	dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, dayCount)
	for dateInTimeFormat.Format(dFormat) <= now.Format(dFormat) {
		dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, dayCount)
	}

	return dateInTimeFormat.Format(dFormat), nil
}
