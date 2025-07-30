package controller

import (
	"lgc/src/infraestructure/di"
	usecase "lgc/src/usecase/reportes"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DescargarRelacionIngresosExcel(c *gin.Context) {
	useCase := usecase.NewGenerarReporteRelacionDeIngresosUseCase(
		di.GetContainer().GetEstadisticasRepository(),
	)

	fileBytes, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo generar el reporte",
		})
		return
	}

	loc, _ := time.LoadLocation("America/Bogota")
	fechaBogota := time.Now().In(loc)
	fileName := "reporte_ingresos_inscripciones_aniversario_" + fechaBogota.Format("2006-01-02_15-04-05") + ".xlsx"

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}

func ReporteRelacionDeIngresos(c *gin.Context) {

	useCase := usecase.NewObtenerReporteContadorUseCase(
		di.GetContainer().GetEstadisticasRepository(),
	)

	reporte := useCase.Execute()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    reporte,
		"message": "Reporte generado correctamente",
	})
}

func DescargarReporteLogistica(c *gin.Context) {
	usecase := usecase.NewGenerarReporteLogisticaUseCase(di.GetContainer().GetParticipanteRepository())

	excelBytes, err := usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el reporte"})
		return
	}

	loc, _ := time.LoadLocation("America/Bogota")
	fechaBogota := time.Now().In(loc)
	fileName := "reporte_logistica_" + fechaBogota.Format("2006-01-02_15-04-05") + ".xlsx"

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", excelBytes)
}
