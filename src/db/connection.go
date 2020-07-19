package db

import (
	"context"

	pg "github.com/go-pg/pg/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, err := q.FormattedQuery()
	if err != nil {
		logrus.Trace("Could not get the query")
	}
	logrus.Tracef("[DB] %s", query)
	return nil
}

// Connect connects to the database
func Connect() *pg.DB {
	logrus.Trace("[DB] Establishing a connection to the database")

	// define the database connection
	db := pg.Connect(&pg.Options{
		Addr:     viper.GetString("db.addr"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
	})

	db.AddQueryHook(dbLogger{})

	logrus.Trace("[DB] Done")

	return db
}

// Disconnect closes the database connection
func Disconnect(db *pg.DB) {
	logrus.Trace("[DB] Closing the database connection")
	db.Close()
	return
}
