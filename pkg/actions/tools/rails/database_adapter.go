package rails

import "github.com/carapace-sh/carapace"

// ActionDatabaseAdapters completes db:system:change --to values
//
//	mysql (MySQL adapter)
//	trilogy (MySQL via Trilogy)
func ActionDatabaseAdapters() carapace.Action {
	return carapace.ActionValuesDescribed(
		"mariadb-mysql", "MariaDB via mysql2",
		"mariadb-trilogy", "MariaDB via Trilogy",
		"mysql", "MySQL adapter",
		"postgresql", "PostgreSQL adapter",
		"sqlite3", "SQLite3 adapter",
		"trilogy", "MySQL via Trilogy",
	).Tag("database adapters").Uid("rails", "database-adapter")
}
