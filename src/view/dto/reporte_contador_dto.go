package dto

type ReporteContadorParticipanteDTO struct {
	NombreCompleto  string `json:"nombre_completo"`
	NumeroDocumento string `json:"numero_documento"`
	Telefono        string `json:"telefono"`
}

type ReporteContadorInscripcionDTO struct {
	ID             int                              `json:"id"`
	FormaPago      string                           `json:"forma_pago"`
	MontoPagadoCOP float64                          `json:"monto_pagado_cop"`
	MontoPagadoUSD float64                          `json:"monto_pagado_usd"`
	SoportePagoURL string                           `json:"soporte_pago_url"`
	Participantes  []ReporteContadorParticipanteDTO `json:"participantes"`
}

type ReporteInscritosVirtualDTO struct {
	FormaPago      string  `json:"forma_pago"`
	Modalidad      string  `json:"modalidad"`
	MontoPagadoUSD float64 `json:"monto_pagado_usd"`
	NombreCompleto string  `json:"nombre_completo"`
	SoportePagoURL string  `json:"soporte_pago_url"`
}
