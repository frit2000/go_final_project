package nextdate

import "time"

// функция поиска числа для переноса задачи, если надо переносить на год
func addYear(now time.Time, dateInTimeFormat time.Time) (string, error) {
	//устанавливаем новую дату для задачи
	dateInTimeFormat = dateInTimeFormat.AddDate(1, 0, 0)
	for dateInTimeFormat.Format(dFormat) <= now.Format(dFormat) {
		dateInTimeFormat = dateInTimeFormat.AddDate(1, 0, 0)
	}
	return dateInTimeFormat.Format(dFormat), nil
}
