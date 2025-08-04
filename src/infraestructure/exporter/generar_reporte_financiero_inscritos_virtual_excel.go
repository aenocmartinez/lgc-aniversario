package exporter

import (
	"fmt"
	"lgc/src/view/dto"

	"github.com/xuri/excelize/v2"
)

func GenerarReporteFinancieroVirtual(data []dto.ReporteInscritosVirtualDTO) (*excelize.File, error) {
	f := excelize.NewFile()
	sheet := "Reporte"
	f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	// Encabezados ajustados al nuevo DTO
	headers := []string{
		"Forma de Pago", "Modalidad", "Monto USD", "Nombre Completo", "Soporte de Pago",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Escribir los datos
	for row, item := range data {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row+2), item.FormaPago)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row+2), item.Modalidad)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row+2), item.MontoPagadoUSD)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row+2), item.NombreCompleto)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row+2), item.SoportePagoURL)
	}

	return f, nil
}
