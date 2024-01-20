package utils

import (
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"system_mananagement_migrate/config"
	"system_mananagement_migrate/shared/constant"
)

func MigrateUp(cfg config.Config, dbMigrate constant.MigrateDB) error {
	db, migrateFiles := ConnectToDB(cfg, dbMigrate)
	migrations := &migrate.FileMigrationSource{
		Dir: migrateFiles,
	}

	count, err := migrate.Exec(db.DB, constant.PostgresDialect, migrations, migrate.Up)
	if err != nil {
		return err
	}
	logrus.Infof("migrate-up: applied %d migrations", count)
	return nil
}

func MigrateDown(cfg config.Config, dbMigrate constant.MigrateDB) error {
	db, migrateFiles := ConnectToDB(cfg, dbMigrate)
	migrations := &migrate.FileMigrationSource{
		Dir: migrateFiles,
	}

	count, err := migrate.Exec(db.DB, constant.PostgresDialect, migrations, migrate.Down)
	if err != nil {
		return err
	}
	logrus.Infof("migrate-up: applied %d migrations", count)
	return nil
}

func MigrateFresh(cfg config.Config, dbMigrate constant.MigrateDB) error {
	db, _ := ConnectToDB(cfg, dbMigrate)

	logrus.Info("Drop database")
	_, err := db.Exec(`DROP SCHEMA public CASCADE; CREATE SCHEMA public;`)
	if err != nil {
		return err
	}

	db, migrateFiles := ConnectToDB(cfg, dbMigrate)
	migrations := &migrate.FileMigrationSource{
		Dir: migrateFiles,
	}

	count, err := migrate.Exec(db.DB, constant.PostgresDialect, migrations, migrate.Up)
	if err != nil {
		return err
	}
	logrus.Infof("migrate-up: applied %d migrations", count)
	return nil
}
