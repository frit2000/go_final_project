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

	switch repeat[0] {
	case 'd':
		return addDays(now, dateInTimeFormat, repeat)
	case 'y':
		return addYear(now, dateInTimeFormat)
	case 'w':
		return addWeek(now, dateInTimeFormat, repeat)
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

func addWeek(now time.Time, dateInTimeFormat time.Time, repeat string) (string, error) {
	daySlice := strings.Split(repeat, " ")
	if len(daySlice) != 2 {
		return "", fmt.Errorf("ошибка формата repeat, дни недели не заданы")
	}
	weekDaySlice := strings.Split(daySlice[1], ",")
	//делаем мапу доспустимых значений дней, на которые можно перенести задачу
	validDays := make(map[int]string)
	for _, day := range weekDaySlice {
		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return "", fmt.Errorf("ошибка формата repeat, день недели не число")
		}
		if dayInt > 7 || dayInt < 1 {
			return "", fmt.Errorf("ошибка формата repeat, день недели не от 1 до 7")
		}
		validDays[dayInt] = day
	}

	//проверяем, что считаем от даты которая дальше чем сегодня, иначе считаем с сегодня
	if dateInTimeFormat.Format("20060102") < now.Format("20060102") {
		dateInTimeFormat = now
	}

	//если в мапе есть допустимый ключ для переноса дня, переносим на этот день
	for i := 1; i <= 7; i++ {
		dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, 1)
		weekDay := weekDay(dateInTimeFormat)
		_, ok := validDays[weekDay]
		if ok {
			return dateInTimeFormat.Format("20060102"), nil
		}
	}
	return "", fmt.Errorf("что-то глобально пошло не так в переносе на день недели")
}

func weekDay(day time.Time) int {
	weekDay := day.Weekday()
	if int(weekDay) == 0 {
		return 7
	}
	return int(weekDay)
}
