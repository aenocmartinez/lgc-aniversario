package controller

import (
	"lgc/src/infraestructure/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckDBConnection(c *gin.Context) {
	db := database.GetDB()

	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "No se pudo obtener la conexión a la base de datos",
			"error":   err.Error(),
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "No se pudo conectar a la base de datos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Conexión a la base de datos establecida correctamente",
	})
}

func Mutant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Version 1.0.0 - Mutant endpoint",
	})
}
