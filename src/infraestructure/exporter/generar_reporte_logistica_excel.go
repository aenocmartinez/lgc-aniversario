package exporter

import (
	"bytes"
	"fmt"
	"lgc/src/view/dto"

	"github.com/xuri/excelize/v2"
)

func GenerarReporteLogisticaExcel(data []dto.ReporteLogisticaDTO) ([]byte, error) {
	f := excelize.NewFile()
	sheet := "Logistica"
	f.NewSheet(sheet)

	headers := []string{"#", "NOMBRE COMPLETO", "NUMERO DOCUMENTO", "CORREO ELECTRONICO", "TELEFONO", "DIAS DE ASISTENCIA"}
	for i, h := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheet, cell, h)
	}

	for i, p := range data {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), fmt.Sprintf("%d.", i+1))
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), p.NombreCompleto)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), p.NumeroDocumento)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), p.CorreoElectronico)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), p.Telefono)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), p.DiasAsistencia)
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
