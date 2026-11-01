package action

import "github.com/carapace-sh/carapace"

// ActionTestSubcommands completes test:* subcommands
func ActionTestSubcommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "Run all tests including system tests",
		"channels", "Run channel tests",
		"controllers", "Run controller tests",
		"db", "Reset test database and run tests",
		"functionals", "Run functional tests",
		"generators", "Run generator tests",
		"helpers", "Run helper tests",
		"integration", "Run integration tests",
		"jobs", "Run job tests",
		"mailboxes", "Run mailbox tests",
		"mailers", "Run mailer tests",
		"models", "Run model tests",
		"system", "Run system tests only",
		"units", "Run unit tests",
	).Tag("test tasks").UidF(taskUid("test"))
}
