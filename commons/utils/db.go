package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"system_mananagement_migrate/config"
	"system_mananagement_migrate/shared/constant"
)

func getConnectionConfig(cfg config.Config, migrateDB constant.MigrateDB) config.DBConnection {
	var dbMigrate = map[constant.MigrateDB]config.DBConnection{
		constant.MigrateDBManagementSystem: cfg.DBManagementSystem,
	}
	return dbMigrate[migrateDB]
}

func ConnectToDB(cfg config.Config, migrateDb constant.MigrateDB) (*bun.DB, string) {
	connectionConfig := getConnectionConfig(cfg, migrateDb)
	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connectionConfig.Host,
		connectionConfig.Port,
		connectionConfig.Username,
		connectionConfig.Password,
		connectionConfig.DBName,
	)
	hsqldb, err := sql.Open("postgres", connection)
	if err != nil {
		logrus.Fatal(err)
	}

	// Create a Bun db on top of it.
	db := bun.NewDB(hsqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return db, connectionConfig.MigrationsFile
}
