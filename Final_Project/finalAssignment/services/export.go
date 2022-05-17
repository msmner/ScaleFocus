package services

import (
	"encoding/csv"
	"final/persistence"
	"fmt"
	"os"
)

type ExportService struct {
	listRepository *persistence.ListRepository
	taskRepository *persistence.TaskRepository
	userRepository *persistence.UserRepository
}

func NewExportService(lr *persistence.ListRepository, tr *persistence.TaskRepository, ur *persistence.UserRepository) *ExportService {
	return &ExportService{listRepository: lr, taskRepository: tr, userRepository: ur}
}

func (es *ExportService) CreateFile(username interface{}) (*os.File, error) {
	user, err := es.userRepository.GetUser(username.(string))
	if err != nil {
		return nil, fmt.Errorf("cant get user %w", err)
	}

	csvFile, err := os.Create("todo.csv")
	if err != nil {
		return nil, fmt.Errorf("error creating csv file: %w", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	lists, err := es.listRepository.GetLists(user)
	if err != nil {
		return nil, fmt.Errorf("error getting lists for user %w", err)
	}

	taskData := make([][]string, len(lists))
	for i, list := range lists {
		tasks, err := es.taskRepository.GetTasks(list.ID)
		if err != nil {
			return nil, fmt.Errorf("error getting tasks for list %w", err)
		}

		taskData[i] = make([]string, len(tasks))
		for j, task := range tasks {
			taskData[i][j] = task.Text
		}
	}

	for _, taskRow := range taskData {
		err := csvwriter.Write(taskRow)
		if err != nil {
			return nil, fmt.Errorf("error writing tasks to csv %w", err)
		}
	}

	csvwriter.Flush()
	csvFile.Close()

	return csvFile, nil
}
