package usecase

import (
	"lgc/src/domain"
	"lgc/src/infraestructure/util"
	"lgc/src/view/dto"
	"os"
)

type ObtenerResumenEstadisticasUseCase struct {
	estadisticasRepo domain.EstadisticasRepository
}

func NewObtenerResumenEstadisticasUseCase(repo domain.EstadisticasRepository) *ObtenerResumenEstadisticasUseCase {
	return &ObtenerResumenEstadisticasUseCase{
		estadisticasRepo: repo,
	}
}

func (uc *ObtenerResumenEstadisticasUseCase) Execute() dto.APIResponse {
	cupoMax, err := util.ConvertStringToInt(os.Getenv("APP_CUPO_MAX"))
	if err != nil {
		cupoMax = 400
	}

	resumen := uc.estadisticasRepo.ObtenerResumenEstadisticas(cupoMax)

	return dto.NewAPIResponse(200, "Resumen generado correctamente", resumen)
}
