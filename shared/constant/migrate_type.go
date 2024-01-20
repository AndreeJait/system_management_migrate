package constant

type MigrateType string

const (
	MigrateUp    MigrateType = "migrate-up"
	MigrateDown  MigrateType = "migrate-down"
	MigrateFresh MigrateType = "migrate-fresh"
)
