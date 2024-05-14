package nextdate

import (
	"fmt"
	"time"
)

const dFormat = "20060102"

func NextDate(now time.Time, date string, repeat string) (string, error) {

	dateInTimeFormat, err := time.Parse(dFormat, date)
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
	case 'm':
		return addMonth(now, dateInTimeFormat, repeat)
	default:
		return "", fmt.Errorf("ошибка формата repeat, первый символ не допустим")
	}

}
