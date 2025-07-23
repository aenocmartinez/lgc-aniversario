package usecase

import (
	"lgc/src/domain"
	"lgc/src/infraestructure/exporter"
)

type GenerarReporteRelacionDeIngresosUseCase struct {
	estadisticaRepo domain.EstadisticasRepository
}

func NewGenerarReporteRelacionDeIngresosUseCase(estadisticaRepo domain.EstadisticasRepository) *GenerarReporteRelacionDeIngresosUseCase {
	return &GenerarReporteRelacionDeIngresosUseCase{
		estadisticaRepo: estadisticaRepo,
	}
}

func (uc *GenerarReporteRelacionDeIngresosUseCase) Execute() ([]byte, error) {
	reporte := uc.estadisticaRepo.ObtenerReporteParaContador()

	excelFile, err := exporter.GenerarReporteContadorExcel(reporte)
	if err != nil {
		return nil, err
	}

	buffer, err := excelFile.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
