package controller

import (
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

	c.JSON(http.StatusOK, gin.H{"message": "Formulario de inscripción"})
}
