package datasource_test

import (
	"path/filepath"
	"testing"

	"github.com/gevs90/invoice-cfdi-validator/pkg/datasource"
	testdataloader "github.com/peteole/testdata-loader"
	"github.com/stretchr/testify/assert"
)

func TestNewFileExcel(t *testing.T) {
	filename := "myexcel.xls"

	newFileExcel := datasource.NewFileExcel(filename)

	assert.Equal(t, newFileExcel.Filename, filename)
}

func TestFileExcel_Connect(t *testing.T) {
	catalogSatFilePath := filepath.Join(testdataloader.GetBasePath(), "testdata/catCFDI_V_33_code_challenge.xls")
	newFileExcel := datasource.NewFileExcel(catalogSatFilePath)

	err := newFileExcel.Connect()

	assert.Nil(t, err)
}
