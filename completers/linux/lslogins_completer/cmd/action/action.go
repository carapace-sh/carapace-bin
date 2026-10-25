package action

import "github.com/carapace-sh/carapace"

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"CONTEXT", "the user's security context",
		"FAILED-TTY", "where did the login fail?",
		"GECOS", "full user name",
		"HUSHED", "user's hush settings",
		"PROC", "number of processes run by the user",
		"PWD-CHANGE", "date of last password change",
		"PWD-EXPIR", "password expiration date",
		"PWD-MAX", "max number of days a password may remain unchanged",
		"PWD-MIN", "number of days required between changes",
		"PWD-WARN", "days user is warned of password expiration",
		"UID", "user ID",
		"USER", "user name",
	)
}
