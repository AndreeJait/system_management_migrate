package constant

type MigrateDB string

const (
	MigrateDBManagementSystem MigrateDB = "management_system"
)

func GetMigrateDB(migrateDb string) MigrateDB {
	switch migrateDb {
	default:
		return MigrateDBManagementSystem
	}
}
