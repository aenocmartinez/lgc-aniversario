package dto

type InscripcionConParticipantesDTO struct {
	ID             int64             `json:"id"`
	FormaPago      string            `json:"forma_pago"`
	MontoPagoCOP   int               `json:"monto_pago_cop"`
	MontoPagoUSD   int               `json:"monto_pago_usd"`
	UrlSoportePago string            `json:"url_soporte_pago"`
	Estado         string            `json:"estado"`
	FechaCreacion  string            `json:"fecha_creacion"`
	Participantes  []ParticipanteDTO `json:"participantes"`
}
