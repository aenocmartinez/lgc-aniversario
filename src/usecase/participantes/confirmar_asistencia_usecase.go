package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type ConfirmaAsistenciaUseCase struct {
	participanteRepo domain.ParticipanteRepository
}

func NewConfirmaAsistenciaUseCase(participanteRepo domain.ParticipanteRepository) *ConfirmaAsistenciaUseCase {
	return &ConfirmaAsistenciaUseCase{
		participanteRepo: participanteRepo,
	}
}

func (uc *ConfirmaAsistenciaUseCase) Execute(documento string) dto.APIResponse {

	participante := uc.participanteRepo.BuscarParticipantePorDocumento(documento)
	if !participante.Existe() {
		return dto.NewAPIResponse(404, "Participante no encontrado", nil)
	}

	if participante.AsistenciaConfirmada() {
		return dto.NewAPIResponse(409, "El participante ya ha confirmado su asistencia", nil)
	}

	if !participante.RegistrarAsistencia() {
		return dto.NewAPIResponse(500, "Ha ocurrido un error al confirmar la asistencia", nil)
	}

	return dto.NewAPIResponse(200, "Asistencia confirmada exitosamente", nil)
}
