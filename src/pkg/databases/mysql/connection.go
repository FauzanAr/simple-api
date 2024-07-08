package mysql

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"simple-api.com/m/src/config"
	"simple-api.com/m/src/pkg/logger"
)

type Mysql struct {
	ctx context.Context
	cfg config.MySql
	log logger.Logger
	db  *gorm.DB
}

func NewMysql(ctx context.Context, cfg config.MySql, log logger.Logger) *Mysql {
	return &Mysql{
		ctx: ctx,
		cfg: cfg,
		log: log,
	}
}

func (Mysqldb *Mysql) Connect() (*Mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Mysqldb.cfg.Username,
		Mysqldb.cfg.Password,
		Mysqldb.cfg.Host,
		Mysqldb.cfg.Port,
		Mysqldb.cfg.DatabaseName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize: 256,
	}))

	if err != nil {
		Mysqldb.log.Error(Mysqldb.ctx, "Error connecting to the database", err, nil)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		Mysqldb.log.Error(Mysqldb.ctx, "Error getting database instance", err, nil)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(15)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	Mysqldb.db = db

	Mysqldb.log.Info(Mysqldb.ctx, "Success connect to database!", nil)

	return Mysqldb, nil
}

func (mysql *Mysql) Close() error {
	sqlDB, err := mysql.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}