package gorm

import (
	"fmt"

	"github.com/jnjam6681/go-clean-architecture-rest-api/config"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func ConnectionDB(c *config.Config) (*gorm.DB, error) {

	fmt.Printf("host: %v\n", c.Postgres.Host)
	fmt.Printf("post: %v\n", c.Postgres.Port)
	fmt.Printf("username: %v\n", c.Postgres.Username)
	fmt.Printf("password: %v\n", c.Postgres.Password)
	fmt.Printf("dbname: %v\n", c.Postgres.Dbname)
	fmt.Printf("sslmode: %v\n", c.Postgres.SSLMode)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Dbname,
		c.Postgres.SSLMode,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, nil
}

func Migrate() {
	db.AutoMigrate(&entity.Todo{})
}
