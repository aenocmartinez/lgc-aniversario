package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		fmt.Println("DB_USER: ", os.Getenv("DB_USER"))
		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
			"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
			os.Getenv("DB_NAME") + "?parseTime=true"

		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,   // Límite de consultas lentas
					LogLevel:      logger.Silent, // 🔇 Desactiva logs de consultas
					Colorful:      false,
				},
			),
		})
		if err != nil {
			log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
		}

		log.Println("✅ Conexión a la base de datos establecida con GORM.")
	})

	return db
}
