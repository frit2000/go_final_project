package servicetask

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/frit2000/go_final_project/nextdate"
	"github.com/frit2000/go_final_project/params"
)

const limit = 15

// инициализация базы
func NewTaskStore(db *sql.DB) TaskStore {
	return TaskStore{db: db}
}

// добавляем в базу задачу, возвращаем номер добавленной записи в виде строки и ошибку
func (ts TaskStore) Add(t Task) (string, error) {
	//записываем поля структуры Task в БД
	res, err := ts.db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat))
	if err != nil {
		return "", fmt.Errorf("ошибка при добавлении записи БД: %w", err)
	}

	//получаем ID последней добавленной записи
	lastID, err := res.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("ошибка получении последнего ID: %w", err)
	}

	return strconv.Itoa(int(lastID)), nil
}

// получаем задачу по ID
func (ts TaskStore) GetOneTask(id int) (Task, params.RespTaskError, error) {
	var task Task
	var respTask params.RespTaskError

	err := ts.db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return Task{}, params.RespTaskError{}, fmt.Errorf("ошибка чтении данных по id: %w", err)
	}
	//проверяем, есть ли такой ID задачи
	if len(task.Id) == 0 {
		respTask.Err = "Ошибка, нет такого ID"
	}

	return task, respTask, nil
}

// получаем все задачи из базы
func (ts TaskStore) GetAll() (map[string][]Task, error) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task
	rows, err := ts.db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit",
		sql.Named("limit", limit))

	if err != nil {
		return tasks, fmt.Errorf("ошибка запроса в базу: %w", err)

	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, fmt.Errorf("ошибка разбора строк после чтения из базы: %w", err)
		}
		tasks["tasks"] = append(tasks["tasks"], task)
	}
	return tasks, nil
}

// получаем задачи если в поиске строка (не дата)
func (ts TaskStore) GetSearch(searchString string) (map[string][]Task, error) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task
	rows, err := ts.db.Query("SELECT * FROM scheduler WHERE title LIKE :searchString OR comment LIKE :searchString ORDER BY date LIMIT :limit",
		sql.Named("searchString", "%"+searchString+"%"),
		sql.Named("limit", limit))

	if err != nil {
		return tasks, fmt.Errorf("ошибка запроса в базу: %w", err)

	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, fmt.Errorf("ошибка разбора строк после чтения из базы: %w", err)
		}
		tasks["tasks"] = append(tasks["tasks"], task)
	}
	return tasks, nil
}

// получаем задачи если в поиске дата
func (ts TaskStore) GetSearchDate(searchDate time.Time) (map[string][]Task, error) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task
	rows, err := ts.db.Query("SELECT * FROM scheduler WHERE date = :searchString LIMIT :limit",
		sql.Named("searchString", searchDate.Format(params.DFormat)),
		sql.Named("limit", limit))

	if err != nil {
		return tasks, fmt.Errorf("ошибка запроса в базу: %w", err)

	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, fmt.Errorf("ошибка разбора строк после чтения из базы: %w", err)
		}
		tasks["tasks"] = append(tasks["tasks"], task)
	}
	return tasks, nil
}

// удаляем задачу
func (ts TaskStore) Delete(id int) (params.RespTaskError, error) {
	var checkedID string
	var respTask params.RespTaskError

	err := ts.db.QueryRow("SELECT id FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&checkedID)
	if err != nil {
		return params.RespTaskError{}, fmt.Errorf("ошибка чтении данных по id: %w", err)
	}

	//проверяем, есть ли такой ID задачи
	if len(checkedID) == 0 {
		respTask.Err = "Ошибка, нет такого ID"
	}

	_, err = ts.db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return params.RespTaskError{}, fmt.Errorf("ошибка при обновлении записи БД: %w", err)
	}

	return respTask, nil
}

// отметить задачу выполненной
func (ts TaskStore) Done(id int) (params.RespTaskError, error) {
	var task Task
	var respTask params.RespTaskError

	err := ts.db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return params.RespTaskError{}, fmt.Errorf("ошибка чтении данных по id: %w", err)
	}

	//проверяем, есть ли такой ID задачи
	if len(task.Id) == 0 {
		respTask.Err = "Ошибка, нет такого ID"
	}

	if task.Repeat == "" {
		_, err = ts.db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
		if err != nil {
			return params.RespTaskError{}, fmt.Errorf("ошибка при обновлении записи БД: %w", err)
		}
	} else {
		newDate, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			return params.RespTaskError{}, fmt.Errorf("ошибка при вычислении новой даты: %w", err)
		}
		_, err = ts.db.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
			sql.Named("date", newDate),
			sql.Named("id", task.Id))
		if err != nil {
			return params.RespTaskError{}, fmt.Errorf("ошибка при обновлении записи БД: %w", err)
		}
	}
	return respTask, nil
}

// обновитьзадачу
func (ts TaskStore) Update(t Task) (params.RespTaskError, error) {
	var respTask params.RespTaskError

	//обновляем поля структуры task в БД
	result, err := ts.db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
		sql.Named("id", t.Id))
	if err != nil {
		return params.RespTaskError{}, fmt.Errorf("ошибка при обновлении записи БД: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return params.RespTaskError{}, fmt.Errorf("ошибка при обновлении записи БД: %w", err)
	}
	if rowsAffected == 0 {
		respTask.Err = "задача не найдена"
	}

	return respTask, nil
}
