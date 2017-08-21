package database

import (
	"time"

	"github.com/albert-widi/sqlt"
	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/log"
	_ "github.com/lib/pq"
)

type dsn struct {
	Master   string
	Slave    string
	Retry    int
	Interval string
}

// Config of database
type Config struct {
	DSN map[string]*dsn
}

// DB of database
type db struct {
	connectedDbs map[dbType]*sqlt.DB
}

var dbObject *db

// Init database connection
func Init(cfg Config) error {
	dbObject = &db{connectedDbs: make(map[dbType]*sqlt.DB)}
	for dbName, dsn := range cfg.DSN {
		var (
			interval time.Duration
			err      error
		)
		if dsn.Interval != "" {
			interval, err = time.ParseDuration(dsn.Interval)
			if err != nil {
				return err
			}
		}

		for counter := 0; counter < dsn.Retry; counter++ {
			log.Debugf("[Database] Connecting to database [%s]...", dbName)
			newDB, err := sqlt.Open("postgres", dsn.Master+";"+dsn.Slave)
			if err != nil {
				log.Errorf("[Database] Failed to connect to db %s. Error: %s. Retrying in %s", dbName, err.Error(), dsn.Interval)
				// exit when retry is max
				if counter == dsn.Retry-1 {
					return err
				}
				// wait for retry interval
				time.Sleep(interval)
				continue
			}
			dbObject.connectedDbs[dbType(dbName)] = newDB
			break
		}
	}
	return nil
}

// dbType is type of database
type dbType string

// const of database type
const (
	TxDB dbType = "TxDB"
)

// Get database
func Get(dType dbType) (*sqlt.DB, error) {
	if dbConn, ok := dbObject.connectedDbs[dType]; ok {
		return dbConn, nil
	}
	return nil, errors.New(errors.DatabaseTypeNotExists)
}

// GetFatal database
func GetFatal(dType dbType) *sqlt.DB {
	if dbConn, ok := dbObject.connectedDbs[dType]; ok {
		return dbConn
	}
	log.Fatalf("Database with type %s is not exists", dType)
	return nil
}
