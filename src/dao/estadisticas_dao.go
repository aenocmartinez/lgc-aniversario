package dao

import (
	"lgc/src/view/dto"

	"gorm.io/gorm"
)

type EstadisticasDao struct {
	db *gorm.DB
}

func NewEstadisticasDao(db *gorm.DB) *EstadisticasDao {
	return &EstadisticasDao{db: db}
}

func (i *EstadisticasDao) ObtenerResumenEstadisticasEvento(cupoMax int) dto.EstadisticaEventoDTO {
	db := i.db

	// 1. Cupo presencial ocupado
	var totalPresenciales int64
	db.Table("participantes").
		Where("modalidad = ? AND dias_asistencia = ? AND inscripcion_id IN (SELECT id FROM inscripciones WHERE estado != ?)",
			"presencial", "sabado", "Rechazada").
		Count(&totalPresenciales)

	// 2. Total por modalidad
	var modalidadResults []struct {
		Modalidad string
		Total     int
	}
	db.Table("participantes").
		Select("modalidad, COUNT(*) as total").
		Group("modalidad").
		Scan(&modalidadResults)

	totalPorModalidad := map[string]int{}
	for _, r := range modalidadResults {
		totalPorModalidad[r.Modalidad] = r.Total
	}

	// 3. Total por días de asistencia
	var diasAsistenciaResults []struct {
		Dia   string
		Total int
	}
	db.Table("participantes").
		Select("dias_asistencia as dia, COUNT(*) as total").
		Group("dias_asistencia").
		Scan(&diasAsistenciaResults)

	totalPorDiaAsistencia := map[string]int{}
	for _, r := range diasAsistenciaResults {
		totalPorDiaAsistencia[r.Dia] = r.Total
	}

	// 4. Estado por forma de pago
	var estadoPagoResults []struct {
		FormaPago string
		Estado    string
		Total     int
	}
	db.Table("inscripciones").
		Select("forma_pago, estado, COUNT(*) as total").
		Group("forma_pago, estado").
		Scan(&estadoPagoResults)

	estadoPorFormaPago := map[string]map[string]int{}
	for _, r := range estadoPagoResults {
		if _, ok := estadoPorFormaPago[r.FormaPago]; !ok {
			estadoPorFormaPago[r.FormaPago] = map[string]int{}
		}
		estadoPorFormaPago[r.FormaPago][r.Estado] = r.Total
	}

	// 5. Inscripciones por día
	var inscripcionesPorDia []dto.InscripcionesDiaDTO
	db.Table("inscripciones").
		Select("DATE(created_at) as fecha, COUNT(*) as total").
		Group("DATE(created_at)").
		Order("fecha").
		Scan(&inscripcionesPorDia)

	// 6. Participantes sin iglesia
	var totalSinIglesia int64
	db.Table("participantes").
		Where("TRIM(iglesia) = 'No asiste a una iglesia'").
		Count(&totalSinIglesia)

	// 7. Total Recaudo COP y USD (solo inscripciones aprobadas o preaprobadas)
	var totalCOP, totalUSD float64
	db.Table("inscripciones").
		Select("SUM(monto_pagado_cop), SUM(monto_pagado_usd)").
		Where("estado IN ?", []string{"Aprobada", "PreAprobada"}).
		Row().
		Scan(&totalCOP, &totalUSD)

	// 8. Recaudo por modalidad
	type RecaudoPorModalidad struct {
		Modalidad string
		TotalCOP  float64
		TotalUSD  float64
	}

	var recaudos []RecaudoPorModalidad

	db.Table("participantes AS p").
		Select("p.modalidad, SUM(i.monto_pagado_cop) AS total_cop, SUM(i.monto_pagado_usd) AS total_usd").
		Joins("JOIN inscripciones i ON i.id = p.inscripcion_id").
		Where("i.estado IN ?", []string{"Aprobada", "PreAprobada"}).
		Group("p.modalidad").
		Scan(&recaudos)

	recaudoPorModalidad := map[string]map[string]float64{}
	for _, r := range recaudos {
		recaudoPorModalidad[r.Modalidad] = map[string]float64{
			"cop": r.TotalCOP,
			"usd": r.TotalUSD,
		}
	}

	return dto.EstadisticaEventoDTO{
		CupoMaximoPresencial:    cupoMax,
		CupoUtilizadoPresencial: int(totalPresenciales),
		CupoRestantePresencial:  cupoMax - int(totalPresenciales),
		TotalPorModalidad:       totalPorModalidad,
		TotalPorDiaAsistencia:   totalPorDiaAsistencia,
		EstadoPorFormaPago:      estadoPorFormaPago,
		TotalSinIglesia:         int(totalSinIglesia),
		InscripcionesPorDia:     inscripcionesPorDia,
		TotalRecaudoCOP:         totalCOP,
		TotalRecaudoUSD:         totalUSD,
		RecaudoPorModalidad:     recaudoPorModalidad,
	}
}

func (i *EstadisticasDao) ObtenerReporteParaContador() []dto.ReporteContadorInscripcionDTO {
	db := i.db

	// 1. Obtener todas las inscripciones que no son gratuitas
	var inscripciones []struct {
		ID             int
		FormaPago      string
		MontoPagadoCOP float64
		MontoPagadoUSD float64
		SoportePagoURL string
	}
	db.Table("inscripciones").
		Select("id, forma_pago, monto_pagado_cop, monto_pagado_usd, soporte_pago_url").
		Where("forma_pago != ?", "gratuito").
		Find(&inscripciones)

	// 2. Obtener todos los participantes relacionados
	var participantes []struct {
		InscripcionID   int
		NombreCompleto  string
		NumeroDocumento string
		Telefono        string
	}
	db.Table("participantes").
		Select("inscripcion_id, nombre_completo, numero_documento, telefono").
		Where("inscripcion_id IN (?)", db.Table("inscripciones").Select("id").Where("forma_pago != ?", "gratuito")).
		Find(&participantes)

	// 3. Agrupar participantes por inscripción
	participantesPorInscripcion := make(map[int][]dto.ReporteContadorParticipanteDTO)
	for _, p := range participantes {
		dtoPart := dto.ReporteContadorParticipanteDTO{
			NombreCompleto:  p.NombreCompleto,
			NumeroDocumento: p.NumeroDocumento,
			Telefono:        p.Telefono,
		}
		participantesPorInscripcion[p.InscripcionID] = append(participantesPorInscripcion[p.InscripcionID], dtoPart)
	}

	// 4. Construir el DTO final
	var resultado []dto.ReporteContadorInscripcionDTO
	for _, insc := range inscripciones {
		resultado = append(resultado, dto.ReporteContadorInscripcionDTO{
			ID:             insc.ID,
			FormaPago:      insc.FormaPago,
			MontoPagadoCOP: insc.MontoPagadoCOP,
			MontoPagadoUSD: insc.MontoPagadoUSD,
			SoportePagoURL: insc.SoportePagoURL,
			Participantes:  participantesPorInscripcion[insc.ID],
		})
	}

	return resultado
}
