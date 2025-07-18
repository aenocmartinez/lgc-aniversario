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

	// Pruebas
	r.GET("/check-db", controller.CheckDBConnection)
	r.GET("/mutant", controller.Mutant)

	// Formulario inscripcion
	r.POST("/formulario-inscripcion", controller.CrearInscripcion)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// 	protected.GET("/secure-data", controller.SecureData)
		// 	protected.POST("/logout", controller.Logout)

		// 	cursoGroup := protected.Group("/cursos")
		// 	{
		// 		cursoGroup.GET("", controller.ListarCursos)
		// 		cursoGroup.POST("", controller.CrearCurso)
		// 		cursoGroup.PUT("/:id", controller.ActualizarCurso)
		// 		cursoGroup.DELETE("/:id", controller.EliminarCurso)
		// 	}

		// 	periodoGroup := protected.Group("/periodos")
		// 	{
		// 		periodoGroup.GET("", controller.ListarPeriodos)
		// 		periodoGroup.POST("", controller.CrearPeriodo)
		// 		periodoGroup.PUT("/:id", controller.ActualizarPeriodo)
		// 		periodoGroup.DELETE("/:id", controller.EliminarPeriodo)
		// 		periodoGroup.GET("/:id/cursos", controller.ListarCursosDePeriodo)
		// 		periodoGroup.GET("/:id/cursos/:curso_id/alumnos", controller.ListarAlumnosMatriculados)
		// 		periodoGroup.POST("/curso-periodo/:id/contenido-tematico", controller.AgregarContenidoTematico)
		// 		periodoGroup.DELETE("/curso-periodo/:id/contenido-tematico/:contenido_tematico_id", controller.QuitarContenidoTematico)
		// 		periodoGroup.GET("/curso-periodo/:id/contenido-tematico", controller.ObtenerContenidoTematicoDeUnCursoPeriodo)
		// 	}

		// 	alumnoGroup := protected.Group("/alumnos")
		// 	{
		// 		alumnoGroup.GET("", controller.ListarAlumnos)
		// 		alumnoGroup.GET("/:id", controller.BuscarAlumnoPorId)
		// 		alumnoGroup.POST("", controller.CrearAlumno)
		// 		alumnoGroup.PUT("/:id", controller.ActualizarAlumno)
		// 		alumnoGroup.DELETE("/:id", controller.EliminarAlumno)
		// 		alumnoGroup.POST("/:id/matriculas", controller.MatricularAlumno)
		// 		alumnoGroup.PUT("/:id/cambiar-estado", controller.CambiarEstadoAlumno)
		// 	}

		// 	maestroGroup := protected.Group("/maestros")
		// 	{
		// 		maestroGroup.GET("", controller.ListarMaestros)
		// 		maestroGroup.GET("/:id", controller.BuscarMaestroPorId)
		// 		maestroGroup.POST("", controller.CrearMaestro)
		// 		maestroGroup.PUT("/:id", controller.ActualizarMaestro)
		// 		maestroGroup.DELETE("/:id", controller.EliminarMaestro)
		// 	}

		// 	celebracionGroup := protected.Group("/celebraciones")
		// 	{
		// 		celebracionGroup.GET("", controller.ListarCelebraciones)
		// 		celebracionGroup.POST("", controller.CrearCelebracion)
		// 		celebracionGroup.PUT("/:id", controller.ActualizarCelebracion)
		// 		celebracionGroup.DELETE("/:id", controller.EliminarCelebracion)
		// 	}

		// 	grupoGroup := protected.Group("/grupos")
		// 	{
		// 		grupoGroup.GET("", controller.ListarGrupos)
		// 		// grupoGroup.POST("", controller.CrearGrupo)
		// 		grupoGroup.POST("", controller.GuardarGrupo)
		// 		// grupoGroup.PUT("/:id", controller.ActualizarGrupo)

		// 		grupoGroup.DELETE("/:id", controller.EliminarGrupo)
		// 	}

		// 	claseGroup := protected.Group("/clases")
		// 	{
		// 		claseGroup.POST("/asistencia", controller.RegistrarAsistencia)
		// 	}

	}

	r.Run(":8586")
}
