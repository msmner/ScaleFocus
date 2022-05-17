package interfaces

import "os"

type IExportService interface {
	CreateFile(username interface{}) (*os.File, error)
}
