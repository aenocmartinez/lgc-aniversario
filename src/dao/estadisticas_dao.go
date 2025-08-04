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
	var cupoMaximoSabado int = cupoMax
	db := i.db
	var resultado dto.EstadisticaEventoDTO
	resultado.CupoMaximoPresencialSabado = cupoMaximoSabado

	// Variables auxiliares para Count
	var totalInscritos int64
	var totalSabado int64
	var totalViernes int64
	var totalPresencialesSabado int64
	var totalVirtualesSabado int64

	// 1. Total de inscritos (excluyendo Rechazada)
	db.Table("participantes AS p").
		Joins("JOIN inscripciones i ON i.id = p.inscripcion_id").
		Where("i.estado <> ?", "Rechazada").
		Count(&totalInscritos)

	// 2. Total inscritos sábado
	db.Table("participantes AS p").
		Joins("JOIN inscripciones i ON i.id = p.inscripcion_id").
		Where("p.dias_asistencia = ? AND i.estado <> ?", "sabado", "Rechazada").
		Count(&totalSabado)

	// 3. Total inscritos viernes_y_domingo
	db.Table("participantes AS p").
		Joins("JOIN inscripciones i ON i.id = p.inscripcion_id").
		Where("p.dias_asistencia = ? AND i.estado <> ?", "viernes_y_domingo", "Rechazada").
		Count(&totalViernes)

	// 4. Total PRESENCIALES sábado
	db.Table("participantes AS p").
		Joins("JOIN inscripciones i ON i.id = p.inscripcion_id").
		Where("p.dias_asistencia = ? AND p.modalidad = ? AND i.estado <> ?", "sabado", "presencial", "Rechazada").
		Count(&totalPresencialesSabado)

	// 5. Total VIRTUALES sábado
	db.Table("participantes AS p").
		Joins("JOIN inscripciones i ON i.id = p.inscripcion_id").
		Where("p.dias_asistencia = ? AND p.modalidad = ? AND i.estado <> ?", "sabado", "virtual", "Rechazada").
		Count(&totalVirtualesSabado)

	// Asignar al DTO (de int64 a int)
	resultado.TotalInscritos = int(totalInscritos)
	resultado.TotalInscritosSabado = int(totalSabado)
	resultado.TotalInscritosViernes = int(totalViernes)
	resultado.TotalPresencialesSabado = int(totalPresencialesSabado)
	resultado.TotalVirtualesSabado = int(totalVirtualesSabado)

	// 6. Cupo restante
	resultado.CupoRestanteSabado = cupoMaximoSabado - resultado.TotalPresencialesSabado

	// 7. Porcentaje de avance a la meta
	if cupoMaximoSabado > 0 {
		resultado.PorcentajeAvanceMeta = (float64(resultado.TotalPresencialesSabado) / float64(cupoMaximoSabado)) * 100
	}

	var totalCOP, presencialCOP int64
	var virtualUSD float64

	// Total COP (todos los que pagaron en COP y no están Rechazados)
	db.Table("inscripciones").
		Where("estado <> ? AND forma_pago != ?", "Rechazada", "gratuito").
		Select("SUM(monto_pagado_cop)").Scan(&totalCOP)

	// Recaudo presencial COP (solo presencial)
	db.Table("inscripciones").
		Where("estado <> ?", "Rechazada").
		Select("SUM(monto_pagado_cop)").
		Scan(&presencialCOP)

	// Recaudo virtual USD (solo virtual)
	db.Table("inscripciones").
		Where("estado <> ?", "Rechazada").
		Select("SUM(monto_pagado_usd)").
		Scan(&virtualUSD)

	resultado.TotalRecaudo = int(totalCOP)
	resultado.RecaudoPresencial = int(presencialCOP)
	resultado.RecaudoVirtual = virtualUSD

	return resultado
}

func (i *EstadisticasDao) ObtenerReporteParaContador() []dto.ReporteContadorInscripcionDTO {
	db := i.db

	// 1. Obtener inscripciones PreAprobadas que no son gratuitas
	var inscripciones []struct {
		ID             int
		FormaPago      string
		MontoPagadoCOP float64
		MontoPagadoUSD float64
		SoportePagoURL string
	}
	db.Table("inscripciones").
		Select("id, forma_pago, monto_pagado_cop, monto_pagado_usd, soporte_pago_url").
		Where("forma_pago != ? AND estado = ?", "gratuito", "PreAprobada").
		Find(&inscripciones)

	// 2. Obtener participantes de esas inscripciones
	var participantes []struct {
		InscripcionID   int
		NombreCompleto  string
		NumeroDocumento string
		Telefono        string
	}
	db.Table("participantes").
		Select("inscripcion_id, nombre_completo, numero_documento, telefono").
		Where("inscripcion_id IN (?)",
			db.Table("inscripciones").
				Select("id").
				Where("forma_pago != ? AND estado = ?", "gratuito", "PreAprobada")).
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

func (i *EstadisticasDao) ObtenerReporteFinancieroInscritosVirtual() []dto.ReporteInscritosVirtualDTO {
	db := i.db

	var resultado []dto.ReporteInscritosVirtualDTO

	db.Table("inscripciones AS i").
		Select("i.forma_pago, p.modalidad, i.monto_pagado_usd, p.nombre_completo, i.soporte_pago_url").
		Joins("INNER JOIN participantes p ON i.id = p.inscripcion_id").
		Where("i.estado <> ? AND i.forma_pago <> ? AND p.modalidad = ?", "Rechazada", "gratuito", "virtual").
		Scan(&resultado)

	return resultado
}
