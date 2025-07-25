package formrequest

type CrearUsuarioFormRequest struct {
	Email    string `json:"email" binding:"required"`
	Nombre   string `json:"nombre" binding:"required"`
	Password string `json:"password" bingind:"required"`
}
