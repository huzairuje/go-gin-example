package database

import (
	"database/sql"
	"fmt"
	"github.com/go-gin-example/config"
	"github.com/go-gin-example/utils"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var conn *sql.DB

func ConnectDB(config *config.Config) *sql.DB {
	dbHost := config.DatabaseHost
	dbPort := config.DatabasePort
	dbUser := config.DatabaseUser
	dbPass := config.DatabasePassword
	dbName := config.DatabaseName

	var connection string
	if dbHost != "" ||  dbPort != ""  || dbUser != "" || dbPass != "" || dbName != "" {
		logrus.Info(utils.DatabaseConfigSet)
		connection = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	} else {
		logrus.Errorf(utils.DatabaseConfigNotSet)
		connection = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "postgres", "localhost", "5432", "loan_process")
	}

	db, err := sql.Open("postgres", connection)
	if err != nil {
		logrus.Errorf(utils.FailedOpenDb, err)
	}
	err = db.Ping()
	if err != nil {
		logrus.Errorf(utils.FailedPingDb, err)
	}
	SetConnection(db)
	return db
}

//GetConnection : Get Available Connection
func GetConnection() *sql.DB {
	return conn
}

//SetConnection : Set Available Connection
func SetConnection(connection *sql.DB) {
	conn = connection
}
