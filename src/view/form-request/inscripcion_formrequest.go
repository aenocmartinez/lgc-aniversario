package formrequest

type InscripcionFormRequest struct {
	Nombre         string `json:"nombre" binding:"required"`
	Documento      string `json:"documento" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Telefono       string `json:"telefono" binding:"required"`
	Ciudad         string `json:"ciudad,omitempty"`
	Iglesia        string `json:"iglesia,omitempty"`
	HabeasData     bool   `json:"habeas_data" binding:"required"`
	Asistencia     string `json:"asistencia" binding:"required,oneof=Virtual Presencial"`
	ComprobatePago string `json:"comprobante_pago" binding:"required,url"`
}
