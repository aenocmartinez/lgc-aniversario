package exporter

import (
	"fmt"
	"lgc/src/view/dto"

	"github.com/xuri/excelize/v2"
)

func GenerarReporteContadorExcel(data []dto.ReporteContadorInscripcionDTO) (*excelize.File, error) {
	f := excelize.NewFile()
	sheet := "Reporte"
	f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	// Encabezados
	headers := []string{"Inscripción ID", "Forma de Pago", "Monto COP", "Monto USD", "Soporte", "Nombre", "Documento", "Teléfono"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	row := 2
	for _, insc := range data {
		for i, p := range insc.Participantes {
			if i == 0 {
				f.SetCellValue(sheet, fmt.Sprintf("A%d", row), insc.ID)
				f.SetCellValue(sheet, fmt.Sprintf("B%d", row), insc.FormaPago)
				f.SetCellValue(sheet, fmt.Sprintf("C%d", row), insc.MontoPagadoCOP)
				f.SetCellValue(sheet, fmt.Sprintf("D%d", row), insc.MontoPagadoUSD)
				f.SetCellValue(sheet, fmt.Sprintf("E%d", row), insc.SoportePagoURL)
			}

			f.SetCellValue(sheet, fmt.Sprintf("F%d", row), p.NombreCompleto)
			f.SetCellValue(sheet, fmt.Sprintf("G%d", row), p.NumeroDocumento)
			f.SetCellValue(sheet, fmt.Sprintf("H%d", row), p.Telefono)
			row++
		}
	}

	return f, nil
}
