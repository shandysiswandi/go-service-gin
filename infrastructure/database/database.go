package database

import (
	"context"
	"errors"
	"fmt"
	"go-service-gin/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// all errros from database
var (
	ErrDatabaseConfigIsNil      = errors.New("database configuration is nil")
	ErrNotUseDatabase           = errors.New("application not using any database")
	ErrNotConnectMysql          = errors.New("can't connect database mysql")
	ErrNotConnectPostgres       = errors.New("can't connect database postgres")
	ErrNotConnectMongo          = errors.New("can't connect database mongo")
	ErrDatabaseDriverNotSupport = errors.New("driver database not support")
	ErrDialectNotSupport        = errors.New("dialect not support")
)

// Database is
type Database struct {
	SQL   *gorm.DB
	Mongo *mongo.Database
}

// NewDatabase is
func NewDatabase(dc *config.DatabaseConfig) (*Database, error) {
	if dc == nil {
		return nil, ErrDatabaseConfigIsNil
	}

	if dc.Driver == "" {
		return nil, ErrNotUseDatabase
	}

	var err error
	var db = &Database{}

	switch dc.Driver {
	case "mysql":
		temp := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf(temp, dc.Username, dc.Password, dc.Host, dc.Port, dc.Name)
		db.SQL, err = gormConnection(dsn, dc.Driver)
		if err != nil {
			return nil, ErrNotConnectMysql
		}
	case "postgres":
		temp := "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s"
		dsn := fmt.Sprintf(temp, dc.Username, dc.Password, dc.Host, dc.Port, dc.Name, dc.Timezone)

		db.SQL, err = gormConnection(dsn, dc.Driver)
		if err != nil {
			return nil, ErrNotConnectPostgres
		}
	case "mongo":
		temp := "mongodb://%s:%s@%s:%s/?readPreference=primary&ssl=false"
		dsn := fmt.Sprintf(temp, dc.Username, dc.Password, dc.Host, dc.Port)

		db.Mongo, err = mongoConnection(dsn, dc.Name)
		if err != nil {
			return nil, ErrNotConnectMongo
		}
	default:
		return nil, ErrDatabaseDriverNotSupport
	}

	return db, nil
}

func gormConnection(dsn, driver string) (db *gorm.DB, err error) {
	var dialect gorm.Dialector

	if driver == "postgres" {
		dialect = postgres.Open(dsn)
	} else {
		dialect = mysql.Open(dsn)
	}

	db, err = gorm.Open(dialect, &gorm.Config{PrepareStmt: true, Logger: logger.Default.LogMode(logger.Silent), SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}

	// pooling connection
	sqlCon, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlCon.SetMaxIdleConns(10 / 2)
	sqlCon.SetMaxOpenConns(100 / 2)
	sqlCon.SetConnMaxLifetime(time.Hour / 2)

	return db, nil
}

func mongoConnection(uri, database string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(database), nil
}
