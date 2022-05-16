package persistence

import (
	"database/sql"
	"final/models"
	"fmt"
	"log"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTasks(listId int64) ([]models.Task, error) {
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
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *TaskRepository) DeleteTask(id int64) error {
	deleteQuery := `DELETE FROM tasks WHERE ID=$1`
	_, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("error deleting %s", err)
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) UpdateTask(id int64) (models.Task, error) {
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
		return task, err
	}

	return task, nil
}

func (r *TaskRepository) GetTask(id int64) (models.Task, error) {
	task := models.Task{}
	query := `SELECT * FROM tasks WHERE id=$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return task, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Text, &task.ListID, &task.Completed)
		if err != nil {
			return task, err
		}
	}

	err = rows.Err()
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepository) InsertTask(task models.Task) (int64, error) {
	var id int64
	insertQuery := `INSERT INTO tasks(Text, ListId, Completed) VALUES($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(insertQuery, task.Text, task.ListID, task.Completed).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting new Task: %w", err)
	}

	return id, nil
}
