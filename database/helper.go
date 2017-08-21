package database

import (
	"context"
	"log"

	"github.com/albert-widi/sqlt"
)

// Watch db connection with heartbeat
func HeartBeat() {
	if dbObject == nil {
		return
	}
	for _, val := range dbObject.connectedDbs {
		if val == nil {
			continue
		}
		val.DoHeartBeat()
	}
}

// StopWatch will stop all heartbeat
func StopBeat() {
	if dbObject == nil {
		return
	}
	for _, val := range dbObject.connectedDbs {
		if val == nil {
			continue
		}
		val.StopBeat()
	}
}

func SetMaxConnection(ctx context.Context, maxnumber int) error {
	for _, val := range dbObject.connectedDbs {
		if val == nil {
			continue
		}
		val.SetMaxOpenConnections(maxnumber)
	}
	return nil
}

// Prepare will fatal all failed prepared query
func Prepare(ctx context.Context, db *sqlt.DB, query string) *sqlt.Stmt {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatalf("Failed to prepare query %s. Error: %s", query, err.Error())
	}
	return stmt
}

// Preparex will fatal all failed preparedx query
func Preparex(ctx context.Context, db *sqlt.DB, query string) *sqlt.Stmtx {
	stmtx, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Fatalf("Failed to preparex query %s. Error: %s", query, err.Error())
	}
	return stmtx
}
