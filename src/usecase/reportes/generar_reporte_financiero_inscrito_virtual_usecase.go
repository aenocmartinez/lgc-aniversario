package usecase

import (
	"lgc/src/domain"
	"lgc/src/infraestructure/exporter"
)

type GenerarReporteFinancieroInscritoVirtualUsecase struct {
	estadisticaRepo domain.EstadisticasRepository
}

func NewGenerarReporteFinancieroInscritoVirtualUsecase(estadisticaRepo domain.EstadisticasRepository) *GenerarReporteFinancieroInscritoVirtualUsecase {
	return &GenerarReporteFinancieroInscritoVirtualUsecase{
		estadisticaRepo: estadisticaRepo,
	}
}

func (uc *GenerarReporteFinancieroInscritoVirtualUsecase) Execute() ([]byte, error) {
	reporte := uc.estadisticaRepo.ObtenerReporteFinancieroInscritosVirtual()

	excelFile, err := exporter.GenerarReporteFinancieroVirtual(reporte)
	if err != nil {
		return nil, err
	}

	buffer, err := excelFile.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
