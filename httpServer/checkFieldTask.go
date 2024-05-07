package httpServer

import (
	"fmt"
	"time"

	"github.com/frit2000/go_final_project/nextdate"
)

func checkFieldsTask(task *Task) error {
	if task.Title == "" {
		return fmt.Errorf("не указан заголовок задачи")
	}

	if task.Date == "" {
		task.Date = time.Now().Format("20060102")
		return nil
	}
	_, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("дата неверного формата")
	}

	newDate := time.Now().Format("20060102")
	err = nil
	if task.Repeat != "" {
		newDate, err = nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	}

	if task.Date < time.Now().Format("20060102") {
		task.Date = newDate
	}

	return err
}
