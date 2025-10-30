package config

import (
	"fmt"
	"log"

	"github.com/apextrade/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func ConnectDB() *DB {
	cnf, _ := Load()
	db_cnf := cnf.Database

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=require",
		db_cnf.Host, db_cnf.User, db_cnf.Password, db_cnf.Name, db_cnf.Port,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	postgresDB, _ := db.DB()
	postgresDB.SetMaxIdleConns(10)

	if err := db.AutoMigrate(&models.Stock{}, &models.Order{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("DB ready: Migrated models tables!")
	return &DB{DB: db}
}
