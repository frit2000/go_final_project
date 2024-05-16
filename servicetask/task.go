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
func NewTaskStore(Db *sql.DB) TaskStore {
	return TaskStore{Db: Db}
}

// добавляем в базу задачу, возвращаем номер добавленной записи в виде строки и ошибку
func (ts TaskStore) Add(t *Task) error {
	//записываем поля структуры Task в БД
	res, err := ts.Db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat))
	if err != nil {
		return fmt.Errorf("ошибка при добавлении записи БД: %w", err)
	}

	//получаем ID последней добавленной записи
	lastID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("ошибка получении последнего ID: %w", err)
	}
	t = &Task{}
	t.IdErr = strconv.Itoa(int(lastID))
	return nil
}

// получаем задачу по ID
func (ts TaskStore) GetOneTask(id int) (Task, error) {
	var task = Task{}

	err := ts.Db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return task, fmt.Errorf("ошибка чтении данных по id: %w", err)
	}

	//проверяем, есть ли такой ID задачи
	if len(task.Id) == 0 {
		task.Err = "Ошибка, нет такого ID"
	}

	return task, nil
}

// получаем все задачи из базы
func (ts TaskStore) GetAll() (map[string][]Task, error) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task
	fmt.Println("ПОПАЛИ В GETALL")
	rows, err := ts.Db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit",
		sql.Named("limit", limit))

	if err != nil {
		return tasks, fmt.Errorf("ошибка запроса в базу: %w", err)
	}
	defer rows.Close()

	fmt.Println("ЧТОТО ПРОЧИТАЛИ ИЗ БАЗЫ")

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
	rows, err := ts.Db.Query("SELECT * FROM scheduler WHERE title LIKE :searchString OR comment LIKE :searchString ORDER BY date LIMIT :limit",
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
	rows, err := ts.Db.Query("SELECT * FROM scheduler WHERE date = :searchString LIMIT :limit",
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
func (ts TaskStore) Delete(id int) (Task, error) {
	var task = Task{}
	var checkedID string

	err := ts.Db.QueryRow("SELECT id FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&checkedID)
	if err != nil {
		return task, fmt.Errorf("ошибка чтении данных по id: %w", err)
	}

	//проверяем, есть ли такой ID задачи
	if len(checkedID) == 0 {
		task.Err = "Ошибка, нет такого ID"
	}

	_, err = ts.Db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return task, fmt.Errorf("ошибка при обновлении записи БД: %w", err)
	}

	return task, nil
}

// отметить задачу выполненной
func (ts TaskStore) Done(id int) (Task, error) {
	var task = Task{}

	err := ts.Db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return Task{}, fmt.Errorf("ошибка чтении данных по id: %w", err)
	}

	//проверяем, есть ли такой ID задачи
	if len(task.Id) == 0 {
		task.Err = "Ошибка, нет такого ID"
	}

	if task.Repeat == "" {
		_, err = ts.Db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
		if err != nil {
			return Task{}, fmt.Errorf("ошибка при обновлении записи БД: %w", err)
		}
	} else {
		newDate, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			return Task{}, fmt.Errorf("ошибка при вычислении новой даты: %w", err)
		}
		_, err = ts.Db.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
			sql.Named("date", newDate),
			sql.Named("id", task.Id))
		if err != nil {
			return Task{}, fmt.Errorf("ошибка при обновлении записи БД: %w", err)
		}
	}
	return task, nil
}

// обновитьзадачу
func (ts TaskStore) Update(t *Task) error {
	var task = Task{}

	//обновляем поля структуры task в БД
	result, err := ts.Db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
		sql.Named("id", t.Id))
	if err != nil {
		return fmt.Errorf("ошибка при обновлении записи БД: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при обновлении записи БД: %w", err)
	}
	if rowsAffected == 0 {
		task = Task{}
		task.Err = "задача не найдена"
	}

	return nil
}
