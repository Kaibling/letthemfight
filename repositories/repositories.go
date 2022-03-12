package repositories

type DBMigrator interface {
	Migrate() error
}
