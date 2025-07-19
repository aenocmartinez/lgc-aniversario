package controller

import (
	usecase "lgc/src/usecase/inscripcion"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CargarArchivoDePago(c *gin.Context) {
	file, err := c.FormFile("ruta_comprobante_pago")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comprobante de pago no encontrado"})
		return
	}

	useCase := usecase.NewUploadS3UseCase()
	path, err := useCase.Execute(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"path": path})
}
