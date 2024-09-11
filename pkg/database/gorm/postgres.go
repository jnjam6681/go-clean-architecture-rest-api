package gorm

import (
	"fmt"
	"log"

	"github.com/jnjam6681/go-clean-architecture-rest-api/config"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormClient(cfg *config.Config) (*gorm.DB, error) {
	sslmode := map[bool]string{true: "enable", false: "disable"}[cfg.Postgres.SSLMode]

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		sslmode,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// // ตั้งค่าการจัดการ connection pool
	// postgres, err := db.DB()
	// if err != nil {
	// 	log.Fatalf("Underlying database connection is not sql.DB err: %v", err)
	// }
	// // กำหนดค่า Max Open Connections
	// postgres.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	// // กำหนดค่า Max Idle Connections (connection ที่เปิดรอแต่ยังไม่ได้ใช้)
	// postgres.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	// // กำหนดเวลาชีวิตของ connection (ระยะเวลาที่ connection สามารถใช้งานได้)
	// postgres.SetConnMaxLifetime(time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Minute)

	return db, nil
}

func RunMigrate(db *gorm.DB) {
	log.Print("Starting database migrations...")

	// Add all model migrates here
	modelsToMigrate := []interface{}{
		&entity.Todo{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatal("Migration failed", err)
		}
	}

	log.Print("Database migrations completed successfully")
}
