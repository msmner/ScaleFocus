package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	mockListRepository := MockListRepository{}
	mockUserRepository := MockUserRepository{}
	mockTaskRepository := MockTaskRepository{}
	exportService := NewExportService(&mockListRepository, &mockTaskRepository, &mockUserRepository)

	_, err := exportService.CreateFile("error")
	assert.EqualError(t, err, "cant get user error")

	_, err = exportService.CreateFile("error2")
	assert.EqualError(t, err, "error getting lists for user error")

	_, err = exportService.CreateFile("test")
	assert.EqualError(t, err, "error getting tasks for list error")
}
