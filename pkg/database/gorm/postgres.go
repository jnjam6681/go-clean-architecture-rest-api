package gorm

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jnjam6681/go-clean-architecture-rest-api/config"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormClient(cfg *config.Config) (*gorm.DB, func(), error) {
	sslmode := map[bool]string{true: "enable", false: "disable"}[cfg.Postgres.SSLMode]

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		sslmode,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		// NowFunc: func() time.Time {
		// 	return time.Now().UTC()
		// },
		NowFunc: func() time.Time {
			utc, _ := time.LoadLocation("")
			return time.Now().In(utc)
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to retrieve SQL DB: %v", err)
	}

	if err := configureConnectionPool(sqlDB); err != nil {
		return nil, nil, err
	}

	runMigrate(db)

	return db, func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Failed to close the database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}, nil
}

func runMigrate(db *gorm.DB) {
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

func configureConnectionPool(sqlDB *sql.DB) error {
	// // กำหนดค่า Max Open Connections
	// postgres.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	// // กำหนดค่า Max Idle Connections (connection ที่เปิดรอแต่ยังไม่ได้ใช้)
	// postgres.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	// // กำหนดเวลาชีวิตของ connection (ระยะเวลาที่ connection สามารถใช้งานได้)
	// postgres.SetConnMaxLifetime(time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Minute)

	// ตรวจสอบค่าการตั้งค่า
	log.Printf("Database connection pool configured with MaxOpenConns: 20, MaxIdleConns: 10, ConnMaxLifetime: 5m")

	return nil
}
