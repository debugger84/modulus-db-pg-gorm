package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func NewDb(cfg *ModuleConfig, log *GormLogger) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		PrepareStmt: false,
	}

	gormConfig.Logger = log

	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: "host=" + cfg.host +
					" port=" + strconv.Itoa(cfg.port) +
					" user=" + cfg.user +
					" dbname=" + cfg.name +
					" password=" + cfg.pass +
					" sslmode=" + cfg.sslMode,
				PreferSimpleProtocol: *cfg.preferSimpleProtocol,
			},
		), gormConfig,
	)

	if err != nil {
		return nil, err
	}

	if db != nil {
		sqlDB, err := db.DB()

		if err != nil {
			return nil, err
		}

		sqlDB.SetMaxIdleConns(cfg.maxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.maxOpenConns)
		sqlDB.SetConnMaxLifetime(cfg.connMaxLifetime)
	}

	if db != nil && *cfg.loggingEnabled {
		db = db.Debug()
	}

	if err != nil {
		return nil, err
	}

	return db, err
}
