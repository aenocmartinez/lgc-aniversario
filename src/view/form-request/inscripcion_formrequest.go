package formrequest

type ParticipanteRequest struct {
	Nombre         string `json:"nombre" binding:"required"`
	Documento      string `json:"documento" binding:"required"`
	Email          string `json:"email" binding:"omitempty,email"`
	Telefono       string `json:"telefono" binding:"required"`
	Modalidad      string `json:"modalidad" binding:"required,oneof=presencial virtual"`
	DiasAsistencia string `json:"dias_asistencia" binding:"omitempty,oneof=viernes_y_domingo sabado"`
	Iglesia        string `json:"iglesia" binding:"omitempty"`
	Ciudad         string `json:"ciudad" binding:"omitempty"`
	HabeasData     bool   `json:"habeas_data" binding:"required"`
}

type InscripcionFormRequest struct {
	FormaPago      string                `json:"forma_pago" binding:"required,oneof=efectivo transaccion gratuito"`
	MontoCOP       int                   `json:"monto_cop" binding:"omitempty,min=0"`
	MontoUSD       float32               `json:"monto_usd" binding:"omitempty,min=0"`
	UrlSoportePago string                `json:"url_soporte_pago" binding:"omitempty,url"`
	Participantes  []ParticipanteRequest `json:"participantes" binding:"required,dive"`
}
