package rails

import "github.com/carapace-sh/carapace"

// ActionDatabaseAdapters completes db:system:change --to values
//
//	mysql (MySQL adapter)
//	trilogy (MySQL via Trilogy)
func ActionDatabaseAdapters() carapace.Action {
	return carapace.ActionValuesDescribed(
		"mysql", "MySQL adapter",
		"trilogy", "MySQL via Trilogy",
		"postgresql", "PostgreSQL adapter",
		"sqlite3", "SQLite3 adapter",
		"mariadb-mysql", "MariaDB via mysql2",
		"mariadb-trilogy", "MariaDB via Trilogy",
	).Tag("database adapters").Uid("rails", "database-adapter")
}
