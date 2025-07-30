package controller

import (
	"lgc/src/infraestructure/di"
	usecase "lgc/src/usecase/usuarios"
	formrequest "lgc/src/view/form-request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearUsuario(c *gin.Context) {

	var req formrequest.CrearUsuarioFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	crearUsuario := usecase.NewCrearUsuarioUseCase(di.GetContainer().GetUserRepository())
	response := crearUsuario.Execute(req.Nombre, req.Email, req.Password)

	c.JSON(response.StatusCode, response)
}

func ActualizarUsuario(c *gin.Context) {

	var req formrequest.ActualizarUsuarioFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actualizarUusario := usecase.NewActualizarUsuarioUseCase(di.GetContainer().GetUserRepository())
	response := actualizarUusario.Executar(req)

	c.JSON(response.StatusCode, response)
}
