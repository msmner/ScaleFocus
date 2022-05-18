package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockExportService struct{}

func (es *MockExportService) CreateFile(username interface{}) (*os.File, error) {
	csvFile, err := os.Create("todo.csv")
	if err != nil {
		return nil, fmt.Errorf("error creating csv file: %w", err)
	}
	return csvFile, nil
}

func TestExportFile(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockService := MockExportService{}
	exportController := NewExportController(&mockService)
	if assert.NoError(t, exportController.ExportFile(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}
