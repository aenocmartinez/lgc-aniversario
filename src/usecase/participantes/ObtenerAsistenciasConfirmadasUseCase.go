package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type ObtenerAsistenciasConfirmadasUseCase struct {
	participanteRepo domain.ParticipanteRepository
}

func NewObtenerAsistenciasConfirmadasUseCase(participanteRepo domain.ParticipanteRepository) *ObtenerAsistenciasConfirmadasUseCase {
	return &ObtenerAsistenciasConfirmadasUseCase{
		participanteRepo: participanteRepo,
	}
}

func (uc *ObtenerAsistenciasConfirmadasUseCase) Execute() dto.APIResponse {
	asistencias := uc.participanteRepo.GetAsistenciasConfirmadas()
	return dto.APIResponse{
		StatusCode: 200,
		Data:       asistencias,
	}
}
