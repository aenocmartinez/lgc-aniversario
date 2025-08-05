package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type BuscarParticipantePorDocumentoUseCase struct {
	participanteRepo domain.ParticipanteRepository
}

func NewBuscarParticipantePorDocumentoUseCase(participanteRepo domain.ParticipanteRepository) *BuscarParticipantePorDocumentoUseCase {
	return &BuscarParticipantePorDocumentoUseCase{
		participanteRepo: participanteRepo,
	}
}

func (uc *BuscarParticipantePorDocumentoUseCase) Execute(documento string) dto.APIResponse {

	participante := uc.participanteRepo.BuscarParticipantePorDocumento(documento)

	if !participante.Existe() {
		return dto.NewAPIResponse(404, "Participante no encontrado", nil)
	}

	return dto.NewAPIResponse(200, "Participante encontrado", participante.ToDTO())
}
