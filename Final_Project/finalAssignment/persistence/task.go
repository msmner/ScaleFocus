package persistence

import (
	"database/sql"
	"final/models"
	"fmt"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTasks(listId int) ([]models.Task, error) {
	tasks := []models.Task{}
	rows, err := r.db.Query("select * from tasks where listId=$1", listId)
	if err != nil {
		return tasks, fmt.Errorf("error getting tasks from the database: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.ID, &task.Text, &task.ListID, &task.Completed)
		if err != nil {
			return tasks, fmt.Errorf("error in rows getting tasks: %w", err)
		}
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return tasks, fmt.Errorf("error in rows getting tasks: %w", err)
	}
	return tasks, nil
}

func (r *TaskRepository) DeleteTask(id int) error {
	deleteQuery := `DELETE FROM tasks WHERE ID=$1`
	_, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}

	return nil
}

func (r *TaskRepository) UpdateTask(id int) (models.Task, error) {
	task, err := r.GetTask(id)
	if err != nil {
		return task, err
	}

	if task.Completed {
		task.Completed = false
	} else {
		task.Completed = true
	}

	query := `UPDATE tasks SET completed=$1 WHERE id=$2`
	_, err = r.db.Exec(query, task.Completed, id)
	if err != nil {
		return task, fmt.Errorf("error updating task: %w", err)
	}

	return task, nil
}

func (r *TaskRepository) GetTask(id int) (models.Task, error) {
	task := models.Task{}
	query := `SELECT * FROM tasks WHERE id=$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return task, fmt.Errorf("error getting task: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Text, &task.ListID, &task.Completed)
		if err != nil {
			return task, fmt.Errorf("error in rows getting task: %w", err)
		}
	}

	err = rows.Err()
	if err != nil {
		return task, fmt.Errorf("error in rows getting task: %w", err)
	}
	return task, nil
}

func (r *TaskRepository) InsertTask(task models.Task) (int, error) {
	var id int
	insertQuery := `INSERT INTO tasks(Text, ListId, Completed) VALUES($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(insertQuery, task.Text, task.ListID, task.Completed).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting new Task: %w", err)
	}

	return id, nil
}
