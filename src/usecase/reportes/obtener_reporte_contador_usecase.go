package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type ObtenerReporteContadorUseCase struct {
	estadisticaRepo domain.EstadisticasRepository
}

func NewObtenerReporteContadorUseCase(estadisticaRepo domain.EstadisticasRepository) *ObtenerReporteContadorUseCase {
	return &ObtenerReporteContadorUseCase{
		estadisticaRepo: estadisticaRepo,
	}
}

func (uc *ObtenerReporteContadorUseCase) Execute() []dto.ReporteContadorInscripcionDTO {
	return uc.estadisticaRepo.ObtenerReporteParaContador()
}
