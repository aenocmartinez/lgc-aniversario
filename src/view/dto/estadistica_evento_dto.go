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
	TotalRecaudo               int     `json:"TotalRecaudo"`      // COP
	RecaudoVirtual             float64 `json:"RecaudoVirtual"`    // USD
	RecaudoPresencial          int     `json:"RecaudoPresencial"` // COP
}
