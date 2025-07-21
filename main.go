package main

import (
	"lgc/src/infraestructure/middleware"
	"lgc/src/view/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Login
	r.POST("/login", controller.Login)

	// Pruebas
	r.GET("/check-db", controller.CheckDBConnection)
	r.GET("/mutant", controller.Mutant)

	// Formulario inscripcion
	r.POST("/realizar-inscripcion", controller.RealizarInscripcion)
	r.POST("/cargar-soporte-pago", controller.CargarArchivoDePago)
	r.GET("/cupos-disponibles", controller.ConsultarCuposDisponibles)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/logout", controller.Logout)

		// inscripcionGroup := protected.Group("/inscripciones")
		// {
		// 	inscripcionGroup.GET("", controller.ListarInscripciones)
		// 	inscripcionGroup.GET("/pendientes", controller.ListarInscripcionesPendientes)
		// 	inscripcionGroup.GET("/aprobadas", controller.ListarInscripcionesAprobadas)
		// 	inscripcionGroup.PUT("/anular/:id", controller.AnularInscripcion)
		// 	inscripcionGroup.PUT("/aprobar/:id", controller.AprobarInscripcion)
		// 	inscripcionGroup.POST("/pago-efectivo", controller.RealizarInscripcionPagoEfectivo)
		// }

		estadisticaGroup := protected.Group("/estadisticas")
		{
			estadisticaGroup.GET("/resumen", controller.ObtenerResumenEstadisticas)
		}
	}

	r.Run(":8586")
}
