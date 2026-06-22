package rails

import "github.com/carapace-sh/carapace"

// ActionDatabases completes rails new --database values
//
//	sqlite3
//	postgresql
func ActionDatabases() carapace.Action {
	return carapace.ActionValues(
		"sqlite3",
		"postgresql",
		"mysql",
		"oracle",
		"sqlserver",
		"jdbc",
		"none",
	).Tag("databases").Uid("rails", "database")
}
