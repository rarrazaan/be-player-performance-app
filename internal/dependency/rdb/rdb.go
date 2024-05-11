package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RDB struct {
	PostgresDB *gorm.DB
}

func NewRDB(ctx context.Context, config config.Config) *RDB {
	return &RDB{
		PostgresDB: newPostgresConnection(ctx, config),
	}
}

func (r *RDB) Close(ctx context.Context) error {
	sqlDB, err := r.PostgresDB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func newPostgresConnection(ctx context.Context, config config.Config) *gorm.DB {
	var (
		host     = config.PostgresDB.DBHost
		port     = config.PostgresDB.DBPort
		user     = config.PostgresDB.DBUser
		password = config.PostgresDB.DBPassword
		dbName   = config.PostgresDB.DBName
		// timezone = config.Timezone
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect postgres, err: %s", err))
	}

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.PingContext(ctx); err != nil {
		panic(fmt.Sprintf("failed to ping postgres, err: %s", err))
	}

	var logLevel logger.LogLevel
	switch config.DBLogLevel {
	case "info":
		logLevel = logger.Info
	case "error":
		logLevel = logger.Error
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.LogLevel(logLevel),
				Colorful:      true,
			},
		),
		TranslateError: true,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to instantiate gorm, err: %s", err))
	}

	return db
}
