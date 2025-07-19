package controller

import (
	"lgc/src/infraestructure/di"
	usecase "lgc/src/usecase/inscripcion"

	"github.com/gin-gonic/gin"
)

func ObtenerResumenEstadisticas(c *gin.Context) {

	resumenEstadisticas := usecase.NewObtenerResumenEstadisticasUseCase(
		di.GetContainer().GetEstadisticasRepository(),
	)

	response := resumenEstadisticas.Execute()

	c.JSON(response.StatusCode, response)
}
