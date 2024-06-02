package storage

import (
	"database/sql"
	"fmt"

	"github.com/codadept/schdl/internal/util"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(dbFile string) (*Storage, error) {
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		due_date DATETIME,
		priority INTEGER
	);
	`

	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, fmt.Errorf("failed to create tasks table: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) AddTask(task *util.Task) (int, error) {
	query := `INSERT INTO tasks (title, description, due_date, priority) VALUES (?, ?, ?, ?)`
	res, err := s.db.Exec(query, task.Title, task.Description, task.DueDate, task.Priority)
	if err != nil {
		return -1, fmt.Errorf("failed to add task: %w", err)
	}
	return util.ConvertInt64ToInt(res.LastInsertId)
}

func (s *Storage) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (s *Storage) UpdateTask(task *util.Task) error {
	query := `UPDATE tasks SET title = ?, description = ?, due_date = ?, priority = ? WHERE id = ?`
	_, err := s.db.Exec(query, task.Title, task.Description, task.DueDate, task.Priority, task.ID)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (s *Storage) GetTask(id int) (*util.Task, error) {
	query := `SELECT id, title, description, due_date, priority FROM tasks WHERE id = ?`
	row := s.db.QueryRow(query, id)

	task := &util.Task{}
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No task found
		}
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return task, nil
}

func (s *Storage) ListTasks() ([]util.Task, error) {
	query := `SELECT id, title, description, due_date, priority FROM tasks`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	defer rows.Close()

	var tasks []util.Task
	for rows.Next() {
		task := util.Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
