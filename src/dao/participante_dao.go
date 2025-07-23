package dao

import (
	"lgc/src/view/dto"
	"strings"

	"gorm.io/gorm"
)

type ParticipanteDao struct {
	db *gorm.DB
}

func NewParticipanteDao(db *gorm.DB) *ParticipanteDao {
	return &ParticipanteDao{db: db}
}

func (r *ParticipanteDao) ObtenerParticipantesParaLogistica() []dto.ReporteLogisticaDTO {
	var results []dto.ReporteLogisticaDTO

	r.db.Table("participantes").
		Select("nombre_completo, numero_documento, correo_electronico, dias_asistencia, telefono").
		Where("modalidad = ?", "presencial").
		Scan(&results)

	for i := range results {
		switch strings.ToLower(strings.TrimSpace(results[i].DiasAsistencia)) {
		case "sabado":
			results[i].DiasAsistencia = "viernes, sábado y domingo"
		case "viernes_y_domingo":
			results[i].DiasAsistencia = "viernes y domingo"
		}
	}

	return results
}
