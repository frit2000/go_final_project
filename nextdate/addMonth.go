package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// функция поиска числа для переноса задачи, если надо переносить на числа месяца
func addMonth(now time.Time, dateInTimeFormat time.Time, repeat string) (string, error) {
	monthSlice := strings.Split(repeat, " ")
	if len(monthSlice) != 2 {
		return "", fmt.Errorf("ошибка формата repeat, дни в месяце некорректно заданы")
	}
	monthDaySlice := strings.Split(monthSlice[1], ",")
	//делаем мапу доспустимых значений дней, на которые можно перенести задачу
	validDays := make(map[int]string)
	for _, day := range monthDaySlice {
		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return "", fmt.Errorf("ошибка формата repeat, день недели не число")
		}
		if dayInt > 31 || dayInt < -2 || dayInt == 0 {
			return "", fmt.Errorf("ошибка формата repeat, день задан некорректным числом")
		}
		validDays[dayInt] = day
	}

	//проверяем, что считаем от даты которая дальше чем сегодня, иначе считаем с сегодня
	if dateInTimeFormat.Format("20060102") < now.Format("20060102") {
		dateInTimeFormat = now
	}

	//	amountOfDaysInMonth := daysInMonth(int(dateInTimeFormat.Month()), int(dateInTimeFormat.Year()))
	amountOfDaysInMonth := daysInMonth(dateInTimeFormat)
	//если в мапе есть допустимый ключ для переноса дня, заполняем данные для этого ключа правильной строкой с датой
	for i := 1; i <= 31; i++ {
		//		   dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, 1)
		_, ok := validDays[i]
		if ok {
			if (i <= int(dateInTimeFormat.Day())) || (i > amountOfDaysInMonth) {
				varDate := dateInTimeFormat.AddDate(0, 1, 0)
				validDays[i] = varDate.Format("200601") + fmt.Sprintf("%02d", i)
			} else {
				validDays[i] = dateInTimeFormat.Format("200601") + fmt.Sprintf("%+02d", i)
			}
		}
	}

	if _, ok := validDays[-1]; ok {
		validDays[-1] = dateInTimeFormat.AddDate(0, 1, -dateInTimeFormat.Day()).Format("20060102")
		fmt.Println("-1, validdays=", validDays[-1])
	}

	if _, ok := validDays[-2]; ok {
		validDays[-2] = dateInTimeFormat.AddDate(0, 1, -dateInTimeFormat.Day()-1).Format("20060102")
		fmt.Println("-2, validdays=", validDays[-2])
	}
	fmt.Println("validDays=", validDays)

	//ищем ближайшую дату из мапы
	targetDay := dateInTimeFormat.AddDate(1, 0, 0).Format("20060102")
	for _, validDay := range validDays {
		if validDay < targetDay {
			targetDay = validDay
		}
	}
	return targetDay, nil
}

func daysInMonth(date time.Time) int {
	return int(date.AddDate(0, 1, -date.Day()).Day())
}
