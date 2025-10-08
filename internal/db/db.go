package db

import (
	"api-service/internal/config"
	"api-service/internal/model"
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB(cfg *config.DBConfig) *gorm.DB {
	sqlDB := setupDB(cfg)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	err = gormDB.AutoMigrate(
		&model.Blockchain{},
		&model.CryptoPrice{},
	)
	if err != nil {
		logrus.Fatalf("Failed to open GORM DB: %v", err)
	}

	if cfg.AppEnv == "debug" {
		gormDB = gormDB.Debug()
		logrus.Info("GORM debug mode enabled")
	}
	return gormDB
}

func setupDB(cfg *config.DBConfig) *sql.DB {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SslMode,
	)

	if cfg.SslMode != "disable" {
		authPemPath := filepath.Join(".", "config", "auth.pem")
		dataSourceName += fmt.Sprintf(" sslrootcert=%s", authPemPath)
	}

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		logrus.Fatalf("Failed to connect to DB: %v", err)
	}

	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLife) * time.Second)

	return db
}
