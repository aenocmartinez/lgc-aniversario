package controller

import (
	"lgc/src/infraestructure/di"
	usecase "lgc/src/usecase/inscripcion"
	formrequest "lgc/src/view/form-request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RealizarInscripcion(c *gin.Context) {

	var req formrequest.InscripcionFormRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	realizarInscripcion := usecase.NewRealizarInscripcionUseCase(
		di.GetContainer().GetInscripcionRepository(),
	)

	response := realizarInscripcion.Execute(req)

	c.JSON(response.StatusCode, response)
}
