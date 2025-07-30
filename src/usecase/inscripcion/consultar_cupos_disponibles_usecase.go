package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
	"os"
	"strconv"
)

type ConsultarCuposDisponiblesUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewConsultarCuposDisponiblesUseCase(inscripcionRepo domain.InscripcionRepository) *ConsultarCuposDisponiblesUseCase {
	return &ConsultarCuposDisponiblesUseCase{inscripcionRepo: inscripcionRepo}
}

func (uc *ConsultarCuposDisponiblesUseCase) Execute() dto.APIResponse {
	cupoMax, err := strconv.Atoi(os.Getenv("APP_CUPO_MAX"))
	if err != nil {
		cupoMax = 400
	}

	ocupados, disponibles := uc.inscripcionRepo.CuposDisponibles(cupoMax)

	data := map[string]int{
		"total":       cupoMax,
		"ocupados":    ocupados,
		"disponibles": disponibles,
	}

	return dto.NewAPIResponse(200, "Consulta exitosa de cupos disponibles.", data)
}
