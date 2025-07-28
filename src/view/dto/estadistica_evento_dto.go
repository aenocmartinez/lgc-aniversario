package dto

type EstadisticaEventoDTO struct {
	TotalInscritos             int
	TotalInscritosSabado       int
	TotalInscritosViernes      int
	TotalPresencialesSabado    int
	TotalVirtualesSabado       int
	CupoMaximoPresencialSabado int
	CupoRestanteSabado         int
	PorcentajeAvanceMeta       float64
}
