package action

import "github.com/carapace-sh/carapace"

// ActionTestSubcommands completes test:* subcommands
func ActionTestSubcommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "Run all tests including system tests",
		"system", "Run system tests only",
		"functionals", "Run functional tests",
		"units", "Run unit tests",
		"generators", "Run generator tests",
		"models", "Run model tests",
		"helpers", "Run helper tests",
		"channels", "Run channel tests",
		"controllers", "Run controller tests",
		"mailers", "Run mailer tests",
		"integration", "Run integration tests",
		"jobs", "Run job tests",
		"mailboxes", "Run mailbox tests",
		"db", "Reset test database and run tests",
	).Tag("test tasks").UidF(taskUid("test"))
}
