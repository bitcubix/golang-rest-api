package db

import (
	"github.com/bitcubix/go-rest-api-boilerplate/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Connection wrapper for gorm with own logging
type Connection struct {
	*gorm.DB
	log log.Logger
}

type Driver string

const (
	DriverMySQL Driver = "mysql"
)

//New returns a Connection with a opened database connection with gorm
func New(driver Driver, dsn string, logger log.Logger) (*Connection, error) {
	logger.WithFields(log.Fields{"driver": driver, "dsn": dsn}).Traceln("establish database connection")

	var db *gorm.DB
	var err error

	switch driver {
	case DriverMySQL:
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		return nil, ErrUnknownDriver
	}

	if err != nil {
		return nil, err
	}

	return &Connection{DB: db, log: logger}, nil
}

//Close closing the database connection
func (c *Connection) Close() error {
	c.log.Traceln("closing database connection")
	return c.Close()
}
