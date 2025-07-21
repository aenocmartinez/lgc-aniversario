package dto

type InscripcionDTO struct {
	ID             int64  `json:"id"`
	FormaPago      string `json:"forma_pago"`
	MontoPagoCOP   int    `json:"monto_pago_cop"`
	MontoPagoUSD   int    `json:"monto_pago_usd"`
	UrlSoportePago string `json:"url_soporte_pago"`
	FechaCreacion  string `json:"fecha_creacion"`
	Estado         string `json:"estado"`
}
