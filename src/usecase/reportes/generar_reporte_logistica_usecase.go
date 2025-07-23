package usecase

import (
	"lgc/src/domain"
	"lgc/src/infraestructure/exporter"
)

type GenerarReporteLogisticaUseCase struct {
	participanteRepo domain.ParticipanteRepository
}

func NewGenerarReporteLogisticaUseCase(participanteRepo domain.ParticipanteRepository) *GenerarReporteLogisticaUseCase {
	return &GenerarReporteLogisticaUseCase{participanteRepo: participanteRepo}
}

func (uc *GenerarReporteLogisticaUseCase) Execute() ([]byte, error) {
	participantes := uc.participanteRepo.ObtenerParticipantesParaLogistica()

	return exporter.GenerarReporteLogisticaExcel(participantes)
}
