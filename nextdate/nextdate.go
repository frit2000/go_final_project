package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {

	dateInTimeFormat, err := time.Parse("20060102", date)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга формата заданной даты: %w", err)
	}
	if repeat == "" {
		return "", fmt.Errorf("ошибка формата repeat, пустая строка")
	}

	// if date < now.Format("20060102") {
	// 	date = now.Format("20060102")
	// }

	// if repeat == nil {
	// 	return now.Format("20060102"), nil
	// }
	switch repeat[0] {
	case 'd':
		return addDays(now, dateInTimeFormat, repeat)
	case 'y':
		return addYear(now, dateInTimeFormat)
	default:
		return "", fmt.Errorf("ошибка формата repeat, первый символ не допустим")
	}

}

func addDays(now time.Time, dateInTimeFormat time.Time, repeat string) (string, error) {
	daySlice := strings.Split(repeat, " ")
	if len(daySlice) != 2 {
		return "", fmt.Errorf("ошибка формата repeat, количество дней не задано")
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
	for dateInTimeFormat.Format("20060102") <= now.Format("20060102") {
		dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, dayCount)
	}

	return dateInTimeFormat.Format("20060102"), nil
}

func addYear(now time.Time, dateInTimeFormat time.Time) (string, error) {
	//устанавливаем новую дату для задачи
	dateInTimeFormat = dateInTimeFormat.AddDate(1, 0, 0)
	for dateInTimeFormat.Format("20060102") <= now.Format("20060102") {
		dateInTimeFormat = dateInTimeFormat.AddDate(1, 0, 0)
	}
	return dateInTimeFormat.Format("20060102"), nil
}
