package serverservice

import (
	"fmt"
	"time"

	"github.com/frit2000/go_final_project/nextdate"
	"github.com/frit2000/go_final_project/params"
	"github.com/frit2000/go_final_project/servicetask"
)

func checkFieldsTask(task *servicetask.Task) error {
	if task.Title == "" {
		return fmt.Errorf("не указан заголовок задачи")
	}

	if task.Date == "" {
		task.Date = time.Now().Format(params.DFormat)
		return nil
	}
	_, err := time.Parse(params.DFormat, task.Date)
	if err != nil {
		return fmt.Errorf("дата неверного формата")
	}

	newDate := time.Now().Format(params.DFormat)
	err = nil
	if task.Repeat != "" {
		newDate, err = nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	}

	if task.Date < time.Now().Format(params.DFormat) {
		task.Date = newDate
	}

	return err
}
