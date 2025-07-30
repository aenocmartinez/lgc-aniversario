package formrequest

type ActualizarUsuarioFormRequest struct {
	ID       int64  `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Nombre   string `json:"nombre" binding:"required"`
	Password string `json:"password" bingind:"required"`
}
