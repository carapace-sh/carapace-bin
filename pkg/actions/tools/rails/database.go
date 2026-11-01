package rails

import "github.com/carapace-sh/carapace"

// ActionDatabases completes rails new --database values
//
//	sqlite3
//	postgresql
func ActionDatabases() carapace.Action {
	return carapace.ActionValues(
		"jdbc",
		"mysql",
		"none",
		"oracle",
		"postgresql",
		"sqlite3",
		"sqlserver",
	).Tag("databases").Uid("rails", "database")
}
