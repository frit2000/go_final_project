package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// функция поиска числа для переноса задачи, если надо переносить на заданные дни недели
func addWeek(now time.Time, dateInTimeFormat time.Time, repeat string) (string, error) {
	daySlice := strings.Split(repeat, " ")
	if len(daySlice) != 2 {
		return "", fmt.Errorf("ошибка формата repeat, дни недели некорректно заданы")
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
	if dateInTimeFormat.Format(dFormat) < now.Format(dFormat) {
		dateInTimeFormat = now
	}

	//если в мапе есть допустимый ключ для переноса дня, заполняем данные для этого ключа правильной строкой с датой
	for i := 1; i <= 7; i++ {
		dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, 1)
		weekDay := weekDayNumber(dateInTimeFormat)
		_, ok := validDays[weekDay]
		if ok {
			validDays[weekDay] = dateInTimeFormat.Format(dFormat)
		}
	}
	//ищем ближайшую дату из мапы
	targetDay := dateInTimeFormat.Format(dFormat)
	for _, validDay := range validDays {
		if validDay < targetDay {
			targetDay = validDay
		}
	}
	return targetDay, nil
}

// Назначаем дням недели порядковые номера, где пн = 1, а вс = 7
func weekDayNumber(day time.Time) int {
	weekDay := day.Weekday()
	if int(weekDay) == 0 {
		return 7
	}
	return int(weekDay)
}
