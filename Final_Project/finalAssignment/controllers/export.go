package controllers

import (
	"final/services"
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type ExportController struct {
	exportService *services.ExportService
}

func NewExportController(es *services.ExportService) *ExportController {
	return &ExportController{exportService: es}
}

func (ec *ExportController) ExportFile(c echo.Context) error {
	user := c.Get("user")
	file, err := ec.exportService.CreateFile(user)
	if err != nil {
		return err
	}

	thepath, err := filepath.Abs(filepath.Dir(file.Name()))
	if err != nil {
		return err
	}

	fileStat, err := os.Stat(thepath)
	if err != nil {
		return err
	}

	c.Response().Header().Set("Content-Disposition", "attachment; filename="+file.Name())
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", fileStat.Size()))

	return c.File(file.Name())
}
