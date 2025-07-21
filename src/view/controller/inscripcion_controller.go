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

func ConsultarCuposDisponibles(c *gin.Context) {
	consultarCupos := usecase.NewConsultarCuposDisponiblesUseCase(di.GetContainer().GetInscripcionRepository())
	response := consultarCupos.Execute()

	c.JSON(response.StatusCode, response)
}

// func RealizarInscripcionPagoEfectivo(c *gin.Context) {

// 	var req formrequest.InscripcionPagoEfectivoFormRequest

// 	err := c.ShouldBindJSON(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	realizarInscripcion := usecase.NewRealizarInscripcionPagoEnEfectivoUseCase(
// 		di.GetContainer().GetInscripcionRepository(),
// 	)

// 	response := realizarInscripcion.Execute(req)

// 	c.JSON(response.StatusCode, response)
// }

// func ListarInscripciones(c *gin.Context) {
// 	listarInscripciones := usecase.NewListarInscripcionesUseCase(
// 		di.GetContainer().GetInscripcionRepository(),
// 	)

// 	response := listarInscripciones.Execute()

// 	c.JSON(response.StatusCode, response)
// }

// func ListarInscripcionesPendientes(c *gin.Context) {
// 	listarInscripciones := usecase.NewListarInscripcionesPendientesUseCase(
// 		di.GetContainer().GetInscripcionRepository(),
// 	)

// 	response := listarInscripciones.Execute()

// 	c.JSON(response.StatusCode, response)
// }

// func ListarInscripcionesAprobadas(c *gin.Context) {
// 	listarInscripciones := usecase.NewListarInscripcionesAprobadasUseCase(
// 		di.GetContainer().GetInscripcionRepository(),
// 	)

// 	response := listarInscripciones.Execute()

// 	c.JSON(response.StatusCode, response)
// }

// func AnularInscripcion(c *gin.Context) {
// 	id, err := util.ConvertStringToID(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
// 		return
// 	}

// 	anularInscripcion := usecase.NewAnularInscripcionUseCase(
// 		di.GetContainer().GetInscripcionRepository(),
// 	)

// 	response := anularInscripcion.Execute(id)

// 	c.JSON(response.StatusCode, response)
// }

// func AprobarInscripcion(c *gin.Context) {
// 	id, err := util.ConvertStringToID(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, dto.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
// 		return
// 	}

// 	aprobarInscripcion := usecase.NewAprobarInscripcionUseCase(
// 		di.GetContainer().GetInscripcionRepository(),
// 	)

// 	response := aprobarInscripcion.Execute(id)

// 	c.JSON(response.StatusCode, response)
// }
